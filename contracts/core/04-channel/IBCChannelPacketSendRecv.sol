// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel} from "../../proto/Channel.sol";
import {ILightClient, ClientStatus} from "../02-client/ILightClient.sol";
import {IBCHeight} from "../02-client/IBCHeight.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IBCChannelLib} from "./IBCChannelLib.sol";
import {IIBCChannelPacketSendRecv} from "./IIBCChannel.sol";

/**
 * @dev IBCChannelPacketSendRecv is a contract that implements [ICS-4](https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics).
 */
contract IBCChannelPacketSendRecv is IBCModuleManager, IIBCChannelPacketSendRecv {
    using IBCHeight for Height.Data;

    bytes32 internal constant SUCCESSFUL_RECEIPT_COMMITMENT =
        keccak256(abi.encodePacked(IBCChannelLib.PacketReceipt.SUCCESSFUL));

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
        require(authenticateCapability(channelCapabilityPath(sourcePort, sourceChannel)), "unauthorized");

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
        emit SendPacket(packetSequence, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data);
        return packetSequence;
    }

    /**
     * @dev writeAcknowledgement writes the packet execution acknowledgement to the state,
     * which will be verified by the counterparty chain using AcknowledgePacket.
     */
    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes memory acknowledgement
    ) public {
        require(authenticateCapability(channelCapabilityPath(destinationPortId, destinationChannel)), "unauthorized");
        _writeAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function _writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes memory acknowledgement
    ) internal {
        require(acknowledgement.length > 0, "acknowledgement cannot be empty");

        Channel.Data storage channel = channels[destinationPortId][destinationChannel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        bytes32 ackCommitmentKey =
            IBCCommitment.packetAcknowledgementCommitmentKey(destinationPortId, destinationChannel, sequence);
        bytes32 ackCommitment = commitments[ackCommitmentKey];
        require(ackCommitment == bytes32(0), "acknowledgement for packet already exists");
        commitments[ackCommitmentKey] = keccak256(abi.encodePacked(sha256(acknowledgement)));
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(MsgPacketRecv calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.destination_port][msg_.packet.destination_channel];
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        require(
            keccak256(bytes(msg_.packet.source_port)) == keccak256(bytes(channel.counterparty.port_id)),
            "packet source port doesn't match the counterparty's port"
        );
        require(
            keccak256(bytes(msg_.packet.source_channel)) == keccak256(bytes(channel.counterparty.channel_id)),
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
            commitments[commitmentKey] = SUCCESSFUL_RECEIPT_COMMITMENT;
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
        bytes memory acknowledgement = lookupModuleByChannel(
            msg_.packet.destination_port, msg_.packet.destination_channel
        ).onRecvPacket(msg_.packet, _msgSender());
        if (acknowledgement.length > 0) {
            _writeAcknowledgement(
                msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement
            );
        }
        emit RecvPacket(msg_.packet);
    }

    /**
     * @dev AcknowledgePacket is called by a module to process the acknowledgement of a
     * packet previously sent by the calling module on a channel to a counterparty
     * module on the counterparty chain. Its intended usage is within the ante
     * handler. AcknowledgePacket will clean up the packet commitment,
     * which is no longer necessary since the packet has been received and acted upon.
     * It will also increment NextSequenceAck in case of ORDERED channels.
     */
    function acknowledgePacket(MsgPacketAcknowledgement calldata msg_) external {
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
        lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel).onAcknowledgementPacket(
            msg_.packet, msg_.acknowledgement, _msgSender()
        );
        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
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
