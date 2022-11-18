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

    mapping(string => ClientState.Data) clientStates;
    mapping(string => mapping(uint128 => ConsensusState.Data)) consensusStates;

    function createClient(
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata clientStateBytes,
        bytes calldata consensusStateBytes
    ) external override returns (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;

        (clientState, ok) = unmarshalClientState(clientStateBytes);
        if (!ok) {
            return (clientStateCommitment, updates, false);
        }
        (consensusState, ok) = unmarshalConsensusState(consensusStateBytes);
        if (!ok) {
            return (clientStateCommitment, updates, false);
        }
        if (
            clientState.latest_height.revision_number != 0 || clientState.latest_height.revision_height == 0
                || consensusState.timestamp == 0
        ) {
            return (clientStateCommitment, updates, false);
        }
        clientStates[clientId] = clientState;
        consensusStates[clientId][height.toUint128()] = consensusState;
        updates = new ConsensusStateUpdates[](1);
        updates[0] = ConsensusStateUpdates({consensusStateCommitment: keccak256(consensusStateBytes), height: height});
        return (keccak256(clientStateBytes), updates, true);
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(string calldata clientId, Height.Data calldata height)
        external
        view
        override
        returns (uint64, bool)
    {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        return (consensusState.timestamp, consensusState.timestamp != 0);
    }

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(string calldata clientId) external view override returns (Height.Data memory, bool) {
        ClientState.Data storage clientState = clientStates[clientId];
        return (
            clientState.latest_height,
            clientState.latest_height.revision_number != 0 && clientState.latest_height.revision_height != 0
        );
    }

    /**
     * @dev verifyClientMessageAndUpdateState is intended to perform the followings:
     * 1. verify a given client message(e.g. header)
     * 2. check misbehaviour such like duplicate block height
     * 3. if misbehaviour is found, update state accordingly and return
     * 4. update state(s) with the client message
     * 5. persist the state(s) on the host
     */
    function verifyClientMessageAndUpdateState(string calldata clientId, bytes calldata clientMessageBytes)
        external
        override
        returns (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok)
    {
        Height.Data memory height;
        uint64 timestamp;
        Any.Data memory anyClientState;
        Any.Data memory anyConsensusState;

        (height, timestamp) = parseHeader(clientMessageBytes);
        if (height.gt(clientStates[clientId].latest_height)) {
            clientStates[clientId].latest_height = height;
        }
        anyClientState.type_url = "/ibc.lightclients.mock.v1.ClientState";
        anyClientState.value = ClientState.encode(clientStates[clientId]);

        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        consensusState.timestamp = timestamp;

        anyConsensusState.type_url = "/ibc.lightclients.mock.v1.ConsensusState";
        anyConsensusState.value = ConsensusState.encode(consensusState);

        updates = new ConsensusStateUpdates[](1);
        updates[0] =
            ConsensusStateUpdates({consensusStateCommitment: keccak256(Any.encode(anyConsensusState)), height: height});
        return (keccak256(Any.encode(anyClientState)), updates, true);
    }

    /**
     * @dev verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyMembership(
        string calldata clientId,
        Height.Data calldata height,
        uint64,
        uint64,
        bytes calldata proof,
        bytes memory,
        bytes memory,
        bytes calldata value
    ) external view override returns (bool) {
        require(consensusStates[clientId][height.toUint128()].timestamp != 0, "consensus state not found");
        return sha256(value) == proof.toBytes32();
    }

    /* Internal functions */

    function parseHeader(bytes memory bz) internal pure returns (Height.Data memory, uint64) {
        Any.Data memory any = Any.decode(bz);
        require(keccak256(abi.encodePacked(any.type_url)) == HEADER_TYPE_URL_HASH, "invalid header type");
        Header.Data memory header = Header.decode(any.value);
        require(
            header.height.revision_number == 0 && header.height.revision_height != 0 && header.timestamp != 0,
            "invalid header"
        );
        return (header.height, header.timestamp);
    }

    function unmarshalClientState(bytes memory bz)
        internal
        pure
        returns (ClientState.Data memory clientState, bool ok)
    {
        Any.Data memory anyClientState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyClientState.type_url)) != CLIENT_STATE_TYPE_URL_HASH) {
            return (clientState, false);
        }
        return (ClientState.decode(anyClientState.value), true);
    }

    function unmarshalConsensusState(bytes memory bz)
        internal
        pure
        returns (ConsensusState.Data memory consensusState, bool ok)
    {
        Any.Data memory anyConsensusState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyConsensusState.type_url)) != CONSENSUS_STATE_TYPE_URL_HASH) {
            return (consensusState, false);
        }
        return (ConsensusState.decode(anyConsensusState.value), true);
    }
}
