// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Channel, UpgradeFields, Timeout} from "../../../../contracts/proto/Channel.sol";
import {IIBCHandler} from "../../../../contracts/core/25-handler/IIBCHandler.sol";
import {IIBCModuleUpgrade} from "../../../../contracts/core/26-router/IIBCModuleUpgrade.sol";
import {AppBase} from "../../../../contracts/apps/commons/IBCAppBase.sol";

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

abstract contract IBCChannelUpgradableModuleBase is
    AppBase,
    IIBCModuleUpgrade,
    IIBCChannelUpgradableModule,
    IIBCChannelUpgradableModuleErrors
{
    // ------------------- Storage ------------------- //

    /**
     * @dev Proposed upgrades for each channel
     */
    mapping(string portId => mapping(string channelId => UpgradeProposal)) internal upgradeProposals;
    /**
     * @dev Allowed transitions for each upgrade sequence
     */
    mapping(string portId => mapping(string channelId => mapping(uint64 upgradeSequence => AllowedTransition))) internal
        allowedTransitions;

    // ------------------- Modifiers ------------------- //

    /**
     * @dev Throws if the sender is not an authorized upgrader
     * @param portId Port identifier
     * @param channelId Channel identifier
     */
    modifier onlyAuthorizedUpgrader(string calldata portId, string calldata channelId) {
        if (!_isAuthorizedUpgrader(portId, channelId, _msgSender())) {
            revert IBCChannelUpgradableModuleUnauthorizedUpgrader();
        }
        _;
    }

    // ------------------- Public Functions ------------------- //

    /**
     * @dev See {IIBCChannelUpgradableModule-getUpgradeProposal}
     */
    function getUpgradeProposal(string calldata portId, string calldata channelId)
        public
        view
        virtual
        override
        returns (UpgradeProposal memory)
    {
        return upgradeProposals[portId][channelId];
    }

    /**
     * @dev See {IIBCChannelUpgradableModule-proposeUpgrade}
     */
    function proposeUpgrade(
        string calldata portId,
        string calldata channelId,
        UpgradeFields.Data calldata upgradeFields,
        Timeout.Data calldata timeout
    ) public virtual override onlyAuthorizedUpgrader(portId, channelId) {
        if (timeout.height.revision_number == 0 && timeout.height.revision_height == 0 && timeout.timestamp == 0) {
            revert IBCChannelUpgradableModuleInvalidTimeout();
        }
        if (upgradeFields.ordering == Channel.Order.ORDER_NONE_UNSPECIFIED || upgradeFields.connection_hops.length == 0)
        {
            revert IBCChannelUpgradableModuleInvalidConnectionHops();
        }
        (Channel.Data memory channel, bool found) = IIBCHandler(ibcAddress()).getChannel(portId, channelId);
        if (!found) {
            revert IBCChannelUpgradableModuleChannelNotFound();
        }
        UpgradeProposal storage upgrade = upgradeProposals[portId][channelId];
        if (upgrade.fields.connection_hops.length != 0) {
            // re-proposal is allowed as long as it does not transition to FLUSHING state yet
            if (channel.state != Channel.State.STATE_OPEN) {
                revert IBCChannelUpgradableModuleCannotOverwriteUpgrade();
            }
        }
        upgrade.fields = upgradeFields;
        upgrade.timeout = timeout;
    }

    /**
     * @dev See {IIBCChannelUpgradableModule-allowTransitionToFlushComplete}
     */
    function allowTransitionToFlushComplete(string calldata portId, string calldata channelId, uint64 upgradeSequence)
        public
        virtual
        override
        onlyAuthorizedUpgrader(portId, channelId)
    {
        UpgradeProposal storage upgrade = upgradeProposals[portId][channelId];
        if (upgrade.fields.connection_hops.length == 0) {
            revert IBCChannelUpgradableModuleUpgradeNotFound();
        }
        (, bool found) = IIBCHandler(ibcAddress()).getChannelUpgrade(portId, channelId);
        if (!found) {
            revert IBCChannelUpgradableModuleUpgradeNotFound();
        }
        (Channel.Data memory channel,) = IIBCHandler(ibcAddress()).getChannel(portId, channelId);
        if (channel.state != Channel.State.STATE_FLUSHING) {
            revert IBCChannelUpgradableModuleChannelNotFlushingState(channel.state);
        }
        if (channel.upgrade_sequence != upgradeSequence) {
            revert IBCChannelUpgradableModuleSequenceMismatch(channel.upgrade_sequence);
        }
        allowedTransitions[portId][channelId][upgradeSequence].flushComplete = true;
    }

    /**
     * @dev See {IIBCChannelUpgradableModule-removeUpgradeProposal}
     */
    function removeUpgradeProposal(string calldata portId, string calldata channelId)
        public
        virtual
        onlyAuthorizedUpgrader(portId, channelId)
    {
        _removeUpgradeProposal(portId, channelId);
    }

    // ------------------- IIBCModuleUpgrade ------------------- //

    /**
     * @dev See {IIBCModuleUpgrade-isAuthorizedUpgrader}
     */
    function isAuthorizedUpgrader(string calldata portId, string calldata channelId, address msgSender)
        public
        view
        virtual
        override
        returns (bool)
    {
        return _isAuthorizedUpgrader(portId, channelId, msgSender);
    }

    /**
     * @dev See {IIBCModuleUpgrade-canTransitionToFlushComplete}
     */
    function canTransitionToFlushComplete(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        address
    ) public view virtual override returns (bool) {
        return allowedTransitions[portId][channelId][upgradeSequence].flushComplete;
    }

    /**
     * @dev See {IIBCModuleUpgrade-getUpgradeTimeout}
     */
    function getUpgradeTimeout(string calldata portId, string calldata channelId)
        public
        view
        virtual
        override
        returns (Timeout.Data memory)
    {
        if (upgradeProposals[portId][channelId].fields.connection_hops.length == 0) {
            revert IBCChannelUpgradableModuleUpgradeNotFound();
        }
        return upgradeProposals[portId][channelId].timeout;
    }

    /**
     * @dev See {IIBCModuleUpgrade-onChanUpgradeInit}
     */
    function onChanUpgradeInit(
        string calldata portId,
        string calldata channelId,
        uint64,
        UpgradeFields.Data calldata proposedUpgradeFields
    ) public view virtual override onlyIBC returns (string calldata version) {
        UpgradeProposal storage upgrade = upgradeProposals[portId][channelId];
        if (upgrade.fields.connection_hops.length == 0) {
            revert IBCChannelUpgradableModuleUpgradeNotFound();
        }
        if (!equals(upgrade.fields, proposedUpgradeFields)) {
            revert IBCChannelUpgradableModuleInvalidUpgrade();
        }
        return proposedUpgradeFields.version;
    }

    /**
     * @dev See {IIBCModuleUpgrade-onChanUpgradeTry}
     */
    function onChanUpgradeTry(
        string calldata portId,
        string calldata channelId,
        uint64,
        UpgradeFields.Data calldata proposedUpgradeFields
    ) public view virtual override onlyIBC returns (string calldata version) {
        UpgradeProposal storage upgrade = upgradeProposals[portId][channelId];
        if (upgrade.fields.connection_hops.length == 0) {
            revert IBCChannelUpgradableModuleUpgradeNotFound();
        }
        if (!equals(upgrade.fields, proposedUpgradeFields)) {
            revert IBCChannelUpgradableModuleInvalidUpgrade();
        }
        return proposedUpgradeFields.version;
    }

    /**
     * @dev See {IIBCModuleUpgrade-onChanUpgradeAck}
     */
    function onChanUpgradeAck(string calldata, string calldata, uint64, string calldata counterpartyVersion)
        public
        view
        virtual
        override
        onlyIBC
    {}

    /**
     * @dev See {IIBCModuleUpgrade-onChanUpgradeOpen}
     */
    function onChanUpgradeOpen(string calldata portId, string calldata channelId, uint64 upgradeSequence)
        public
        virtual
        override
        onlyIBC
    {
        delete upgradeProposals[portId][channelId];
        delete allowedTransitions[portId][channelId][upgradeSequence];
    }

    /**
     * @dev See {IERC165-supportsInterface}
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return
            super.supportsInterface(interfaceId) ||
            interfaceId == type(IIBCModuleUpgrade).interfaceId ||
            interfaceId == type(IIBCChannelUpgradableModule).interfaceId;
    }

    // ------------------- Internal Functions ------------------- //

    /**
     * @dev Returns whether the given address is authorized to upgrade the channel
     */
    function _isAuthorizedUpgrader(string calldata portId, string calldata channelId, address msgSender)
        internal
        view
        virtual
        returns (bool);

    /**
     * @dev Removes the proposed upgrade for the given port and channel
     */
    function _removeUpgradeProposal(string calldata portId, string calldata channelId) internal {
        if (upgradeProposals[portId][channelId].fields.connection_hops.length == 0) {
            revert IBCChannelUpgradableModuleUpgradeNotFound();
        }
        IIBCHandler handler = IIBCHandler(ibcAddress());
        (, bool found) = handler.getChannelUpgrade(portId, channelId);
        if (found) {
            Channel.Data memory channel;
            (channel, found) = handler.getChannel(portId, channelId);
            if (!found) {
                revert IBCChannelUpgradableModuleChannelNotFound();
            }
            if (channel.state != Channel.State.STATE_OPEN) {
                revert IBCChannelUpgradableModuleCannotRemoveInProgressUpgrade();
            }
        }
        delete upgradeProposals[portId][channelId];
    }

    /**
     * @dev Compares two UpgradeFields structs
     */
    function equals(UpgradeFields.Data storage a, UpgradeFields.Data calldata b) internal view returns (bool) {
        if (a.ordering != b.ordering) {
            return false;
        }
        if (a.connection_hops.length != b.connection_hops.length) {
            return false;
        }
        for (uint256 i = 0; i < a.connection_hops.length; i++) {
            if (keccak256(abi.encodePacked(a.connection_hops[i])) != keccak256(abi.encodePacked(b.connection_hops[i])))
            {
                return false;
            }
        }
        return keccak256(abi.encodePacked(a.version)) == keccak256(abi.encodePacked(b.version));
    }
}
