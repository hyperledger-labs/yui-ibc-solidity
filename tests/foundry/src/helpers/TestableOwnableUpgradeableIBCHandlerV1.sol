// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Upgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";
import {Options} from "openzeppelin-foundry-upgrades/Options.sol";
import {IIBCClient} from "../../../../contracts/core/02-client/IIBCClient.sol";
import {IIBCConnection} from "../../../../contracts/core/03-connection/IIBCConnection.sol";
import {
    IIBCChannelHandshake, IIBCChannelPacketSendRecv, IIBCChannelPacketTimeout
} from "../../../../contracts/core/04-channel/IIBCChannel.sol";
import {
    IIBCChannelUpgradeInitTryAck,
    IIBCChannelUpgradeConfirmOpenTimeoutCancel
} from "../../../../contracts/core/04-channel/IIBCChannelUpgrade.sol";
import {OwnableUpgradeableIBCHandler} from "../../../../contracts/core/25-handler/OwnableUpgradeableIBCHandler.sol";

contract TestableOwnableUpgradeableIBCHandlerV1 is OwnableUpgradeableIBCHandler {
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_,
        IIBCChannelUpgradeInitTryAck ibcChannelUpgradeInitTryAck_,
        IIBCChannelUpgradeConfirmOpenTimeoutCancel ibcChannelUpgradeConfirmOpenTimeoutCancel_
    )
        OwnableUpgradeableIBCHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_,
            ibcChannelUpgradeInitTryAck_,
            ibcChannelUpgradeConfirmOpenTimeoutCancel_
        )
    {}
}