// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ILightClient} from "../core/02-client/ILightClient.sol";
import {ILightClientErrors} from "../core/02-client/ILightClientErrors.sol";
import {IBCHeight} from "../core/02-client/IBCHeight.sol";
import {IIBCHandler} from "../core/25-handler/IIBCHandler.sol";
import {Height} from "../proto/Client.sol";
import {IbcLightclientsLocalhostV2ClientState as ClientState} from "../proto/Localhost.sol";
import {GoogleProtobufAny as Any} from "../proto/GoogleProtobufAny.sol";

/**
 * @title LocalhostClient
 * @notice LocalhostClient is a light client to facilitate testing of IBC Apps on a single chain
 * @dev LocalhostClient implements [09-localhost](https://github.com/cosmos/ibc/tree/main/spec/client/ics-009-loopback-cilent), but the following differences:
 * - The client identifier is `09-localhost-0`, not `09-localhost`
 * - `getLatestHeight` always returns the current block number
 * - `verifyMembership` checks the proof height is not greater than the current block height
 */
contract LocalhostClient is ILightClient, ILightClientErrors {
    using IBCHeight for Height.Data;

    error InvalidClientID();
    error InvalidConsensusState();
    error InvalidHeightRevisionNumber();
    error InvalidHeightRevisionHeight();
    error InvalidProof();
    error InvalidPrefix();

    address public immutable ibcHandler;

    constructor(address ibcHandler_) {
        ibcHandler = ibcHandler_;
    }

    /**
     * @dev initializeClient initializes a new localhost client with the given client identifier, client state, and consensus state.
     * @param clientId the client identifier must be match with `LocalhostClientLib.CLIENT_ID`
     * @param protoClientState the client state's latest height must be match with the current block number
     * @param protoConsensusState the consensus state must be match with the sentinel consensus state
     */
    function initializeClient(
        string calldata clientId,
        bytes calldata protoClientState,
        bytes calldata protoConsensusState
    ) public virtual override onlyIBC returns (Height.Data memory height) {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            revert InvalidClientID();
        }
        if (keccak256(protoConsensusState) != keccak256(LocalhostClientLib.sentinelConsensusState())) {
            revert InvalidConsensusState();
        }
        ClientState.Data memory clientState = LocalhostClientLib.unmarshalClientState(protoClientState);
        if (clientState.latest_height.revision_number != 0) {
            revert InvalidHeightRevisionNumber();
        } else if (clientState.latest_height.revision_height != uint64(block.number)) {
            revert InvalidHeightRevisionHeight();
        }
        return Height.Data({revision_number: 0, revision_height: uint64(block.number)});
    }

    /**
     * @dev routeUpdateClient returns the calldata to the receiving function of the client message.
     *      The light client encodes a client message as ethereum ABI.
     */
    function routeUpdateClient(string calldata clientId, bytes calldata)
        public
        pure
        virtual
        override
        returns (bytes4, bytes memory)
    {
        return (this.updateClient.selector, abi.encode(clientId));
    }

    /**
     * @dev updateClient updates the client state commitment with the current block number.
     * @param clientId the client identifier must be match with `LocalhostClientLib.CLIENT_ID`
     */
    function updateClient(string calldata clientId) public returns (Height.Data[] memory heights) {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            revert InvalidClientID();
        }
        IIBCHandler(ibcHandler).updateClientCommitments(clientId, new Height.Data[](0));
        return heights;
    }

    /**
     * @dev getTimestampAtHeight always returns the current block timestamp.
     */
    function getTimestampAtHeight(string calldata, Height.Data calldata height) public view returns (uint64) {
        if (height.revision_number != 0) {
            revert InvalidHeightRevisionNumber();
        } else if (height.revision_height > block.number) {
            revert InvalidHeightRevisionHeight();
        }
        return uint64(block.timestamp);
    }

    /**
     * @dev getLatestHeight always returns the current block height.
     */
    function getLatestHeight(string calldata) public view returns (Height.Data memory) {
        return Height.Data({revision_number: 0, revision_height: uint64(block.number)});
    }

    /**
     * @dev getStatus returns the status of the client corresponding to `clientId`.
     */
    function getStatus(string calldata clientId) public view virtual override returns (ClientStatus) {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            revert InvalidClientID();
        }
        return ClientStatus.Active;
    }

    /**
     * @dev verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyMembership(
        string memory clientId,
        Height.Data memory proofHeight,
        uint64,
        uint64,
        bytes memory proof,
        bytes memory prefix,
        bytes memory path,
        bytes memory value
    ) public view virtual override returns (bool) {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            revert InvalidClientID();
        }
        if (proofHeight.revision_number != 0) {
            revert InvalidHeightRevisionNumber();
        }
        if (proofHeight.revision_height > block.number) {
            revert InvalidHeightRevisionHeight();
        }
        if (keccak256(proof) != keccak256(LocalhostClientLib.sentinelProof())) {
            revert InvalidProof();
        }
        if (keccak256(IIBCHandler(ibcHandler).getCommitmentPrefix()) != keccak256(prefix)) {
            revert InvalidPrefix();
        }
        return IIBCHandler(ibcHandler).getCommitment(keccak256(path)) == keccak256(value);
    }

    /**
     * @dev verifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyNonMembership(
        string memory clientId,
        Height.Data memory proofHeight,
        uint64,
        uint64,
        bytes memory proof,
        bytes memory prefix,
        bytes memory path
    ) public view returns (bool) {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            revert InvalidClientID();
        }
        if (proofHeight.revision_number != 0) {
            revert InvalidHeightRevisionNumber();
        }
        if (proofHeight.revision_height > block.number) {
            revert InvalidHeightRevisionHeight();
        }
        if (keccak256(proof) != keccak256(LocalhostClientLib.sentinelProof())) {
            revert InvalidProof();
        }
        if (keccak256(IIBCHandler(ibcHandler).getCommitmentPrefix()) != keccak256(prefix)) {
            revert InvalidPrefix();
        }
        return IIBCHandler(ibcHandler).getCommitment(keccak256(path)) == bytes32(0);
    }

    /**
     * @dev getClientState returns the client state corresponding to `clientId`.
     * @param clientId the client identifier must be match with `LocalhostClientLib.CLIENT_ID`
     */
    function getClientState(string calldata clientId) public view returns (bytes memory, bool) {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            return (new bytes(0), false);
        }
        return (
            LocalhostClientLib.marshalClientState(
                ClientState.Data({
                    latest_height: Height.Data({revision_number: 0, revision_height: uint64(block.number)})
                })
                ),
            true
        );
    }

    /**
     * @dev getConsensusState always returns the sentinel consensus state.
     * @param clientId the client identifier must be match with `LocalhostClientLib.CLIENT_ID`
     */
    function getConsensusState(string calldata clientId, Height.Data calldata)
        public
        pure
        returns (bytes memory, bool)
    {
        if (keccak256(abi.encodePacked(clientId)) != keccak256(abi.encodePacked(LocalhostClientLib.CLIENT_ID))) {
            revert InvalidClientID();
        }
        return (LocalhostClientLib.sentinelConsensusState(), true);
    }

    modifier onlyIBC() {
        if (msg.sender != ibcHandler) {
            revert InvalidCaller(msg.sender);
        }
        _;
    }
}

