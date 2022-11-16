// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../proto/Client.sol";
import "../proto/Connection.sol";
import "../proto/Channel.sol";
import "./IBCIdentifier.sol";
import "./IBCHeight.sol";

contract IBCHost {
    using IBCHeight for Height.Data;

    // Commitments
    mapping(bytes32 => bytes32) commitments;

    // Store
    mapping(string => address) clientRegistry; // clientType => clientImpl
    mapping(string => string) clientTypes; // clientID => clientType
    mapping(string => bytes) clientStates;
    mapping(string => mapping(uint128 => bytes)) consensusStates;
    mapping(string => mapping(uint128 => uint256)) processedTimes;
    mapping(string => mapping(uint128 => uint256)) processedHeights;
    mapping(string => ConnectionEnd.Data) connections;
    mapping(string => mapping(string => Channel.Data)) channels;
    mapping(string => mapping(string => uint64)) nextSequenceSends;
    mapping(string => mapping(string => uint64)) nextSequenceRecvs;
    mapping(string => mapping(string => uint64)) nextSequenceAcks;
    mapping(string => mapping(string => mapping(uint64 => uint8))) packetReceipts;
    mapping(bytes => address[]) capabilities;
    uint64 nextClientSequence;
    uint64 nextConnectionSequence;
    uint64 nextChannelSequence;

    event GeneratedClientIdentifier(string);
    event GeneratedConnectionIdentifier(string);
    event GeneratedChannelIdentifier(string);

    address owner;
    address ibcModule;

    uint64 expectedTimePerBlock;

    constructor() {
        owner = msg.sender;
    }

    function setIBCModule(address ibcModule_) external {
        require(msg.sender == owner);
        ibcModule = ibcModule_;
    }

    function getIBCModule() external view returns (address) {
        return ibcModule;
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
        commitments[keccak256(IBCIdentifier.clientStatePath(clientId))] = keccak256(clientStateBytes);
    }

    function getClientState(string calldata clientId) external view returns (bytes memory, bool) {
        return (clientStates[clientId], clientStates[clientId].length > 0);
    }

    // ConsensusState

    function setConsensusState(
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata consensusStateBytes
    ) external {
        onlyIBCModule();
        consensusStates[clientId][height.toUint128()] = consensusStateBytes;
        commitments[keccak256(
            IBCIdentifier.consensusStatePath(clientId, height.revision_number, height.revision_height)
        )] = keccak256(consensusStateBytes);
    }

    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (bytes memory, bool)
    {
        uint128 h = height.toUint128();
        return (consensusStates[clientId][h], consensusStates[clientId][h].length > 0);
    }

    // Processed Time/Block

    function setProcessedTime(string calldata clientId, Height.Data calldata height, uint256 processedTime) external {
        onlyIBCModule();
        processedTimes[clientId][height.toUint128()] = processedTime;
    }

    function getProcessedTime(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (uint256, bool)
    {
        uint256 processedTime = processedTimes[clientId][height.toUint128()];
        return (processedTime, processedTime != 0);
    }

    function setProcessedHeight(string calldata clientId, Height.Data calldata height, uint256 processedHeight)
        external
    {
        onlyIBCModule();
        processedHeights[clientId][height.toUint128()] = processedHeight;
    }

    function getProcessedHeight(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (uint256, bool)
    {
        uint256 processedHeight = processedHeights[clientId][height.toUint128()];
        return (processedHeight, processedHeight != 0);
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
        commitments[keccak256(IBCIdentifier.connectionPath(connectionId))] = keccak256(ConnectionEnd.encode(connection));
    }

    function getConnection(string calldata connectionId)
        external
        view
        returns (ConnectionEnd.Data memory connection, bool)
    {
        return (
            connections[connectionId],
            connections[connectionId].state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED
        );
    }

    // Channel

    function setChannel(string memory portId, string memory channelId, Channel.Data memory channel) public {
        onlyIBCModule();
        channels[portId][channelId] = channel;
        commitments[keccak256(IBCIdentifier.channelPath(portId, channelId))] = keccak256(Channel.encode(channel));
    }

    function getChannel(string calldata portId, string calldata channelId)
        external
        view
        returns (Channel.Data memory channel, bool)
    {
        return (
            channels[portId][channelId],
            channels[portId][channelId].state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED
        );
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
        commitments[keccak256(IBCIdentifier.nextSequenceRecvCommitmentPath(portId, channelId))] =
            keccak256(abi.encodePacked(sequence));
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

    function setPacketCommitment(
        string memory portId,
        string memory channelId,
        uint64 sequence,
        Packet.Data memory packet
    ) public {
        onlyIBCModule();
        commitments[keccak256(IBCIdentifier.packetCommitmentPath(portId, channelId, sequence))] =
            keccak256(abi.encodePacked(makePacketCommitment(packet)));
    }

    function deletePacketCommitment(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        delete commitments[keccak256(IBCIdentifier.packetCommitmentPath(portId, channelId, sequence))];
    }

    function getPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bytes32, bool)
    {
        bytes32 commitment = commitments[keccak256(IBCIdentifier.packetCommitmentPath(portId, channelId, sequence))];
        return (commitment, commitment != bytes32(0));
    }

    function makePacketCommitment(Packet.Data memory packet) public pure returns (bytes32) {
        return sha256(
            abi.encodePacked(
                packet.timeout_timestamp,
                packet.timeout_height.revision_number,
                packet.timeout_height.revision_height,
                sha256(packet.data)
            )
        );
    }

    function setPacketAcknowledgementCommitment(
        string calldata portId,
        string calldata channelId,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external {
        onlyIBCModule();
        commitments[keccak256(IBCIdentifier.packetAcknowledgementCommitmentPath(portId, channelId, sequence))] =
            keccak256(abi.encodePacked(sha256(acknowledgement)));
    }

    function getPacketAcknowledgementCommitment(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bytes32, bool)
    {
        bytes32 commitment =
            commitments[keccak256(IBCIdentifier.packetAcknowledgementCommitmentPath(portId, channelId, sequence))];
        return (commitment, commitment != bytes32(0));
    }

    function setPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence) external {
        onlyIBCModule();
        packetReceipts[portId][channelId][sequence] = 1;
        commitments[keccak256(IBCIdentifier.packetReceiptCommitmentPath(portId, channelId, sequence))] =
            keccak256(abi.encodePacked(uint8(1)));
    }

    function hasPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bool)
    {
        return packetReceipts[portId][channelId][sequence] == 1;
    }

    function getExpectedTimePerBlock() external view returns (uint64) {
        return expectedTimePerBlock;
    }

    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) external {
        onlyIBCModule();
        expectedTimePerBlock = expectedTimePerBlock_;
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

    /// Identifier generators ///

    function generateClientIdentifier(string calldata clientType) external returns (string memory) {
        onlyIBCModule();
        string memory identifier = string(abi.encodePacked(clientType, "-", uint2str(nextClientSequence)));
        nextClientSequence++;
        emit GeneratedClientIdentifier(identifier);
        return identifier;
    }

    function generateConnectionIdentifier() external returns (string memory) {
        onlyIBCModule();
        string memory identifier = string(abi.encodePacked("connection-", uint2str(nextConnectionSequence)));
        nextConnectionSequence++;
        emit GeneratedConnectionIdentifier(identifier);
        return identifier;
    }

    function generateChannelIdentifier() external returns (string memory) {
        onlyIBCModule();
        string memory identifier = string(abi.encodePacked("channel-", uint2str(nextChannelSequence)));
        nextChannelSequence++;
        emit GeneratedChannelIdentifier(identifier);
        return identifier;
    }

    function uint2str(uint64 _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint64 j = _i;
        uint64 len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint64 k = len;
        while (_i != 0) {
            k = k - 1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
