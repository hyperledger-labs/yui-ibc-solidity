// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {Channel, Upgrade, UpgradeFields, ErrorReceipt, Timeout} from "../../proto/Channel.sol";

interface IIBCChannelUpgrade {
    // --------------------- Events --------------------- //

    /// @notice emitted when channelUpgradeInit is successfully executed
    event ChannelUpgradeInit(
        string portId, string channelId, uint64 upgradeSequence, UpgradeFields.Data proposedUpgradeFields
    );

    /// @notice emitted when channelUpgradeTry is successfully executed
    event ChannelUpgradeTry(
        string portId,
        string channelId,
        uint64 upgradeSequence,
        UpgradeFields.Data upgradeFields,
        Timeout.Data timeout,
        uint64 nextSequenceSend
    );

    /// @notice emitted when channelUpgradeAck is successfully executed
    /// @param channelState post channel state (FLUSHING or FLUSHCOMPLETE)
    event ChannelUpgradeAck(
        string portId,
        string channelId,
        uint64 upgradeSequence,
        Channel.State channelState,
        UpgradeFields.Data upgradeFields,
        Timeout.Data timeout,
        uint64 nextSequenceSend
    );

    /// @notice emitted when channelUpgradeConfirm is successfully executed
    /// @param channelState post channel state (FLUSHING or FLUSHCOMPLETE or OPEN)
    event ChannelUpgradeConfirm(string portId, string channelId, uint64 upgradeSequence, Channel.State channelState);

    /// @notice emitted when channelUpgradeOpen is successfully executed
    event ChannelUpgradeOpen(string portId, string channelId, uint64 upgradeSequence);

    /// @notice error when the upgrade attempt fails
    /// @param portId port identifier
    /// @param channelId channel identifier
    /// @param upgradeSequence upgrade sequence
    event WriteErrorReceipt(string portId, string channelId, uint64 upgradeSequence, string message);

    // --------------------- Data Structure --------------------- //

    enum UpgradeHandshakeError {
        None,
        Overwritten,
        OutOfSync,
        Timeout,
        Cancel,
        IncompatibleProposal,
        AckCallbackFailed
    }

    struct MsgChannelUpgradeInit {
        string portId;
        string channelId;
        UpgradeFields.Data proposedUpgradeFields;
    }

    struct MsgChannelUpgradeTry {
        string portId;
        string channelId;
        uint64 counterpartyUpgradeSequence;
        UpgradeFields.Data counterpartyUpgradeFields;
        string[] proposedConnectionHops;
        ChannelUpgradeProofs proofs;
    }

    struct MsgChannelUpgradeAck {
        string portId;
        string channelId;
        Upgrade.Data counterpartyUpgrade;
        ChannelUpgradeProofs proofs;
    }

    struct MsgChannelUpgradeConfirm {
        string portId;
        string channelId;
        Channel.State counterpartyChannelState;
        Upgrade.Data counterpartyUpgrade;
        ChannelUpgradeProofs proofs;
    }

    struct MsgChannelUpgradeOpen {
        string portId;
        string channelId;
        Channel.State counterpartyChannelState;
        uint64 counterpartyUpgradeSequence;
        bytes proofChannel;
        Height.Data proofHeight;
    }

    struct MsgCancelChannelUpgrade {
        string portId;
        string channelId;
        ErrorReceipt.Data errorReceipt;
        bytes proofUpgradeError;
        Height.Data proofHeight;
    }

    struct MsgTimeoutChannelUpgrade {
        string portId;
        string channelId;
        Channel.Data counterpartyChannel;
        bytes proofChannel;
        Height.Data proofHeight;
    }

    struct ChannelUpgradeProofs {
        bytes proofChannel;
        bytes proofUpgrade;
        Height.Data proofHeight;
    }

    // --------------------- External Functions --------------------- //

    /**
     * @dev channelUpgradeInit is called by a module to initiate a channel upgrade handshake with
     * a module on another chain.
     * The current channel state must be OPEN. The post channel state is always OPEN.
     */
    function channelUpgradeInit(MsgChannelUpgradeInit calldata msg_) external returns (uint64 upgradeSequence);

    /**
     * @dev channelUpgradeTry is called by a module to accept the first step of a channel upgrade handshake initiated by
     * a module on another chain.
     * The current channel state must be OPEN. The post channel state will be OPEN or FLUSHING.
     * If this function is successful, the proposed upgrade will be emitted as event.
     * If the upgrade fails, the upgrade sequence will still be incremented but an error will be returned.
     */
    function channelUpgradeTry(MsgChannelUpgradeTry calldata msg_) external returns (bool ok, uint64 upgradeSequence);

    /**
     * @dev channelUpgradeAck is called by a module to accept the ACKUPGRADE handshake step of the channel upgrade protocol.
     * This method will verify that the counterparty has called the channelUpgradeTry handler.
     * and that its own upgrade is compatible with the selected counterparty version.
     *
     * NOTE: The current channel may be in either the OPEN or FLUSHING state.
     * The channel may be in OPEN if we are in the happy path.
     *   A -> Init (OPEN), B -> Try (FLUSHING), A -> Ack (begins in OPEN, ends in FLUSHING or FLUSHCOMPLETE)
     *
     * The channel may be in FLUSHING if we are in a crossing hellos situation.
     *   A -> Init (OPEN), B -> Init (OPEN) -> A -> Try (FLUSHING), B -> Try (FLUSHING), A -> Ack (begins in FLUSHING, ends in FLUSHING or FLUSHCOMPLETE)
     */
    function channelUpgradeAck(MsgChannelUpgradeAck calldata msg_) external returns (bool);

    /**
     * @dev channelUpgradeConfirm is called on the chain which is on FLUSHING after channelUpgradeAck is called on the counterparty.
     * This will inform the TRY chain of the timeout set on ACK by the counterparty. If the timeout has already exceeded,
     * we will write an error receipt and restore the channel to OPEN state.
     */
    function channelUpgradeConfirm(MsgChannelUpgradeConfirm calldata msg_) external returns (bool);

    /**
     * @dev channelUpgradeOpen is called by a module to complete the channel upgrade handshake and move the channel back to an OPEN state.
     * This method should only be called after both channels have flushed any in-flight packets.
     */
    function channelUpgradeOpen(MsgChannelUpgradeOpen calldata msg_) external;

    /**
     * @dev cancelChannelUpgrade proves that an error receipt was written on the counterparty
     * which constitutes a valid situation where the upgrade should be cancelled.
     * The tx reverts if sufficient evidence for cancelling the upgrade has not been provided.
     */
    function cancelChannelUpgrade(MsgCancelChannelUpgrade calldata msg_) external;

    /**
     * @dev timeoutChannelUpgrade times out an outstanding upgrade.
     * This should be used by the initialising chain when the counterparty chain has not responded to an upgrade proposal within the specified timeout period.
     */
    function timeoutChannelUpgrade(MsgTimeoutChannelUpgrade calldata msg_) external;
}
