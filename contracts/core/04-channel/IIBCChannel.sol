// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {Channel} from "../../proto/Channel.sol";

/// @notice Packet defines a type that carries data across different chains through IBC.
/// @param sequence corresponds to the order of sends and receives, where a packet with an earlier sequence number must be sent and received before a packet with a later sequence number
/// @param sourcePort identifies the port on the sending chain
/// @param sourceChannel identifies the channel end on the sending chain
/// @param destPort identifies the port on the receiving chain
/// @param destChannel identifies the channel end on the receiving chain
/// @param data is an opaque value which can be defined by the application logic of the associated modules
/// @param timeoutHeight indicates a consensus height on the destination chain after which the packet will no longer be processed, and will instead count as having timed-out
/// @param timeoutTimestamp indicates a timestamp on the destination chain after which the packet will no longer be processed, and will instead count as having timed-out
struct Packet {
    uint64 sequence;
    string sourcePort;
    string sourceChannel;
    string destinationPort;
    string destinationChannel;
    bytes data;
    Height.Data timeoutHeight;
    uint64 timeoutTimestamp;
}

interface IIBCChannelHandshake {
    // --------------------- Data Structure --------------------- //

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

    // --------------------- Events --------------------- //

    /// @param channelId channel identifier
    event GeneratedChannelIdentifier(string channelId);

    // --------------------- Functions --------------------- //

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
    // --------------------- Events --------------------- //

    /// @notice event emitted upon sending a packet
    event SendPacket(
        uint64 sequence,
        string sourcePort,
        string sourceChannel,
        Height.Data timeoutHeight,
        uint64 timeoutTimestamp,
        bytes data
    );

    // --------------------- Functions --------------------- //

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
    // --------------------- Events --------------------- //

    /// @notice event emitted upon writing an acknowledgement
    /// @param destinationPortId destination port
    /// @param destinationChannel destination channel
    /// @param sequence packet sequence
    /// @param acknowledgement acknowledgement
    event WriteAcknowledgement(
        string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement
    );

    // --------------------- Functions --------------------- //

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
    // --------------------- Data Structure --------------------- //

    struct MsgPacketRecv {
        Packet packet;
        bytes proof;
        Height.Data proofHeight;
    }

    // --------------------- Events --------------------- //

    /// @notice event emitted upon receiving a packet
    event RecvPacket(Packet packet);

    // --------------------- Functions --------------------- //

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(MsgPacketRecv calldata msg_) external;
}

interface IIBCChannelAcknowledgePacket {
    // --------------------- Data Structure --------------------- //

    struct MsgPacketAcknowledgement {
        Packet packet;
        bytes acknowledgement;
        bytes proof;
        Height.Data proofHeight;
    }

    // --------------------- Events --------------------- //

    /// @notice event emitted upon acknowledging a packet
    event AcknowledgePacket(Packet packet, bytes acknowledgement);

    // --------------------- Functions --------------------- //

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
    // --------------------- Data Structure --------------------- //

    struct MsgTimeoutPacket {
        Packet packet;
        bytes proof;
        Height.Data proofHeight;
        uint64 nextSequenceRecv;
    }

    struct MsgTimeoutOnClose {
        Packet packet;
        bytes proofUnreceived;
        bytes proofClose;
        Height.Data proofHeight;
        uint64 nextSequenceRecv;
        uint64 counterpartyUpgradeSequence;
    }

    // --------------------- Events --------------------- //

    /// @notice event emitted upon timeout of a packet
    event TimeoutPacket(Packet packet);

    // --------------------- Functions --------------------- //

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
