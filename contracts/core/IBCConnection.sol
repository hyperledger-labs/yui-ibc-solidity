pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Connection.sol";
import "./ProvableStore.sol";
import "./IBCClient.sol";

contract IBCConnection {
    ProvableStore provableStore;
    IBCClient client;

    // types
    struct MsgConnectionOpenTry {
        string connectionId;
        Counterparty.Data counterparty; // counterpartyConnectionIdentifier, counterpartyPrefix and counterpartyClientIdentifier
        uint64 delayPeriod;
        string clientId; // clientID of chainA
        ClientState.Data clientState; // clientState that chainA has for chainB
        Version.Data[] counterpartyVersions; // supported versions of chain A
        bytes proofInit; // proof that chainA stored connectionEnd in state (on ConnOpenInit)
        bytes proofClient; // proof that chainA stored a light client of chainB
        bytes proofConsensus; // proof that chainA stored chainB's consensus state at consensus height
        uint64 proofHeight; // height at which relayer constructs proof of A storing connectionEnd in state
        uint64 consensusHeight; // latest height of chain B which chain A has stored in its chain B client
    }

    struct MsgConnectionOpenAck {
        string connectionId;
        ClientState.Data clientState; // client state for chainA on chainB
        Version.Data version; // version that ChainB chose in ConnOpenTry
        string counterpartyConnectionID;
        bytes proofTry; // proof that connectionEnd was added to ChainB state in ConnOpenTry
        bytes proofClient; // proof of client state on chainB for chainA
        bytes proofConsensus; // proof that chainB has stored ConsensusState of chainA on its client
        uint64 proofHeight; // height that relayer constructed proofTry
        uint64 consensusHeight; // latest height of chainA that chainB has stored on its chainA client
    }

    struct MsgConnectionOpenConfirm {
        string connectionId;
        bytes proofAck;
        uint64 proofHeight;
    }

    struct ClientConnectionPaths {
        string[] paths;
    }

    // constant values
    Version.Data[] versions;
    string[] features;
    bytes commitmentPrefix;

    // storage
    mapping(string => ClientConnectionPaths) clientConnectionPaths;

    constructor(ProvableStore store, IBCClient client_) public {
        // initialize
        commitmentPrefix = bytes("ibc");
        features.push("ORDER_ORDERED");
        features.push("ORDER_UNORDERED");
        versions.push(Version.Data({
            identifier: "1",
            features: features
        }));

        provableStore = store;
        client = client_;
    }

    /* Public functions */

    // ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
    // is returned.
    function connectionOpenInit(
        string memory clientId,
        string memory connectionId,
        Counterparty.Data memory counterparty,
        uint64 delayPeriod) public returns (string memory) {
 
        ConnectionEnd.Data memory connection;
        bool found;
        (connection, found) = provableStore.getConnection(connectionId);
        require(!found, "connection already exists");

        connection = ConnectionEnd.Data({
            client_id: clientId,
            versions: versions,
            state: CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_INIT,
            delay_period: delayPeriod,
            counterparty: counterparty
        });
        provableStore.setConnection(connectionId, connection);
        addConnectionToClient(clientId, connectionId);
        return connectionId;
    }

    // ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
    // code is executed on chain B).
    function connectionOpenTry(
        MsgConnectionOpenTry memory msg_
    ) public returns (string memory) {
        require(msg_.consensusHeight < block.number, "consensus height is greater than or equal to the current block height");
        require(client.validateSelfClient(msg_.clientState), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        // TODO add a support for selfConsensusState getter
        // (ConsensusState.Data memory expectedConsensusState, bool found) = client.getSelfConsensusState(consensusHeight);

        ConnectionEnd.Data memory connection = ConnectionEnd.Data({
            client_id: msg_.clientId,
            versions: versions,
            state: CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_TRYOPEN,
            delay_period: msg_.delayPeriod,
            counterparty: msg_.counterparty
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: msg_.counterparty.client_id,
            versions: msg_.counterpartyVersions,
            state: CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: Counterparty.Data({
                client_id: msg_.clientId,
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: commitmentPrefix})
            })
        });

        require(verifyConnectionState(connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection), "failed to verify connection state");
        // TODO commentout this after connectionState validation is passed
        // require(verifyClientState(connection, msg_.proofHeight, msg_.proofClient, msg_.clientState), "failed to verify clientState");
        // TODO commentout this
        // require(verifyClientConsensusState(connection, proofHeight, consensusHeight, proofConsensus, expectedConsensusState), "failed to verify consensusState");

        addConnectionToClient(msg_.clientId, msg_.connectionId);
        provableStore.setConnection(msg_.connectionId, connection);
        return msg_.connectionId;
    }

    function connectionOpenAck(
        MsgConnectionOpenAck memory msg_
    ) public {
        require(msg_.consensusHeight < block.number, "consensus height is greater than or equal to the current block height");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(msg_.connectionId);
        require(found, "connection not found");

        if (connection.state != CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_INIT && connection.state != CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_TRYOPEN) {
            revert("connection state is not INIT or TRYOPEN");
        } else if (connection.state == CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_INIT && !isSupportedVersion(msg_.version)) {
            revert("connection state is in INIT but the provided version is not supported");
        } else if (connection.state == CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_TRYOPEN && (connection.versions.length != 1 || !isEqualVersion(connection.versions[0], msg_.version))) {
            revert("connection state is in TRYOPEN but the provided version is not set in the previous connection versions"); 
        }

        require(client.validateSelfClient(msg_.clientState), "failed to validate self client state");

        // TODO add a support for selfConsensusState getter
        // (ConsensusState.Data memory expectedConsensusState, bool found) = client.getSelfConsensusState(consensusHeight);
        
        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: commitmentPrefix}) 
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: getVersionArray(msg_.version),
            state: CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_TRYOPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(verifyConnectionState(connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionID, expectedConnection), "failed to verify connection state");

        connection.state = CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_OPEN;
        connection.versions = expectedConnection.versions;
        connection.counterparty.connection_id = msg_.counterpartyConnectionID;
        provableStore.setConnection(msg_.connectionId, connection);
    }

    function connectionOpenConfirm(
        MsgConnectionOpenConfirm memory msg_
    ) public {
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(msg_.connectionId);
        require(found, "connection not found");

        require(connection.state == CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_TRYOPEN, "connection state is not TRYOPEN");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: commitmentPrefix}) 
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: connection.versions,
            state: CONNECTION_PROTO_GLOBAL_ENUMS.State.STATE_OPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(verifyConnectionState(connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection), "failed to verify connection state");

        provableStore.setConnection(msg_.connectionId, connection);
    }

    function isSupportedVersion(Version.Data memory proposedVersion) internal view returns (bool) {
        return true;
    }

    function isEqualVersion(Version.Data memory a, Version.Data memory b) internal view returns (bool) {
        return keccak256(Version.encode(a)) == keccak256(Version.encode(b));
    }

    function getVersionArray(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }   

    // Internal functions

    function addConnectionToClient(
        string memory clientId,
        string memory connectionId
    ) internal {
        require(provableStore.hasClientState(clientId), "client not found");
        clientConnectionPaths[clientId].paths.push(connectionId);
    }

    // Verification functions

    function verifyConnectionState(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory connectionId, ConnectionEnd.Data memory counterpartyConnection) public view returns (bool) {
        (ClientState.Data memory clientState, bool found) = provableStore.getClientState(connection.client_id);
        require(found, "clientState not found");
        return client.verifyConnectionState(clientState, connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, connectionId, ConnectionEnd.encode(counterpartyConnection));
    }

    function verifyConnectionStateAndGet(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, string memory connectionId, ConnectionEnd.Data memory counterpartyConnection) public view returns (bytes32, bytes32) {
        (ClientState.Data memory clientState, bool found) = provableStore.getClientState(connection.client_id);
        require(found, "clientState not found");
        return client.verifyConnectionStateAndGet(clientState, connection.client_id, height, connection.counterparty.prefix.key_prefix, proof, connectionId, ConnectionEnd.encode(counterpartyConnection));
    }

    function verifyClientState(ConnectionEnd.Data memory connection, uint64 height, bytes memory proof, ClientState.Data memory clientState) internal view returns (bool) {
        (ClientState.Data memory targetClient, bool found) = provableStore.getClientState(connection.client_id);
        require(found, "clientState not found");
        return client.verifyClientState(targetClient, connection.client_id, height, connection.counterparty.prefix.key_prefix, connection.counterparty.client_id, proof, clientState);
    }

    function verifyClientConsensusState(ConnectionEnd.Data memory connection, uint64 height, uint64 consensusHeight, bytes memory proof, ConsensusState.Data memory consensusState) internal view returns (bool) {
        (ClientState.Data memory clientState, bool found) = provableStore.getClientState(connection.client_id);
        require(found, "clientState not found");
        return client.verifyClientConsensusState(clientState, connection.client_id, height, connection.counterparty.client_id, consensusHeight, connection.counterparty.prefix.key_prefix, proof, ConsensusState.encode(consensusState));
    }
}
