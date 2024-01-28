// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCClientErrors} from "../02-client/IIBCClientErrors.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {IIBCConnectionErrors} from "../03-connection/IIBCConnectionErrors.sol";
import {IIBCChannelHandshake, IIBCChannelPacket} from "../04-channel/IIBCChannel.sol";
import {IIBCChannelErrors} from "../04-channel/IIBCChannelErrors.sol";
import {IIBCChannelUpgrade} from "../04-channel/IIBCChannelUpgrade.sol";
import {IIBCChannelUpgradeErrors} from "../04-channel/IIBCChannelUpgradeErrors.sol";
import {IIBCHostConfigurator} from "../24-host/IIBCHostConfigurator.sol";
import {IIBCHostErrors} from "../24-host/IIBCHostErrors.sol";
import {IIBCQuerier} from "./IIBCQuerier.sol";

/**
 * @dev IIBCHandler is a handler interface supports [ICS-25](https://github.com/cosmos/ibc/tree/main/spec/core/ics-025-handler-interface),
 */
interface IIBCHandler is
    IIBCClient,
    IIBCClientErrors,
    IIBCConnection,
    IIBCConnectionErrors,
    IIBCChannelHandshake,
    IIBCChannelPacket,
    IIBCChannelErrors,
    IIBCChannelUpgrade,
    IIBCChannelUpgradeErrors,
    IIBCHostConfigurator,
    IIBCHostErrors,
    IIBCQuerier
{}
