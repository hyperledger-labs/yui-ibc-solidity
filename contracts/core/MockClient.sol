pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import {MockClientState as ClientState, MockConsensusState as ConsensusState, MockHeader as Header} from "./types/MockClient.sol";
import "../lib/Bytes.sol";

contract MockClient is IClient {
    using Bytes for bytes;

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string memory clientId,
        uint64 height
    ) public override view returns (uint64, bool) {
        (bytes memory consensusStateBytes, bool found) = host.getConsensusState(clientId, height);
        if (!found) {
            return (0, false);
        }
        return (ConsensusState.decode(consensusStateBytes).timestamp, true);
    }

    function getLatestHeight(
        IBCHost host,
        string memory clientId
    ) public override view returns (uint64, bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        return (clientState.latest_height, true);
    }

    /**
     * @dev checkHeaderAndUpdateState checks if the provided header is valid
     */
    function checkHeaderAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory headerBytes
    ) public override view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height) {
        uint64 timestamp;
        ClientState.Data memory clientState = ClientState.decode(clientStateBytes);
        (height, timestamp) = parseHeader(headerBytes);
        if (height > clientState.latest_height) {
            clientState.latest_height = height;
        }
        ConsensusState.Data memory consensusState = ConsensusState.Data({timestamp: timestamp});
        return (ClientState.encode(clientState), ConsensusState.encode(consensusState), height);
    }

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        bytes memory clientStateBytes // serialized with pb
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(clientStateBytes) == proof.toBytes32();
    }

    function verifyClientConsensusState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        string memory counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes memory prefix,
        bytes memory proof,
        bytes memory consensusStateBytes // serialized with pb
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(consensusStateBytes) == proof.toBytes32();
    }

    function verifyConnectionState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory connectionBytes // serialized with pb
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(connectionBytes) == proof.toBytes32();
    }

    function verifyChannelState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes // serialized with pb
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(channelBytes) == proof.toBytes32();
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return commitmentBytes == proof.toBytes32();
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes memory acknowledgement
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return host.makePacketAcknowledgementCommitment(acknowledgement) == proof.toBytes32();
    }

    function getClientState(IBCHost host, string memory clientId) public view returns (ClientState.Data memory clientState) {
        (bytes memory clientStateBytes, bool found) = host.getClientState(clientId);
        require(found, "client state not found");
        return ClientState.decode(clientStateBytes);
    }

    function getConsensusState(IBCHost host, string memory clientId, uint64 height) public view returns (ConsensusState.Data memory) {
        (bytes memory consensusStateBytes, bool found) = host.getConsensusState(clientId, height);
        require(found, "clientState not found");
        return ConsensusState.decode(consensusStateBytes);
    }

    function parseHeader(bytes memory headerBytes) internal pure returns (uint64, uint64) {
        Header.Data memory header = Header.decode(headerBytes);
        return (header.height, header.timestamp);
    }
}
