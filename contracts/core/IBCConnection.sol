pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Connection.sol";
import "./IBCStore.sol";
import "./IBCClient.sol";
import "./IBCMsgs.sol";
import "./IHandler.sol";

abstract contract IBCConnection is IHandler, IBCClient {

    string constant commitmentPrefix = "ibc";

    /* Public functions */

    // ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
    // is returned.
    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit memory msg_) public override returns (string memory) {
        ConnectionEnd.Data memory connection;
        bool found;
        (connection, found) = ibcStore.getConnection(msg_.connectionId);
        require(!found, "connection already exists");

        connection = ConnectionEnd.Data({
            client_id: msg_.clientId,
            versions: getVersions(),
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: msg_.counterparty
        });
        ibcStore.setConnection(msg_.connectionId, connection);
        addConnectionToClient(msg_.clientId, msg_.connectionId);
        return msg_.connectionId;
    }

    // ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
    // code is executed on chain B).
    function connectionOpenTry(
        IBCMsgs.MsgConnectionOpenTry memory msg_
    ) public override returns (string memory) {
        require(msg_.consensusHeight < block.number, "consensus height is greater than or equal to the current block height");
        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        // TODO add a support for selfConsensusState getter
        // (ConsensusState.Data memory expectedConsensusState, bool found) = client.getSelfConsensusState(consensusHeight);

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

        require(verifyConnectionState(connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection), "failed to verify connection state");
        // TODO commentout this after connectionState validation is passed
        // require(verifyClientState(connection, msg_.proofHeight, msg_.proofClient, msg_.clientState), "failed to verify clientState");
        // TODO commentout this
        // require(verifyClientConsensusState(connection, proofHeight, consensusHeight, proofConsensus, expectedConsensusState), "failed to verify consensusState");

        addConnectionToClient(msg_.clientId, msg_.connectionId);
        ibcStore.setConnection(msg_.connectionId, connection);
        return msg_.connectionId;
    }

    function connectionOpenAck(
        IBCMsgs.MsgConnectionOpenAck memory msg_
    ) public override {
        require(msg_.consensusHeight < block.number, "consensus height is greater than or equal to the current block height");
        (ConnectionEnd.Data memory connection, bool found) = ibcStore.getConnection(msg_.connectionId);
        require(found, "connection not found");

        if (connection.state != ConnectionEnd.State.STATE_INIT && connection.state != ConnectionEnd.State.STATE_TRYOPEN) {
            revert("connection state is not INIT or TRYOPEN");
        } else if (connection.state == ConnectionEnd.State.STATE_INIT && !isSupportedVersion(msg_.version)) {
            revert("connection state is in INIT but the provided version is not supported");
        } else if (connection.state == ConnectionEnd.State.STATE_TRYOPEN && (connection.versions.length != 1 || !isEqualVersion(connection.versions[0], msg_.version))) {
            revert("connection state is in TRYOPEN but the provided version is not set in the previous connection versions"); 
        }

        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");

        // TODO add a support for selfConsensusState getter
        // (ConsensusState.Data memory expectedConsensusState, bool found) = client.getSelfConsensusState(consensusHeight);
        
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

        require(verifyConnectionState(connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionID, expectedConnection), "failed to verify connection state");

        connection.state = ConnectionEnd.State.STATE_OPEN;
        connection.versions = expectedConnection.versions;
        connection.counterparty.connection_id = msg_.counterpartyConnectionID;
        ibcStore.setConnection(msg_.connectionId, connection);
    }

    function connectionOpenConfirm(
        IBCMsgs.MsgConnectionOpenConfirm memory msg_
    ) public override {
        (ConnectionEnd.Data memory connection, bool found) = ibcStore.getConnection(msg_.connectionId);
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

        require(verifyConnectionState(connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection), "failed to verify connection state");

        connection.state = ConnectionEnd.State.STATE_OPEN;
        ibcStore.setConnection(msg_.connectionId, connection);
    }

    // Verification functions

    function verifyConnectionState(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory connectionId, ConnectionEnd.Data memory counterpartyConnection) internal view returns (bool) {
        return getClient(connection.client_id).verifyConnectionState(connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, connectionId, ConnectionEnd.encode(counterpartyConnection));
    }

    function verifyClientState(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, bytes memory clientStateBytes) internal view returns (bool) {
        return getClient(connection.client_id).verifyClientState(connection.client_id, height, connection.counterparty.prefix.key_prefix, connection.counterparty.client_id, proof, clientStateBytes);
    }

    function verifyClientConsensusStateWithConnection(ConnectionEnd.Data memory connection, uint64 height, uint64 consensusHeight, bytes memory proof, bytes memory consensusStateBytes) internal view returns (bool) {
        return getClient(connection.client_id).verifyClientConsensusState(connection.client_id, height, connection.counterparty.client_id, consensusHeight, connection.counterparty.prefix.key_prefix, proof, consensusStateBytes);
    }

    function verifyChannelState(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory portId, string memory channelId, bytes memory channelBytes) internal view returns (bool) {
        return getClient(connection.client_id).verifyChannelState(connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, portId, channelId, channelBytes);
    }

    function verifyPacketCommitment(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory portId, string memory channelId, uint64 sequence, bytes32 commitmentBytes) internal view returns (bool) {
        return getClient(connection.client_id).verifyPacketCommitment(connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, portId, channelId, sequence, commitmentBytes);
    }

    function verifyPacketAcknowledgement(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory portId, string memory channelId, uint64 sequence, bytes32 ackCommitmentBytes) internal view returns (bool) {
        return getClient(connection.client_id).verifyPacketAcknowledgement(connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, portId, channelId, sequence, ackCommitmentBytes);
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
    function isSupportedVersion(Version.Data memory proposedVersion) internal view returns (bool) {
        return true;
    }

    function isEqualVersion(Version.Data memory a, Version.Data memory b) internal view returns (bool) {
        return keccak256(Version.encode(a)) == keccak256(Version.encode(b));
    }

    function makeVersionArray(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }

    function addConnectionToClient(
        string memory clientId,
        string memory connectionId
    ) internal {
        require(ibcStore.hasClientState(clientId), "client not found");
        ibcStore.addConnectionPath(clientId, connectionId);
    }
}
