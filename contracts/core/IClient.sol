pragma solidity ^0.8.9;

import "./IBCHost.sol";

interface IClient {

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string calldata clientId,
        uint64 height
    ) external view returns (uint64, bool);

    function getLatestHeight(
        IBCHost host,
        string calldata clientId
    ) external view returns (uint64, bool);

    function checkHeaderAndUpdateState(
        IBCHost host,
        string calldata clientId, 
        bytes calldata clientStateBytes,
        bytes calldata headerBytes
    ) external view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height);

    function verifyClientState(
        IBCHost host,
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        string calldata counterpartyClientIdentifier,
        bytes calldata proof,
        bytes calldata clientStateBytes // serialized with pb
    ) external returns (bool);

    function verifyClientConsensusState(
        IBCHost host,
        string calldata clientId,
        uint64 height,
        string calldata counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes calldata prefix,
        bytes calldata proof,
        bytes calldata consensusStateBytes // serialized with pb
    ) external returns (bool);

    function verifyConnectionState(
        IBCHost host,
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata connectionId,
        bytes calldata connectionBytes // serialized with pb
    ) external returns (bool);

    function verifyChannelState(
        IBCHost host,
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        bytes calldata channelBytes // serialized with pb
    ) external returns (bool);

    function verifyPacketCommitment(
        IBCHost host,
        string calldata clientId,
        uint64 height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        uint64 sequence,
        bytes32 commitmentBytes // serialized with pb
    ) external returns (bool);

    function verifyPacketAcknowledgement(
        IBCHost host,
        string calldata clientId,
        uint64 height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        uint64 sequence,
        bytes calldata acknowledgement // serialized with pb
    ) external returns (bool);
}
