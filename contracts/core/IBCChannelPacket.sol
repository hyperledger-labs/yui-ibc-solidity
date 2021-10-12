pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Channel.sol";
import "./IBCConnection.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";
import "../lib/strings.sol";

library IBCChannelPacket {
    using strings for *;

    function sendPacket(IBCHost host, Packet.Data calldata packet) external {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        IClient client;
        uint64 latestHeight;
        uint64 latestTimestamp;
        uint64 nextSequenceSend;
        bool found;

        (channel, found) = host.getChannel(packet.source_port, packet.source_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");
        require(packet.destination_port.toSlice().equals(channel.counterparty.port_id.toSlice()), "packet destination port doesn't match the counterparty's port");
        require(packet.destination_channel.toSlice().equals(channel.counterparty.channel_id.toSlice()), "packet destination channel doesn't match the counterparty's channel");
        (connection, found) = host.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        client = IBCClient.getClient(host, connection.client_id);
        (latestHeight, found) = client.getLatestHeight(host, connection.client_id);
        require(packet.timeout_height.revision_height == 0 || latestHeight < packet.timeout_height.revision_height, "receiving chain block height >= packet timeout height");
        (latestTimestamp, found) = client.getTimestampAtHeight(host, connection.client_id, latestHeight);
        require(found, "consensusState not found");
        require(packet.timeout_timestamp == 0 || latestTimestamp < packet.timeout_timestamp, "receiving chain block timestamp >= packet timeout timestamp");

        nextSequenceSend = host.getNextSequenceSend(packet.source_port, packet.source_channel);
        require(nextSequenceSend > 0, "sequenceSend not found");
        require(packet.sequence == nextSequenceSend, "packet sequence ≠ next send sequence");

        nextSequenceSend++;
        host.setNextSequenceSend(packet.source_port, packet.source_channel, nextSequenceSend);
        host.setPacketCommitment(packet.source_port, packet.source_channel, packet.sequence, packet);

        // TODO emit an event that includes a packet
    }

    function recvPacket(IBCHost host, IBCMsgs.MsgPacketRecv calldata msg_) external {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;
        (channel, found) = host.getChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        // TODO
        // Authenticate capability to ensure caller has authority to receive packet on this channel

        require(msg_.packet.source_port.toSlice().equals(channel.counterparty.port_id.toSlice()), "packet source port doesn't match the counterparty's port");
        require(msg_.packet.source_channel.toSlice().equals(channel.counterparty.channel_id.toSlice()), "packet source channel doesn't match the counterparty's channel");

        (connection, found) = host.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        require(msg_.packet.timeout_height.revision_height == 0 || block.number < msg_.packet.timeout_height.revision_height, "block height >= packet timeout height");
        require(msg_.packet.timeout_timestamp == 0 || block.timestamp < msg_.packet.timeout_timestamp, "block timestamp >= packet timeout timestamp");

        bytes32 commitment = host.makePacketCommitment(msg_.packet);
        require(IBCConnection.verifyPacketCommitment(host, connection, msg_.proofHeight, msg_.proof, msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence, commitment), "failed to verify packet commitment");

        if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            require(!host.hasPacketReceipt(msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence), "packet sequence already has been received");
            host.setPacketReceipt(msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence);
        } else if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            uint64 nextSequenceRecv = host.getNextSequenceRecv(msg_.packet.destination_port, msg_.packet.destination_channel);
            require(nextSequenceRecv > 0 && nextSequenceRecv == msg_.packet.sequence, "packet sequence ≠ next receive sequence");
            host.setNextSequenceRecv(msg_.packet.destination_port, msg_.packet.destination_channel, nextSequenceRecv+1);
        } else {
            revert("unknown ordering type");
        }
    }

    // WriteAcknowledgement writes the packet execution acknowledgement to the state,
    // which will be verified by the counterparty chain using AcknowledgePacket.
    function writeAcknowledgement(IBCHost host, string calldata destinationPortId, string calldata destinationChannel, uint64 sequence, bytes calldata acknowledgement) external {
        host.onlyIBCModule();
        Channel.Data memory channel;
        bytes32 ackHash;
        bool found;
        (channel, found) = host.getChannel(destinationPortId, destinationChannel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        (ackHash, found) = host.getPacketAcknowledgementCommitment(destinationPortId, destinationChannel, sequence);
        require(!found, "acknowledgement for packet already exists");

        require(acknowledgement.length > 0, "acknowledgement cannot be empty");
        host.setPacketAcknowledgementCommitment(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    // TODO use calldata
    function acknowledgePacket(IBCHost host, IBCMsgs.MsgPacketAcknowledgement memory msg_) public {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bytes32 commitment;
        bool found;
        (channel, found) = host.getChannel(msg_.packet.source_port, msg_.packet.source_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        require(msg_.packet.destination_port.toSlice().equals(channel.counterparty.port_id.toSlice()), "packet destination port doesn't match the counterparty's port");
        require(msg_.packet.destination_channel.toSlice().equals(channel.counterparty.channel_id.toSlice()), "packet destination channel doesn't match the counterparty's channel");

        (connection, found) = host.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        (commitment, found) = host.getPacketCommitment(msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence);
        require(found, "packet commitment not found");

        require(commitment == host.makePacketCommitment(msg_.packet), "commitment bytes are not equal");

        require(IBCConnection.verifyPacketAcknowledgement(host, connection, msg_.proofHeight, msg_.proof, msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, msg_.acknowledgement), "failed to verify packet acknowledgement commitment");

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            uint64 nextSequenceAck = host.getNextSequenceAck(msg_.packet.source_port, msg_.packet.source_channel);
            require(nextSequenceAck == 0, "sequence ack not found");
            require(msg_.packet.sequence == nextSequenceAck, "packet sequence ≠ next ack sequence");
            nextSequenceAck++;
            host.setNextSequenceAck(msg_.packet.source_port, msg_.packet.source_channel, nextSequenceAck);
        }

        host.deletePacketCommitment(msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence);
    }

}
