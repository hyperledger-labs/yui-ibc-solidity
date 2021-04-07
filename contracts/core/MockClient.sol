pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import {MockClientState as ClientState, MockConsensusState as ConsensusState} from "./types/MockClient.sol";
import "../lib/ECRecovery.sol";
import "../lib/Bytes.sol";
import "../lib/TrieProofs.sol";
import "../lib/RLP.sol";

contract MockClient is IClient {
    using RLP for RLP.RLPItem;
    using RLP for bytes;
    using Bytes for bytes;

    struct Header {
        bytes32 stateRoot;
        uint64 height;
        uint64 time;
    }

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
        Header memory header;
        ConsensusState.Data memory consensusState;
        ClientState.Data memory clientState = ClientState.decode(clientStateBytes);

        header = parseETHHeader(headerBytes);
        if (header.height > clientState.latest_height) {
            clientState.latest_height = header.height;
        }
        consensusState = ConsensusState.Data({timestamp: header.time});
        return (ClientState.encode(clientState), ConsensusState.encode(consensusState), header.height);
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
        return true;
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
        return true;
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
        return true;
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
        return true;
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
        bytes32 commitmentBytes // serialized with pb
    ) public override view returns (bool) {
        return true;
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
        bytes32 ackCommitmentBytes // serialized with pb
    ) public override view returns (bool) {
        return true;
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

    function parseETHHeader(bytes memory headerBytes) internal pure returns (Header memory header) {
        RLP.RLPItem[] memory items = headerBytes.toRLPItem().toList();
        require(items.length == 15, "items length must be 15");
        header.stateRoot = items[3].toBytes().toBytes32();
        header.height = uint64(items[8].toUint());
        header.time = uint64(items[11].toUint());
        return header;
    }
}
