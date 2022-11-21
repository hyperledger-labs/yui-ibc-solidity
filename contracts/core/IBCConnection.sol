// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../proto/Client.sol";
import "../proto/Connection.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";
import "./IBCCommitment.sol";

contract IBCConnection is IBCHost {
    string private constant commitmentPrefix = "ibc";

    /* Handshake functions */

    // ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
    // is returned.
    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit calldata msg_) public returns (string memory) {
        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = connections[connectionId];
        connection.client_id = msg_.clientId;
        setVersions(connection.versions);
        connection.state = ConnectionEnd.State.STATE_INIT;
        connection.delay_period = msg_.delayPeriod;
        connection.counterparty = msg_.counterparty;
        return connectionId;
    }

    // ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
    // code is executed on chain B).
    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry calldata msg_) public returns (string memory) {
        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = connections[connectionId];
        connection.client_id = msg_.clientId;
        setVersions(connection.versions);
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
                prefix: MerklePrefix.Data({key_prefix: bytes(commitmentPrefix)})
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

        return connectionId;
    }

    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck calldata msg_) public {
        ConnectionEnd.Data storage connection = connections[msg_.connectionId];
        if (connection.state != ConnectionEnd.State.STATE_INIT && connection.state != ConnectionEnd.State.STATE_TRYOPEN)
        {
            revert("connection state is not INIT or TRYOPEN");
        } else if (connection.state == ConnectionEnd.State.STATE_INIT && !isSupportedVersion(msg_.version)) {
            revert("connection state is in INIT but the provided version is not supported");
        } else if (
            connection.state == ConnectionEnd.State.STATE_TRYOPEN
                && (connection.versions.length != 1 || !isEqualVersion(connection.versions[0], msg_.version))
        ) {
            revert(
                "connection state is in TRYOPEN but the provided version is not set in the previous connection versions"
            );
        }

        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: bytes(commitmentPrefix)})
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: makeVersionArray(msg_.version),
            state: ConnectionEnd.State.STATE_TRYOPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionID, expectedConnection
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
        copyVersions(expectedConnection.versions, connection.versions);
        connection.counterparty.connection_id = msg_.counterpartyConnectionID;
    }

    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm calldata msg_) public {
        ConnectionEnd.Data storage connection = connections[msg_.connectionId];
        require(connection.state == ConnectionEnd.State.STATE_TRYOPEN, "connection state is not TRYOPEN");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: bytes(commitmentPrefix)})
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: connection.versions,
            state: ConnectionEnd.State.STATE_OPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection
            ),
            "failed to verify connection state"
        );

        connection.state = ConnectionEnd.State.STATE_OPEN;
    }

    /* Verification functions */

    function verifyClientState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory path,
        bytes memory proof,
        bytes memory clientStateBytes
    ) private returns (bool) {
        return getClient(connection.client_id).verifyMembership(
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
        return getClient(connection.client_id).verifyMembership(
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
        return getClient(connection.client_id).verifyMembership(
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

    /* Internal functions */

    function setVersions(Version.Data[] storage versions) internal {
        Version.Data storage version = versions[0];
        version.identifier = "1";
        version.features.push("ORDER_ORDERED");
        version.features.push("ORDER_UNORDERED");
    }

    // TODO implements
    function isSupportedVersion(Version.Data memory) internal pure returns (bool) {
        return true;
    }

    function isEqualVersion(Version.Data memory a, Version.Data memory b) internal pure returns (bool) {
        return keccak256(Version.encode(a)) == keccak256(Version.encode(b));
    }

    function makeVersionArray(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }

    function copyVersions(Version.Data[] memory src, Version.Data[] storage dst) internal {
        for (uint256 i = 0; i < src.length; i++) {
            copyVersion(src[i], dst[i]);
        }
    }

    function copyVersion(Version.Data memory src, Version.Data storage dst) internal {
        dst.identifier = src.identifier;
        for (uint256 i = 0; i < src.features.length; i++) {
            dst.features.push(src.features[i]);
        }
    }
}
