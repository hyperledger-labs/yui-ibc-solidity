// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel} from "../../proto/Channel.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCClientErrors} from "../02-client/IIBCClientErrors.sol";
import {IBCHeight} from "../02-client/IBCHeight.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IBCChannelLib} from "./IBCChannelLib.sol";
import {IIBCChannelPacketSendRecv} from "./IIBCChannel.sol";
import {IIBCChannelErrors} from "./IIBCChannelErrors.sol";

/**
 * @dev IBCChannelPacketSendRecv is a contract that implements [ICS-4](https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics).
 */
contract IBCChannelPacketSendRecv is
    IBCModuleManager,
    IIBCChannelPacketSendRecv,
    IIBCChannelErrors,
    IIBCClientErrors
{
    using IBCHeight for Height.Data;

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
        authenticateChannelCapability(sourcePort, sourceChannel);

        Channel.Data storage channel = channels[sourcePort][sourceChannel];
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        if (timeoutHeight.isZero() && timeoutTimestamp == 0) {
            revert IBCChannelZeroPacketTimeout();
        }

        {
            // NOTE: We can assume here that the connection state is OPEN because the channel state is OPEN
            ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];
            (Height.Data memory latestHeight, uint64 latestTimestamp, ILightClient.ClientStatus status) =
                ILightClient(clientImpls[connection.client_id]).getLatestInfo(connection.client_id);
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

        uint64 packetSequence = nextSequenceSends[sourcePort][sourceChannel];
        nextSequenceSends[sourcePort][sourceChannel] = packetSequence + 1;
        commitments[IBCCommitment.packetCommitmentKeyCalldata(sourcePort, sourceChannel, packetSequence)] = keccak256(
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
    ) public {
        authenticateChannelCapability(destinationPortId, destinationChannel);
        Channel.Data storage channel = channels[destinationPortId][destinationChannel];
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        if (acknowledgement.length == 0) {
            revert IBCChannelEmptyAcknowledgement();
        }
        _writeAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function _writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes memory acknowledgement
    ) internal {
        bytes32 ackCommitmentKey =
            IBCCommitment.packetAcknowledgementCommitmentKeyCalldata(destinationPortId, destinationChannel, sequence);
        if (commitments[ackCommitmentKey] != bytes32(0)) {
            revert IBCChannelAcknowledgementAlreadyWritten(destinationPortId, destinationChannel, sequence);
        }
        commitments[ackCommitmentKey] = keccak256(abi.encodePacked(sha256(acknowledgement)));
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    /**
     * @dev recvPacket is called by a module in order to receive & process an IBC packet
     * sent on the corresponding channel end on the counterparty chain.
     */
    function recvPacket(MsgPacketRecv calldata msg_) external {
        Channel.Data storage channel = channels[msg_.packet.destinationPort][msg_.packet.destinationChannel];
        if (channel.state == Channel.State.STATE_OPEN) {} else if (
            channel.state == Channel.State.STATE_FLUSHING || channel.state == Channel.State.STATE_FLUSHCOMPLETE
        ) {
            RecvStartSequence storage rseq =
                recvStartSequences[msg_.packet.destinationPort][msg_.packet.destinationChannel];
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

        if (msg_.packet.timeoutHeight.revision_height != 0 && block.number >= msg_.packet.timeoutHeight.revision_height)
        {
            revert IBCChannelTimeoutPacketHeight(block.number, msg_.packet.timeoutHeight.revision_height);
        }
        if (msg_.packet.timeoutTimestamp != 0 && block.timestamp * 1e9 >= msg_.packet.timeoutTimestamp) {
            revert IBCChannelTimeoutPacketTimestamp(block.timestamp * 1e9, msg_.packet.timeoutTimestamp);
        }

        verifyPacketCommitment(
            connections[channel.connection_hops[0]],
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
                if (
                    msg_.packet.sequence
                        < recvStartSequences[msg_.packet.destinationPort][msg_.packet.destinationChannel].sequence
                ) {
                    revert IBCChannelPacketAlreadyProcessInPreviousUpgrade(
                        msg_.packet.sequence,
                        recvStartSequences[msg_.packet.destinationPort][msg_.packet.destinationChannel].sequence
                    );
                }
            }
            bytes32 commitmentKey = IBCCommitment.packetReceiptCommitmentKeyCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
            );
            if (commitments[commitmentKey] != bytes32(0)) {
                revert IBCChannelPacketReceiptAlreadyExists(
                    msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
                );
            }
            commitments[commitmentKey] = IBCChannelLib.PACKET_RECEIPT_SUCCESSFUL_KECCAK256;
        } else if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            if (nextSequenceRecvs[msg_.packet.destinationPort][msg_.packet.destinationChannel] != msg_.packet.sequence)
            {
                revert IBCChannelUnexpectedNextSequenceRecv(
                    nextSequenceRecvs[msg_.packet.destinationPort][msg_.packet.destinationChannel]
                );
            }
            nextSequenceRecvs[msg_.packet.destinationPort][msg_.packet.destinationChannel]++;
            commitments[IBCCommitment.nextSequenceRecvCommitmentKeyCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel
            )] = keccak256(
                uint64ToBigEndianBytes(nextSequenceRecvs[msg_.packet.destinationPort][msg_.packet.destinationChannel])
            );
        } else {
            revert IBCChannelUnknownChannelOrder(channel.ordering);
        }
        bytes memory acknowledgement = lookupModuleByChannel(
            msg_.packet.destinationPort, msg_.packet.destinationChannel
        ).onRecvPacket(msg_.packet, _msgSender());
        if (acknowledgement.length > 0) {
            _writeAcknowledgement(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence, acknowledgement
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
        Channel.Data storage channel = channels[msg_.packet.sourcePort][msg_.packet.sourceChannel];
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

        // NOTE: We can assume here that the connection state is OPEN because the channel state is OPEN
        ConnectionEnd.Data storage connection = connections[channel.connection_hops[0]];

        bytes32 packetCommitmentKey = IBCCommitment.packetCommitmentKeyCalldata(
            msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
        );
        bytes32 packetCommitment = commitments[packetCommitmentKey];
        if (packetCommitment == bytes32(0)) {
            revert IBCChannelPacketCommitmentNotFound(
                msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
            );
        }
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

        verifyPacketAcknowledgement(
            connection,
            msg_.proofHeight,
            msg_.proof,
            IBCCommitment.packetAcknowledgementCommitmentPathCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
            ),
            sha256(msg_.acknowledgement)
        );

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            if (msg_.packet.sequence != nextSequenceAcks[msg_.packet.sourcePort][msg_.packet.sourceChannel]) {
                revert IBCChannelUnexpectedNextSequenceAck(
                    nextSequenceAcks[msg_.packet.sourcePort][msg_.packet.sourceChannel]
                );
            }
            nextSequenceAcks[msg_.packet.sourcePort][msg_.packet.sourceChannel]++;
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            if (msg_.packet.sequence < ackStartSequences[msg_.packet.sourcePort][msg_.packet.sourceChannel]) {
                revert IBCChannelAckAlreadyProcessedInPreviousUpgrade(
                    msg_.packet.sequence, ackStartSequences[msg_.packet.sourcePort][msg_.packet.sourceChannel]
                );
            }
        } else {
            revert IBCChannelUnknownChannelOrder(channel.ordering);
        }

        delete commitments[packetCommitmentKey];
        lookupModuleByChannel(msg_.packet.sourcePort, msg_.packet.sourceChannel).onAcknowledgementPacket(
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
        bytes32 commitment
    ) private {
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
            return;
        }
        revert IBCChannelFailedVerifyPacketAcknowledgement(
            connection.client_id, path, acknowledgementCommitment, proof, height
        );
    }

    // private functions

    function calcBlockDelay(uint64 timeDelay) private view returns (uint64) {
        if (timeDelay == 0) {
            return 0;
        } else if (expectedTimePerBlock == 0) {
            return 0;
        } else {
            return (timeDelay + expectedTimePerBlock - 1) / expectedTimePerBlock;
        }
    }

    function uint64ToBigEndianBytes(uint64 v) private pure returns (bytes memory) {
        return abi.encodePacked(bytes8(v));
    }
}
