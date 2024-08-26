// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Upgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";
import {Options} from "openzeppelin-foundry-upgrades/Options.sol";
import {IBCClient} from "../../../contracts/core/02-client/IBCClient.sol";
import {IBCConnectionSelfStateNoValidation} from "../../../contracts/core/03-connection/IBCConnectionSelfStateNoValidation.sol";
import {IBCChannelHandshake} from "../../../contracts/core/04-channel/IBCChannelHandshake.sol";
import {IBCChannelPacketSendRecv} from "../../../contracts/core/04-channel/IBCChannelPacketSendRecv.sol";
import {IBCChannelPacketTimeout} from "../../../contracts/core/04-channel/IBCChannelPacketTimeout.sol";
import {IBCChannelUpgradeInitTryAck, IBCChannelUpgradeConfirmOpenTimeoutCancel} from "../../../contracts/core/04-channel/IBCChannelUpgrade.sol";
import {OwnableUpgradeableIBCHandler} from "../../../contracts/core/25-handler/OwnableUpgradeableIBCHandler.sol";
import {IBCTestHelper} from "./helpers/IBCTestHelper.t.sol";
import {TestableOwnableUpgradeableIBCHandlerV1} from "./helpers/TestableOwnableUpgradeableIBCHandlerV1.sol";
import {TestableOwnableUpgradeableIBCHandlerV2} from "./helpers/TestableOwnableUpgradeableIBCHandlerV2.sol";

contract ContractUpgrade is IBCTestHelper {
    function testUpgrade() public {
        if (!vm.envOr("TEST_UPGRADEABLE", false)) {
            return;
        }
        Options memory opts;
        opts.constructorData = abi.encode(
            new IBCClient(),
            new IBCConnectionSelfStateNoValidation(),
            new IBCChannelHandshake(),
            new IBCChannelPacketSendRecv(),
            new IBCChannelPacketTimeout(),
            new IBCChannelUpgradeInitTryAck(),
            new IBCChannelUpgradeConfirmOpenTimeoutCancel()
        );
        address proxy = Upgrades.deployUUPSProxy(
            "TestableOwnableUpgradeableIBCHandlerV1.sol",
            abi.encodePacked(OwnableUpgradeableIBCHandler.initialize.selector),
            opts
        );
        Upgrades.upgradeProxy(
            proxy,
            "TestableOwnableUpgradeableIBCHandlerV2.sol",
            abi.encodePacked(OwnableUpgradeableIBCHandler.initialize.selector),
            opts
        );
    }
}