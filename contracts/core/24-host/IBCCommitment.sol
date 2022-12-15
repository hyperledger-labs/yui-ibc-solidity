// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";

library IBCCommitment {
    // Commitment path generators that comply with https://github.com/cosmos/ibc/tree/main/spec/core/ics-024-host-requirements#path-space

    function clientStatePath(string calldata clientId) public pure returns (bytes memory) {
        return abi.encodePacked("clients/", clientId, "/clientState");
    }

    function consensusStatePath(string calldata clientId, uint64 revisionNumber, uint64 revisionHeight)
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            "clients/",
            clientId,
            "/consensusStates/",
            Strings.toString(revisionNumber),
            "-",
            Strings.toString(revisionHeight)
        );
    }

    function connectionPath(string calldata connectionId) public pure returns (bytes memory) {
        return abi.encodePacked("connections/", connectionId);
    }

    function channelPath(string calldata portId, string calldata channelId) public pure returns (bytes memory) {
        return abi.encodePacked("channelEnds/ports/", portId, "/channels/", channelId);
    }

    function packetCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            "commitments/ports/", portId, "/channels/", channelId, "/sequences/", Strings.toString(sequence)
        );
    }

    function packetAcknowledgementCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        public
        pure
        returns (bytes memory)
    {
        return
            abi.encodePacked("acks/ports/", portId, "/channels/", channelId, "/sequences/", Strings.toString(sequence));
    }

    function packetReceiptCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            "receipts/ports/", portId, "/channels/", channelId, "/sequences/", Strings.toString(sequence)
        );
    }

    function nextSequenceRecvCommitmentPath(string calldata portId, string calldata channelId)
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("nextSequenceRecv/ports/", portId, "/channels/", channelId);
    }

    // Key generators for Commitment mapping

    function clientStateCommitmentKey(string calldata clientId) external pure returns (bytes32) {
        return keccak256(clientStatePath(clientId));
    }

    function consensusStateCommitmentKey(string calldata clientId, uint64 revisionNumber, uint64 revisionHeight)
        external
        pure
        returns (bytes32)
    {
        return keccak256(consensusStatePath(clientId, revisionNumber, revisionHeight));
    }

    function connectionCommitmentKey(string calldata connectionId) external pure returns (bytes32) {
        return keccak256(connectionPath(connectionId));
    }

    function channelCommitmentKey(string calldata portId, string calldata channelId) external pure returns (bytes32) {
        return keccak256(channelPath(portId, channelId));
    }

    function packetCommitmentKey(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes32)
    {
        return keccak256(packetCommitmentPath(portId, channelId, sequence));
    }

    function packetAcknowledgementCommitmentKey(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes32)
    {
        return keccak256(packetAcknowledgementCommitmentPath(portId, channelId, sequence));
    }

    function packetReceiptCommitmentKey(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes32)
    {
        return keccak256(packetReceiptCommitmentPath(portId, channelId, sequence));
    }

    function nextSequenceRecvCommitmentKey(string calldata portId, string calldata channelId)
        external
        pure
        returns (bytes32)
    {
        return keccak256(nextSequenceRecvCommitmentPath(portId, channelId));
    }
}
