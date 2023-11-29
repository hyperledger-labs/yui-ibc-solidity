// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {IIBCChannelHandshake, IIBCChannelPacket} from "../04-channel/IIBCChannel.sol";
import {IIBCHostConfigurator} from "../24-host/IIBCHostConfigurator.sol";
import {IIBCQuerier} from "./IIBCQuerier.sol";

/**
 * @dev IIBCHandler is a handler interface supports [ICS-25](https://github.com/cosmos/ibc/tree/main/spec/core/ics-025-handler-interface),
 */
interface IIBCHandler is
    IIBCClient,
    IIBCConnection,
    IIBCChannelHandshake,
    IIBCChannelPacket,
    IIBCQuerier,
    IIBCHostConfigurator
{}
