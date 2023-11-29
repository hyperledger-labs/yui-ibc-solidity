// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, ChannelCounterparty} from "../../proto/Channel.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCHeight} from "../02-client/IBCHeight.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCChannelPacketTimeout} from "./IIBCChannel.sol";

contract IBCChannelPacketTimeout is IBCModuleManager, IIBCChannelPacketTimeout {
    using IBCHeight for Height.Data;

    function timeoutPacket(MsgTimeoutPacket calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.source_port][msg_.packet.source_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        require(
            keccak256(bytes(msg_.packet.destination_port)) == keccak256(bytes(channel.counterparty.port_id)),
            "packet destination port doesn't match the counterparty's port"
        );
        require(
            keccak256(bytes(msg_.packet.destination_channel)) == keccak256(bytes(channel.counterparty.channel_id)),
            "packet destination channel doesn't match the counterparty's channel"
        );

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(bytes(connection.client_id).length != 0, "connection not found");
        ILightClient client = ILightClient(clientImpls[connection.client_id]);
        require(address(client) != address(0), "client not found");
        {
            uint64 proofTimestamp;
            (Height.Data memory latestHeight, bool found) = client.getLatestHeight(connection.client_id);
            require(found, "clientState not found");
            (proofTimestamp, found) = client.getTimestampAtHeight(connection.client_id, latestHeight);
            require(found, "consensusState not found");
            if (
                (msg_.packet.timeout_height.isZero() || msg_.proofHeight.lt(msg_.packet.timeout_height))
                    && (msg_.packet.timeout_timestamp == 0 || proofTimestamp < msg_.packet.timeout_timestamp)
            ) {
                revert("packet timeout has not been reached for height or timestamp");
            }
        }

        {
            bytes32 commitment = commitments[IBCCommitment.packetCommitmentKey(
                msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence
            )];
            // NOTE: if false, this indicates that the timeoutPacket already been executed
            require(commitment != bytes32(0), "packet commitment not found");
            require(
                commitment
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
        }

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            // check that packet has not been received
            require(msg_.nextSequenceRecv <= msg_.packet.sequence, "packet sequence > next receive sequence");
            require(
                client.verifyMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proof,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.nextSequenceRecvCommitmentPath(
                        msg_.packet.destination_port, msg_.packet.destination_channel
                    ),
                    uint64ToBigEndianBytes(msg_.nextSequenceRecv)
                ),
                "failed to verify next sequence receive"
            );
            channel.state = Channel.State.STATE_CLOSED;
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            require(
                client.verifyNonMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proof,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.packetReceiptCommitmentPath(
                        msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence
                    )
                ),
                "failed to verify packet receipt absense"
            );
        } else {
            revert("unknown ordering type");
        }

        delete commitments[IBCCommitment.packetCommitmentKey(
            msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence
        )];

        lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel).onTimeoutPacket(
            msg_.packet, _msgSender()
        );
        emit TimeoutPacket(msg_.packet);
    }

    function timeoutOnClose(MsgTimeoutOnClose calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.source_port][msg_.packet.source_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        require(
            keccak256(bytes(msg_.packet.destination_port)) == keccak256(bytes(channel.counterparty.port_id)),
            "packet destination port doesn't match the counterparty's port"
        );
        require(
            keccak256(bytes(msg_.packet.destination_channel)) == keccak256(bytes(channel.counterparty.channel_id)),
            "packet destination channel doesn't match the counterparty's channel"
        );

        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
        require(bytes(connection.client_id).length != 0, "connection not found");
        ILightClient client = ILightClient(clientImpls[connection.client_id]);
        require(address(client) != address(0), "client not found");

        {
            bytes32 commitment = commitments[IBCCommitment.packetCommitmentKey(
                msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence
            )];
            // NOTE: if false, this indicates that the timeoutPacket already been executed
            require(commitment != bytes32(0), "packet commitment not found");
            require(
                commitment
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
        }

        {
            Channel.Data memory expectedChannel = Channel.Data({
                state: Channel.State.STATE_CLOSED,
                ordering: channel.ordering,
                counterparty: ChannelCounterparty.Data({
                    port_id: msg_.packet.source_port,
                    channel_id: msg_.packet.source_channel
                }),
                connection_hops: buildConnectionHops(connection.counterparty.connection_id),
                version: channel.version
            });
            require(
                client.verifyMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proofClose,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.channelPath(msg_.packet.destination_port, msg_.packet.destination_channel),
                    Channel.encode(expectedChannel)
                ),
                "failed to verify channel state"
            );
        }

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            // check that packet has not been received
            require(msg_.nextSequenceRecv <= msg_.packet.sequence, "packet sequence > next receive sequence");
            require(
                client.verifyMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proofUnreceived,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.nextSequenceRecvCommitmentPath(
                        msg_.packet.destination_port, msg_.packet.destination_channel
                    ),
                    uint64ToBigEndianBytes(msg_.nextSequenceRecv)
                ),
                "failed to verify next sequence receive"
            );
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            require(
                client.verifyNonMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proofUnreceived,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.packetReceiptCommitmentPath(
                        msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence
                    )
                ),
                "failed to verify packet receipt absense"
            );
        } else {
            revert("unknown ordering type");
        }
        lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel).onTimeoutPacket(
            msg_.packet, _msgSender()
        );
        emit TimeoutPacket(msg_.packet);
    }

    // private functions

    function calcBlockDelay(uint64 timeDelay) private view returns (uint64) {
        uint64 blockDelay = 0;
        if (expectedTimePerBlock != 0) {
            blockDelay = (timeDelay + expectedTimePerBlock - 1) / expectedTimePerBlock;
        }
        return blockDelay;
    }

    function buildConnectionHops(string memory connectionId) private pure returns (string[] memory hops) {
        hops = new string[](1);
        hops[0] = connectionId;
        return hops;
    }

    function uint64ToBigEndianBytes(uint64 v) private pure returns (bytes memory) {
        return abi.encodePacked(bytes8(v));
    }
}
