// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

library IBCCommitment {
    // Commitment path generators that comply with https://github.com/cosmos/ibc/tree/main/spec/core/ics-024-host-requirements#path-space

    function clientStatePath(string memory clientId) internal pure returns (bytes memory) {
        return abi.encodePacked("clients/", clientId, "/clientState");
    }

    function consensusStatePath(string memory clientId, uint64 revisionNumber, uint64 revisionHeight)
        internal
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

    function connectionPath(string memory connectionId) internal pure returns (bytes memory) {
        return abi.encodePacked("connections/", connectionId);
    }

    function channelPath(string memory portId, string memory channelId) internal pure returns (bytes memory) {
        return abi.encodePacked("channelEnds/ports/", portId, "/channels/", channelId);
    }

    function packetCommitmentPathCalldata(string calldata portId, string calldata channelId, uint64 sequence)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            "commitments/ports/", portId, "/channels/", channelId, "/sequences/", Strings.toString(sequence)
        );
    }

    function packetAcknowledgementCommitmentPathCalldata(
        string calldata portId,
        string calldata channelId,
        uint64 sequence
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked("acks/ports/", portId, "/channels/", channelId, "/sequences/", Strings.toString(sequence));
    }

    function packetReceiptCommitmentPathCalldata(string calldata portId, string calldata channelId, uint64 sequence)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            "receipts/ports/", portId, "/channels/", channelId, "/sequences/", Strings.toString(sequence)
        );
    }

    function nextSequenceRecvCommitmentPath(string memory portId, string memory channelId)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("nextSequenceRecv/ports/", portId, "/channels/", channelId);
    }

    function nextSequenceRecvCommitmentPathCalldata(string calldata portId, string calldata channelId)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("nextSequenceRecv/ports/", portId, "/channels/", channelId);
    }

    function channelUpgradePath(string memory portId, string memory channelId) internal pure returns (bytes memory) {
        return abi.encodePacked("channelUpgrades/upgrades/ports/", portId, "/channels/", channelId);
    }

    function channelUpgradeErrorPath(string memory portId, string memory channelId)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("channelUpgrades/upgradeError/ports/", portId, "/channels/", channelId);
    }

    // Key generators for Commitment mapping

    function clientStateCommitmentKey(string memory clientId) internal pure returns (bytes32) {
        return keccak256(clientStatePath(clientId));
    }

    function consensusStateCommitmentKey(string memory clientId, uint64 revisionNumber, uint64 revisionHeight)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(consensusStatePath(clientId, revisionNumber, revisionHeight));
    }

    function connectionCommitmentKey(string memory connectionId) internal pure returns (bytes32) {
        return keccak256(connectionPath(connectionId));
    }

    function channelCommitmentKey(string memory portId, string memory channelId) internal pure returns (bytes32) {
        return keccak256(channelPath(portId, channelId));
    }

    function nextSequenceRecvCommitmentKey(string memory portId, string memory channelId)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(nextSequenceRecvCommitmentPath(portId, channelId));
    }

    function packetCommitmentKeyCalldata(string calldata portId, string calldata channelId, uint64 sequence)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(packetCommitmentPathCalldata(portId, channelId, sequence));
    }

    function packetAcknowledgementCommitmentKeyCalldata(
        string calldata portId,
        string calldata channelId,
        uint64 sequence
    ) internal pure returns (bytes32) {
        return keccak256(packetAcknowledgementCommitmentPathCalldata(portId, channelId, sequence));
    }

    function packetReceiptCommitmentKeyCalldata(string calldata portId, string calldata channelId, uint64 sequence)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(packetReceiptCommitmentPathCalldata(portId, channelId, sequence));
    }

    function nextSequenceRecvCommitmentKeyCalldata(string calldata portId, string calldata channelId)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(nextSequenceRecvCommitmentPathCalldata(portId, channelId));
    }

    function channelUpgradeCommitmentKey(string memory portId, string memory channelId)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(channelUpgradePath(portId, channelId));
    }

    function channelUpgradeErrorCommitmentKey(string memory portId, string memory channelId)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(channelUpgradeErrorPath(portId, channelId));
    }
}
