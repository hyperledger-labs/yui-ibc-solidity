// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Height} from "../../proto/Client.sol";
import {ConnectionEnd, Version, Counterparty} from "../../proto/Connection.sol";
import {MerklePrefix} from "../../proto/Commitment.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {IBCSelfStateValidator} from "../24-host/IBCSelfStateValidator.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCConnectionLib} from "./IBCConnectionLib.sol";

/**
 * @dev IBCConnection is a contract that implements [ICS-3](https://github.com/cosmos/ibc/tree/main/spec/core/ics-003-connection-semantics).
 */
abstract contract IBCConnection is IBCHost, IBCSelfStateValidator, IIBCConnection {
    /**
     * @dev connectionOpenInit initialises a connection attempt on chain A. The generated connection identifier
     * is returned.
     */
    function connectionOpenInit(IIBCConnection.MsgConnectionOpenInit calldata msg_)
        external
        override
        returns (string memory)
    {
        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = connections[connectionId];
        require(connection.state == ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED, "connectionId already exists");
        // ensure the client exists
        checkAndGetClient(msg_.clientId);
        require(bytes(msg_.counterparty.connection_id).length == 0, "counterparty connectionId must be empty");
        connection.client_id = msg_.clientId;

        if (msg_.version.features.length > 0) {
            require(
                IBCConnectionLib.isSupportedVersion(getCompatibleVersions(), msg_.version),
                "the selected version is not supported"
            );
            connection.versions.push(msg_.version);
        } else {
            IBCConnectionLib.setSupportedVersions(getCompatibleVersions(), connection.versions);
        }

        connection.state = ConnectionEnd.State.STATE_INIT;
        connection.delay_period = msg_.delayPeriod;
        connection.counterparty = msg_.counterparty;
        updateConnectionCommitment(connectionId);
        emit GeneratedConnectionIdentifier(connectionId);
        return connectionId;
    }

    /**
     * @dev connectionOpenTry relays notice of a connection attempt on chain A to chain B (this
     * code is executed on chain B).
     */
    function connectionOpenTry(IIBCConnection.MsgConnectionOpenTry calldata msg_)
        external
        override
        returns (string memory)
    {
        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = connections[connectionId];
        require(connection.state == ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED, "connectionId already exists");
        // ensure the client exists
        checkAndGetClient(msg_.clientId);

        connection.versions.push(IBCConnectionLib.pickVersion(getCompatibleVersions(), msg_.counterpartyVersions));
        connection.client_id = msg_.clientId;
        connection.state = ConnectionEnd.State.STATE_TRYOPEN;
        connection.delay_period = msg_.delayPeriod;
        connection.counterparty = msg_.counterparty;

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: msg_.counterparty.client_id,
            versions: msg_.counterpartyVersions,
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: Counterparty.Data({
                client_id: msg_.clientId,
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: _getCommitmentPrefix()})
            })
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection
            ),
            "failed to verify connection state"
        );
        require(
            verifyClientState(
                connection,
                msg_.proofHeight,
                IBCCommitment.clientStatePath(connection.counterparty.client_id),
                msg_.proofClient,
                msg_.clientStateBytes
            ),
            "failed to verify clientState"
        );
        // TODO we should also verify a consensus state

        updateConnectionCommitment(connectionId);
        emit GeneratedConnectionIdentifier(connectionId);
        return connectionId;
    }

    /**
     * @dev connectionOpenAck relays acceptance of a connection open attempt from chain B back
     * to chain A (this code is executed on chain A).
     */
    function connectionOpenAck(IIBCConnection.MsgConnectionOpenAck calldata msg_) external override {
        ConnectionEnd.Data storage connection = connections[msg_.connectionId];
        require(connection.state == ConnectionEnd.State.STATE_INIT, "connection state is not INIT");
        require(
            IBCConnectionLib.isSupportedVersion(connection.versions, msg_.version),
            "the counterparty selected version is not supported by versions selected on INIT"
        );
        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: IBCConnectionLib.newVersions(msg_.version),
            state: ConnectionEnd.State.STATE_TRYOPEN,
            delay_period: connection.delay_period,
            counterparty: Counterparty.Data({
                client_id: connection.client_id,
                connection_id: msg_.connectionId,
                prefix: MerklePrefix.Data({key_prefix: _getCommitmentPrefix()})
            })
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionId, expectedConnection
            ),
            "failed to verify connection state"
        );
        require(
            verifyClientState(
                connection,
                msg_.proofHeight,
                IBCCommitment.clientStatePath(connection.counterparty.client_id),
                msg_.proofClient,
                msg_.clientStateBytes
            ),
            "failed to verify clientState"
        );
        // TODO we should also verify a consensus state

        connection.state = ConnectionEnd.State.STATE_OPEN;
        connection.counterparty.connection_id = msg_.counterpartyConnectionId;
        IBCConnectionLib.copyVersions(expectedConnection.versions, connection.versions);
        updateConnectionCommitment(msg_.connectionId);
    }

    /**
     * @dev connectionOpenConfirm confirms opening of a connection on chain A to chain B, after
     * which the connection is open on both chains (this code is executed on chain B).
     */
    function connectionOpenConfirm(IIBCConnection.MsgConnectionOpenConfirm calldata msg_) external override {
        ConnectionEnd.Data storage connection = connections[msg_.connectionId];
        require(connection.state == ConnectionEnd.State.STATE_TRYOPEN, "connection state is not TRYOPEN");

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: connection.versions,
            state: ConnectionEnd.State.STATE_OPEN,
            delay_period: connection.delay_period,
            counterparty: Counterparty.Data({
                client_id: connection.client_id,
                connection_id: msg_.connectionId,
                prefix: MerklePrefix.Data({key_prefix: _getCommitmentPrefix()})
            })
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection
            ),
            "failed to verify connection state"
        );

        connection.state = ConnectionEnd.State.STATE_OPEN;
        updateConnectionCommitment(msg_.connectionId);
    }

    function updateConnectionCommitment(string memory connectionId) private {
        commitments[IBCCommitment.connectionCommitmentKey(connectionId)] =
            keccak256(ConnectionEnd.encode(connections[connectionId]));
    }

    /* Verification functions */

    function verifyClientState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory path,
        bytes memory proof,
        bytes memory clientStateBytes
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
            connection.client_id, height, 0, 0, proof, connection.counterparty.prefix.key_prefix, path, clientStateBytes
        );
    }

    function verifyClientConsensusState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        Height.Data memory consensusHeight,
        bytes memory proof,
        bytes memory consensusStateBytes
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCCommitment.consensusStatePath(
                connection.counterparty.client_id, consensusHeight.revision_number, consensusHeight.revision_height
            ),
            consensusStateBytes
        );
    }

    function verifyConnectionState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory proof,
        string memory connectionId,
        ConnectionEnd.Data memory counterpartyConnection
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCCommitment.connectionPath(connectionId),
            ConnectionEnd.encode(counterpartyConnection)
        );
    }

    /**
     * @dev getCompatibleVersions returns the supported versions of the host chain.
     */
    function getCompatibleVersions() public pure virtual returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        versions[0] = IBCConnectionLib.defaultIBCVersion();
        return versions;
    }

    /* Internal functions */

    function generateConnectionIdentifier() private returns (string memory) {
        string memory identifier = string(abi.encodePacked("connection-", Strings.toString(nextConnectionSequence)));
        nextConnectionSequence++;
        return identifier;
    }
}
