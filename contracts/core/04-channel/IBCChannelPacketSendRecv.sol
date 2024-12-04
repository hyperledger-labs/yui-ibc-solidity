// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, Timeout} from "../../proto/Channel.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCClientErrors} from "../02-client/IIBCClientErrors.sol";
import {IBCHeight} from "../02-client/IBCHeight.sol";
import {IBCChannelUpgradeBase} from "./IBCChannelUpgrade.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCChannelLib} from "./IBCChannelLib.sol";
import {IIBCChannelPacketSendRecv} from "./IIBCChannel.sol";
import {IIBCChannelErrors} from "./IIBCChannelErrors.sol";

/**
 * @dev IBCChannelPacketSendRecv is a contract that implements [ICS-4](https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics).
 */
contract IBCChannelPacketSendRecv is
    IBCChannelUpgradeBase,
    IIBCChannelPacketSendRecv,
    IIBCChannelErrors,
    IIBCClientErrors
{
    using IBCHeight for Height.Data;

    // --------- IIBCChannelPacketSendRecv Implementation --------- //

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
    ) public override returns (uint64) {
        authenticateChannelCapability(sourcePort, sourceChannel);

        ChannelStorage storage channelStorage = getChannelStorage()[sourcePort][sourceChannel];
        Channel.Data storage channel = channelStorage.channel;
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        if (timeoutHeight.isZero() && timeoutTimestamp == 0) {
            revert IBCChannelZeroPacketTimeout();
        }

        {
            // NOTE: We can assume here that the connection state is OPEN because the channel state is OPEN
            ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
            (Height.Data memory latestHeight, uint64 latestTimestamp, ILightClient.ClientStatus status) =
                ILightClient(getClientStorage()[connection.client_id].clientImpl).getLatestInfo(connection.client_id);
            if (status != ILightClient.ClientStatus.Active) {
                revert IBCClientNotActiveClient(connection.client_id);
            }
            if (!timeoutHeight.isZero() && latestHeight.gte(timeoutHeight)) {
                revert IBCChannelPastPacketTimeoutHeight(timeoutHeight, latestHeight);
            }
            if (timeoutTimestamp != 0 && latestTimestamp >= timeoutTimestamp) {
                revert IBCChannelPastPacketTimeoutTimestamp(timeoutTimestamp, latestTimestamp);
            }
        }

        uint64 packetSequence = channelStorage.nextSequenceSend;
        channelStorage.nextSequenceSend = packetSequence + 1;
        getCommitments()[IBCCommitment.packetCommitmentKeyCalldata(sourcePort, sourceChannel, packetSequence)] =
        keccak256(
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
        bytes calldata acknowledgement
    ) public override {
        authenticateChannelCapability(destinationPortId, destinationChannel);
        Channel.Data storage channel = getChannelStorage()[destinationPortId][destinationChannel].channel;
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        if (acknowledgement.length == 0) {
            revert IBCChannelEmptyAcknowledgement();
        }
        _writeAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(MsgPacketRecv calldata msg_) public override {
        ChannelStorage storage channelStorage =
            getChannelStorage()[msg_.packet.destinationPort][msg_.packet.destinationChannel];
        Channel.Data storage channel = channelStorage.channel;
        if (channel.state == Channel.State.STATE_OPEN) {} else if (
            channel.state == Channel.State.STATE_FLUSHING || channel.state == Channel.State.STATE_FLUSHCOMPLETE
        ) {
            RecvStartSequence storage rseq = channelStorage.recvStartSequence;
            // prevSequence=0 means the channel is not in the process of being upgraded or counterparty has not been upgraded yet
            if (rseq.prevSequence != 0) {
                if (msg_.packet.sequence >= rseq.sequence) {
                    revert IBCChannelCannotRecvNextUpgradePacket(msg_.packet.sequence, rseq.sequence);
                } else if (msg_.packet.sequence < rseq.prevSequence) {
                    revert IBCChannelPacketAlreadyProcessInPreviousUpgrade(msg_.packet.sequence, rseq.prevSequence);
                }
            }
        } else {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }

        if (keccak256(bytes(msg_.packet.sourcePort)) != keccak256(bytes(channel.counterparty.port_id))) {
            revert IBCChannelUnexpectedPacketSource(msg_.packet.sourcePort, msg_.packet.sourceChannel);
        } else if (keccak256(bytes(msg_.packet.sourceChannel)) != keccak256(bytes(channel.counterparty.channel_id))) {
            revert IBCChannelUnexpectedPacketSource(msg_.packet.sourcePort, msg_.packet.sourceChannel);
        }

        if (!msg_.packet.timeoutHeight.isZero() && hostHeight().gte(msg_.packet.timeoutHeight)) {
            revert IBCChannelTimeoutPacketHeight(hostHeight(), msg_.packet.timeoutHeight);
        }
        if (msg_.packet.timeoutTimestamp != 0 && hostTimestamp() >= msg_.packet.timeoutTimestamp) {
            revert IBCChannelTimeoutPacketTimestamp(hostTimestamp(), msg_.packet.timeoutTimestamp);
        }

        verifyPacketCommitment(
            getConnectionStorage()[channel.connection_hops[0]].connection,
            msg_.proofHeight,
            msg_.proof,
            IBCCommitment.packetCommitmentPathCalldata(
                msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
            ),
            sha256(
                abi.encodePacked(
                    msg_.packet.timeoutTimestamp,
                    msg_.packet.timeoutHeight.revision_number,
                    msg_.packet.timeoutHeight.revision_height,
                    sha256(msg_.packet.data)
                )
            )
        );

        if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            if (channel.state == Channel.State.STATE_OPEN) {
                RecvStartSequence storage rseq = channelStorage.recvStartSequence;
                if (msg_.packet.sequence < rseq.sequence) {
                    revert IBCChannelPacketAlreadyProcessInPreviousUpgrade(msg_.packet.sequence, rseq.sequence);
                }
            }
            bytes32 commitmentKey = IBCCommitment.packetReceiptCommitmentKeyCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
            );
            mapping(bytes32 => bytes32) storage commitments = getCommitments();
            if (commitments[commitmentKey] != bytes32(0)) {
                revert IBCChannelPacketReceiptAlreadyExists(
                    msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
                );
            }
            commitments[commitmentKey] = IBCChannelLib.PACKET_RECEIPT_SUCCESSFUL_KECCAK256;
        } else if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            if (channelStorage.nextSequenceRecv != msg_.packet.sequence) {
                revert IBCChannelUnexpectedNextSequenceRecv(
                    msg_.packet.destinationPort,
                    msg_.packet.destinationChannel,
                    msg_.packet.sequence,
                    channelStorage.nextSequenceRecv
                );
            }
            channelStorage.nextSequenceRecv++;
            getCommitments()[IBCCommitment.nextSequenceRecvCommitmentKeyCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel
            )] = keccak256(IBCChannelLib.uint64ToBigEndianBytes(channelStorage.nextSequenceRecv));
        } else {
            revert IBCChannelUnknownChannelOrder(channel.ordering);
        }
        emit RecvPacket(msg_.packet);
        bytes memory acknowledgement = lookupModuleByChannel(
            msg_.packet.destinationPort, msg_.packet.destinationChannel
        ).onRecvPacket(msg_.packet, _msgSender());
        if (acknowledgement.length > 0) {
            _writeAcknowledgement(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence, acknowledgement
            );
        }
    }

    /**
     * @dev AcknowledgePacket is called by a module to process the acknowledgement of a
     * packet previously sent by the calling module on a channel to a counterparty
     * module on the counterparty chain. Its intended usage is within the ante
     * handler. AcknowledgePacket will clean up the packet commitment,
     * which is no longer necessary since the packet has been received and acted upon.
     * It will also increment NextSequenceAck in case of ORDERED channels.
     */
    function acknowledgePacket(MsgPacketAcknowledgement calldata msg_) public override {
        ChannelStorage storage channelStorage = getChannelStorage()[msg_.packet.sourcePort][msg_.packet.sourceChannel];
        Channel.Data storage channel = channelStorage.channel;
        if (channel.state != Channel.State.STATE_OPEN) {
            if (channel.state != Channel.State.STATE_FLUSHING) {
                revert IBCChannelUnexpectedChannelState(channel.state);
            }
        }

        if (keccak256(bytes(msg_.packet.destinationPort)) != keccak256(bytes(channel.counterparty.port_id))) {
            revert IBCChannelUnexpectedPacketDestination(msg_.packet.destinationPort, msg_.packet.destinationChannel);
        } else if (
            keccak256(bytes(msg_.packet.destinationChannel)) != keccak256(bytes(channel.counterparty.channel_id))
        ) {
            revert IBCChannelUnexpectedPacketDestination(msg_.packet.destinationPort, msg_.packet.destinationChannel);
        }

        bytes32 packetCommitmentKey = IBCCommitment.packetCommitmentKeyCalldata(
            msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
        );
        mapping(bytes32 => bytes32) storage commitments = getCommitments();
        bytes32 packetCommitment = commitments[packetCommitmentKey];
        if (packetCommitment == bytes32(0)) {
            revert IBCChannelPacketCommitmentNotFound(
                msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
            );
        }
        {
            bytes32 commitment = keccak256(
                abi.encodePacked(
                    sha256(
                        abi.encodePacked(
                            msg_.packet.timeoutTimestamp,
                            msg_.packet.timeoutHeight.revision_number,
                            msg_.packet.timeoutHeight.revision_height,
                            sha256(msg_.packet.data)
                        )
                    )
                )
            );
            if (packetCommitment != commitment) {
                revert IBCChannelPacketCommitmentMismatch(packetCommitment, commitment);
            }
        }

        verifyPacketAcknowledgement(
            // NOTE: We can assume here that the connection state is OPEN because the channel state is OPEN
            getConnectionStorage()[channel.connection_hops[0]].connection,
            msg_.proofHeight,
            msg_.proof,
            IBCCommitment.packetAcknowledgementCommitmentPathCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
            ),
            sha256(msg_.acknowledgement)
        );

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            if (msg_.packet.sequence != channelStorage.nextSequenceAck) {
                revert IBCChannelUnexpectedNextSequenceAck(channelStorage.nextSequenceAck);
            }
            channelStorage.nextSequenceAck++;
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            if (msg_.packet.sequence < channelStorage.ackStartSequence) {
                revert IBCChannelAckAlreadyProcessedInPreviousUpgrade(
                    msg_.packet.sequence, channelStorage.ackStartSequence
                );
            }
        } else {
            revert IBCChannelUnknownChannelOrder(channel.ordering);
        }

        delete commitments[packetCommitmentKey];

        if (channel.state == Channel.State.STATE_FLUSHING) {
            Timeout.Data memory timeout = channelStorage.counterpartyUpgradeTimeout;
            if (!timeout.height.isZero() || timeout.timestamp != 0) {
                if (
                    !timeout.height.isZero() && hostHeight().gte(timeout.height)
                        || timeout.timestamp != 0 && hostTimestamp() >= timeout.timestamp
                ) {
                    restoreChannel(msg_.packet.sourcePort, msg_.packet.sourceChannel, UpgradeHandshakeError.Timeout);
                } else if (
                    canTransitionToFlushComplete(
                        channel.ordering, msg_.packet.sourcePort, msg_.packet.sourceChannel, channel.upgrade_sequence
                    )
                ) {
                    channel.state = Channel.State.STATE_FLUSHCOMPLETE;
                    updateChannelCommitment(msg_.packet.sourcePort, msg_.packet.sourceChannel, channel);
                }
            }
        }

        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
        lookupModuleByChannel(msg_.packet.sourcePort, msg_.packet.sourceChannel).onAcknowledgementPacket(
            msg_.packet, msg_.acknowledgement, _msgSender()
        );
    }

    // --------- Private Functions --------- //

    function _writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes memory acknowledgement
    ) private {
        bytes32 ackCommitmentKey =
            IBCCommitment.packetAcknowledgementCommitmentKeyCalldata(destinationPortId, destinationChannel, sequence);
        if (getCommitments()[ackCommitmentKey] != bytes32(0)) {
            revert IBCChannelAcknowledgementAlreadyWritten(destinationPortId, destinationChannel, sequence);
        }
        getCommitments()[ackCommitmentKey] = keccak256(abi.encodePacked(sha256(acknowledgement)));
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    /**
     * @dev calcBlockDelay calculates the block delay based on the expected time per block
     */
    function calcBlockDelay(uint64 timeDelay) private view returns (uint64) {
        HostStorage storage hostStorage = getHostStorage();
        if (timeDelay == 0) {
            return 0;
        } else if (hostStorage.expectedTimePerBlock == 0) {
            return 0;
        } else {
            return (timeDelay + hostStorage.expectedTimePerBlock - 1) / hostStorage.expectedTimePerBlock;
        }
    }

    function verifyPacketCommitment(
        ConnectionEnd.Data storage connection,
        Height.Data calldata height,
        bytes calldata proof,
        bytes memory path,
        bytes32 commitment
    ) private {
        // slither-disable-start reentrancy-no-eth
        if (
            checkAndGetClient(connection.client_id).verifyMembership(
                connection.client_id,
                height,
                connection.delay_period,
                calcBlockDelay(connection.delay_period),
                proof,
                connection.counterparty.prefix.key_prefix,
                path,
                abi.encodePacked(commitment)
            )
        ) {
            // slither-disable-end reentrancy-no-eth
            return;
        }
        revert IBCChannelFailedVerifyPacketCommitment(connection.client_id, path, commitment, proof, height);
    }

    function verifyPacketAcknowledgement(
        ConnectionEnd.Data storage connection,
        Height.Data calldata height,
        bytes calldata proof,
        bytes memory path,
        bytes32 acknowledgementCommitment
    ) private {
        // slither-disable-start reentrancy-no-eth
        if (
            checkAndGetClient(connection.client_id).verifyMembership(
                connection.client_id,
                height,
                connection.delay_period,
                calcBlockDelay(connection.delay_period),
                proof,
                connection.counterparty.prefix.key_prefix,
                path,
                abi.encodePacked(acknowledgementCommitment)
            )
        ) {
            // slither-disable-end reentrancy-no-eth
            return;
        }
        revert IBCChannelFailedVerifyPacketAcknowledgement(
            connection.client_id, path, acknowledgementCommitment, proof, height
        );
    }
}
