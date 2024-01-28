// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {UpgradeFields, Timeout} from "../../proto/Channel.sol";

interface IIBCModuleUpgrade {
    /**
     * @dev Returns the absolute timeout for the upgrade
     * @param portId Port identifier
     * @param channelId Channel identifier
     */
    function getUpgradeTimeout(string calldata portId, string calldata channelId)
        external
        view
        returns (Timeout.Data memory);

    /**
     * @dev Returns whether the `msgSender` is an authorized upgrader for the given channel
     * @param portId Port identifier
     * @param channelId Channel identifier
     * @param msgSender sender of the upgrade message
     */
    function isAuthorizedUpgrader(string calldata portId, string calldata channelId, address msgSender)
        external
        view
        returns (bool);

    /**
     * @dev Returns whether the given channel can transition to flush complete at the given upgrade sequence
     * NOTE: this function is never called in some cases where the core contract can guarantee all inflight packets are already acknolwedged
     * @param portId Port identifier
     * @param channelId Channel identifier
     * @param upgradeSequence Upgrade sequence
     * @param msgSender sender of the upgrade message
     */
    function canTransitionToFlushComplete(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        address msgSender
    ) external view returns (bool);

    /**
     * @dev OnChanUpgradeInit enables additional custom logic to be executed when the channel upgrade is initialized.
     * It must validate the proposed version, order, and connection hops.
     * NOTE: in the case of crossing hellos, this callback may be executed on both chains.
     * @param portId Port identifier
     * @param channelId Channel identifier
     * @param upgradeSequence Upgrade sequence
     * @param proposedUpgradeFields Proposed upgrade fields
     */
    function onChanUpgradeInit(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        UpgradeFields.Data calldata proposedUpgradeFields
    ) external view returns (string memory version);

    /**
     * @dev OnChanUpgradeTry enables additional custom logic to be executed in the ChannelUpgradeTry step of the
     * channel upgrade handshake. It must validate the proposed version (provided by the counterparty), order,
     * and connection hops.
     * @param portId Port identifier
     * @param channelId Channel identifier
     * @param upgradeSequence Upgrade sequence
     * @param proposedUpgradeFields Proposed upgrade fields
     */
    function onChanUpgradeTry(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        UpgradeFields.Data calldata proposedUpgradeFields
    ) external view returns (string memory version);

    /**
     * @dev OnChanUpgradeAck enables additional custom logic to be executed in the ChannelUpgradeAck step of the
     * channel upgrade handshake. It must validate the version proposed by the counterparty.
     * @param portId Port identifier
     * @param channelId Channel identifier
     * @param upgradeSequence Upgrade sequence
     * @param counterpartyVersion Counterparty version
     */
    function onChanUpgradeAck(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        string calldata counterpartyVersion
    ) external view;

    /**
     * @dev OnChanUpgradeOpen enables additional custom logic to be executed when the channel upgrade has successfully completed, and the channel
     * has returned to the OPEN state. Any logic associated with changing of the channel fields should be performed
     * in this callback.
     *  @param portId Port identifier
     * @param channelId Channel identifier
     * @param upgradeSequence Upgrade sequence
     */
    function onChanUpgradeOpen(string calldata portId, string calldata channelId, uint64 upgradeSequence) external;
}
