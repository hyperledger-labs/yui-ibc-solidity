pragma solidity ^0.6.8;

interface IClient {

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        string calldata clientId,
        uint64 height
    ) external view returns (uint64, bool);

    function getLatestHeight(
        string calldata clientId
    ) external view returns (uint64, bool);

    function checkHeaderAndUpdateState(
        string calldata clientId, 
        bytes calldata clientStateBytes,
        bytes calldata headerBytes
    ) external view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height);

    function verifyClientState(
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        string calldata counterpartyClientIdentifier,
        bytes calldata proof,
        bytes calldata clientStateBytes // serialized with pb
    ) external view returns (bool);

    function verifyClientConsensusState(
        string calldata clientId,
        uint64 height,
        string calldata counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes calldata prefix,
        bytes calldata proof,
        bytes calldata consensusStateBytes // serialized with pb
    ) external view returns (bool);

    function verifyConnectionState(
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata connectionId,
        bytes calldata connectionBytes // serialized with pb
    ) external view returns (bool);

    function verifyChannelState(
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        bytes calldata channelBytes // serialized with pb
    ) external view returns (bool);

    function verifyPacketCommitment(
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        uint64 sequence,
        bytes32 commitmentBytes // serialized with pb
    ) external view returns (bool);

    function verifyPacketAcknowledgement(
        string calldata clientId,
        uint64 height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        uint64 sequence,
        bytes32 ackCommitmentBytes // serialized with pb
    ) external view returns (bool);
}
