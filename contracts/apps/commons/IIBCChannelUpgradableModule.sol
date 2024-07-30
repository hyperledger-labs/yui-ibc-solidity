// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Channel, UpgradeFields, Timeout} from "../../proto/Channel.sol";

interface IIBCChannelUpgradableModuleErrors {
    // ------------------- Errors ------------------- //

    error IBCChannelUpgradableModuleUnauthorizedUpgrader();
    error IBCChannelUpgradableModuleInvalidTimeout();
    error IBCChannelUpgradableModuleInvalidConnectionHops();
    error IBCChannelUpgradableModuleUpgradeAlreadyExists();
    error IBCChannelUpgradableModuleUpgradeNotFound();
    error IBCChannelUpgradableModuleInvalidUpgrade();

    error IBCChannelUpgradableModuleCannotRemoveInProgressUpgrade();
    /// @param state The current state of the channel
    error IBCChannelUpgradableModuleChannelNotFlushingState(Channel.State state);
    /// @param actual The actual upgrade sequence
    error IBCChannelUpgradableModuleSequenceMismatch(uint64 actual);

    error IBCChannelUpgradableModuleChannelNotFound();
    error IBCChannelUpgradableModuleCannotOverwriteUpgrade();
}

interface IIBCChannelUpgradableModule {
    // ------------------- Data Structures ------------------- //

    /**
     * @dev Proposed upgrade fields
     * @param fields Upgrade fields
     * @param timeout Absolute timeout for the upgrade
     */
    struct UpgradeProposal {
        UpgradeFields.Data fields;
        Timeout.Data timeout;
    }

    /**
     * @dev Allowed transition for the channel upgrade
     * @param flushComplete Whether the upgrade is allowed to transition to the flush complete state
     */
    struct AllowedTransition {
        bool flushComplete;
    }

    // ------------------- Functions ------------------- //

    /**
     * @dev Returns the proposed upgrade for the given port, channel, and sequence
     */
    function getUpgradeProposal(string calldata portId, string calldata channelId)
        external
        view
        returns (UpgradeProposal memory);

    /**
     * @dev Propose an upgrade for the given port, channel, and sequence
     * @notice This function is only callable by an authorized upgrader
     * The upgrader must call this function before calling `channelUpgradeInit` or `channelUpgradeTry` of the IBC handler
     */
    function proposeUpgrade(
        string calldata portId,
        string calldata channelId,
        UpgradeFields.Data calldata upgradeFields,
        Timeout.Data calldata timeout
    ) external;

    /**
     * @dev Removes the proposed upgrade for the given port and channel
     * @notice This function is only callable by an authorized upgrader
     * @param portId Port identifier
     * @param channelId Channel identifier
     */
    function removeUpgradeProposal(string calldata portId, string calldata channelId) external;

    /**
     * @dev Allow the upgrade to transition to the flush complete state
     * @notice This function is only callable by an authorized upgrader
     * WARNING: Before calling this function, the upgrader must ensure that all inflight packets have been received on the receiving chain,
     * and all acknowledgements written have been acknowledged on the sending chain
     */
    function allowTransitionToFlushComplete(string calldata portId, string calldata channelId, uint64 upgradeSequence)
        external;
}
