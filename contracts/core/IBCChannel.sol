// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../proto/Client.sol";
import "../proto/Channel.sol";
import "./IBCMsgs.sol";
import "./IBCHeight.sol";
import "./IBCHost.sol";
import "./IBCIdentifier.sol";

contract IBCChannel is IBCHost {
    using IBCHeight for Height.Data;

    /* Handshake functions */

    function channelOpenInit(IBCMsgs.MsgChannelOpenInit calldata msg_) external returns (string memory) {
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        ConnectionEnd.Data storage connection = connections[msg_.channel.connection_hops[0]];
        require(
            connection.versions.length == 1, "single version must be negotiated on connection before opening channel"
        );
        require(msg_.channel.state == Channel.State.STATE_INIT, "channel state must STATE_INIT");

        // TODO verifySupportedFeature

        // TODO authenticates a port binding

        string memory channelId = generateChannelIdentifier();
        channels[msg_.portId][channelId] = msg_.channel;
        nextSequenceSends[msg_.portId][channelId] = 1;
        nextSequenceRecvs[msg_.portId][channelId] = 1;
        nextSequenceAcks[msg_.portId][channelId] = 1;
        return channelId;
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry calldata msg_) external returns (string memory) {
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        ConnectionEnd.Data storage connection = connections[msg_.channel.connection_hops[0]];
        require(
            connection.versions.length == 1, "single version must be negotiated on connection before opening channel"
        );
        require(msg_.channel.state == Channel.State.STATE_TRYOPEN, "channel state must be STATE_TRYOPEN");
        require(msg_.channel.connection_hops.length == 1);

        // TODO verifySupportedFeature

        // TODO authenticates a port binding

        ChannelCounterparty.Data memory expectedCounterparty =
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: ""});
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_INIT,
            ordering: msg_.channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(msg_.channel.connection_hops[0]),
            version: msg_.counterpartyVersion
        });
        require(
            verifyChannelState(
                connection,
                msg_.proofHeight,
                msg_.proofInit,
                msg_.channel.counterparty.port_id,
                msg_.channel.counterparty.channel_id,
                Channel.encode(expectedChannel)
            ),
            "failed to verify channel state"
        );

        string memory channelId = generateChannelIdentifier();
        channels[msg_.portId][channelId] = msg_.channel;
        nextSequenceSends[msg_.portId][channelId] = 1;
        nextSequenceRecvs[msg_.portId][channelId] = 1;
        nextSequenceAcks[msg_.portId][channelId] = 1;
        return channelId;
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck calldata msg_) external {
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        require(
            channel.state == Channel.State.STATE_INIT || channel.state == Channel.State.STATE_TRYOPEN,
            "invalid channel state"
        );

        // TODO authenticates a port binding

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");
        require(channel.connection_hops.length == 1);

        ChannelCounterparty.Data memory expectedCounterparty =
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId});
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_TRYOPEN,
            ordering: channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(channel.connection_hops[0]),
            version: msg_.counterpartyVersion
        });
        require(
            verifyChannelState(
                connection,
                msg_.proofHeight,
                msg_.proofTry,
                channel.counterparty.port_id,
                msg_.counterpartyChannelId,
                Channel.encode(expectedChannel)
            ),
            "failed to verify channel state"
        );
        channel.state = Channel.State.STATE_OPEN;
        channel.version = msg_.counterpartyVersion;
        channel.counterparty.channel_id = msg_.counterpartyChannelId;
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm calldata msg_) external {
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        require(channel.state == Channel.State.STATE_TRYOPEN, "channel state is not TRYOPEN");

        // TODO authenticates a port binding

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");
        require(channel.connection_hops.length == 1);

        ChannelCounterparty.Data memory expectedCounterparty =
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId});
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_OPEN,
            ordering: channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(channel.connection_hops[0]),
            version: channel.version
        });
        require(
            verifyChannelState(
                connection,
                msg_.proofHeight,
                msg_.proofAck,
                channel.counterparty.port_id,
                channel.counterparty.channel_id,
                Channel.encode(expectedChannel)
            ),
            "failed to verify channel state"
        );
        channel.state = Channel.State.STATE_OPEN;
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit calldata msg_) external {
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        require(channel.state != Channel.State.STATE_CLOSED, "channel state is already CLOSED");

        // TODO authenticates a port binding

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        channel.state = Channel.State.STATE_CLOSED;
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm calldata msg_) external {
        Channel.Data storage channel = channels[msg_.portId][msg_.channelId];
        require(channel.state != Channel.State.STATE_CLOSED, "channel state is already CLOSED");

        // TODO authenticates a port binding

        require(channel.connection_hops.length == 1);
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        ChannelCounterparty.Data memory expectedCounterparty =
            ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId});
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_CLOSED,
            ordering: channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(channel.connection_hops[0]),
            version: channel.version
        });
        require(
            verifyChannelState(
                connection,
                msg_.proofHeight,
                msg_.proofInit,
                channel.counterparty.port_id,
                channel.counterparty.channel_id,
                Channel.encode(expectedChannel)
            ),
            "failed to verify channel state"
        );
        channel.state = Channel.State.STATE_CLOSED;
    }

    /* Packet handlers */

    function sendPacket(Packet.Data calldata packet) external {
        uint64 latestTimestamp;

        Channel.Data storage channel = channels[packet.source_port][packet.source_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");
        require(
            hashString(packet.destination_port) == hashString(channel.counterparty.port_id),
            "packet destination port doesn't match the counterparty's port"
        );
        require(
            hashString(packet.destination_channel) == hashString(channel.counterparty.channel_id),
            "packet destination channel doesn't match the counterparty's channel"
        );
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        IClient client = IClient(clientImpls[connection.client_id]);
        (Height.Data memory latestHeight, bool found) = client.getLatestHeight(connection.client_id);
        require(
            packet.timeout_height.isZero() || latestHeight.lt(packet.timeout_height),
            "receiving chain block height >= packet timeout height"
        );
        (latestTimestamp, found) = client.getTimestampAtHeight(connection.client_id, latestHeight);
        require(found, "consensusState not found");
        require(
            packet.timeout_timestamp == 0 || latestTimestamp < packet.timeout_timestamp,
            "receiving chain block timestamp >= packet timeout timestamp"
        );

        require(
            packet.sequence == nextSequenceSends[packet.source_port][packet.source_channel],
            "packet sequence != next send sequence"
        );

        nextSequenceSends[packet.source_port][packet.source_channel]++;
        commitments[keccak256(
            IBCIdentifier.packetCommitmentPath(packet.source_port, packet.source_channel, packet.sequence)
        )] = keccak256(
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
        // TODO emit an event that includes a packet
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.destination_port][msg_.packet.destination_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        // TODO
        // Authenticate capability to ensure caller has authority to receive packet on this channel

        require(
            hashString(msg_.packet.source_port) == hashString(channel.counterparty.port_id),
            "packet source port doesn't match the counterparty's port"
        );
        require(
            hashString(msg_.packet.source_channel) == hashString(channel.counterparty.channel_id),
            "packet source channel doesn't match the counterparty's channel"
        );

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        require(
            msg_.packet.timeout_height.revision_height == 0 || block.number < msg_.packet.timeout_height.revision_height,
            "block height >= packet timeout height"
        );
        require(
            msg_.packet.timeout_timestamp == 0 || block.timestamp < msg_.packet.timeout_timestamp,
            "block timestamp >= packet timeout timestamp"
        );

        require(
            verifyPacketCommitment(
                connection,
                msg_.proofHeight,
                msg_.proof,
                IBCIdentifier.packetCommitmentPath(
                    msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence
                ),
                sha256(
                    abi.encodePacked(
                        msg_.packet.timeout_timestamp,
                        msg_.packet.timeout_height.revision_number,
                        msg_.packet.timeout_height.revision_height,
                        sha256(msg_.packet.data)
                    )
                )
            ),
            "failed to verify packet commitment"
        );

        if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            require(
                packetReceipts[msg_.packet.destination_port][msg_.packet.destination_channel][msg_.packet.sequence] == 0,
                "packet sequence already has been received"
            );
            packetReceipts[msg_.packet.destination_port][msg_.packet.destination_channel][msg_.packet.sequence] = 1;
        } else if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            require(
                nextSequenceRecvs[msg_.packet.destination_port][msg_.packet.destination_channel] == msg_.packet.sequence,
                "packet sequence != next receive sequence"
            );
            nextSequenceRecvs[msg_.packet.destination_port][msg_.packet.destination_channel]++;
        } else {
            revert("unknown ordering type");
        }
    }

    // WriteAcknowledgement writes the packet execution acknowledgement to the state,
    // which will be verified by the counterparty chain using AcknowledgePacket.
    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external {
        require(acknowledgement.length > 0, "acknowledgement cannot be empty");

        Channel.Data storage channel = channels[destinationPortId][destinationChannel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        bytes32 ackCommitmentKey = keccak256(
            IBCIdentifier.packetAcknowledgementCommitmentPath(destinationPortId, destinationChannel, sequence)
        );
        bytes32 ackCommitment = commitments[ackCommitmentKey];
        require(ackCommitment == bytes32(0), "acknowledgement for packet already exists");
        commitments[ackCommitmentKey] = keccak256(abi.encodePacked(sha256(acknowledgement)));
    }

    // TODO use calldata
    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.source_port][msg_.packet.source_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        require(
            hashString(msg_.packet.destination_port) == hashString(channel.counterparty.port_id),
            "packet destination port doesn't match the counterparty's port"
        );
        require(
            hashString(msg_.packet.destination_channel) == hashString(channel.counterparty.channel_id),
            "packet destination channel doesn't match the counterparty's channel"
        );

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        bytes32 packetCommitmentKey = keccak256(
            IBCIdentifier.packetCommitmentPath(
                msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence
            )
        );
        bytes32 packetCommitment = commitments[packetCommitmentKey];
        require(packetCommitment != bytes32(0), "packet commitment not found");
        require(
            packetCommitment
                == keccak256(
                    abi.encodePacked(
                        sha256(
                            abi.encodePacked(
                                msg_.packet.timeout_timestamp,
                                msg_.packet.timeout_height.revision_number,
                                msg_.packet.timeout_height.revision_height,
                                sha256(msg_.packet.data)
                            )
                        )
                    )
                ),
            "commitment bytes are not equal"
        );

        require(
            verifyPacketAcknowledgement(
                connection,
                msg_.proofHeight,
                msg_.proof,
                IBCIdentifier.packetAcknowledgementCommitmentPath(
                    msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence
                ),
                sha256(msg_.acknowledgement)
            ),
            "failed to verify packet acknowledgement commitment"
        );

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            require(
                msg_.packet.sequence == nextSequenceAcks[msg_.packet.source_port][msg_.packet.source_channel],
                "packet sequence != next ack sequence"
            );
            nextSequenceAcks[msg_.packet.source_port][msg_.packet.source_channel]++;
        }

        delete commitments[packetCommitmentKey];
    }

    /* Verification functions */

    function verifyChannelState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes
    ) private returns (bool) {
        return getClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.channelPath(portId, channelId),
            channelBytes
        );
    }

    function verifyPacketCommitment(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory proof,
        bytes memory path,
        bytes32 commitmentBytes
    ) private returns (bool) {
        return getClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            connection.delay_period,
            calcBlockDelay(connection.delay_period),
            proof,
            connection.counterparty.prefix.key_prefix,
            path,
            abi.encodePacked(commitmentBytes)
        );
    }

    function verifyPacketAcknowledgement(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory proof,
        bytes memory path,
        bytes32 acknowledgementCommitmentBytes
    ) private returns (bool) {
        return getClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            connection.delay_period,
            calcBlockDelay(connection.delay_period),
            proof,
            connection.counterparty.prefix.key_prefix,
            path,
            abi.encodePacked(acknowledgementCommitmentBytes)
        );
    }

    /* Internal functions */

    function getCounterpartyHops(string memory connectionId) internal view returns (string[] memory hops) {
        hops = new string[](1);
        hops[0] = connections[connectionId].counterparty.connection_id;
        return hops;
    }

    function calcBlockDelay(uint64 timeDelay) private view returns (uint64) {
        uint64 blockDelay = 0;
        if (expectedTimePerBlock != 0) {
            blockDelay = (timeDelay + expectedTimePerBlock - 1) / expectedTimePerBlock;
        }
        return blockDelay;
    }

    function hashString(string memory s) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(s));
    }
}
