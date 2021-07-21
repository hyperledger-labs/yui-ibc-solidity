pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import {
    IbcLightclientsMockV1ClientState as ClientState,
    IbcLightclientsMockV1ConsensusState as ConsensusState,
    IbcLightclientsMockV1Header as Header
} from "./types/MockClient.sol";
import {GoogleProtobufAny as Any} from "./types/GoogleProtobufAny.sol";
import "../lib/Bytes.sol";

contract MockClient is IClient {
    using Bytes for bytes;

    struct protoTypes {
        bytes32 clientState;
        bytes32 consensusState;
        bytes32 header;
    }

    protoTypes pts;

    constructor() public {
        // TODO The typeUrl should be defined in types/MockClient.sol
        // The schema of typeUrl follows cosmos/cosmos-sdk/codec/types/any.go
        pts = protoTypes({
            clientState: keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.ClientState")),
            consensusState: keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.ConsensusState")),
            header: keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.Header"))
        });
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string memory clientId,
        uint64 height
    ) public override view returns (uint64, bool) {
        (ConsensusState.Data memory consensusState, bool found) = getConsensusState(host, clientId, height);
        if (!found) {
            return (0, false);
        }
        return (consensusState.timestamp, true);
    }

    function getLatestHeight(
        IBCHost host,
        string memory clientId
    ) public override view returns (uint64, bool) {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        if (!found) {
            return (0, false);
        }
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
        Any.Data memory anyClientState;
        Any.Data memory anyConsensusState;

        anyClientState = Any.decode(clientStateBytes);
        require(keccak256(abi.encodePacked(anyClientState.type_url)) == pts.clientState, "invalid client type");
        ClientState.Data memory clientState = ClientState.decode(anyClientState.value);
        (height, timestamp) = parseHeader(headerBytes);
        if (height > clientState.latest_height) {
            clientState.latest_height = height;
        }

        anyClientState.value = ClientState.encode(clientState);
        anyConsensusState.type_url = "/ibc.lightclients.mock.v1.ConsensusState";
        anyConsensusState.value = ConsensusState.encode(ConsensusState.Data({timestamp: timestamp}));
        return (Any.encode(anyClientState), Any.encode(anyConsensusState), height);
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

    function getClientState(IBCHost host, string memory clientId) public view returns (ClientState.Data memory clientState, bool found) {
        bytes memory clientStateBytes;
        (clientStateBytes, found) = host.getClientState(clientId);
        if (!found) {
            return (clientState, false);
        }
        return (ClientState.decode(Any.decode(clientStateBytes).value), true);
    }

    function getConsensusState(IBCHost host, string memory clientId, uint64 height) public view returns (ConsensusState.Data memory consensusState, bool found) {
        bytes memory consensusStateBytes;
        (consensusStateBytes, found) = host.getConsensusState(clientId, height);
        if (!found) {
            return (consensusState, false);
        }
        return (ConsensusState.decode(Any.decode(consensusStateBytes).value), true);
    }

    function parseHeader(bytes memory headerBytes) internal view returns (uint64, uint64) {
        Any.Data memory any = Any.decode(headerBytes);
        require(keccak256(abi.encodePacked(any.type_url)) == pts.header, "invalid header type");
        Header.Data memory header = Header.decode(any.value);
        return (header.height, header.timestamp);
    }
}
