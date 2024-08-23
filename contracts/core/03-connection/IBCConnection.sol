// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Height} from "../../proto/Client.sol";
import {ConnectionEnd, Version, Counterparty} from "../../proto/Connection.sol";
import {MerklePrefix} from "../../proto/Commitment.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {IIBCConnectionErrors} from "../03-connection/IIBCConnectionErrors.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {IBCSelfStateValidator} from "../24-host/IBCSelfStateValidator.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCConnectionLib} from "./IBCConnectionLib.sol";

/**
 * @dev IBCConnection is a contract that implements [ICS-3](https://github.com/cosmos/ibc/tree/main/spec/core/ics-003-connection-semantics).
 */
abstract contract IBCConnection is IBCHost, IBCSelfStateValidator, IIBCConnection, IIBCConnectionErrors {
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
        ConnectionEnd.Data storage connection = getConnectionStorage()[connectionId].connection;
        if (connection.state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED) {
            revert IBCConnectionAlreadyConnectionExists();
        }
        // ensure the client exists
        checkAndGetClient(msg_.clientId);
        if (bytes(msg_.counterparty.connection_id).length > 0) {
            revert IBCConnectionInvalidCounterpartyConnectionIdentifier(msg_.counterparty.connection_id);
        }
        connection.client_id = msg_.clientId;

        if (msg_.version.features.length > 0) {
            if (!IBCConnectionLib.isSupportedVersion(getCompatibleVersions(), msg_.version)) {
                revert IBCConnectionIBCVersionNotSupported();
            }
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
        if (msg_.counterpartyVersions.length == 0) {
            revert IBCConnectionEmptyConnectionCounterpartyVersions();
        } else if (!validateSelfClient(msg_.clientStateBytes)) {
            revert IBCConnectionInvalidSelfClientState();
        }
        bytes memory selfConsensusState = getSelfConsensusState(msg_.consensusHeight, msg_.hostConsensusStateProof);

        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = getConnectionStorage()[connectionId].connection;
        if (connection.state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED) {
            revert IBCConnectionAlreadyConnectionExists();
        }
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

        verifyConnectionState(
            connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection
        );
        verifyClientState(
            connection,
            msg_.proofHeight,
            IBCCommitment.clientStatePath(connection.counterparty.client_id),
            msg_.proofClient,
            msg_.clientStateBytes
        );
        verifyClientConsensusState(
            connection, msg_.proofHeight, msg_.consensusHeight, msg_.proofConsensus, selfConsensusState
        );

        updateConnectionCommitment(connectionId);
        emit GeneratedConnectionIdentifier(connectionId);
        return connectionId;
    }

    /**
     * @dev connectionOpenAck relays acceptance of a connection open attempt from chain B back
     * to chain A (this code is executed on chain A).
     */
    function connectionOpenAck(IIBCConnection.MsgConnectionOpenAck calldata msg_) external override {
        ConnectionEnd.Data storage connection = getConnectionStorage()[msg_.connectionId].connection;
        if (connection.state != ConnectionEnd.State.STATE_INIT) {
            revert IBCConnectionUnexpectedConnectionState(connection.state);
        }
        if (!IBCConnectionLib.isSupportedVersion(connection.versions, msg_.version)) {
            revert IBCConnectionIBCVersionNotSupported();
        }
        if (!validateSelfClient(msg_.clientStateBytes)) {
            revert IBCConnectionInvalidSelfClientState();
        }
        bytes memory selfConsensusState = getSelfConsensusState(msg_.consensusHeight, msg_.hostConsensusStateProof);

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

        verifyConnectionState(
            connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionId, expectedConnection
        );
        verifyClientState(
            connection,
            msg_.proofHeight,
            IBCCommitment.clientStatePath(connection.counterparty.client_id),
            msg_.proofClient,
            msg_.clientStateBytes
        );
        verifyClientConsensusState(
            connection, msg_.proofHeight, msg_.consensusHeight, msg_.proofConsensus, selfConsensusState
        );

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
        ConnectionEnd.Data storage connection = getConnectionStorage()[msg_.connectionId].connection;
        if (connection.state != ConnectionEnd.State.STATE_TRYOPEN) {
            revert IBCConnectionUnexpectedConnectionState(connection.state);
        }

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

        verifyConnectionState(
            connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection
        );

        connection.state = ConnectionEnd.State.STATE_OPEN;
        updateConnectionCommitment(msg_.connectionId);
    }

    /**
     * @dev getCompatibleVersions returns the supported versions of the host chain.
     */
    function getCompatibleVersions() public pure virtual returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        versions[0] = IBCConnectionLib.defaultIBCVersion();
        return versions;
    }

    // --------- Private Functions --------- //

    function generateConnectionIdentifier() private returns (string memory) {
        HostStorage storage hostStorage = getHostStorage();
        string memory identifier =
            string(abi.encodePacked("connection-", Strings.toString(hostStorage.nextConnectionSequence)));
        hostStorage.nextConnectionSequence++;
        return identifier;
    }

    function updateConnectionCommitment(string memory connectionId) private {
        getCommitments()[IBCCommitment.connectionCommitmentKey(connectionId)] =
            keccak256(ConnectionEnd.encode(getConnectionStorage()[connectionId].connection));
    }

    function verifyClientState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory path,
        bytes memory proof,
        bytes memory clientStateBytes
    ) private {
        // slither-disable-start reentrancy-no-eth
        if (
            checkAndGetClient(connection.client_id).verifyMembership(
                connection.client_id,
                height,
                0,
                0,
                proof,
                connection.counterparty.prefix.key_prefix,
                path,
                clientStateBytes
            )
        ) {
            // slither-disable-end reentrancy-no-eth
            return;
        }
        revert IBCConnectionFailedVerifyClientState(connection.client_id, path, clientStateBytes, proof, height);
    }

    function verifyClientConsensusState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        Height.Data memory consensusHeight,
        bytes memory proof,
        bytes memory consensusStateBytes
    ) private {
        // slither-disable-start reentrancy-no-eth
        if (
            checkAndGetClient(connection.client_id).verifyMembership(
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
            )
        ) {
            // slither-disable-end reentrancy-no-eth
            return;
        }
        revert IBCConnectionFailedVerifyClientConsensusState(
            connection.client_id,
            IBCCommitment.consensusStatePath(
                connection.counterparty.client_id, consensusHeight.revision_number, consensusHeight.revision_height
            ),
            consensusStateBytes,
            proof,
            height
        );
    }

    function verifyConnectionState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory proof,
        string memory counterpartyConnectionId,
        ConnectionEnd.Data memory counterpartyConnection
    ) private {
        // slither-disable-start reentrancy-no-eth
        if (
            checkAndGetClient(connection.client_id).verifyMembership(
                connection.client_id,
                height,
                0,
                0,
                proof,
                connection.counterparty.prefix.key_prefix,
                IBCCommitment.connectionPath(counterpartyConnectionId),
                ConnectionEnd.encode(counterpartyConnection)
            )
        ) {
            // slither-disable-end reentrancy-no-eth
            return;
        }
        revert IBCConnectionFailedVerifyConnectionState(
            connection.client_id,
            IBCCommitment.connectionPath(counterpartyConnectionId),
            ConnectionEnd.encode(counterpartyConnection),
            proof,
            height
        );
    }
}
