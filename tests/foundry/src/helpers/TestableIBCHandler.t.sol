// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "../../../../contracts/proto/Connection.sol";
import "../../../../contracts/proto/Channel.sol";
import "../../../../contracts/core/03-connection/IBCConnectionLib.sol";
import "../../../../contracts/core/04-channel/IIBCChannel.sol";
import "../../../../contracts/core/04-channel/IBCChannelLib.sol";
import "../../../../contracts/core/24-host/IBCHostLib.sol";
import "../../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../../contracts/core/25-handler/OwnableIBCHandler.sol";

contract TestableIBCHandler is OwnableIBCHandler {
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_,
        IIBCChannelUpgradeInitTryAck ibcChannelUpgradeInitTryAck_,
        IIBCChannelUpgradeConfirmOpenTimeoutCancel ibcChannelUpgradeConfirmOpenTimeoutCancel_
    )
        OwnableIBCHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_,
            ibcChannelUpgradeInitTryAck_,
            ibcChannelUpgradeConfirmOpenTimeoutCancel_
        )
    {}

    function setConnection(string memory connectionId, ConnectionEnd.Data memory connection) external {
        ConnectionEnd.Data storage conn = getConnectionStorage()[connectionId].connection;
        conn.client_id = connection.client_id;
        conn.state = connection.state;
        conn.delay_period = connection.delay_period;
        delete conn.versions;
        for (uint8 i = 0; i < connection.versions.length; i++) {
            conn.versions.push(connection.versions[i]);
        }
        conn.counterparty = connection.counterparty;
        getCommitments()[keccak256(IBCCommitment.connectionPath(connectionId))] = keccak256(ConnectionEnd.encode(connection));
    }

    function setChannel(string memory portId, string memory channelId, Channel.Data memory channel) external {
        getChannelStorage()[portId][channelId].channel = channel;
        getCommitments()[keccak256(IBCCommitment.channelPath(portId, channelId))] = keccak256(Channel.encode(channel));
    }

    function setNextSequenceSend(string calldata portId, string calldata channelId, uint64 sequence) external {
        getChannelStorage()[portId][channelId].nextSequenceSend = sequence;
    }

    function setNextSequenceRecv(string calldata portId, string calldata channelId, uint64 sequence) external {
        getChannelStorage()[portId][channelId].nextSequenceRecv = sequence;
        getCommitments()[keccak256(IBCCommitment.nextSequenceRecvCommitmentPath(portId, channelId))] =
            keccak256(abi.encodePacked(sequence));
    }

    function setNextSequenceAck(string calldata portId, string calldata channelId, uint64 sequence) external {
        getChannelStorage()[portId][channelId].nextSequenceAck = sequence;
    }

    function setNextClientSequence(uint64 sequence) external {
        getHostStorage().nextClientSequence = sequence;
    }

    function setNextConnectionSequence(uint64 sequence) external {
        getHostStorage().nextConnectionSequence = sequence;
    }

    function setNextChannelSequence(uint64 sequence) external {
        getHostStorage().nextChannelSequence = sequence;
    }

    function setPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence, bytes32 commitment)
        external
    {
        getCommitments()[IBCCommitment.packetCommitmentKeyCalldata(portId, channelId, sequence)] = commitment;
    }

    function setPacketCommitment(Packet calldata packet) external {
        getCommitments()[IBCCommitment.packetCommitmentKeyCalldata(packet.sourcePort, packet.sourceChannel, packet.sequence)]
        = keccak256(
            abi.encodePacked(
                sha256(
                    abi.encodePacked(
                        packet.timeoutTimestamp,
                        packet.timeoutHeight.revision_number,
                        packet.timeoutHeight.revision_height,
                        sha256(packet.data)
                    )
                )
            )
        );
    }

    function setChannelCapability(string calldata portId, string calldata channelId, address addr) external {
        getHostStorage().channelCapabilities[portId][channelId] = addr;
    }

    function getPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bytes32)
    {
        return getCommitment(IBCCommitment.packetCommitmentKeyCalldata(portId, channelId, sequence));
    }

    function hasPacketReceipt(string calldata destinationPortId, string calldata destinationChannelId, uint64 sequence)
        external
        view
        returns (bool)
    {
        return IBCChannelLib.receiptCommitmentToReceipt(
            getCommitments()[IBCCommitment.packetReceiptCommitmentKeyCalldata(
                destinationPortId, destinationChannelId, sequence
            )]
        ) == IBCChannelLib.PacketReceipt.SUCCESSFUL;
    }
}
