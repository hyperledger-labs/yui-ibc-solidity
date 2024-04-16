// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {Channel} from "../../proto/Channel.sol";
import {IIBCChannelErrors} from "./IIBCChannelErrors.sol";

interface IIBCChannelUpgradeErrors is IIBCChannelErrors {
    error IBCChannelUpgradeNoChanges();

    error IBCChannelUpgradeInvalidUpgradeFields();

    /// @notice proposal must be empty if upgrade exists
    error IBCChannelUpgradeTryProposalMustEmptyIfExist();

    /// @param upgrader address of the upgrader
    error IBCChannelUpgradeUnauthorizedChannelUpgrader(address upgrader);

    error IBCChannelUpgradeIncompatibleProposal();

    error IBCChannelUpgradeErrorReceiptEmpty();

    /// @param latestSequence latest upgrade sequence
    /// @param sequence upgrade sequence
    error IBCChannelUpgradeWriteOldErrorReceiptSequence(uint64 latestSequence, uint64 sequence);

    error IBCChannelUpgradeNoExistingUpgrade();

    /// @param state channel state
    error IBCChannelUpgradeNotFlushing(Channel.State state);

    /// @param state channel state
    error IBCChannelUpgradeNotOpenOrFlushing(Channel.State state);

    /// @param state channel state
    error IBCChannelUpgradeCounterpartyNotOpenOrFlushcomplete(Channel.State state);

    /// @notice counterparty channel not flushing or flushcomplete
    /// @param state channel state
    error IBCChannelUpgradeCounterpartyNotFlushingOrFlushcomplete(Channel.State state);

    /// @param state channel state
    error IBCChannelUpgradeNotFlushcomplete(Channel.State state);

    /// @param clientId client identifier
    /// @param path key path
    /// @param value key value
    /// @param proof proof of membership
    /// @param height proof height
    error IBCChannelUpgradeFailedVerifyMembership(
        string clientId, string path, bytes value, bytes proof, Height.Data height
    );

    error IBCChannelUpgradeTimeoutHeightNotReached();

    error IBCChannelUpgradeTimeoutTimestampNotReached();

    error IBCChannelUpgradeInvalidTimeout();

    error IBCChannelUpgradeCounterpartyAlreadyFlushCompleted();

    error IBCChannelUpgradeCounterpartyAlreadyUpgraded();

    error IBCChannelUpgradeTimeoutUnallowedState();

    error IBCChannelUpgradeOldErrorReceiptSequence();

    error IBCChannelUpgradeInvalidErrorReceiptSequence();

    error IBCChannelUpgradeOldCounterpartyUpgradeSequence();

    error IBCChannelUpgradeUnsupportedOrdering(Channel.Order ordering);
}
