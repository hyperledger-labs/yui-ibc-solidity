// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {UpgradeFields} from "../../../../contracts/proto/Channel.sol";
import {TestableIBCHandler} from "./TestableIBCHandler.t.sol";
import {IBCChannelUpgradableMockApp} from "../../../../contracts/apps/mock/IBCChannelUpgradableMockApp.sol";

contract TestIBCChannelUpgradableMockAppInconsistentVersions is IBCChannelUpgradableMockApp {
    constructor(TestableIBCHandler _ibcHandler) IBCChannelUpgradableMockApp(_ibcHandler) {}

    function onChanUpgradeInit(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        UpgradeFields.Data calldata proposedUpgradeFields
    ) public view virtual override onlyIBC returns (string memory version) {
        return super.onChanUpgradeInit(portId, channelId, upgradeSequence, proposedUpgradeFields);
    }

    function onChanUpgradeTry(
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence,
        UpgradeFields.Data calldata proposedUpgradeFields
    ) public view virtual override onlyIBC returns (string memory version) {
        return string(abi.encodePacked(
            super.onChanUpgradeTry(portId, channelId, upgradeSequence, proposedUpgradeFields),
            "-inconsistent"
        ));
    }
}