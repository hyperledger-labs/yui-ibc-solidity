// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../25-handler/IBCMsgs.sol";

interface IIBCChannelHandshake {
    /**
     * @dev channelOpenInit is called by a module to initiate a channel opening handshake with a module on another chain.
     */
    function channelOpenInit(IBCMsgs.MsgChannelOpenInit calldata msg_) external returns (string memory channelId);

    /**
     * @dev channelOpenTry is called by a module to accept the first step of a channel opening handshake initiated by a module on another chain.
     */
    function channelOpenTry(IBCMsgs.MsgChannelOpenTry calldata msg_) external returns (string memory channelId);

    /**
     * @dev channelOpenAck is called by the handshake-originating module to acknowledge the acceptance of the initial request by the counterparty module on the other chain.
     */
    function channelOpenAck(IBCMsgs.MsgChannelOpenAck calldata msg_) external;

    /**
     * @dev channelOpenConfirm is called by the counterparty module to close their end of the channel, since the other end has been closed.
     */
    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm calldata msg_) external;

    /**
     * @dev channelCloseInit is called by either module to close their end of the channel. Once closed, channels cannot be reopened.
     */
    function channelCloseInit(IBCMsgs.MsgChannelCloseInit calldata msg_) external;

    /**
     * @dev channelCloseConfirm is called by the counterparty module to close their end of the
     * channel, since the other end has been closed.
     */
    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm calldata msg_) external;

    /**
     * @dev writeChannel writes a channel which has successfully passed the OpenInit or OpenTry handshake step.
     */
    function writeChannel(
        string calldata portId,
        string calldata channelId,
        Channel.State state,
        Channel.Order order,
        ChannelCounterparty.Data calldata counterparty,
        string[] calldata connectionHops,
        string calldata version
    ) external;
}

interface IIBCPacket {
    /**
     * @dev sendPacket is called by a module in order to send an IBC packet on a channel.
     * The packet sequence generated for the packet to be sent is returned. An error
     * is returned if one occurs. Also, `timeoutTimestamp` is given in nanoseconds since unix epoch.
     */
    function sendPacket(
        string calldata sourcePort,
        string calldata sourceChannel,
        Height.Data calldata timeoutHeight,
        uint64 timeoutTimestamp,
        bytes calldata data
    ) external returns (uint64);

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external;

    /**
     * @dev writeAcknowledgement writes the packet execution acknowledgement to the state,
     * which will be verified by the counterparty chain using AcknowledgePacket.
     */
    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external;

    /**
     * @dev AcknowledgePacket is called by a module to process the acknowledgement of a
     * packet previously sent by the calling module on a channel to a counterparty
     * module on the counterparty chain. Its intended usage is within the ante
     * handler. AcknowledgePacket will clean up the packet commitment,
     * which is no longer necessary since the packet has been received and acted upon.
     * It will also increment NextSequenceAck in case of ORDERED channels.
     */
    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external;

    /**
     * @dev TimeoutPacket is called by a module which originally attempted to send a
     * packet to a counterparty module, where the timeout height has passed on the
     * counterparty chain without the packet being committed, to prove that the
     * packet can no longer be executed and to allow the calling module to safely
     * perform appropriate state transitions. Its intended usage is within the
     * ante handler.
     */
    function timeoutPacket(IBCMsgs.MsgTimeoutPacket calldata msg_) external;

    /**
     * @dev TimeoutOnClose is called by a module in order to prove that the channel to
     * which an unreceived packet was addressed has been closed, so the packet will
     * never be received (even if the timeoutHeight has not yet been reached).
     */
    function timeoutOnClose(IBCMsgs.MsgTimeoutOnClose calldata msg_) external;
}