/**
 * @title LocalhostClientLib
 * @notice LocalhostClientLib is a library that provides the client type, client identifier, client state type URL, consensus state type URL, and helper functions for the localhost client.
 */
library LocalhostClientLib {
    /// @param url the unexpected type URL
    error UnexpectedProtoAnyTypeURL(string url);

    string internal constant CLIENT_TYPE = "09-localhost";
    string internal constant CLIENT_ID = string(abi.encodePacked(CLIENT_TYPE, "-0"));
    string internal constant CLIENT_STATE_TYPE_URL = "/ibc.lightclients.localhost.v2.ClientState";
    string internal constant CONSENSUS_STATE_TYPE_URL = "/ibc.lightclients.localhost.v2.ConsensusState";
    bytes32 internal constant CLIENT_STATE_TYPE_URL_HASH = keccak256(abi.encodePacked(CLIENT_STATE_TYPE_URL));

    function marshalClientState(ClientState.Data memory clientState) internal pure returns (bytes memory) {
        return Any.encode(Any.Data({type_url: CLIENT_STATE_TYPE_URL, value: ClientState.encode(clientState)}));
    }

    function unmarshalClientState(bytes calldata protoClientState) internal pure returns (ClientState.Data memory) {
        Any.Data memory any = Any.decode(protoClientState);
        if (keccak256(abi.encodePacked(any.type_url)) != CLIENT_STATE_TYPE_URL_HASH) {
            revert UnexpectedProtoAnyTypeURL(any.type_url);
        }
        return ClientState.decode(any.value);
    }

    function sentinelConsensusState() internal pure returns (bytes memory) {
        return Any.encode(Any.Data({type_url: CONSENSUS_STATE_TYPE_URL, value: new bytes(0)}));
    }

    function sentinelProof() internal pure returns (bytes memory) {
        return hex"01";
    }
}
