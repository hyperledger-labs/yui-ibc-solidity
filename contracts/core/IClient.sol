// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IBCHost.sol";
import "./types/Client.sol";

interface IClient {

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string calldata clientId,
        Height.Data calldata height
    ) external view returns (uint64, bool);

    function getLatestHeight(
        IBCHost host,
        string calldata clientId
    ) external view returns (Height.Data memory, bool);

    function checkHeaderAndUpdateState(
        IBCHost host,
        string calldata clientId, 
        bytes calldata clientStateBytes,
        bytes calldata headerBytes
    ) external returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, Height.Data memory height);

    function verifyClientState(
        IBCHost host,
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata prefix,
        string calldata counterpartyClientIdentifier,
        bytes calldata proof,
        bytes calldata clientStateBytes // serialized with pb
    ) external returns (bool);

    function verifyClientConsensusState(
        IBCHost host,
        string calldata clientId,
        Height.Data calldata height,
        string calldata counterpartyClientIdentifier,
        Height.Data calldata consensusHeight,
        bytes calldata prefix,
        bytes calldata proof,
        bytes calldata consensusStateBytes // serialized with pb
    ) external returns (bool);

    function verifyConnectionState(
        IBCHost host,
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata connectionId,
        bytes calldata connectionBytes // serialized with pb
    ) external returns (bool);

    function verifyChannelState(
        IBCHost host,
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata prefix,
        bytes calldata proof,
        string calldata portId,
        string calldata channelId,
        bytes calldata channelBytes // serialized with pb
    ) external returns (bool);

    function verifyPacketCommitment(
        IBCHost host,
        string calldata clientId,
        Height.Data calldata height,
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
        Height.Data calldata height,
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
