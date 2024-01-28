// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "../../../../contracts/proto/Connection.sol";
import "../../../../contracts/proto/Channel.sol";
import "../../../../contracts/core/04-channel/IIBCChannel.sol";
import "../../../../contracts/core/04-channel/IBCChannelLib.sol";
import "../../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../../contracts/core/25-handler/OwnableIBCHandler.sol";

contract TestableIBCHandler is OwnableIBCHandler {
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_,
        IIBCChannelUpgrade ibcChannelUpgrade_
    )
        OwnableIBCHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_,
            ibcChannelUpgrade_
        )
    {}

    function setConnection(string memory connectionId, ConnectionEnd.Data memory connection) external {
        connections[connectionId].client_id = connection.client_id;
        connections[connectionId].state = connection.state;
        connections[connectionId].delay_period = connection.delay_period;
        delete connections[connectionId].versions;
        for (uint8 i = 0; i < connection.versions.length; i++) {
            connections[connectionId].versions.push(connection.versions[i]);
        }
        connections[connectionId].counterparty = connection.counterparty;
        commitments[keccak256(IBCCommitment.connectionPath(connectionId))] = keccak256(ConnectionEnd.encode(connection));
    }

    function setChannel(string memory portId, string memory channelId, Channel.Data memory channel) external {
        channels[portId][channelId] = channel;
        commitments[keccak256(IBCCommitment.channelPath(portId, channelId))] = keccak256(Channel.encode(channel));
    }

    function setNextSequenceSend(string calldata portId, string calldata channelId, uint64 sequence) external {
        nextSequenceSends[portId][channelId] = sequence;
    }

    function setNextSequenceRecv(string calldata portId, string calldata channelId, uint64 sequence) external {
        nextSequenceRecvs[portId][channelId] = sequence;
        commitments[keccak256(IBCCommitment.nextSequenceRecvCommitmentPath(portId, channelId))] =
            keccak256(abi.encodePacked(sequence));
    }

    function setNextSequenceAck(string calldata portId, string calldata channelId, uint64 sequence) external {
        nextSequenceAcks[portId][channelId] = sequence;
    }

    function setNextClientSequence(uint64 sequence) external {
        nextClientSequence = sequence;
    }

    function setNextConnectionSequence(uint64 sequence) external {
        nextConnectionSequence = sequence;
    }

    function setNextChannelSequence(uint64 sequence) external {
        nextChannelSequence = sequence;
    }

    function setPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence, bytes32 commitment)
        external
    {
        commitments[IBCCommitment.packetCommitmentKeyCalldata(portId, channelId, sequence)] = commitment;
    }

    function setPacketCommitment(Packet calldata packet) external {
        commitments[IBCCommitment.packetCommitmentKeyCalldata(packet.sourcePort, packet.sourceChannel, packet.sequence)]
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
        channelCapabilities[portId][channelId] = addr;
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
            commitments[IBCCommitment.packetReceiptCommitmentKeyCalldata(
                destinationPortId, destinationChannelId, sequence
            )]
        ) == IBCChannelLib.PacketReceipt.SUCCESSFUL;
    }
}
