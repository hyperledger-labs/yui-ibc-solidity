// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ILightClient} from "../core/02-client/ILightClient.sol";
import {IBCHeight} from "../core/02-client/IBCHeight.sol";
import {IIBCHandler} from "../core/25-handler/IIBCHandler.sol";
import {Height} from "../proto/Client.sol";
import {
    IbcLightclientsMockV1ClientState as ClientState,
    IbcLightclientsMockV1ConsensusState as ConsensusState,
    IbcLightclientsMockV1Header as Header
} from "../proto/MockClient.sol";
import {GoogleProtobufAny as Any} from "../proto/GoogleProtobufAny.sol";

/// @notice MockClient implements https://github.com/datachainlab/ibc-mock-client
/// WARNING: This client is intended to be used for testing purpose.
contract MockClient is Ownable, ILightClient {
    using IBCHeight for Height.Data;

    /// @param caller the caller of the function
    error InvalidCaller(address caller);
    error InvalidClientState();
    error InvalidConsensusState();
    error InvalidHeader();
    error InvalidProof();
    error InvalidPrefix();
    /// @param clientId client identifier
    error ClientStateNotFound(string clientId);
    /// @param clientId client identifier
    /// @param height consensus height
    error ConsensusStateNotFound(string clientId, Height.Data height);
    error NotActiveClient();
    /// @param url type url of the any
    error UnexpectedProtoAnyTypeURL(string url);

    string internal constant HEADER_TYPE_URL = "/ibc.lightclients.mock.v1.Header";
    string internal constant CLIENT_STATE_TYPE_URL = "/ibc.lightclients.mock.v1.ClientState";
    string internal constant CONSENSUS_STATE_TYPE_URL = "/ibc.lightclients.mock.v1.ConsensusState";

    bytes32 internal constant HEADER_TYPE_URL_HASH = keccak256(abi.encodePacked(HEADER_TYPE_URL));
    bytes32 internal constant CLIENT_STATE_TYPE_URL_HASH = keccak256(abi.encodePacked(CLIENT_STATE_TYPE_URL));
    bytes32 internal constant CONSENSUS_STATE_TYPE_URL_HASH = keccak256(abi.encodePacked(CONSENSUS_STATE_TYPE_URL));

    address internal immutable ibcHandler;

    mapping(string => ClientState.Data) internal clientStates;
    mapping(string => mapping(uint128 => ConsensusState.Data)) internal consensusStates;
    mapping(string => ClientStatus) internal statuses;

    constructor(address ibcHandler_) {
        ibcHandler = ibcHandler_;
    }

    /**
     * @dev initializeClient creates a new client with the given state
     */
    function initializeClient(
        string calldata clientId,
        bytes calldata protoClientState,
        bytes calldata protoConsensusState
    ) external virtual override onlyIBC returns (Height.Data memory height) {
        ClientState.Data memory clientState = unmarshalClientState(protoClientState);
        ConsensusState.Data memory consensusState = unmarshalConsensusState(protoConsensusState);
        if (clientState.latest_height.revision_number != 0 || clientState.latest_height.revision_height == 0) {
            revert InvalidClientState();
        }
        if (consensusState.timestamp == 0) {
            revert InvalidConsensusState();
        }
        clientStates[clientId] = clientState;
        consensusStates[clientId][clientState.latest_height.toUint128()] = consensusState;
        statuses[clientId] = ClientStatus.Active;
        return clientState.latest_height;
    }

    /**
     * @dev routeUpdateClient returns the calldata to the receiving function of the client message.
     *      The light client encodes a client message as ethereum ABI.
     */
    function routeUpdateClient(string calldata clientId, bytes calldata protoClientMessage)
        external
        pure
        virtual
        override
        returns (bytes4, bytes memory)
    {
        Any.Data memory any = Any.decode(protoClientMessage);
        if (keccak256(abi.encodePacked(any.type_url)) != HEADER_TYPE_URL_HASH) {
            revert UnexpectedProtoAnyTypeURL(any.type_url);
        }
        Header.Data memory header = Header.decode(any.value);
        if (header.height.revision_number != 0 || header.height.revision_height == 0 || header.timestamp == 0) {
            revert InvalidHeader();
        }
        return (this.updateClient.selector, abi.encode(clientId, header));
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     *      The timestamp is nanoseconds since unix epoch.
     */
    function getTimestampAtHeight(string calldata clientId, Height.Data calldata height)
        external
        view
        virtual
        override
        returns (uint64)
    {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        if (consensusState.timestamp == 0) {
            revert ConsensusStateNotFound(clientId, height);
        }
        return consensusState.timestamp;
    }

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(string calldata clientId) external view virtual override returns (Height.Data memory) {
        ClientState.Data storage clientState = clientStates[clientId];
        if (clientState.latest_height.revision_height == 0) {
            revert ClientStateNotFound(clientId);
        }
        return clientState.latest_height;
    }

    /**
     * @dev getStatus returns the status of the client corresponding to `clientId`.
     */
    function getStatus(string calldata clientId) external view virtual override returns (ClientStatus) {
        return statuses[clientId];
    }

    /**
     * @dev setStatus sets the status of the client corresponding to `clientId`.
     */
    function setStatus(string calldata clientId, ClientStatus status) external virtual onlyOwner {
        statuses[clientId] = status;
    }

    /**
     * @dev updateClient updates the client state and returns the updated heights.
     */
    function updateClient(string calldata clientId, Header.Data calldata header)
        public
        returns (Height.Data[] memory heights)
    {
        if (statuses[clientId] != ClientStatus.Active) {
            revert NotActiveClient();
        }
        if (header.height.revision_number != 0 || header.height.revision_height == 0 || header.timestamp == 0) {
            revert InvalidHeader();
        }
        if (header.height.gt(clientStates[clientId].latest_height)) {
            clientStates[clientId].latest_height = header.height;
        }
        consensusStates[clientId][header.height.toUint128()].timestamp = header.timestamp;
        heights = new Height.Data[](1);
        heights[0] = header.height;
        return heights;
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
        bytes memory prefix,
        bytes memory path,
        bytes calldata value
    ) external view virtual override returns (bool) {
        if (consensusStates[clientId][height.toUint128()].timestamp == 0) {
            revert ConsensusStateNotFound(clientId, height);
        }
        if (keccak256(IIBCHandler(ibcHandler).getCommitmentPrefix()) != keccak256(prefix)) {
            revert InvalidPrefix();
        }
        if (proof.length != 32) {
            revert InvalidProof();
        }
        return
            sha256(abi.encodePacked(height.toUint128(), sha256(prefix), sha256(path), sha256(value))) == bytes32(proof);
    }

    /**
     * @dev verifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyNonMembership(
        string calldata clientId,
        Height.Data calldata height,
        uint64,
        uint64,
        bytes calldata proof,
        bytes memory prefix,
        bytes memory path
    ) external view virtual override returns (bool) {
        if (consensusStates[clientId][height.toUint128()].timestamp == 0) {
            revert ConsensusStateNotFound(clientId, height);
        }
        if (keccak256(IIBCHandler(ibcHandler).getCommitmentPrefix()) != keccak256(prefix)) {
            revert InvalidPrefix();
        }
        return sha256(abi.encodePacked(height.toUint128(), sha256(prefix), sha256(path), sha256(""))) == bytes32(proof);
    }

    /* State accessors */

    /**
     * @dev getClientState returns the clientState corresponding to `clientId`.
     *      If it's not found, the function returns false.
     */
    function getClientState(string calldata clientId)
        external
        view
        virtual
        returns (bytes memory clientStateBytes, bool)
    {
        ClientState.Data storage clientState = clientStates[clientId];
        if (clientState.latest_height.revision_height == 0) {
            return (clientStateBytes, false);
        }
        return (Any.encode(Any.Data({type_url: CLIENT_STATE_TYPE_URL, value: ClientState.encode(clientState)})), true);
    }

    /**
     * @dev getConsensusState returns the consensusState corresponding to `clientId` and `height`.
     *      If it's not found, the function returns false.
     */
    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        virtual
        returns (bytes memory consensusStateBytes, bool)
    {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        if (consensusState.timestamp == 0) {
            return (consensusStateBytes, false);
        }
        return (
            Any.encode(Any.Data({type_url: CONSENSUS_STATE_TYPE_URL, value: ConsensusState.encode(consensusState)})),
            true
        );
    }

    /* Internal functions */

    function unmarshalClientState(bytes calldata bz) internal pure returns (ClientState.Data memory clientState) {
        Any.Data memory anyClientState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyClientState.type_url)) != CLIENT_STATE_TYPE_URL_HASH) {
            revert UnexpectedProtoAnyTypeURL(anyClientState.type_url);
        }
        return ClientState.decode(anyClientState.value);
    }

    function unmarshalConsensusState(bytes calldata bz)
        internal
        pure
        returns (ConsensusState.Data memory consensusState)
    {
        Any.Data memory anyConsensusState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyConsensusState.type_url)) != CONSENSUS_STATE_TYPE_URL_HASH) {
            revert UnexpectedProtoAnyTypeURL(anyConsensusState.type_url);
        }
        return ConsensusState.decode(anyConsensusState.value);
    }

    modifier onlyIBC() {
        if (msg.sender != ibcHandler) {
            revert InvalidCaller(msg.sender);
        }
        _;
    }
}
