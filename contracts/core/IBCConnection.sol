pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Connection.sol";
import "./IBCClient.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";

library IBCConnection {

    string constant commitmentPrefix = "ibc";

    /* Public functions */

    // ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
    // is returned.
    function connectionOpenInit(IBCHost host, IBCMsgs.MsgConnectionOpenInit memory msg_) public returns (string memory) {
        host.onlyIBCModule();
        ConnectionEnd.Data memory connection = ConnectionEnd.Data({
            client_id: msg_.clientId,
            versions: getVersions(),
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: msg_.counterparty
        });
        string memory connectionId = host.generateConnectionIdentifier();
        host.setConnection(connectionId, connection);
        return connectionId;
    }

    // ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
    // code is executed on chain B).
    function connectionOpenTry(
        IBCHost host,
        IBCMsgs.MsgConnectionOpenTry memory msg_
    ) public returns (string memory) {
        host.onlyIBCModule();
        require(IBCClient.validateSelfClient(host, msg_.clientStateBytes), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        ConnectionEnd.Data memory connection = ConnectionEnd.Data({
            client_id: msg_.clientId,
            versions: getVersions(),
            state: ConnectionEnd.State.STATE_TRYOPEN,
            delay_period: msg_.delayPeriod,
            counterparty: msg_.counterparty
        });

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

        require(verifyConnectionState(host, connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection), "failed to verify connection state");
        require(verifyClientState(host, connection, msg_.proofHeight, msg_.proofClient, msg_.clientStateBytes), "failed to verify clientState");
        // TODO we should also verify a consensus state

        string memory connectionId = host.generateConnectionIdentifier();
        host.setConnection(connectionId, connection);
        return connectionId;
    }

    function connectionOpenAck(
        IBCHost host,
        IBCMsgs.MsgConnectionOpenAck memory msg_
    ) public {
        host.onlyIBCModule();
        (ConnectionEnd.Data memory connection, bool found) = host.getConnection(msg_.connectionId);
        require(found, "connection not found");

        if (connection.state != ConnectionEnd.State.STATE_INIT && connection.state != ConnectionEnd.State.STATE_TRYOPEN) {
            revert("connection state is not INIT or TRYOPEN");
        } else if (connection.state == ConnectionEnd.State.STATE_INIT && !isSupportedVersion(msg_.version)) {
            revert("connection state is in INIT but the provided version is not supported");
        } else if (connection.state == ConnectionEnd.State.STATE_TRYOPEN && (connection.versions.length != 1 || !isEqualVersion(connection.versions[0], msg_.version))) {
            revert("connection state is in TRYOPEN but the provided version is not set in the previous connection versions"); 
        }

        require(IBCClient.validateSelfClient(host, msg_.clientStateBytes), "failed to validate self client state");

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

        require(verifyConnectionState(host, connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionID, expectedConnection), "failed to verify connection state");
        require(verifyClientState(host, connection, msg_.proofHeight, msg_.proofClient, msg_.clientStateBytes), "failed to verify clientState");
        // TODO we should also verify a consensus state

        connection.state = ConnectionEnd.State.STATE_OPEN;
        connection.versions = expectedConnection.versions;
        connection.counterparty.connection_id = msg_.counterpartyConnectionID;
        host.setConnection(msg_.connectionId, connection);
    }

    function connectionOpenConfirm(
        IBCHost host,
        IBCMsgs.MsgConnectionOpenConfirm memory msg_
    ) public {
        host.onlyIBCModule();
        (ConnectionEnd.Data memory connection, bool found) = host.getConnection(msg_.connectionId);
        require(found, "connection not found");

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

        require(verifyConnectionState(host, connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection), "failed to verify connection state");

        connection.state = ConnectionEnd.State.STATE_OPEN;
        host.setConnection(msg_.connectionId, connection);
    }

    // Verification functions

    function verifyConnectionState(IBCHost host, ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory connectionId, ConnectionEnd.Data memory counterpartyConnection) internal view returns (bool) {
        return IBCClient.getClient(host, connection.client_id).verifyConnectionState(host, connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, connectionId, ConnectionEnd.encode(counterpartyConnection));
    }

    function verifyClientState(IBCHost host, ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, bytes memory clientStateBytes) internal view returns (bool) {
        return IBCClient.getClient(host, connection.client_id).verifyClientState(host, connection.client_id, height, connection.counterparty.prefix.key_prefix, connection.counterparty.client_id, proof, clientStateBytes);
    }

    function verifyClientConsensusStateWithConnection(IBCHost host, ConnectionEnd.Data memory connection, uint64 height, uint64 consensusHeight, bytes memory proof, bytes memory consensusStateBytes) internal view returns (bool) {
        return IBCClient.getClient(host, connection.client_id).verifyClientConsensusState(host, connection.client_id, height, connection.counterparty.client_id, consensusHeight, connection.counterparty.prefix.key_prefix, proof, consensusStateBytes);
    }

    function verifyChannelState(IBCHost host, ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory portId, string memory channelId, bytes memory channelBytes) public view returns (bool) {
        return IBCClient.getClient(host, connection.client_id).verifyChannelState(host, connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, portId, channelId, channelBytes);
    }

    function verifyPacketCommitment(IBCHost host, ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory portId, string memory channelId, uint64 sequence, bytes32 commitmentBytes) public view returns (bool) {
        return IBCClient.getClient(host, connection.client_id).verifyPacketCommitment(host, connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, portId, channelId, sequence, commitmentBytes);
    }

    function verifyPacketAcknowledgement(IBCHost host, ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory portId, string memory channelId, uint64 sequence, bytes memory acknowledgement) public view returns (bool) {
        return IBCClient.getClient(host, connection.client_id).verifyPacketAcknowledgement(host, connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, portId, channelId, sequence, acknowledgement);
    }

    // Internal functions

    function getVersions() internal pure returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        string[] memory features = new string[](2);
        features[0] = "ORDER_ORDERED";
        features[1] = "ORDER_UNORDERED";
        versions[0] = Version.Data({
            identifier: "1",
            features: features
        });
        return versions;
    }

    // TODO implements
    function isSupportedVersion(Version.Data memory proposedVersion) internal pure returns (bool) {
        return true;
    }

    function isEqualVersion(Version.Data memory a, Version.Data memory b) internal pure returns (bool) {
        return keccak256(Version.encode(a)) == keccak256(Version.encode(b));
    }

    function makeVersionArray(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }
}
