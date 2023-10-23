// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";
import "../../proto/Client.sol";
import "../../proto/Channel.sol";
import "../25-handler/IBCMsgs.sol";
import "../02-client/IBCHeight.sol";
import "../24-host/IBCStore.sol";
import "../24-host/IBCCommitment.sol";
import "../04-channel/IIBCChannel.sol";

/**
 * @dev IBCPacket is a contract that implements [ICS-4](https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics).
 */
contract IBCPacket is IBCStore, IIBCPacket {
    using IBCHeight for Height.Data;

    bytes internal constant SUCCESSFUL_RECEIPT = hex"01";
    bytes32 internal constant HASHED_SUCCESSFUL_RECEIPT = keccak256(SUCCESSFUL_RECEIPT);

    /* Packet handlers */

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
    ) external returns (uint64) {
        Channel.Data storage channel = channels[sourcePort][sourceChannel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        {
            uint64 latestTimestamp;
            ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
            ILightClient client = ILightClient(clientImpls[connection.client_id]);
            require(address(client) != address(0), "client not found");
            require(client.getStatus(connection.client_id) == ClientStatus.Active, "client state is not active");

            require(!timeoutHeight.isZero() || timeoutTimestamp != 0, "timeout height and timestamp cannot both be 0");
            (Height.Data memory latestHeight, bool found) = client.getLatestHeight(connection.client_id);
            require(found, "clientState not found");
            require(
                timeoutHeight.isZero() || latestHeight.lt(timeoutHeight),
                "receiving chain block height >= packet timeout height"
            );
            (latestTimestamp, found) = client.getTimestampAtHeight(connection.client_id, latestHeight);
            require(found, "consensusState not found");
            require(
                timeoutTimestamp == 0 || latestTimestamp < timeoutTimestamp,
                "receiving chain block timestamp >= packet timeout timestamp"
            );
        }

        uint64 packetSequence = nextSequenceSends[sourcePort][sourceChannel];
        nextSequenceSends[sourcePort][sourceChannel] = packetSequence + 1;
        commitments[IBCCommitment.packetCommitmentKey(sourcePort, sourceChannel, packetSequence)] = keccak256(
            abi.encodePacked(
                sha256(
                    abi.encodePacked(
                        timeoutTimestamp, timeoutHeight.revision_number, timeoutHeight.revision_height, sha256(data)
                    )
                )
            )
        );
        return packetSequence;
    }

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.destination_port][msg_.packet.destination_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

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
            msg_.packet.timeout_timestamp == 0 || block.timestamp * 1e9 < msg_.packet.timeout_timestamp,
            "block timestamp >= packet timeout timestamp"
        );

        require(
            verifyPacketCommitment(
                connection,
                msg_.proofHeight,
                msg_.proof,
                IBCCommitment.packetCommitmentPath(
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
            bytes32 commitmentKey = IBCCommitment.packetReceiptCommitmentKey(
                msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence
            );
            require(commitments[commitmentKey] == bytes32(0), "packet receipt already exists");
            commitments[commitmentKey] = HASHED_SUCCESSFUL_RECEIPT;
        } else if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            require(
                nextSequenceRecvs[msg_.packet.destination_port][msg_.packet.destination_channel] == msg_.packet.sequence,
                "packet sequence != next receive sequence"
            );
            nextSequenceRecvs[msg_.packet.destination_port][msg_.packet.destination_channel]++;
            commitments[IBCCommitment.nextSequenceRecvCommitmentKey(
                msg_.packet.destination_port, msg_.packet.destination_channel
            )] = keccak256(
                uint64ToBigEndianBytes(nextSequenceRecvs[msg_.packet.destination_port][msg_.packet.destination_channel])
            );
        } else {
            revert("unknown ordering type");
        }
    }

    /**
     * @dev writeAcknowledgement writes the packet execution acknowledgement to the state,
     * which will be verified by the counterparty chain using AcknowledgePacket.
     */
    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external {
        require(acknowledgement.length > 0, "acknowledgement cannot be empty");

        Channel.Data storage channel = channels[destinationPortId][destinationChannel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        bytes32 ackCommitmentKey =
            IBCCommitment.packetAcknowledgementCommitmentKey(destinationPortId, destinationChannel, sequence);
        bytes32 ackCommitment = commitments[ackCommitmentKey];
        require(ackCommitment == bytes32(0), "acknowledgement for packet already exists");
        commitments[ackCommitmentKey] = keccak256(abi.encodePacked(sha256(acknowledgement)));
    }

    /**
     * @dev AcknowledgePacket is called by a module to process the acknowledgement of a
     * packet previously sent by the calling module on a channel to a counterparty
     * module on the counterparty chain. Its intended usage is within the ante
     * handler. AcknowledgePacket will clean up the packet commitment,
     * which is no longer necessary since the packet has been received and acted upon.
     * It will also increment NextSequenceAck in case of ORDERED channels.
     */
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

        bytes32 packetCommitmentKey =
            IBCCommitment.packetCommitmentKey(msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence);
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
                IBCCommitment.packetAcknowledgementCommitmentPath(
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

    function timeoutPacket(IBCMsgs.MsgTimeoutPacket calldata msg_) external {
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
    }

    function buildCounterparty(Packet.Data memory packet) private pure returns (ChannelCounterparty.Data memory) {
        return ChannelCounterparty.Data({port_id: packet.source_port, channel_id: packet.source_channel});
    }

    function timeoutOnClose(IBCMsgs.MsgTimeoutOnClose calldata msg_) external {
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
    }

    /* Verification functions */

    function verifyPacketCommitment(
        ConnectionEnd.Data storage connection,
        Height.Data calldata height,
        bytes calldata proof,
        bytes memory path,
        bytes32 commitmentBytes
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
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
        Height.Data calldata height,
        bytes calldata proof,
        bytes memory path,
        bytes32 acknowledgementCommitmentBytes
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
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

    function hashString(string memory s) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(s));
    }

    function uint64ToBigEndianBytes(uint64 v) private pure returns (bytes memory) {
        return abi.encodePacked(bytes8(v));
    }
}
