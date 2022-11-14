// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import "./IBCHeight.sol";
import "./types/Client.sol";
import {
    IbcLightclientsMockV1ClientState as ClientState,
    IbcLightclientsMockV1ConsensusState as ConsensusState,
    IbcLightclientsMockV1Header as Header
} from "./types/MockClient.sol";
import {GoogleProtobufAny as Any} from "./types/GoogleProtobufAny.sol";
import "../lib/Bytes.sol";

// MockClient implements https://github.com/datachainlab/ibc-mock-client
// WARNING: This client is intended to be used for testing purpose. Therefore, it is not generally available in a production, except in a fully trusted environment.
contract MockClient is IClient {
    using Bytes for bytes;
    using IBCHeight for Height.Data;

    bytes32 private constant headerTypeUrlHash = keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.Header"));
    bytes32 private constant clientStateTypeUrlHash = keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.ClientState"));
    bytes32 private constant consensusStateTypeUrlHash = keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.ConsensusState"));

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string memory clientId,
        Height.Data memory height
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
    ) public override view returns (Height.Data memory, bool) {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        if (!found) {
            return (Height.Data(0, 0), false);
        }
        return (clientState.latest_height, true);
    }

    /**
     * @dev verifyClientMessageAndUpdateState is intended to perform:
     * 1. client message verification
     * 2. check for duplicate height misbehaviour
     * 3. if misbehaviour is found, update state accordingly and return
     * 4. update state in a way that verified headers carrying one or more consensus states can be updated
     * 5. persist the state internally
     * 6. return an array of consensus heights
     */
    function verifyClientMessageAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory clientMessageBytes
    ) public override returns (bool) {
        // verify clientMessageBytes

        // check for misbehaviour

        // updates state upon misbehaviour, freezing the ClientState.
        // This method should only be called when misbehaviour is detected
        // as it does not perform any misbehaviour checks.

        // if client message is verified and there is no misbehaviour, update state
        Height.Data memory height;
        uint64 timestamp;
        Any.Data memory anyClientState;
        Any.Data memory anyConsensusState;

        anyClientState = Any.decode(clientStateBytes);
        require(keccak256(abi.encodePacked(anyClientState.type_url)) == clientStateTypeUrlHash, "invalid client type");
        ClientState.Data memory clientState = ClientState.decode(anyClientState.value);
        (height, timestamp) = parseHeader(clientMessageBytes);
        if (height.gt(clientState.latest_height)) {
            clientState.latest_height = height;
        }

        anyClientState.value = ClientState.encode(clientState);
        anyConsensusState.type_url = "/ibc.lightclients.mock.v1.ConsensusState";
        anyConsensusState.value = ConsensusState.encode(ConsensusState.Data({timestamp: timestamp}));

        host.setClientState(clientId, Any.encode(anyClientState));
        host.setConsensusState(clientId, height, Any.encode(anyConsensusState));
        host.setProcessedTime(clientId, height, block.timestamp);
        host.setProcessedHeight(clientId, height, block.number);
        return true;
    }

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        bytes memory,
        string memory,
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
        Height.Data memory height,
        string memory,
        Height.Data memory,
        bytes memory,
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
        Height.Data memory height,
        bytes memory,
        bytes memory proof,
        string memory,
        bytes memory connectionBytes // serialized with pb
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(connectionBytes) == proof.toBytes32();
    }

    function verifyChannelState(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        bytes memory,
        bytes memory proof,
        string memory,
        string memory,
        bytes memory channelBytes // serialized with pb
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(channelBytes) == proof.toBytes32();
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        uint64,
        uint64,
        bytes memory,
        bytes memory proof,
        string memory,
        string memory,
        uint64,
        bytes32 commitmentBytes
    ) public override view returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return commitmentBytes == proof.toBytes32();
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        uint64,
        uint64,
        bytes memory,
        bytes memory proof,
        string memory,
        string memory,
        uint64,
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

    function getConsensusState(IBCHost host, string memory clientId, Height.Data memory height) public view returns (ConsensusState.Data memory consensusState, bool found) {
        bytes memory consensusStateBytes;
        (consensusStateBytes, found) = host.getConsensusState(clientId, height);
        if (!found) {
            return (consensusState, false);
        }
        return (ConsensusState.decode(Any.decode(consensusStateBytes).value), true);
    }

    function parseHeader(bytes memory headerBytes) internal pure returns (Height.Data memory, uint64) {
        Any.Data memory any = Any.decode(headerBytes);
        require(keccak256(abi.encodePacked(any.type_url)) == headerTypeUrlHash, "invalid header type");
        Header.Data memory header = Header.decode(any.value);
        return (header.height, header.timestamp);
    }
}
