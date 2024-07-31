// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {UpgradeFields, Timeout} from "../../proto/Channel.sol";
import {IIBCChannelUpgradeBase} from "../../core/04-channel/IIBCChannelUpgrade.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {IBCMockApp} from "./IBCMockApp.sol";
import {IBCChannelUpgradableModuleBase} from "../commons/IBCChannelUpgradableModule.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";

contract IBCChannelUpgradableMockApp is IBCMockApp, IBCChannelUpgradableModuleBase {
    constructor(IIBCHandler ibcHandler_) IBCMockApp(ibcHandler_) {}

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(IBCChannelUpgradableModuleBase, IBCAppBase)
        returns (bool)
    {
        return super.supportsInterface(interfaceId) || interfaceId == this.proposeAndInitUpgrade.selector;
    }

    /**
     * @dev Propose upgrade and perform chanUpgradeInit.
     */
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
