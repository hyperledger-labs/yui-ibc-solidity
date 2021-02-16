pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Client.sol";
import "./types/Connection.sol";
import "./types/Channel.sol";

contract ProvableStore {
    // Commitments
    mapping (bytes32 => bytes32) commitments;

    // constant values
    uint256 constant commitmentSlot = 0;
    string constant clientPrefix = "client/";
    string constant consensusStatePrefix = "consensus/";
    string constant connectionPrefix = "connection/";
    string constant channelPrefix = "channel/";
    string constant packetPrefix = "packet/";

    // TODO provides ACL
    address[] internal allowedAccessors;

    // TODO use RLP instead of pb?
    // Store
    mapping (string => bytes) clientStates;
    mapping (string => mapping(uint64 => bytes)) consensusStates;
    mapping (string => bytes) connections;
    mapping (string => mapping(string => bytes)) channels;
    mapping (string => mapping(string => uint64)) nextSequenceSends;
    mapping (string => mapping(string => uint64)) nextSequenceRecvs;
    mapping (string => mapping(string => uint64)) nextSequenceAcks;


    // Commitment key generator

    function clientCommitmentKey(string memory clientId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(clientPrefix, clientId));
    }

    function consensusCommitmentKey(string memory clientId, uint64 height) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(consensusStatePrefix, clientId, "/", height));
    }

    function connectionCommitmentKey(string memory connectionId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(connectionPrefix, connectionId));
    }

    function channelCommitmentKey(string memory portId, string memory channelId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(channelPrefix, portId, "/", channelId));
    }

    function packetCommitmentKey(string memory portId, string memory channelId, uint64 sequence) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(packetPrefix, portId, "/", channelId, "/", sequence));
    }

    // Slot calculator

    function clientStateCommitmentSlot(string memory clientId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(clientCommitmentKey(clientId), commitmentSlot));
    }

    function consensusStateCommitmentSlot(string memory clientId, uint64 height) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(consensusCommitmentKey(clientId, height), commitmentSlot));
    }

    function connectionCommitmentSlot(string memory connectionId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(connectionCommitmentKey(connectionId), commitmentSlot));
    }

    function channelCommitmentSlot(string memory portId, string memory channelId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(channelCommitmentKey(portId, channelId), commitmentSlot));
    }

    function packetCommitmentSlot(string memory portId, string memory channelId, uint64 sequence) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(packetCommitmentKey(portId, channelId, sequence), commitmentSlot));
    }

    /// Storage accessor ///

    // ClientState

    function setClientState(string memory clientId, ClientState.Data memory data) public {
        bytes memory encoded = ClientState.encode(data);
        clientStates[clientId] = encoded;
        commitments[clientCommitmentKey(clientId)] = keccak256(encoded);
    }

    function getClientState(string memory clientId) public view returns (ClientState.Data memory, bool) {
        bytes memory encoded = clientStates[clientId];
        ClientState.Data memory memoryData;
        if (encoded.length == 0) {
            return (memoryData, false);
        }
        memoryData = ClientState.decode(encoded);
        return (memoryData, true);
    }

    function hasClientState(string memory clientId) public view returns (bool) {
        bytes memory encoded = clientStates[clientId];
        return encoded.length != 0;
    }

    // ConsensusState

    function setConsensusState(string memory clientId, uint64 height, ConsensusState.Data memory consensusState) public {
        bytes memory encoded = ConsensusState.encode(consensusState);
        consensusStates[clientId][height] = encoded;
        commitments[consensusCommitmentKey(clientId, height)] = keccak256(encoded);
    }

    function getConsensusState(string memory clientId, uint64 height) public view returns (ConsensusState.Data memory consensusState, bool) {
        bytes memory encoded = consensusStates[clientId][height];
        if (encoded.length == 0) {
            return (consensusState, false);
        }
        consensusState = ConsensusState.decode(encoded);
        return (consensusState, true);
    }

    // Connection

    function setConnection(string memory connectionId, ConnectionEnd.Data memory connection) public {
        connections[connectionId] = ConnectionEnd.encode(connection);
        commitments[connectionCommitmentKey(connectionId)] = keccak256(connections[connectionId]);
    }

    function getConnection(string memory connectionId) public view returns (ConnectionEnd.Data memory connection, bool) {
        bytes memory encoded = connections[connectionId];
        if (encoded.length == 0) {
            return (connection, false);
        }
        connection = ConnectionEnd.decode(encoded);
        return (connection, true);
    }

    // Channel

    function setChannel(string memory portId, string memory channelId, Channel.Data memory channel) public {
        channels[portId][channelId] = Channel.encode(channel);
        commitments[channelCommitmentKey(portId, channelId)] = keccak256(channels[portId][channelId]);
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

    function setNextSequenceSend(string memory portId, string memory channelId, uint64 sequence) public {
        nextSequenceSends[portId][channelId] = sequence;
    }

    function getNextSequenceSend(string memory portId, string memory channelId) public view returns (uint64) {
        return nextSequenceSends[portId][channelId];
    }

    function setNextSequenceRecv(string memory portId, string memory channelId, uint64 sequence) public {
        nextSequenceRecvs[portId][channelId] = sequence;
    }

    function getNextSequenceRecv(string memory portId, string memory channelId) public view returns (uint64) {
        return nextSequenceRecvs[portId][channelId];
    }

    function setNextSequenceAck(string memory portId, string memory channelId, uint64 sequence) public {
        nextSequenceAcks[portId][channelId] = sequence;
    }

    function getNextSequenceAck(string memory portId, string memory channelId) public view returns (uint64) {
        return nextSequenceAcks[portId][channelId];
    }

    function setPacketCommitment(string memory portId, string memory channelId, uint64 sequence, Packet.Data memory packet) public {
        commitments[packetCommitmentKey(portId, channelId, sequence)] = makePacketCommitment(packet);
    }

    function makePacketCommitment(Packet.Data memory packet) public view returns (bytes32) {
        bytes32 dataHash = sha256(packet.data);
        // return sha256(timeoutTimestamp, timeoutHeight.revision_number, timeoutHeight.revision_height, dataHash);
        // revert("not implemented error");
        return dataHash;
    }

    // Debug

    function getClientStateBytes(string memory clientId) public view returns (bytes memory, bool) {
        bytes memory encoded = clientStates[clientId];
        if (encoded.length == 0) {
            return (encoded, false);
        }
        return (encoded, true);
    }

    function getConnectionBytes(string memory connectionId) public view returns (bytes memory, bool) {
        bytes memory encoded = connections[connectionId];
        if (encoded.length == 0) {
            return (encoded, false);
        }
        return (encoded, true);
    }

    function parseConnectionBytes(bytes memory connectionBytes) public view returns (ConnectionEnd.Data memory connection) {
        ConnectionEnd.Data memory data = ConnectionEnd.decode(connectionBytes);
        return data;
    }

    function getCommitment(string memory connectionId) public view returns (bytes32) {
        return commitments[connectionCommitmentKey(connectionId)];
    }
}
