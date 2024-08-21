// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, ChannelCounterparty} from "../../proto/Channel.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCHeight} from "../02-client/IBCHeight.sol";
import {IBCChannelLib} from "./IBCChannelLib.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCChannelPacketTimeout} from "./IIBCChannel.sol";
import {IIBCChannelErrors} from "./IIBCChannelErrors.sol";

contract IBCChannelPacketTimeout is IBCModuleManager, IIBCChannelPacketTimeout, IIBCChannelErrors {
    using IBCHeight for Height.Data;

    function timeoutPacket(MsgTimeoutPacket calldata msg_) external {
        Channel.Data storage channel = getChannelStorage()[msg_.packet.sourcePort][msg_.packet.sourceChannel].channel;
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }

        if (keccak256(bytes(msg_.packet.destinationPort)) != keccak256(bytes(channel.counterparty.port_id))) {
            revert IBCChannelUnexpectedPacketDestination(msg_.packet.destinationPort, msg_.packet.destinationChannel);
        } else if (
            keccak256(bytes(msg_.packet.destinationChannel)) != keccak256(bytes(channel.counterparty.channel_id))
        ) {
            revert IBCChannelUnexpectedPacketDestination(msg_.packet.destinationPort, msg_.packet.destinationChannel);
        }

        // NOTE: we can assume here that the connection exists because the channel is open
        ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
        ILightClient client = ILightClient(getClientStorage()[connection.client_id].clientImpl);

        if (msg_.packet.timeoutHeight.isZero() || msg_.proofHeight.lt(msg_.packet.timeoutHeight)) {
            if (
                msg_.packet.timeoutTimestamp == 0
                    || client.getTimestampAtHeight(connection.client_id, msg_.proofHeight) < msg_.packet.timeoutTimestamp
            ) {
                revert IBCChannelTimeoutNotReached();
            }
        }

        {
            bytes32 commitment = getCommitments()[IBCCommitment.packetCommitmentKeyCalldata(
                msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
            )];
            // NOTE: if false, this indicates that the timeoutPacket already been executed
            if (commitment == bytes32(0)) {
                revert IBCChannelPacketCommitmentNotFound(
                    msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
                );
            }
            bytes32 packetCommitment = keccak256(
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
            if (commitment != packetCommitment) {
                revert IBCChannelPacketCommitmentMismatch(commitment, packetCommitment);
            }
        }

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            // check that packet has not been received
            if (msg_.packet.sequence < msg_.nextSequenceRecv) {
                revert IBCChannelPacketMaybeAlreadyReceived(msg_.packet.sequence, msg_.nextSequenceRecv);
            }
            if (
                // slither-disable-next-line reentrancy-no-eth
                !client.verifyMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proof,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.nextSequenceRecvCommitmentPath(
                        msg_.packet.destinationPort, msg_.packet.destinationChannel
                    ),
                    IBCChannelLib.uint64ToBigEndianBytes(msg_.nextSequenceRecv)
                )
            ) {
                revert IBCChannelFailedVerifyNextSequenceRecv(
                    connection.client_id,
                    IBCCommitment.nextSequenceRecvCommitmentPath(
                        msg_.packet.destinationPort, msg_.packet.destinationChannel
                    ),
                    msg_.nextSequenceRecv,
                    msg_.proof,
                    msg_.proofHeight
                );
            }
            channel.state = Channel.State.STATE_CLOSED;
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            bytes memory path = IBCCommitment.packetReceiptCommitmentPathCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
            );
            if (
                // slither-disable-next-line reentrancy-no-eth
                !client.verifyNonMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proof,
                    connection.counterparty.prefix.key_prefix,
                    path
                )
            ) {
                revert IBCChannelFailedVerifyPacketReceiptAbsence(
                    connection.client_id, path, msg_.proof, msg_.proofHeight
                );
            }
        } else {
            revert IBCChannelUnknownChannelOrder(channel.ordering);
        }

        delete getCommitments()[IBCCommitment.packetCommitmentKeyCalldata(
            msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
        )];

        lookupModuleByChannel(msg_.packet.sourcePort, msg_.packet.sourceChannel).onTimeoutPacket(
            msg_.packet, _msgSender()
        );
        emit TimeoutPacket(msg_.packet);
    }

    function timeoutOnClose(MsgTimeoutOnClose calldata msg_) external {
        Channel.Data storage channel = getChannelStorage()[msg_.packet.sourcePort][msg_.packet.sourceChannel].channel;
        if (channel.state != Channel.State.STATE_OPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }

        if (keccak256(bytes(msg_.packet.destinationPort)) != keccak256(bytes(channel.counterparty.port_id))) {
            revert IBCChannelUnexpectedPacketDestination(msg_.packet.destinationPort, msg_.packet.destinationChannel);
        } else if (
            keccak256(bytes(msg_.packet.destinationChannel)) != keccak256(bytes(channel.counterparty.channel_id))
        ) {
            revert IBCChannelUnexpectedPacketDestination(msg_.packet.destinationPort, msg_.packet.destinationChannel);
        }

        ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
        ILightClient client = ILightClient(getClientStorage()[connection.client_id].clientImpl);
        {
            bytes32 commitment = getCommitments()[IBCCommitment.packetCommitmentKeyCalldata(
                msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
            )];
            // NOTE: if false, this indicates that the timeoutPacket already been executed
            if (commitment == bytes32(0)) {
                revert IBCChannelPacketCommitmentNotFound(
                    msg_.packet.sourcePort, msg_.packet.sourceChannel, msg_.packet.sequence
                );
            }
            bytes32 packetCommitment = keccak256(
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
            if (commitment != packetCommitment) {
                revert IBCChannelPacketCommitmentMismatch(commitment, packetCommitment);
            }
        }

        {
            Channel.Data memory expectedChannel = Channel.Data({
                state: Channel.State.STATE_CLOSED,
                ordering: channel.ordering,
                counterparty: ChannelCounterparty.Data({
                    port_id: msg_.packet.sourcePort,
                    channel_id: msg_.packet.sourceChannel
                }),
                connection_hops: IBCChannelLib.buildConnectionHops(connection.counterparty.connection_id),
                version: channel.version,
                upgrade_sequence: msg_.counterpartyUpgradeSequence
            });
            if (
                // slither-disable-next-line reentrancy-no-eth
                !client.verifyMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proofClose,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.channelPath(msg_.packet.destinationPort, msg_.packet.destinationChannel),
                    Channel.encode(expectedChannel)
                )
            ) {
                revert IBCChannelFailedVerifyChannelState(
                    connection.client_id,
                    IBCCommitment.channelPath(msg_.packet.destinationPort, msg_.packet.destinationChannel),
                    Channel.encode(expectedChannel),
                    msg_.proofClose,
                    msg_.proofHeight
                );
            }
        }

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            // check that packet has not been received
            if (msg_.packet.sequence < msg_.nextSequenceRecv) {
                revert IBCChannelPacketMaybeAlreadyReceived(msg_.packet.sequence, msg_.nextSequenceRecv);
            }
            if (
                // slither-disable-next-line reentrancy-no-eth
                !client.verifyMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proofUnreceived,
                    connection.counterparty.prefix.key_prefix,
                    IBCCommitment.nextSequenceRecvCommitmentPath(
                        msg_.packet.destinationPort, msg_.packet.destinationChannel
                    ),
                    IBCChannelLib.uint64ToBigEndianBytes(msg_.nextSequenceRecv)
                )
            ) {
                revert IBCChannelFailedVerifyNextSequenceRecv(
                    connection.client_id,
                    IBCCommitment.nextSequenceRecvCommitmentPath(
                        msg_.packet.destinationPort, msg_.packet.destinationChannel
                    ),
                    msg_.nextSequenceRecv,
                    msg_.proofUnreceived,
                    msg_.proofHeight
                );
            }
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            bytes memory path = IBCCommitment.packetReceiptCommitmentPathCalldata(
                msg_.packet.destinationPort, msg_.packet.destinationChannel, msg_.packet.sequence
            );
            if (
                // slither-disable-next-line reentrancy-no-eth
                !client.verifyNonMembership(
                    connection.client_id,
                    msg_.proofHeight,
                    connection.delay_period,
                    calcBlockDelay(connection.delay_period),
                    msg_.proofUnreceived,
                    connection.counterparty.prefix.key_prefix,
                    path
                )
            ) {
                revert IBCChannelFailedVerifyPacketReceiptAbsence(
                    connection.client_id, path, msg_.proofUnreceived, msg_.proofHeight
                );
            }
        } else {
            revert IBCChannelUnknownChannelOrder(channel.ordering);
        }
        lookupModuleByChannel(msg_.packet.sourcePort, msg_.packet.sourceChannel).onTimeoutPacket(
            msg_.packet, _msgSender()
        );
        emit TimeoutPacket(msg_.packet);
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
}
