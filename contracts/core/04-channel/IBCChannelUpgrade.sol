// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, ChannelCounterparty, Upgrade, UpgradeFields, ErrorReceipt, Timeout} from "../../proto/Channel.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCConnectionLib} from "../03-connection/IBCConnectionLib.sol";
import {IIBCConnectionErrors} from "../03-connection/IIBCConnectionErrors.sol";
import {IIBCChannelUpgrade} from "../04-channel/IIBCChannelUpgrade.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCModuleUpgrade} from "../26-router/IIBCModuleUpgrade.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IBCChannelLib} from "./IBCChannelLib.sol";
import {IIBCChannelUpgradeErrors} from "./IIBCChannelUpgradeErrors.sol";

/// @dev IBCChannelUpgrade is a module that implements the IBC channel upgrade protocol.
/// https://github.com/cosmos/ibc/blob/main/spec/core/ics-004-channel-and-packet-semantics/UPGRADES.md
contract IBCChannelUpgrade is IBCModuleManager, IIBCChannelUpgrade, IIBCChannelUpgradeErrors, IIBCConnectionErrors {
    /**
     * @dev See {IIBCChannelUpgrade-channelUpgradeInit}
     */
    function channelUpgradeInit(MsgChannelUpgradeInit calldata msg_) external override returns (uint64) {
        if (!isAuthorizedUpgrader(msg_.portId, msg_.channelId, _msgSender())) {
            revert IBCChannelUpgradeUnauthorizedChannelUpgrader(_msgSender());
        }

        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        Upgrade.Data storage upgrade = upgrades[msg_.portId][msg_.channelId];
        if (upgrade.fields.ordering != Channel.Order.ORDER_NONE_UNSPECIFIED) {
            delete upgrades[msg_.portId][msg_.channelId];
            revertCounterpartyUpgrade(msg_.portId, msg_.channelId);
            // NOTE: we do not delete the upgrade commitment here since the new upgrade will overwrite the old one
            writeErrorReceipt(msg_.portId, msg_.channelId, channel.upgrade_sequence, UpgradeHandshakeError.Overwritten);
        }

        validateInitUpgradeHandshake(channel, msg_.proposedUpgradeFields);
        channel.upgrade_sequence++;
        updateChannelCommitment(msg_.portId, msg_.channelId, channel);

        upgrade.fields = msg_.proposedUpgradeFields;
        upgrade.fields.version = lookupUpgradableModuleByPort(msg_.portId).onChanUpgradeInit(
            msg_.portId, msg_.channelId, channel.upgrade_sequence, msg_.proposedUpgradeFields
        );
        updateUpgradeCommitment(msg_.portId, msg_.channelId, upgrade);

        emit ChannelUpgradeInit(msg_.portId, msg_.channelId, channel.upgrade_sequence, msg_.proposedUpgradeFields);

        return channel.upgrade_sequence;
    }

    /**
     * @dev See {IIBCChannelUpgrade-channelUpgradeTry}
     */
    function channelUpgradeTry(MsgChannelUpgradeTry calldata msg_) external override returns (bool, uint64) {
        // current channel must be OPEN (i.e. not in FLUSHING)
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }

        verifyChannelAndUpgradeMembership(
            channel,
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
            msg_.counterpartyUpgradeSequence,
            // NOTE: `timeout` and `nextSequenceSend` will be filled when the counterparty moves to FLUSHING
            Upgrade.Data({
                fields: msg_.counterpartyUpgradeFields,
                timeout: Timeout.Data({height: Height.Data({revision_height: 0, revision_number: 0}), timestamp: 0}),
                next_sequence_send: 0
            }),
            Channel.State.STATE_OPEN,
            msg_.proofs
        );

        Upgrade.Data storage existingUpgrade = upgrades[msg_.portId][msg_.channelId];
        uint64 expectedUpgradeSequence;
        if (existingUpgrade.fields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED) {
            // NON CROSSING HELLO CASE
            expectedUpgradeSequence = channel.upgrade_sequence + 1;
        } else {
            // CROSSING HELLO CASE
            expectedUpgradeSequence = channel.upgrade_sequence;
        }
        if (msg_.counterpartyUpgradeSequence < expectedUpgradeSequence) {
            // NON CROSSING HELLO CASE:
            // if the counterparty sequence is less than the current sequence,
            // then either the counterparty chain is out-of-sync or the message
            // is out-of-sync and we write an error receipt with our sequence
            // so that the counterparty can abort their attempt and resync with our sequence.
            // When the next upgrade attempt is initiated, both sides will move to a fresh
            // never-before-seen sequence number
            // CROSSING HELLO CASE:
            // if the counterparty sequence is less than the current sequence,
            // then either the counterparty chain is out-of-sync or the message
            // is out-of-sync and we write an error receipt with our sequence - 1
            // so that the counterparty can update their sequence as well.
            // This will cause the outdated counterparty to upgrade the sequence
            // and abort their out-of-sync upgrade without aborting our own since
            // the error receipt sequence is lower than ours and higher than the counterparty.
            unchecked {
                writeErrorReceipt(
                    msg_.portId, msg_.channelId, expectedUpgradeSequence - 1, UpgradeHandshakeError.OutOfSync
                );
            }
            return (false, 0);
        }

        UpgradeFields.Data memory upgradeFields;
        if (existingUpgrade.fields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED) {
            upgradeFields = UpgradeFields.Data({
                ordering: msg_.counterpartyUpgradeFields.ordering,
                version: msg_.counterpartyUpgradeFields.version,
                connection_hops: msg_.proposedConnectionHops
            });
            validateInitUpgradeHandshake(channel, upgradeFields);
            channel.upgrade_sequence = expectedUpgradeSequence;
            // nextSequenceSend and timeout will be filled when we move to FLUSHING
            existingUpgrade.fields = upgradeFields;
        } else {
            if (msg_.proposedConnectionHops.length != 0) {
                revert IBCChannelUpgradeTryProposalMustEmptyIfExist();
            }
            upgradeFields = existingUpgrade.fields;
        }

        if (!isCompatibleUpgradeFields(upgradeFields, msg_.counterpartyUpgradeFields)) {
            revert IBCChannelUpgradeIncompatibleProposal();
        }

        // if the counterparty sequence is greater than the current sequence,
        // we fast forward to the counterparty sequence so that both channel
        // ends are using the same sequence for the current upgrade.
        // initUpgradeHandshake will increment the sequence so after that call
        // both sides will have the same upgradeSequence
        if (msg_.counterpartyUpgradeSequence > channel.upgrade_sequence) {
            channel.upgrade_sequence = msg_.counterpartyUpgradeSequence;
        }

        // call startFlushUpgradeHandshake to move channel to FLUSHING, which will block
        // upgrade from progressing to OPEN until flush completes on both ends
        startFlushUpgradeHandshake(channel, existingUpgrade, msg_.portId, msg_.channelId);
        updateChannelCommitment(msg_.portId, msg_.channelId, channel);

        existingUpgrade.fields.version = lookupUpgradableModuleByPort(msg_.portId).onChanUpgradeTry(
            msg_.portId, msg_.channelId, channel.upgrade_sequence, upgradeFields
        );
        updateUpgradeCommitment(msg_.portId, msg_.channelId, existingUpgrade);

        emit ChannelUpgradeTry(
            msg_.portId,
            msg_.channelId,
            channel.upgrade_sequence,
            existingUpgrade.fields,
            existingUpgrade.timeout,
            existingUpgrade.next_sequence_send
        );

        return (true, channel.upgrade_sequence);
    }

    /**
     * @dev See {IIBCChannelUpgrade-channelUpgradeAck}
     */
    function channelUpgradeAck(MsgChannelUpgradeAck calldata msg_) external override returns (bool) {
        // current channel is OPEN or FLUSHING (crossing hellos)
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        if (channel.state != Channel.State.STATE_OPEN && channel.state != Channel.State.STATE_FLUSHING) {
            revert IBCChannelUpgradeNotOpenOrFlushing(channel.state);
        }

        verifyChannelAndUpgradeMembership(
            channel,
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
            channel.upgrade_sequence,
            msg_.counterpartyUpgrade,
            Channel.State.STATE_FLUSHING,
            msg_.proofs
        );

        Upgrade.Data storage existingUpgrade = upgrades[msg_.portId][msg_.channelId];
        if (existingUpgrade.fields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED) {
            revert IBCChannelUpgradeNoExistingUpgrade();
        }

        // optimistically accept version that TRY chain proposes and pass this to callback for confirmation.
        // in the crossing hello case, we do not modify version that our TRY call returned and instead
        // enforce that both TRY calls returned the same version
        if (channel.state == Channel.State.STATE_OPEN) {
            existingUpgrade.fields.version = msg_.counterpartyUpgrade.fields.version;
        }
        // if upgrades are not compatible by ACK step, then we restore the channel
        if (!isCompatibleUpgradeFields(existingUpgrade.fields, msg_.counterpartyUpgrade.fields)) {
            restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.IncompatibleProposal);
            return false;
        }

        if (channel.state == Channel.State.STATE_OPEN) {
            // prove counterparty and move our own state to flushing
            // if we are already at flushing, then no state changes occur
            // upgrade is blocked on this channelEnd from progressing until flush completes on its end
            startFlushUpgradeHandshake(channel, existingUpgrade, msg_.portId, msg_.channelId);
            // NOTE: `upgrade` and `channel` are updated only when channel.state is OPEN
            updateChannelCommitment(msg_.portId, msg_.channelId, channel);
            updateUpgradeCommitment(msg_.portId, msg_.channelId, existingUpgrade);
        }

        // counterparty-specified timeout must not have exceeded
        // if it has, then restore the channel and abort upgrade handshake
        Timeout.Data calldata timeout = msg_.counterpartyUpgrade.timeout;
        if (
            (timeout.height.revision_height != 0 && block.number >= timeout.height.revision_height)
                || (timeout.timestamp != 0 && block.timestamp >= timeout.timestamp)
        ) {
            restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.Timeout);
            return false;
        }

        // if there are no in-flight packets on our end, we can automatically go to FLUSHCOMPLETE
        if (canTransitionToFlushComplete(channel.ordering, msg_.portId, msg_.channelId, channel.upgrade_sequence)) {
            channel.state = Channel.State.STATE_FLUSHCOMPLETE;
            updateChannelCommitment(msg_.portId, msg_.channelId, channel);
        }
        setCounterpartyUpgrade(msg_.portId, msg_.channelId, msg_.counterpartyUpgrade);

        // call modules onChanUpgradeAck callback
        // module can error on counterparty version
        // ACK should not change state to the new parameters yet
        // as that will happen on the onChanUpgradeOpen callback
        try lookupUpgradableModuleByPort(msg_.portId).onChanUpgradeAck(
            msg_.portId, msg_.channelId, channel.upgrade_sequence, msg_.counterpartyUpgrade.fields.version
        ) {} catch {
            restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.AckCallbackFailed);
            return false;
        }

        emit ChannelUpgradeAck(
            msg_.portId,
            msg_.channelId,
            channel.upgrade_sequence,
            channel.state,
            existingUpgrade.fields,
            existingUpgrade.timeout,
            existingUpgrade.next_sequence_send
        );

        return true;
    }

    /**
     * @dev See {IIBCChannelUpgrade-channelUpgradeConfirm}
     */
    function channelUpgradeConfirm(MsgChannelUpgradeConfirm calldata msg_) external override returns (bool) {
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        // current channel is in FLUSHING
        if (channel.state != Channel.State.STATE_FLUSHING) {
            revert IBCChannelUpgradeNotFlushing(channel.state);
        }

        // counterparty channel is either FLUSHING or FLUSHCOMPLETE
        if (
            msg_.counterpartyChannelState != Channel.State.STATE_FLUSHING
                && msg_.counterpartyChannelState != Channel.State.STATE_FLUSHCOMPLETE
        ) {
            revert IBCChannelUpgradeCounterpartyNotFlushingOrFlushcomplete(msg_.counterpartyChannelState);
        }

        verifyChannelAndUpgradeMembership(
            channel,
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
            channel.upgrade_sequence,
            msg_.counterpartyUpgrade,
            msg_.counterpartyChannelState,
            msg_.proofs
        );

        // counterparty-specified timeout must not have exceeded
        // if it has, then restore the channel and abort upgrade handshake
        Timeout.Data calldata timeout = msg_.counterpartyUpgrade.timeout;
        if (
            (timeout.height.revision_height != 0 && block.number >= timeout.height.revision_height)
                || (timeout.timestamp != 0 && block.timestamp >= timeout.timestamp)
        ) {
            restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.Timeout);
            return false;
        }

        // if there are no in-flight packets on our end, we can automatically go to FLUSHCOMPLETE
        if (canTransitionToFlushComplete(channel.ordering, msg_.portId, msg_.channelId, channel.upgrade_sequence)) {
            channel.state = Channel.State.STATE_FLUSHCOMPLETE;
            updateChannelCommitment(msg_.portId, msg_.channelId, channel);
        }
        setCounterpartyUpgrade(msg_.portId, msg_.channelId, msg_.counterpartyUpgrade);

        // if both chains are already in flushcomplete we can move to OPEN
        if (
            channel.state == Channel.State.STATE_FLUSHCOMPLETE
                && msg_.counterpartyChannelState == Channel.State.STATE_FLUSHCOMPLETE
        ) {
            openUpgradeHandshake(msg_.portId, msg_.channelId);
            lookupUpgradableModuleByPort(msg_.portId).onChanUpgradeOpen(
                msg_.portId, msg_.channelId, channel.upgrade_sequence
            );
        }

        emit ChannelUpgradeConfirm(msg_.portId, msg_.channelId, channel.upgrade_sequence, channel.state);

        return true;
    }

    /**
     * @dev See {IIBCChannelUpgrade-channelUpgradeOpen}
     */
    function channelUpgradeOpen(MsgChannelUpgradeOpen calldata msg_) external override {
        // channel must have completed flushing
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        if (channel.state != Channel.State.STATE_FLUSHCOMPLETE) {
            revert IBCChannelUpgradeNotFlushcomplete(channel.state);
        }

        // get connection for proof verification
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        Channel.Data memory counterpartyChannel;
        // counterparty must be in OPEN or FLUSHCOMPLETE state
        if (msg_.counterpartyChannelState == Channel.State.STATE_OPEN) {
            Upgrade.Data storage upgrade = upgrades[msg_.portId][msg_.channelId];
            // get the counterparty's connection hops for the proposed upgrade connection
            ConnectionEnd.Data storage proposedConnection = connections[upgrade.fields.connection_hops[0]];
            // The counterparty upgrade sequence must be greater than or equal to
            // the channel upgrade sequence. It should normally be equivalent, but
            // in the unlikely case that a new upgrade is initiated after it reopens,
            // then the upgrade sequence will be greater than our upgrade sequence.
            if (msg_.counterpartyUpgradeSequence < channel.upgrade_sequence) {
                revert IBCChannelUpgradeOldCounterpartyUpgradeSequence();
            }
            counterpartyChannel = Channel.Data({
                state: Channel.State.STATE_OPEN,
                ordering: upgrade.fields.ordering,
                counterparty: ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
                connection_hops: IBCChannelLib.buildConnectionHops(proposedConnection.counterparty.connection_id),
                version: upgrade.fields.version,
                upgrade_sequence: msg_.counterpartyUpgradeSequence
            });
        } else if (msg_.counterpartyChannelState == Channel.State.STATE_FLUSHCOMPLETE) {
            // counterparty channel does not upgrade to new parameters yet
            counterpartyChannel = Channel.Data({
                state: Channel.State.STATE_FLUSHCOMPLETE,
                ordering: channel.ordering,
                counterparty: ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
                connection_hops: IBCChannelLib.buildConnectionHops(connection.counterparty.connection_id),
                version: channel.version,
                upgrade_sequence: channel.upgrade_sequence
            });
        } else {
            revert IBCChannelUpgradeCounterpartyNotOpenOrFlushcomplete(msg_.counterpartyChannelState);
        }

        verifyChannelState(
            connection,
            msg_.proofHeight,
            msg_.proofChannel,
            channel.counterparty.port_id,
            channel.counterparty.channel_id,
            counterpartyChannel
        );

        // move channel to OPEN and adopt upgrade parameters
        openUpgradeHandshake(msg_.portId, msg_.channelId);
        // open callback must not return error since counterparty successfully upgraded
        // make application state changes based on new channel parameters
        lookupUpgradableModuleByPort(msg_.portId).onChanUpgradeOpen(
            msg_.portId, msg_.channelId, channel.upgrade_sequence
        );

        emit ChannelUpgradeOpen(msg_.portId, msg_.channelId, channel.upgrade_sequence);
    }

    /**
     * @dev See {IIBCChannelUpgrade-cancelChannelUpgrade}
     */
    function cancelChannelUpgrade(MsgCancelChannelUpgrade calldata msg_) external override {
        // current channel has an upgrade stored
        Upgrade.Data storage upgrade = upgrades[msg_.portId][msg_.channelId];
        if (upgrade.fields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED) {
            revert IBCChannelUpgradeNoExistingUpgrade();
        }
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        // if the msgSender is authorized to make and cancel upgrades AND
        // the current channel has not already reached FLUSHCOMPLETE,
        // then we can restore immediately without any additional checks
        // otherwise, we can only cancel if the counterparty wrote an
        // error receipt during the upgrade handshake
        if (
            isAuthorizedUpgrader(msg_.portId, msg_.channelId, _msgSender())
                && channel.state != Channel.State.STATE_FLUSHCOMPLETE
        ) {
            return restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.Cancel);
        }
        if (bytes(msg_.errorReceipt.message).length == 0) {
            revert IBCChannelUpgradeErrorReceiptEmpty();
        }

        if (channel.state == Channel.State.STATE_FLUSHCOMPLETE) {
            // if the channel state is in FLUSHCOMPLETE, it can **only** be aborted if there
            // is an error receipt with the exact same sequence. This ensures that the counterparty
            // did not successfully upgrade and then cancel at a new upgrade to abort our own end,
            // leading to both channel ends being OPEN with different parameters
            if (msg_.errorReceipt.sequence != channel.upgrade_sequence) {
                revert IBCChannelUpgradeInvalidErrorReceiptSequence();
            }
        } else {
            // If counterparty sequence is less than the current sequence,
            // abort transaction since this error receipt is from a previous upgrade
            if (msg_.errorReceipt.sequence < channel.upgrade_sequence) {
                revert IBCChannelUpgradeOldErrorReceiptSequence();
            }
            // fastforward channel sequence to higher sequence so that we can start
            // new handshake on a fresh sequence
            channel.upgrade_sequence = msg_.errorReceipt.sequence;
        }

        // verify that the provided error receipt is written to the upgradeError path with the counterparty sequence
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        verifyMembership(
            connection,
            msg_.proofHeight,
            msg_.proofUpgradeError,
            IBCCommitment.channelUpgradeErrorPath(channel.counterparty.port_id, channel.counterparty.channel_id),
            ErrorReceipt.encode(msg_.errorReceipt)
        );

        // cancel upgrade and write error receipt
        restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.Cancel);
    }

    /**
     * @dev See {IIBCChannelUpgrade-timeoutChannelUpgrade}
     */
    function timeoutChannelUpgrade(MsgTimeoutChannelUpgrade calldata msg_) external override {
        // current channel must have an upgrade that is FLUSHING or FLUSHCOMPLETE
        Upgrade.Data storage upgrade = upgrades[msg_.portId][msg_.channelId];
        if (upgrade.fields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED) {
            revert IBCChannelUpgradeNoExistingUpgrade();
        }
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        if (channel.state != Channel.State.STATE_FLUSHING && channel.state != Channel.State.STATE_FLUSHCOMPLETE) {
            revert IBCChannelUpgradeTimeoutUnallowedState();
        }

        // proof must be from a height after timeout has elapsed.
        // Either timeoutHeight or timeoutTimestamp must be defined.
        // if timeoutHeight is defined and proof is from before
        // timeout height then abort transaction
        if (
            upgrade.timeout.height.revision_height != 0
                && msg_.proofHeight.revision_height < upgrade.timeout.height.revision_height
        ) {
            revert IBCChannelUpgradeTimeoutHeightNotReached();
        }
        // if timeoutTimestamp is defined then the consensus time
        // from proof height must be greater than timeout timestamp
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        if (
            upgrade.timeout.timestamp != 0
                && ILightClient(clientImpls[connection.client_id]).getTimestampAtHeight(
                    connection.client_id, msg_.proofHeight
                ) < upgrade.timeout.timestamp
        ) {
            revert IBCChannelUpgradeTimeoutTimestampNotReached();
        }

        // counterparty channel must be proved to not have completed flushing after timeout has passed
        if (msg_.counterpartyChannel.state == Channel.State.STATE_FLUSHCOMPLETE) {
            revert IBCChannelUpgradeCounterpartyAlreadyFlushCompleted();
        }

        // if counterparty channel state is OPEN, we should abort the tx
        // only if the counterparty has successfully completed upgrade
        if (msg_.counterpartyChannel.state == Channel.State.STATE_OPEN) {
            // counterparty should have upgraded to `upgrade` parameters
            ConnectionEnd.Data storage proposedConnection = connections[upgrade.fields.connection_hops[0]];
            // check that the channel did not upgrade successfully
            require(msg_.counterpartyChannel.connection_hops.length == 1);
            if (
                keccak256(bytes(upgrade.fields.version)) == keccak256(bytes(msg_.counterpartyChannel.version))
                    && upgrade.fields.ordering == msg_.counterpartyChannel.ordering
                    && keccak256(abi.encodePacked(proposedConnection.counterparty.connection_id))
                        == keccak256(abi.encodePacked(msg_.counterpartyChannel.connection_hops[0]))
            ) {
                revert IBCChannelUpgradeCounterpartyAlreadyUpgraded();
            }
        }

        require(msg_.counterpartyChannel.upgrade_sequence >= channel.upgrade_sequence);
        verifyChannelState(
            connection,
            msg_.proofHeight,
            msg_.proofChannel,
            channel.counterparty.port_id,
            channel.counterparty.channel_id,
            msg_.counterpartyChannel
        );
        restoreChannel(msg_.portId, msg_.channelId, UpgradeHandshakeError.Timeout);
    }

    // --------------------- Internal Functions --------------------- //

    /**
     * @dev validateInitUpgradeHandshake will verify that the channel is in the
     * correct precondition to call the initUpgradeHandshake protocol.
     * It will verify the new upgrade field parameters.
     */
    function validateInitUpgradeHandshake(Channel.Data storage channel, UpgradeFields.Data memory proposedUpgradeFields)
        internal
        view
    {
        // current channel must be OPEN
        // If channel already has an upgrade but isn't in FLUSHING,
        // then this will override the previous upgrade attempt
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }

        // proposedUpgradeFields must be valid
        if (
            bytes(proposedUpgradeFields.version).length == 0
                || proposedUpgradeFields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED
                || proposedUpgradeFields.connection_hops.length != 1
        ) {
            revert IBCChannelUpgradeInvalidUpgradeFields();
        }

        // proposedConnection must exist and be in OPEN state for
        // channel upgrade to be accepted
        ConnectionEnd.Data storage proposedConnection = connections[proposedUpgradeFields.connection_hops[0]];
        if (proposedConnection.state != ConnectionEnd.State.STATE_OPEN) {
            revert IBCConnectionUnexpectedConnectionState(proposedConnection.state);
        }

        // new order must be supported by the new connection
        if (
            !IBCConnectionLib.isSupported(
                proposedConnection.versions, IBCChannelLib.toString(proposedUpgradeFields.ordering)
            )
        ) {
            revert IBCConnectionIBCVersionNotSupported();
        }

        // there exists at least one valid proposed change to the existing channel fields
        if (
            channel.ordering == proposedUpgradeFields.ordering
                && keccak256(bytes(channel.version)) == keccak256(bytes(proposedUpgradeFields.version))
                && keccak256(abi.encodePacked(proposedUpgradeFields.connection_hops[0]))
                    == keccak256(abi.encodePacked(channel.connection_hops[0]))
        ) {
            revert IBCChannelUpgradeNoChanges();
        }
    }

    /**
     * @dev startFlushUpgradeHandshake will verify that the channel
     * is in a valid precondition for calling the startFlushUpgradeHandshake.
     * it will set the channel to flushing state.
     * it will store the nextSequenceSend and upgrade timeout in the upgrade state.
     */
    function startFlushUpgradeHandshake(
        Channel.Data storage channel,
        Upgrade.Data storage upgrade,
        string calldata portId,
        string calldata channelId
    ) internal {
        Timeout.Data memory upgradeTimeout = getUpgradeTimeout(portId, channelId);
        if (
            !(
                upgradeTimeout.height.revision_number > 0 || upgradeTimeout.height.revision_height > 0
                    || upgradeTimeout.timestamp > 0
            )
        ) {
            revert IBCChannelUpgradeInvalidTimeout();
        }
        channel.state = Channel.State.STATE_FLUSHING;
        upgrade.timeout = upgradeTimeout;
        upgrade.next_sequence_send = nextSequenceSends[portId][channelId];
    }

    /**
     * @dev restoreChannel will restore the channel state to its pre-upgrade state
     * and delete upgrade auxiliary state so that upgrade is aborted.
     * it writes an error receipt to state so counterparty can restore as well.
     * NOTE: this function signature may be modified by implementors to take a custom error
     */
    // reason
    function restoreChannel(string calldata portId, string calldata channelId, UpgradeHandshakeError err) internal {
        Channel.Data storage channel = channels[portId][channelId];
        channel.state = Channel.State.STATE_OPEN;

        delete upgrades[portId][channelId];
        deleteUpgradeCommitment(portId, channelId);
        revertCounterpartyUpgrade(portId, channelId);

        updateChannelCommitment(portId, channelId, channel);
        writeErrorReceipt(portId, channelId, channel.upgrade_sequence, err);
    }

    /**
     * @dev openUpgradeHandshake will switch the channel fields
     * over to the agreed upon upgrade fields.
     * it will reset the channel state to OPEN.
     * it will delete auxiliary upgrade state.
     * caller must do all relevant checks before calling this function.
     */
    function openUpgradeHandshake(string calldata portId, string calldata channelId) internal {
        Channel.Data storage channel = channels[portId][channelId];
        Upgrade.Data storage upgrade = upgrades[portId][channelId];

        // In ibc-solidity, we need to set the nextSequenceRecv and nextSequenceAck appropriately for each upgrade
        if (upgrade.fields.ordering == Channel.Order.ORDER_ORDERED) {
            // set nextSequenceRecv to the counterparty nextSequenceSend since all packets were flushed
            nextSequenceRecvs[portId][channelId] = recvStartSequences[portId][channelId].sequence;
            // set nextSequenceAck to our own nextSequenceSend since all packets were flushed
            nextSequenceAcks[portId][channelId] = nextSequenceSends[portId][channelId];
        } else if (upgrade.fields.ordering == Channel.Order.ORDER_UNORDERED) {
            // reset recv and ack sequences to 1 for UNORDERED channel
            nextSequenceRecvs[portId][channelId] = 1;
            nextSequenceAcks[portId][channelId] = 1;
        } else {
            revert IBCChannelUpgradeUnsupportedOrdering(upgrade.fields.ordering);
        }
        commitments[IBCCommitment.nextSequenceRecvCommitmentKey(portId, channelId)] =
            keccak256(abi.encodePacked(nextSequenceRecvs[portId][channelId]));
        ackStartSequences[portId][channelId] = nextSequenceSends[portId][channelId];

        // switch channel fields to upgrade fields
        // and set channel state to OPEN
        channel.ordering = upgrade.fields.ordering;
        channel.version = upgrade.fields.version;
        channel.connection_hops = upgrade.fields.connection_hops;
        channel.state = Channel.State.STATE_OPEN;

        // IMPLEMENTATION DETAIL: Implementations may choose to prune stale acknowledgements and receipts at this stage
        // Since flushing has completed, any acknowledgement or receipt written before the chain went into flushing has
        // already been processed by the counterparty and can be removed.
        // Implementations may do this pruning work over multiple blocks for gas reasons. In this case, they should be sure
        // to only prune stale acknowledgements/receipts and not new ones that have been written after the channel has reopened.
        // Implementations may use the counterparty NextSequenceSend as a way to determine which acknowledgement/receipts
        // were already processed by counterparty when flushing completed

        // delete auxiliary state
        delete upgrades[portId][channelId];
        recvStartSequences[portId][channelId].prevSequence = 0;

        updateChannelCommitment(portId, channelId, channel);
        deleteUpgradeCommitment(portId, channelId);
    }

    function isCompatibleUpgradeFields(
        UpgradeFields.Data memory proposedUpgradeFields,
        UpgradeFields.Data calldata counterpartyUpgradeFields
    ) internal view returns (bool) {
        if (proposedUpgradeFields.ordering != counterpartyUpgradeFields.ordering) {
            return false;
        }
        if (keccak256(bytes(proposedUpgradeFields.version)) != keccak256(bytes(counterpartyUpgradeFields.version))) {
            return false;
        }

        // connectionHops can change in a channel upgrade, however both sides must
        // still be each other's counterparty. Since connection hops may be provided
        // by relayer, we will abort to avoid changing state based on relayer-provided value
        // Note: If the proposed connection came from an existing upgrade, then the
        // off-chain authority is responsible for replacing one side's upgrade fields
        // to be compatible so that the upgrade handshake can proceed

        ConnectionEnd.Data storage proposedConnection = connections[proposedUpgradeFields.connection_hops[0]];
        if (proposedConnection.state != ConnectionEnd.State.STATE_OPEN) {
            return false;
        }

        return keccak256(bytes(counterpartyUpgradeFields.connection_hops[0]))
            == keccak256(bytes(proposedConnection.counterparty.connection_id));
    }

    function lookupUpgradableModuleByPort(string calldata portId) internal view returns (IIBCModuleUpgrade) {
        return IIBCModuleUpgrade(address(lookupModuleByPort(portId)));
    }

    function isAuthorizedUpgrader(string calldata portId, string calldata channelId, address msgSender)
        internal
        view
        returns (bool)
    {
        return lookupUpgradableModuleByPort(portId).isAuthorizedUpgrader(portId, channelId, msgSender);
    }

    function getUpgradeTimeout(string calldata portId, string calldata channelId)
        internal
        view
        returns (Timeout.Data memory)
    {
        return lookupUpgradableModuleByPort(portId).getUpgradeTimeout(portId, channelId);
    }

    function canTransitionToFlushComplete(
        Channel.Order ordering,
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence
    ) internal view virtual returns (bool) {
        if (ordering == Channel.Order.ORDER_ORDERED) {
            if (nextSequenceSends[portId][channelId] == nextSequenceAcks[portId][channelId]) {
                return true;
            }
        }
        return lookupUpgradableModuleByPort(portId).canTransitionToFlushComplete(
            portId, channelId, upgradeSequence, _msgSender()
        );
    }

    // --------------------- Private Functions --------------------- //

    function writeErrorReceipt(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        UpgradeHandshakeError err
    ) private {
        if (latestErrorReceiptSequences[portId][channelId] >= upgradeSequence) {
            revert IBCChannelUpgradeWriteOldErrorReceiptSequence(
                latestErrorReceiptSequences[portId][channelId], upgradeSequence
            );
        }
        latestErrorReceiptSequences[portId][channelId] = upgradeSequence;
        string memory message = toString(err);
        commitments[IBCCommitment.channelUpgradeErrorCommitmentKey(portId, channelId)] =
            keccak256(ErrorReceipt.encode(ErrorReceipt.Data({sequence: upgradeSequence, message: message})));
        emit WriteErrorReceipt(portId, channelId, upgradeSequence, message);
    }

    function updateChannelCommitment(string calldata portId, string calldata channelId, Channel.Data storage channel)
        private
    {
        commitments[IBCCommitment.channelCommitmentKey(portId, channelId)] = keccak256(Channel.encode(channel));
    }

    function updateUpgradeCommitment(string calldata portId, string calldata channelId, Upgrade.Data storage upgrade)
        private
    {
        updateUpgradeCommitment(portId, channelId, keccak256(Upgrade.encode(upgrade)));
    }

    function deleteUpgradeCommitment(string calldata portId, string calldata channelId) private {
        updateUpgradeCommitment(portId, channelId, bytes32(0));
    }

    function updateUpgradeCommitment(string calldata portId, string calldata channelId, bytes32 commitment) private {
        commitments[IBCCommitment.channelUpgradeCommitmentKey(portId, channelId)] = commitment;
    }

    function setCounterpartyUpgrade(string calldata portId, string calldata channelId, Upgrade.Data calldata upgrade)
        private
    {
        if (recvStartSequences[portId][channelId].prevSequence != 0) {
            revertCounterpartyUpgrade(portId, channelId);
        }
        recvStartSequences[portId][channelId].prevSequence = recvStartSequences[portId][channelId].sequence;
        recvStartSequences[portId][channelId].sequence = upgrade.next_sequence_send;
    }

    function revertCounterpartyUpgrade(string calldata portId, string calldata channelId) private {
        uint64 prevRecvStartSequence = recvStartSequences[portId][channelId].prevSequence;
        if (prevRecvStartSequence == 0) {
            return;
        }
        recvStartSequences[portId][channelId].prevSequence = 0;
        recvStartSequences[portId][channelId].sequence = prevRecvStartSequence;
    }

    function verifyMembership(
        ConnectionEnd.Data storage connection,
        Height.Data calldata height,
        bytes calldata proof,
        bytes memory path,
        bytes memory value
    ) private {
        if (
            ILightClient(clientImpls[connection.client_id]).verifyMembership(
                connection.client_id, height, 0, 0, proof, connection.counterparty.prefix.key_prefix, path, value
            )
        ) {
            return;
        }
        revert IBCChannelUpgradeFailedVerifyMembership(connection.client_id, string(path), value, proof, height);
    }

    function verifyChannelState(
        ConnectionEnd.Data storage connection,
        Height.Data calldata height,
        bytes calldata proof,
        string memory portId,
        string memory channelId,
        Channel.Data memory channel
    ) private {
        return verifyMembership(
            connection, height, proof, IBCCommitment.channelPath(portId, channelId), Channel.encode(channel)
        );
    }

    function verifyChannelAndUpgradeMembership(
        Channel.Data storage channel,
        ChannelCounterparty.Data memory counterparty,
        uint64 counterpartyUpgradeSequence,
        Upgrade.Data memory counterpartyUpgrade,
        Channel.State counterpartyChannelState,
        ChannelUpgradeProofs calldata proofs
    ) private {
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        {
            Channel.Data memory counterpartyChannel = Channel.Data({
                state: counterpartyChannelState,
                ordering: channel.ordering,
                counterparty: counterparty,
                connection_hops: IBCChannelLib.buildConnectionHops(connection.counterparty.connection_id),
                version: channel.version,
                upgrade_sequence: counterpartyUpgradeSequence
            });
            verifyChannelState(
                connection,
                proofs.proofHeight,
                proofs.proofChannel,
                channel.counterparty.port_id,
                channel.counterparty.channel_id,
                counterpartyChannel
            );
        }

        verifyMembership(
            connection,
            proofs.proofHeight,
            proofs.proofUpgrade,
            IBCCommitment.channelUpgradePath(channel.counterparty.port_id, channel.counterparty.channel_id),
            Upgrade.encode(counterpartyUpgrade)
        );
    }

    function toString(UpgradeHandshakeError err) internal pure returns (string memory) {
        bytes memory result = new bytes(1);
        unchecked {
            result[0] = bytes1(uint8(err) + 48);
        }
        return string(result);
    }
}
