// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";

/**
 * @dev This defines an interface for Light Client contract can be integrated with ibc-solidity.
 * You can register the Light Client contract that implements this through `registerClient` on IBCHandler.
 */
interface ILightClient {
    /**
     * @dev ConsensusStateUpdate represents a consensus state update.
     */
    struct ConsensusStateUpdate {
        // commitment for updated consensusState
        bytes32 consensusStateCommitment;
        // updated height
        Height.Data height;
    }

    /**
     * @dev ClientStatus represents the status of a client.
     */
    enum ClientStatus {
        Active,
        Expired,
        Frozen
    }

    /**
     * @dev initializeClient initializes a new client with the given state.
     *      If succeeded, it returns heights at which the consensus state are stored.
     *      The function must be only called by IBCHandler.
     */
    function initializeClient(
        string calldata clientId,
        bytes calldata protoClientState,
        bytes calldata protoConsensusState
    ) external returns (Height.Data memory height);

    /**
     * @dev routeUpdateClient returns the calldata to the receiving function of the client message.
     *      Light client contract may encode a client message as other encoding scheme(e.g. ethereum ABI)
     *      Check ADR-001 for details.
     */
    function routeUpdateClient(string calldata clientId, bytes calldata protoClientMessage)
        external
        pure
        returns (bytes4 selector, bytes memory args);

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     *      The timestamp is nanoseconds since unix epoch.
     */
    function getTimestampAtHeight(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (uint64);

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(string calldata clientId) external view returns (Height.Data memory);

    /**
     * @dev getStatus returns the status of the client corresponding to `clientId`.
     */
    function getStatus(string calldata clientId) external view returns (ClientStatus);

    /**
     * @dev verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyMembership(
        string calldata clientId,
        Height.Data calldata height,
        uint64 delayTimePeriod,
        uint64 delayBlockPeriod,
        bytes calldata proof,
        bytes calldata prefix,
        bytes calldata path,
        bytes calldata value
    ) external returns (bool);

    /**
     * @dev verifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyNonMembership(
        string calldata clientId,
        Height.Data calldata height,
        uint64 delayTimePeriod,
        uint64 delayBlockPeriod,
        bytes calldata proof,
        bytes calldata prefix,
        bytes calldata path
    ) external returns (bool);

    /**
     * @dev getClientState returns the clientState corresponding to `clientId`.
     *      If it's not found, the function returns false.
     */
    function getClientState(string calldata clientId) external view returns (bytes memory, bool);

    /**
     * @dev getConsensusState returns the consensusState corresponding to `clientId` and `height`.
     *      If it's not found, the function returns false.
     */
    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (bytes memory, bool);
}
