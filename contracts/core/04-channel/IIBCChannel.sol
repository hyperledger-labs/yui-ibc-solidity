// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";
import {Channel, Packet} from "../../proto/Channel.sol";

interface IIBCChannelHandshake {
    struct MsgChannelOpenInit {
        string portId;
        Channel.Data channel;
    }

    struct MsgChannelOpenTry {
        string portId;
        Channel.Data channel;
        string counterpartyVersion;
        bytes proofInit;
        Height.Data proofHeight;
    }

    struct MsgChannelOpenAck {
        string portId;
        string channelId;
        string counterpartyVersion;
        string counterpartyChannelId;
        bytes proofTry;
        Height.Data proofHeight;
    }

    struct MsgChannelOpenConfirm {
        string portId;
        string channelId;
        bytes proofAck;
        Height.Data proofHeight;
    }

    struct MsgChannelCloseInit {
        string portId;
        string channelId;
    }

    struct MsgChannelCloseConfirm {
        string portId;
        string channelId;
        bytes proofInit;
        Height.Data proofHeight;
    }

    event GeneratedChannelIdentifier(string);

    /**
     * @dev channelOpenInit is called by a module to initiate a channel opening handshake with a module on another chain.
     */
    function channelOpenInit(MsgChannelOpenInit calldata msg_)
        external
        returns (string memory channelId, string memory version);

    /**
     * @dev channelOpenTry is called by a module to accept the first step of a channel opening handshake initiated by a module on another chain.
     */
    function channelOpenTry(MsgChannelOpenTry calldata msg_)
        external
        returns (string memory channelId, string memory version);

    /**
     * @dev channelOpenAck is called by the handshake-originating module to acknowledge the acceptance of the initial request by the counterparty module on the other chain.
     */
    function channelOpenAck(MsgChannelOpenAck calldata msg_) external;

    /**
     * @dev channelOpenConfirm is called by the counterparty module to close their end of the channel, since the other end has been closed.
     */
    function channelOpenConfirm(MsgChannelOpenConfirm calldata msg_) external;

    /**
     * @dev channelCloseInit is called by either module to close their end of the channel. Once closed, channels cannot be reopened.
     */
    function channelCloseInit(MsgChannelCloseInit calldata msg_) external;

    /**
     * @dev channelCloseConfirm is called by the counterparty module to close their end of the
     * channel, since the other end has been closed.
     */
    function channelCloseConfirm(MsgChannelCloseConfirm calldata msg_) external;
}

interface IICS04SendPacket {
    event SendPacket(
        uint64 sequence,
        string sourcePort,
        string sourceChannel,
        Height.Data timeoutHeight,
        uint64 timeoutTimestamp,
        bytes data
    );

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
}

interface IICS04WriteAcknowledgement {
    event WriteAcknowledgement(
        string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement
    );

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
}

interface IIBCChannelRecvPacket {
    struct MsgPacketRecv {
        Packet.Data packet;
        bytes proof;
        Height.Data proofHeight;
    }

    event RecvPacket(Packet.Data packet);

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(MsgPacketRecv calldata msg_) external;
}

interface IIBCChannelAcknowledgePacket {
    struct MsgPacketAcknowledgement {
        Packet.Data packet;
        bytes acknowledgement;
        bytes proof;
        Height.Data proofHeight;
    }

    event AcknowledgePacket(Packet.Data packet, bytes acknowledgement);

    /**
     * @dev AcknowledgePacket is called by a module to process the acknowledgement of a
     * packet previously sent by the calling module on a channel to a counterparty
     * module on the counterparty chain. Its intended usage is within the ante
     * handler. AcknowledgePacket will clean up the packet commitment,
     * which is no longer necessary since the packet has been received and acted upon.
     * It will also increment NextSequenceAck in case of ORDERED channels.
     */
    function acknowledgePacket(MsgPacketAcknowledgement calldata msg_) external;
}

interface IIBCChannelPacketTimeout {
    struct MsgTimeoutPacket {
        Packet.Data packet;
        bytes proof;
        Height.Data proofHeight;
        uint64 nextSequenceRecv;
    }

    struct MsgTimeoutOnClose {
        Packet.Data packet;
        bytes proofUnreceived;
        bytes proofClose;
        Height.Data proofHeight;
        uint64 nextSequenceRecv;
    }

    event TimeoutPacket(Packet.Data packet);

    /**
     * @dev TimeoutPacket is called by a module which originally attempted to send a
     * packet to a counterparty module, where the timeout height has passed on the
     * counterparty chain without the packet being committed, to prove that the
     * packet can no longer be executed and to allow the calling module to safely
     * perform appropriate state transitions. Its intended usage is within the
     * ante handler.
     */
    function timeoutPacket(MsgTimeoutPacket calldata msg_) external;

    /**
     * @dev TimeoutOnClose is called by a module in order to prove that the channel to
     * which an unreceived packet was addressed has been closed, so the packet will
     * never be received (even if the timeoutHeight has not yet been reached).
     */
    function timeoutOnClose(MsgTimeoutOnClose calldata msg_) external;
}

interface IICS04Wrapper is IICS04SendPacket, IICS04WriteAcknowledgement {}

interface IIBCChannelPacketSendRecv is IICS04Wrapper, IIBCChannelRecvPacket, IIBCChannelAcknowledgePacket {}

interface IIBCChannelPacket is IIBCChannelPacketSendRecv, IIBCChannelPacketTimeout {}
