// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../proto/Client.sol";
import "../02-client/ILightClient.sol";
import "../04-channel/IBCPacket.sol";
import "../05-port/ModuleManager.sol";
import "../24-host/IBCStore.sol";
import "../24-host/IBCCommitment.sol";
import "../25-handler/IBCPacketHandler.sol";

abstract contract IBCQuerier is IBCStore {
    function getClientState(string calldata clientId) external view returns (bytes memory, bool) {
        return checkAndGetClient(clientId).getClientState(clientId);
    }

    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (bytes memory consensusStateBytes, bool)
    {
        return checkAndGetClient(clientId).getConsensusState(clientId, height);
    }

    function getConnection(string calldata connectionId) external view returns (ConnectionEnd.Data memory, bool) {
        ConnectionEnd.Data storage connection = connections[connectionId];
        return (connection, connection.state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getChannel(string calldata portId, string calldata channelId)
        external
        view
        returns (Channel.Data memory, bool)
    {
        Channel.Data storage channel = channels[portId][channelId];
        return (channel, channel.state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getCommitment(bytes32 hashedPath) public view returns (bytes32) {
        return commitments[hashedPath];
    }

    function getHashedPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bytes32, bool)
    {
        bytes32 commitment = commitments[keccak256(IBCCommitment.packetCommitmentPath(portId, channelId, sequence))];
        return (commitment, commitment != bytes32(0));
    }

    function getHashedPacketAcknowledgementCommitment(
        string calldata portId,
        string calldata channelId,
        uint64 sequence
    ) external view returns (bytes32, bool) {
        bytes32 commitment =
            commitments[keccak256(IBCCommitment.packetAcknowledgementCommitmentPath(portId, channelId, sequence))];
        return (commitment, commitment != bytes32(0));
    }

    function hasPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bool)
    {
        return commitments[IBCCommitment.packetReceiptCommitmentKey(portId, channelId, sequence)] != bytes32(0);
    }

    function getNextSequenceSend(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceSends[portId][channelId];
    }

    function getNextSequenceRecv(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceRecvs[portId][channelId];
    }

    function getNextSequenceAck(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceAcks[portId][channelId];
    }

    function getExpectedTimePerBlock() external view returns (uint64) {
        return expectedTimePerBlock;
    }
}
