// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import "./IBCHeight.sol";
import "../proto/Client.sol";
import {
    IbcLightclientsMockV1ClientState as ClientState,
    IbcLightclientsMockV1ConsensusState as ConsensusState,
    IbcLightclientsMockV1Header as Header
} from "../proto/MockClient.sol";
import {GoogleProtobufAny as Any} from "../proto/GoogleProtobufAny.sol";
import "../lib/Bytes.sol";

// MockClient implements https://github.com/datachainlab/ibc-mock-client
// WARNING: This client is intended to be used for testing purpose. Therefore, it is not generally available in a production, except in a fully trusted environment.
contract MockClient is IClient {
    using Bytes for bytes;
    using IBCHeight for Height.Data;

    bytes32 private constant HEADER_TYPE_URL_HASH = keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.Header"));
    bytes32 private constant CLIENT_STATE_TYPE_URL_HASH =
        keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.ClientState"));
    bytes32 private constant CONSENSUS_STATE_TYPE_URL_HASH =
        keccak256(abi.encodePacked("/ibc.lightclients.mock.v1.ConsensusState"));

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(IBCHost host, string memory clientId, Height.Data memory height)
        public
        view
        override
        returns (uint64, bool)
    {
        (ConsensusState.Data memory consensusState, bool found) = getConsensusState(host, clientId, height);
        if (!found) {
            return (0, false);
        }
        return (consensusState.timestamp, true);
    }

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(IBCHost host, string memory clientId)
        public
        view
        override
        returns (Height.Data memory, bool)
    {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        if (!found) {
            return (Height.Data(0, 0), false);
        }
        return (clientState.latest_height, true);
    }

    /**
     * @dev verifyClientMessageAndUpdateState is intended to perform the followings:
     * 1. verify a given client message(e.g. header)
     * 2. check misbehaviour such like duplicate block height
     * 3. if misbehaviour is found, update state accordingly and return
     * 4. update state(s) with the client message
     * 5. persist the state(s) on the host
     */
    function verifyClientMessageAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory clientMessageBytes
    ) public override returns (bool) {
        Height.Data memory height;
        uint64 timestamp;
        Any.Data memory anyClientState;
        Any.Data memory anyConsensusState;

        anyClientState = Any.decode(clientStateBytes);
        require(
            keccak256(abi.encodePacked(anyClientState.type_url)) == CLIENT_STATE_TYPE_URL_HASH, "invalid client type"
        );
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

    function verifyMembership(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        uint64,
        uint64,
        bytes memory proof,
        bytes memory,
        bytes memory,
        bytes memory value
    ) external view override returns (bool) {
        (, bool found) = host.getConsensusState(clientId, height);
        require(found, "consensus state not found");
        return sha256(value) == proof.toBytes32();
    }

    function getClientState(IBCHost host, string memory clientId)
        public
        view
        returns (ClientState.Data memory clientState, bool found)
    {
        bytes memory clientStateBytes;
        (clientStateBytes, found) = host.getClientState(clientId);
        if (!found) {
            return (clientState, false);
        }
        return (ClientState.decode(Any.decode(clientStateBytes).value), true);
    }

    function getConsensusState(IBCHost host, string memory clientId, Height.Data memory height)
        public
        view
        returns (ConsensusState.Data memory consensusState, bool found)
    {
        bytes memory consensusStateBytes;
        (consensusStateBytes, found) = host.getConsensusState(clientId, height);
        if (!found) {
            return (consensusState, false);
        }
        Any.Data memory any = Any.decode(consensusStateBytes);
        require(
            keccak256(abi.encodePacked(any.type_url)) == CONSENSUS_STATE_TYPE_URL_HASH, "invalid consensus state type"
        );
        return (ConsensusState.decode(any.value), true);
    }

    function parseHeader(bytes memory headerBytes) internal pure returns (Height.Data memory, uint64) {
        Any.Data memory any = Any.decode(headerBytes);
        require(keccak256(abi.encodePacked(any.type_url)) == HEADER_TYPE_URL_HASH, "invalid header type");
        Header.Data memory header = Header.decode(any.value);
        return (header.height, header.timestamp);
    }
}
