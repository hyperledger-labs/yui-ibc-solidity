pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Connection.sol";
import "./types/Channel.sol";
import "../lib/IBCIdentifier.sol";

contract IBCStore {
    // Commitments
    mapping (bytes32 => bytes32) commitments;

    // Store
    mapping (string => address) clientRegistry; // clientType => clientImpl
    mapping (string => string) clientTypes; // clientID => clientType
    mapping (string => bytes) clientStates;
    mapping (string => mapping(uint64 => bytes)) consensusStates;
    mapping (string => bytes) connections;
    mapping (string => string[]) clientConnectionPaths; // clientID => [connectionID]
    mapping (string => mapping(string => bytes)) channels;
    mapping (string => mapping(string => uint64)) nextSequenceSends;
    mapping (string => mapping(string => uint64)) nextSequenceRecvs;
    mapping (string => mapping(string => uint64)) nextSequenceAcks;
    mapping (string => mapping(string => mapping(uint64 => bool))) packetReceipts;
    // TODO remove this storage variable in production. see details `function setPacket`
    mapping (string => mapping(string => mapping(uint64 => bytes))) packets;

    address owner;
    address ibcClient;
    address ibcConnection;
    address ibcChannel;
    address ibcRoutingModule;

    modifier onlyOwner (){
        require(msg.sender == owner);
        _;
    }

    modifier onlyIBCModule (){
        require(msg.sender == ibcRoutingModule || msg.sender == ibcChannel || msg.sender == ibcConnection || msg.sender == ibcClient);
        _;
    }

    constructor() public {
        owner = msg.sender;
    }

    function setIBCModule(address ibcClient_, address ibcConnection_, address ibcChannel_, address ibcRoutingModule_) onlyOwner public {
        ibcClient = ibcClient_;
        ibcConnection = ibcConnection_;
        ibcChannel = ibcChannel_;
        ibcRoutingModule = ibcRoutingModule_;
    }

    /// Storage accessor ///

    // Client registry

    // TODO specify any ACL modifiers to this
    function setClientImpl(string memory clientType, address clientImpl) public {
        require(address(clientRegistry[clientType]) == address(0), "clientImpl already exists");
        clientRegistry[clientType] = clientImpl;
    }

    function getClientImpl(string memory clientType) onlyIBCModule public view returns (address, bool) {
        address clientImpl = clientRegistry[clientType];
        if (clientImpl == address(0)) {
            return (clientImpl, false);
        }
        return (clientImpl, true);
    }

    // Client types

    function setClientType(string memory clientId, string memory clientType) onlyIBCModule public {
        require(bytes(clientTypes[clientId]).length == 0, "clientId already exists");
        require(bytes(clientType).length > 0, "clientType must not be empty string");
        clientTypes[clientId] = clientType;
    }

    function getClientType(string memory clientId) public view returns (string memory) {
        return clientTypes[clientId];
    }

    // ClientState

    function setClientState(string memory clientId, bytes memory clientStateBytes) onlyIBCModule public {
        clientStates[clientId] = clientStateBytes;
        commitments[IBCIdentifier.clientCommitmentKey(clientId)] = keccak256(clientStateBytes);
    }

    function getClientState(string memory clientId) public view returns (bytes memory clientStateBytes, bool found) {
        clientStateBytes = clientStates[clientId];
        if (clientStateBytes.length == 0) {
            return (clientStateBytes, false);
        }
        return (clientStateBytes, true);
    }

    function hasClientState(string memory clientId) public view returns (bool) {
        bytes memory encoded = clientStates[clientId];
        return encoded.length != 0;
    }

    // ConsensusState

    function setConsensusState(string memory clientId, uint64 height, bytes memory consensusStateBytes) onlyIBCModule public {
        consensusStates[clientId][height] = consensusStateBytes;
        commitments[IBCIdentifier.consensusCommitmentKey(clientId, height)] = keccak256(consensusStateBytes);
    }

    function getConsensusState(string memory clientId, uint64 height) public view returns (bytes memory consensusStateBytes, bool found) {
        consensusStateBytes = consensusStates[clientId][height];
        if (consensusStateBytes.length == 0) {
            return (consensusStateBytes, false);
        }
        return (consensusStateBytes, true);
    }

    // Connection

    function setConnection(string memory connectionId, ConnectionEnd.Data memory connection) onlyIBCModule public {
        connections[connectionId] = ConnectionEnd.encode(connection);
        commitments[IBCIdentifier.connectionCommitmentKey(connectionId)] = keccak256(connections[connectionId]);
    }

    function getConnection(string memory connectionId) public view returns (ConnectionEnd.Data memory connection, bool) {
        bytes memory encoded = connections[connectionId];
        if (encoded.length == 0) {
            return (connection, false);
        }
        connection = ConnectionEnd.decode(encoded);
        return (connection, true);
    }

    function addConnectionPath(string memory clientId, string memory connectionId) onlyIBCModule public {
        clientConnectionPaths[clientId].push(connectionId);
    }

    // Channel

    function setChannel(string memory portId, string memory channelId, Channel.Data memory channel) onlyIBCModule public {
        channels[portId][channelId] = Channel.encode(channel);
        commitments[IBCIdentifier.channelCommitmentKey(portId, channelId)] = keccak256(channels[portId][channelId]);
    }

    function getChannel(string memory portId, string memory channelId) public view returns (Channel.Data memory channel, bool) {
        bytes memory encoded = channels[portId][channelId];
        if (encoded.length == 0) {
            return (channel, false);
        }
        channel = Channel.decode(encoded);
        return (channel, true);
    }

    function hasChannel(string memory portId, string memory channelId) public view returns (bool) {
        return channels[portId][channelId].length != 0;
    }

    // Packet

    function setNextSequenceSend(string memory portId, string memory channelId, uint64 sequence) onlyIBCModule public {
        nextSequenceSends[portId][channelId] = sequence;
    }

    function getNextSequenceSend(string memory portId, string memory channelId) public view returns (uint64) {
        return nextSequenceSends[portId][channelId];
    }

    function setNextSequenceRecv(string memory portId, string memory channelId, uint64 sequence) onlyIBCModule public {
        nextSequenceRecvs[portId][channelId] = sequence;
    }

    function getNextSequenceRecv(string memory portId, string memory channelId) public view returns (uint64) {
        return nextSequenceRecvs[portId][channelId];
    }

    function setNextSequenceAck(string memory portId, string memory channelId, uint64 sequence) onlyIBCModule public {
        nextSequenceAcks[portId][channelId] = sequence;
    }

    function getNextSequenceAck(string memory portId, string memory channelId) public view returns (uint64) {
        return nextSequenceAcks[portId][channelId];
    }

    // TODO remove this function in production
    // NOTE: A packet doesn't need to be stored in storage, but this will help development
    function setPacket(string memory portId, string memory channelId, uint64 sequence, Packet.Data memory packet) internal {
        packets[portId][channelId][sequence] = Packet.encode(packet);
    }

    // TODO remove this function in production
    function getPacket(string memory portId, string memory channelId, uint64 sequence) public view returns (Packet.Data memory) {
        return Packet.decode(packets[portId][channelId][sequence]);
    }

    function setPacketCommitment(string memory portId, string memory channelId, uint64 sequence, Packet.Data memory packet) onlyIBCModule public {
        commitments[IBCIdentifier.packetCommitmentKey(portId, channelId, sequence)] = makePacketCommitment(packet);
        setPacket(portId, channelId, sequence, packet);
    }

    function deletePacketCommitment(string memory portId, string memory channelId, uint64 sequence) onlyIBCModule public {
        delete commitments[IBCIdentifier.packetCommitmentKey(portId, channelId, sequence)];
    }

    function getPacketCommitment(string memory portId, string memory channelId, uint64 sequence) public returns (bytes32, bool) {
        bytes32 commitment = commitments[IBCIdentifier.packetCommitmentKey(portId, channelId, sequence)];
        return (commitment, commitment != bytes32(0));
    }

    function makePacketCommitment(Packet.Data memory packet) public view returns (bytes32) {
        bytes32 dataHash = sha256(packet.data);
        // TODO serialize uint64 to bytes(big-endian)
        return sha256(abi.encodePacked(packet.timeout_timestamp, packet.timeout_height.revision_number, packet.timeout_height.revision_height, dataHash));
    }

    function setPacketAcknowledgementCommitment(string memory portId, string memory channelId, uint64 sequence, bytes memory acknowledgement) onlyIBCModule public {
        commitments[IBCIdentifier.packetAcknowledgementCommitmentKey(portId, channelId, sequence)] = makePacketAcknowledgementCommitment(acknowledgement);
    }

    function getPacketAcknowledgementCommitment(string memory portId, string memory channelId, uint64 sequence) public view returns (bytes32, bool) {
        bytes32 commitment = commitments[IBCIdentifier.packetAcknowledgementCommitmentKey(portId, channelId, sequence)];
        return (commitment, commitment != bytes32(0));
    }

    function makePacketAcknowledgementCommitment(bytes memory acknowledgement) public view returns (bytes32) {
        return sha256(acknowledgement);
    }

    function setPacketReceipt(string memory portId, string memory channelId, uint64 sequence) onlyIBCModule public {
        packetReceipts[portId][channelId][sequence] = true;
    }

    function hasPacketReceipt(string memory portId, string memory channelId, uint64 sequence) public view returns (bool) {
        return packetReceipts[portId][channelId][sequence];
    }
}
