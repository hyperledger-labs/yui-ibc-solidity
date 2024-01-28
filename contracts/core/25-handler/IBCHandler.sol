// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {
    IIBCChannelHandshake, IIBCChannelPacketSendRecv, IIBCChannelPacketTimeout
} from "../04-channel/IIBCChannel.sol";
import {IIBCChannelUpgrade} from "../04-channel/IIBCChannelUpgrade.sol";
import {IBCHostConfigurator} from "../24-host/IBCHostConfigurator.sol";
import {IBCClientConnectionChannelHandler} from "./IBCClientConnectionChannelHandler.sol";
import {IBCQuerier} from "./IBCQuerier.sol";
import {IIBCHandler} from "./IIBCHandler.sol";

abstract contract IBCHandler is IBCHostConfigurator, IBCClientConnectionChannelHandler, IBCQuerier, IIBCHandler {
    /**
     * @dev The arguments of constructor must satisfy the followings:
     * @param ibcClient_ is the address of a contract that implements `IIBCClient`.
     * @param ibcConnection_ is the address of a contract that implements `IIBCConnection`.
     * @param ibcChannelHandshake_ is the address of a contract that implements `IIBCChannelHandshake`.
     * @param ibcChannelPacketSendRecv_ is the address of a contract that implements `IIBCChannelPacketSendRecv`.
     * @param ibcChannelPacketTimeout_ is the address of a contract that implements `IIBCChannelPacketTimeout`.
     * @param ibcChannelUpgrade_ is the address of a contract that implements `IIBCChannelUpgrade`.
     */
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_,
        IIBCChannelUpgrade ibcChannelUpgrade_
    )
        IBCClientConnectionChannelHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_,
            ibcChannelUpgrade_
        )
    {}
}
