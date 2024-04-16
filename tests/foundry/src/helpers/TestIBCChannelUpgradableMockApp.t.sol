// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Upgrade, UpgradeFields, Timeout} from "../../../../contracts/proto/Channel.sol";
import {
    IIBCChannelUpgrade, IIBCChannelUpgradeBase
} from "../../../../contracts/core/04-channel/IIBCChannelUpgrade.sol";
import {IIBCHandler} from "../../../../contracts/core/25-handler/IIBCHandler.sol";
import {IBCMockApp} from "../../../../contracts/apps/mock/IBCMockApp.sol";
import {IBCChannelUpgradableModuleBase} from "./IBCChannelUpgradableModule.sol";

contract TestIBCChannelUpgradableMockApp is IBCMockApp, IBCChannelUpgradableModuleBase {
    constructor(IIBCHandler ibcHandler_) IBCMockApp(ibcHandler_) {}

    function proposeAndInitUpgrade(
        string calldata portId,
        string calldata channelId,
        UpgradeFields.Data calldata proposedUpgradeFields,
        Timeout.Data calldata timeout
    ) public virtual returns (uint64) {
        proposeUpgrade(portId, channelId, proposedUpgradeFields, timeout);
        return IIBCHandler(ibcHandler).channelUpgradeInit(
            IIBCChannelUpgradeBase.MsgChannelUpgradeInit({
                portId: portId,
                channelId: channelId,
                proposedUpgradeFields: proposedUpgradeFields
            })
        );
    }

    function _isAuthorizedUpgrader(string calldata, string calldata, address msgSender)
        internal
        view
        override
        returns (bool)
    {
        return msgSender == owner() || msgSender == address(this);
    }
}
