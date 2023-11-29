// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";
import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {
    IIBCChannelHandshake, IIBCChannelPacketSendRecv, IIBCChannelPacketTimeout
} from "../04-channel/IIBCChannel.sol";

/**
 * @dev IBCClientConnectionChannelHandler is a handler implements ICS-02, ICS-03, and ICS-04
 */
abstract contract IBCClientConnectionChannelHandler is
    IIBCClient,
    IIBCConnection,
    IIBCChannelHandshake,
    IIBCChannelPacketSendRecv,
    IIBCChannelPacketTimeout
{
    address internal immutable ibcClient;
    address internal immutable ibcConnection;
    address internal immutable ibcChannelHandshake;
    address internal immutable ibcChannelPacketSendRecv;
    address internal immutable ibcChannelPacketTimeout;

    /**
     * @dev The arguments of constructor must satisfy the followings:
     * @param ibcClient_ is the address of a contract that implements `IIBCClient`.
     * @param ibcConnection_ is the address of a contract that implements `IIBCConnection`.
     * @param ibcChannelHandshake_ is the address of a contract that implements `IIBCChannelHandshake`.
     * @param ibcChannelPacketSendRecv_ is the address of a contract that implements `IICS04Wrapper + IIBCChannelPacketReceiver`.
     * @param ibcChannelPacketTimeout_ is the address of a contract that implements `IIBCChannelPacketTimeout`.
     */
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_
    ) {
        ibcClient = address(ibcClient_);
        ibcConnection = address(ibcConnection_);
        ibcChannelHandshake = address(ibcChannelHandshake_);
        ibcChannelPacketSendRecv = address(ibcChannelPacketSendRecv_);
        ibcChannelPacketTimeout = address(ibcChannelPacketTimeout_);
    }

    function createClient(MsgCreateClient calldata) external returns (string memory) {
        doFallback(ibcClient);
    }

    function updateClient(MsgUpdateClient calldata) external {
        doFallback(ibcClient);
    }

    function connectionOpenInit(IIBCConnection.MsgConnectionOpenInit calldata) external returns (string memory) {
        doFallback(ibcConnection);
    }

    function connectionOpenTry(IIBCConnection.MsgConnectionOpenTry calldata) external returns (string memory) {
        doFallback(ibcConnection);
    }

    function connectionOpenAck(IIBCConnection.MsgConnectionOpenAck calldata) external {
        doFallback(ibcConnection);
    }

    function connectionOpenConfirm(IIBCConnection.MsgConnectionOpenConfirm calldata) external {
        doFallback(ibcConnection);
    }

    function channelOpenInit(IIBCChannelHandshake.MsgChannelOpenInit calldata)
        external
        returns (string memory, string memory)
    {
        doFallback(ibcChannelHandshake);
    }

    function channelOpenTry(IIBCChannelHandshake.MsgChannelOpenTry calldata)
        external
        returns (string memory, string memory)
    {
        doFallback(ibcChannelHandshake);
    }

    function channelOpenAck(IIBCChannelHandshake.MsgChannelOpenAck calldata) external {
        doFallback(ibcChannelHandshake);
    }

    function channelOpenConfirm(IIBCChannelHandshake.MsgChannelOpenConfirm calldata) external {
        doFallback(ibcChannelHandshake);
    }

    function channelCloseInit(IIBCChannelHandshake.MsgChannelCloseInit calldata) external {
        doFallback(ibcChannelHandshake);
    }

    function channelCloseConfirm(IIBCChannelHandshake.MsgChannelCloseConfirm calldata) external {
        doFallback(ibcChannelHandshake);
    }

    function sendPacket(string calldata, string calldata, Height.Data calldata, uint64, bytes calldata)
        external
        returns (uint64)
    {
        doFallback(ibcChannelPacketSendRecv);
    }

    function writeAcknowledgement(string calldata, string calldata, uint64, bytes calldata) external {
        doFallback(ibcChannelPacketSendRecv);
    }

    function recvPacket(MsgPacketRecv calldata) external {
        doFallback(ibcChannelPacketSendRecv);
    }

    function acknowledgePacket(MsgPacketAcknowledgement calldata) external {
        doFallback(ibcChannelPacketSendRecv);
    }

    function timeoutPacket(MsgTimeoutPacket calldata) external {
        doFallback(ibcChannelPacketTimeout);
    }

    function timeoutOnClose(MsgTimeoutOnClose calldata) external {
        doFallback(ibcChannelPacketTimeout);
    }

    function doFallback(address impl) internal virtual {
        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), impl, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 { revert(0, returndatasize()) }
            default { return(0, returndatasize()) }
        }
    }
}
