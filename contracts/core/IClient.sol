// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../proto/Client.sol";

interface IClient {
    function createClient(
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata clientStateBytes,
        bytes calldata consensusStateBytes
    ) external returns (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok);

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (uint64, bool);

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(string calldata clientId) external view returns (Height.Data memory, bool);

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
        returns (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok);

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
}

struct ConsensusStateUpdates {
    bytes32 consensusStateCommitment;
    Height.Data height;
}
