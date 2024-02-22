// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "../../../../contracts/proto/Connection.sol";
import "../../../../contracts/proto/Channel.sol";
import "../../../../contracts/core/04-channel/IBCChannelLib.sol";
import "../../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../../contracts/core/25-handler/OwnableIBCHandler.sol";

contract TestableIBCHandler is OwnableIBCHandler {
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_
    )
        OwnableIBCHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_
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
        commitments[IBCCommitment.packetCommitmentKey(portId, channelId, sequence)] = commitment;
    }

    function setPacketCommitment(Packet.Data memory packet) external {
        commitments[IBCCommitment.packetCommitmentKey(packet.source_port, packet.source_channel, packet.sequence)] =
        keccak256(
            abi.encodePacked(
                sha256(
                    abi.encodePacked(
                        packet.timeout_timestamp,
                        packet.timeout_height.revision_number,
                        packet.timeout_height.revision_height,
                        sha256(packet.data)
                    )
                )
            )
        );
    }

    function setCapability(string calldata name, address addr) external {
        capabilities[name] = addr;
    }

    function getPacketCommitment(string memory portId, string memory channelId, uint64 sequence)
        external
        view
        returns (bytes32)
    {
        return getCommitment(IBCCommitment.packetCommitmentKey(portId, channelId, sequence));
    }

    function hasPacketReceipt(string memory destinationPortId, string memory destinationChannelId, uint64 sequence)
        external
        view
        returns (bool)
    {
        return IBCChannelLib.receiptCommitmentToReceipt(
            commitments[IBCCommitment.packetReceiptCommitmentKey(destinationPortId, destinationChannelId, sequence)]
        ) == IBCChannelLib.PacketReceipt.SUCCESSFUL;
    }
}