pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Connection.sol";
import "./types/Channel.sol";
import "./IBCIdentifier.sol";

contract IBCHost {
    // Commitments
    mapping (bytes32 => bytes32) commitments;

    // Store
    mapping (string => address) clientRegistry; // clientType => clientImpl
    mapping (string => string) clientTypes; // clientID => clientType
    mapping (string => bytes) clientStates;
    mapping (string => mapping(uint64 => bytes)) consensusStates;
    mapping (string => ConnectionEnd.Data) connections;
    mapping (string => mapping(string => Channel.Data)) channels;
    mapping (string => mapping(string => uint64)) nextSequenceSends;
    mapping (string => mapping(string => uint64)) nextSequenceRecvs;
    mapping (string => mapping(string => uint64)) nextSequenceAcks;
    mapping (string => mapping(string => mapping(uint64 => bool))) packetReceipts;
    // TODO remove this storage variable in production. see details `function setPacket`
    mapping (string => mapping(string => mapping(uint64 => Packet.Data))) packets;
    mapping (bytes => address[]) capabilities;

    address owner;
    address ibcModule;

    constructor() public {
        owner = msg.sender;
    }

    function setIBCModule(address ibcModule_) external {
        require(msg.sender == owner);
        ibcModule = ibcModule_;
    }

    function onlyIBCModule() public view {
        require(msg.sender == ibcModule);
    }

    /// Storage accessor ///

    // Client implementation registry

    function setClientImpl(string calldata clientType, address clientImpl) external {
        onlyIBCModule();
        require(address(clientRegistry[clientType]) == address(0), "clientImpl already exists");
        clientRegistry[clientType] = clientImpl;
    }

    function getClientImpl(string calldata clientType) external view returns (address, bool) {
        return (clientRegistry[clientType], clientRegistry[clientType] != address(0));
    }

    // Client types

    function setClientType(string calldata clientId, string calldata clientType) external {
        onlyIBCModule();
        require(bytes(clientTypes[clientId]).length == 0, "clientId already exists");
        require(bytes(clientType).length > 0, "clientType must not be empty string");
        clientTypes[clientId] = clientType;
    }

    function getClientType(string calldata clientId) external view returns (string memory) {
        return clientTypes[clientId];
    }

    // ClientState

    function setClientState(string calldata clientId, bytes calldata clientStateBytes) external {
        onlyIBCModule();
        clientStates[clientId] = clientStateBytes;
        commitments[IBCIdentifier.clientCommitmentKey(clientId)] = keccak256(clientStateBytes);
    }

    function getClientState(string calldata clientId) external view returns (bytes memory, bool) {
        return (clientStates[clientId], clientStates[clientId].length > 0);
    }

    // ConsensusState

    function setConsensusState(string calldata clientId, uint64 height, bytes calldata consensusStateBytes) external {
        onlyIBCModule();
        consensusStates[clientId][height] = consensusStateBytes;
        commitments[IBCIdentifier.consensusCommitmentKey(clientId, height)] = keccak256(consensusStateBytes);
    }

    function getConsensusState(string calldata clientId, uint64 height) external view returns (bytes memory, bool) {
        return (consensusStates[clientId][height], consensusStates[clientId][height].length > 0);
    }

    // Connection

    function setConnection(string memory connectionId, ConnectionEnd.Data memory connection) public {
        onlyIBCModule();
        connections[connectionId].client_id = connection.client_id;
        connections[connectionId].state = connection.state;
        connections[connectionId].delay_period = connection.delay_period;
        delete connections[connectionId].versions;
        for (uint8 i = 0; i < connection.versions.length; i++) {
            connections[connectionId].versions.push(connection.versions[i]);
        }
        connections[connectionId].counterparty = connection.counterparty;
        commitments[IBCIdentifier.connectionCommitmentKey(connectionId)] = keccak256(ConnectionEnd.encode(connection));
    }

    function getConnection(string calldata connectionId) external view returns (ConnectionEnd.Data memory connection, bool) {
        return (connections[connectionId], connections[connectionId].state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    // Channel

    function setChannel(string memory portId, string memory channelId, Channel.Data memory channel) public {
        onlyIBCModule();
        channels[portId][channelId] = channel;
        commitments[IBCIdentifier.channelCommitmentKey(portId, channelId)] = keccak256(Channel.encode(channel));
    }

    function getChannel(string calldata portId, string calldata channelId) external view returns (Channel.Data memory channel, bool) {
        return (channels[portId][channelId], channels[portId][channelId].state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    // Packet

    function setNextSequenceSend(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        nextSequenceSends[portId][channelId] = sequence;
    }

    function getNextSequenceSend(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceSends[portId][channelId];
    }

    function setNextSequenceRecv(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        nextSequenceRecvs[portId][channelId] = sequence;
    }

    function getNextSequenceRecv(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceRecvs[portId][channelId];
    }

    function setNextSequenceAck(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        nextSequenceAcks[portId][channelId] = sequence;
    }

    function getNextSequenceAck(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceAcks[portId][channelId];
    }

    // TODO remove this function in production
    // NOTE: A packet doesn't need to be stored in storage, but this will help development
    function setPacket(string memory portId, string memory channelId, uint64 sequence, Packet.Data memory packet) internal {
        packets[portId][channelId][sequence] = packet;
    }

    // TODO remove this function in production
    function getPacket(string calldata portId, string calldata channelId, uint64 sequence) external view returns (Packet.Data memory) {
        return packets[portId][channelId][sequence];
    }

    function setPacketCommitment(string memory portId, string memory channelId, uint64 sequence, Packet.Data memory packet) public {
        onlyIBCModule();
        commitments[IBCIdentifier.packetCommitmentKey(portId, channelId, sequence)] = makePacketCommitment(packet);
        setPacket(portId, channelId, sequence, packet);
    }

    function deletePacketCommitment(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        delete commitments[IBCIdentifier.packetCommitmentKey(portId, channelId, sequence)];
    }

    function getPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence) external view returns (bytes32, bool) {
        bytes32 commitment = commitments[IBCIdentifier.packetCommitmentKey(portId, channelId, sequence)];
        return (commitment, commitment != bytes32(0));
    }

    function makePacketCommitment(Packet.Data memory packet) public pure returns (bytes32) {
        // TODO serialize uint64 to bytes(big-endian)
        return sha256(abi.encodePacked(packet.timeout_timestamp, packet.timeout_height.revision_number, packet.timeout_height.revision_height, sha256(packet.data)));
    }

    function setPacketAcknowledgementCommitment(string calldata portId, string calldata channelId, uint64 sequence, bytes calldata acknowledgement) external {
        onlyIBCModule();
        commitments[IBCIdentifier.packetAcknowledgementCommitmentKey(portId, channelId, sequence)] = makePacketAcknowledgementCommitment(acknowledgement);
    }

    function getPacketAcknowledgementCommitment(string calldata portId, string calldata channelId, uint64 sequence) external view returns (bytes32, bool) {
        bytes32 commitment = commitments[IBCIdentifier.packetAcknowledgementCommitmentKey(portId, channelId, sequence)];
        return (commitment, commitment != bytes32(0));
    }

    function makePacketAcknowledgementCommitment(bytes memory acknowledgement) public pure returns (bytes32) {
        return sha256(acknowledgement);
    }

    function setPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        packetReceipts[portId][channelId][sequence] = true;
    }

    function hasPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence) external view returns (bool) {
        return packetReceipts[portId][channelId][sequence];
    }

    // capabilities

    function claimCapability(bytes calldata name, address addr) external {
        onlyIBCModule();
        for (uint32 i = 0; i < capabilities[name].length; i++) {
            require(capabilities[name][i] != addr);
        }
        capabilities[name].push(addr);
    }

    function authenticateCapability(bytes calldata name, address addr) external view returns (bool) {
        onlyIBCModule();
        for (uint32 i = 0; i < capabilities[name].length; i++) {
            if (capabilities[name][i] == addr) {
                return true;
            }
        }
        return false;
    }

    function getModuleOwner(bytes calldata name) external view returns (address, bool) {
        if (capabilities[name].length == 0) {
            return (address(0), false);
        }
        return (capabilities[name][0], true);
    }
}
