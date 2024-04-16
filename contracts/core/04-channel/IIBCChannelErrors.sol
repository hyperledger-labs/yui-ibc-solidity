// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {Channel} from "../../proto/Channel.sol";

interface IIBCChannelErrors {
    /// @param state channel state
    error IBCChannelUnexpectedChannelState(Channel.State state);

    /// @param portId port identifier
    /// @param channelId channel identifier
    error IBCChannelChannelNotFound(string portId, string channelId);

    /// @param ordering channel ordering
    error IBCChannelUnknownChannelOrder(Channel.Order ordering);

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param value value
    /// @param proof proof
    /// @param height proof height
    error IBCChannelFailedVerifyChannelState(string clientId, bytes path, bytes value, bytes proof, Height.Data height);

    /// @param connectionId connection identifier
    error IBCChannelConnectionNotOpened(string connectionId);

    /// @param counterpartyChannelId counterparty channel identifier
    error IBCChannelCounterpartyChannelIdNotEmpty(string counterpartyChannelId);

    /// @param length length of the connection hops
    error IBCChannelInvalidConnectionHopsLength(uint256 length);

    /// @param connectionId connection identifier
    /// @param length length of the connection hops
    error IBCChannelConnectionMultipleVersionsFound(string connectionId, uint256 length);

    /// @param ordering channel ordering
    error IBCChannelConnectionFeatureNotSupported(Channel.Order ordering);

    /// @notice timeout height and timestamp are both zero
    error IBCChannelZeroPacketTimeout();

    /// @notice receiving chain block height >= packet timeout height
    /// @param timeoutHeight packet timeout height
    /// @param latestHeight latest height of the receiving chain
    error IBCChannelPastPacketTimeoutHeight(Height.Data timeoutHeight, Height.Data latestHeight);

    /// @notice receiving chain block timestamp >= packet timeout timestamp
    /// @param timeoutTimestamp packet timeout timestamp
    /// @param latestTimestamp latest timestamp of the receiving chain
    error IBCChannelPastPacketTimeoutTimestamp(uint64 timeoutTimestamp, uint64 latestTimestamp);

    /// @notice packet timeout has not been reached for height or timestamp
    error IBCChannelTimeoutNotReached();

    /// @param commitment packet receipt commitment
    error IBCChannelUnknownPacketReceiptCommitment(bytes32 commitment);

    /// @param sourcePort source port
    /// @param sourceChannel source channel
    error IBCChannelUnexpectedPacketSource(string sourcePort, string sourceChannel);

    /// @param destinationPort destination port
    /// @param destinationChannel destination channel
    error IBCChannelUnexpectedPacketDestination(string destinationPort, string destinationChannel);

    error IBCChannelCannotRecvNextUpgradePacket(uint64 sequence, uint64 counterpartyNextSequenceSend);

    error IBCChannelPacketAlreadyProcessInPreviousUpgrade(uint64 sequence, uint64 recvStartSequence);

    error IBCChannelAckAlreadyProcessedInPreviousUpgrade(uint64 sequence, uint64 ackStartSequence);

    /// @param currentBlockNumber current block number
    /// @param timeoutHeight packet timeout height
    error IBCChannelTimeoutPacketHeight(uint256 currentBlockNumber, uint64 timeoutHeight);

    /// @param currentTimestamp current timestamp
    /// @param timeoutTimestamp packet timeout timestamp
    error IBCChannelTimeoutPacketTimestamp(uint256 currentTimestamp, uint64 timeoutTimestamp);

    /// @param destinationPort destination port
    /// @param destinationChannel destination channel
    /// @param sequence packet sequence
    error IBCChannelPacketReceiptAlreadyExists(string destinationPort, string destinationChannel, uint64 sequence);

    /// @param expected expected sequence
    error IBCChannelUnexpectedNextSequenceRecv(uint64 expected);

    /// @param portId port identifier
    /// @param channelId channel identifier
    /// @param sequence packet sequence
    error IBCChannelPacketCommitmentNotFound(string portId, string channelId, uint64 sequence);

    /// @param expected expected sequence
    error IBCChannelUnexpectedNextSequenceAck(uint64 expected);

    /// @param expected expected commitment
    /// @param actual actual commitment
    error IBCChannelPacketCommitmentMismatch(bytes32 expected, bytes32 actual);

    /// @param sequence packet sequence
    /// @param nextSequenceRecv next sequence received
    error IBCChannelPacketMaybeAlreadyReceived(uint64 sequence, uint64 nextSequenceRecv);

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param commitment packet commitment
    /// @param proof proof
    /// @param height proof height
    error IBCChannelFailedVerifyPacketCommitment(
        string clientId, bytes path, bytes32 commitment, bytes proof, Height.Data height
    );

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param commitment acknowledgement commitment
    /// @param proof proof
    /// @param height proof height
    error IBCChannelFailedVerifyPacketAcknowledgement(
        string clientId, bytes path, bytes32 commitment, bytes proof, Height.Data height
    );

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param nextSequenceRecv next sequence received
    /// @param proof proof
    /// @param height proof height
    error IBCChannelFailedVerifyNextSequenceRecv(
        string clientId, bytes path, uint64 nextSequenceRecv, bytes proof, Height.Data height
    );

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param proof proof
    /// @param height proof height
    error IBCChannelFailedVerifyPacketReceiptAbsence(string clientId, bytes path, bytes proof, Height.Data height);

    /// @notice acknowledgement cannot be empty
    error IBCChannelEmptyAcknowledgement();

    /// @param destinationPort destination port
    /// @param destinationChannel destination channel
    /// @param sequence packet sequence
    error IBCChannelAcknowledgementAlreadyWritten(string destinationPort, string destinationChannel, uint64 sequence);
}
