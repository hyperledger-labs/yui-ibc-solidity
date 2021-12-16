// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./types/Client.sol";
import "./IBCHeight.sol";

library IBCIdentifier {
    using IBCHeight for Height.Data;

    // constant values

    uint256 constant commitmentSlot = 0;
    uint8 constant clientPrefix = 0;
    uint8 constant consensusStatePrefix = 1;
    uint8 constant connectionPrefix = 2;
    uint8 constant channelPrefix = 3;
    uint8 constant packetPrefix = 4;
    uint8 constant packetAckPrefix = 5;

    // Commitment key generator

    function clientCommitmentKey(string memory clientId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(clientPrefix, clientId));
    }

    function consensusCommitmentKey(string memory clientId, Height.Data memory height) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(consensusStatePrefix, clientId, "/", height.toUint128()));
    }

    function connectionCommitmentKey(string memory connectionId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(connectionPrefix, connectionId));
    }

    function channelCommitmentKey(string memory portId, string memory channelId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(channelPrefix, portId, "/", channelId));
    }

    function packetCommitmentKey(string memory portId, string memory channelId, uint64 sequence) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(packetPrefix, portId, "/", channelId, "/", sequence));
    }

    function packetAcknowledgementCommitmentKey(string memory portId, string memory channelId, uint64 sequence) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(packetAckPrefix, portId, "/", channelId, "/", sequence));
    }

    // Slot calculator

    function clientStateCommitmentSlot(string calldata clientId) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(clientCommitmentKey(clientId), commitmentSlot));
    }

    function consensusStateCommitmentSlot(string calldata clientId, Height.Data calldata height) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(consensusCommitmentKey(clientId, height), commitmentSlot));
    }

    function connectionCommitmentSlot(string calldata connectionId) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(connectionCommitmentKey(connectionId), commitmentSlot));
    }

    function channelCommitmentSlot(string calldata portId, string calldata channelId) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(channelCommitmentKey(portId, channelId), commitmentSlot));
    }

    function packetCommitmentSlot(string calldata portId, string calldata channelId, uint64 sequence) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(packetCommitmentKey(portId, channelId, sequence), commitmentSlot));
    }

    function packetAcknowledgementCommitmentSlot(string calldata portId, string calldata channelId, uint64 sequence) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(packetAcknowledgementCommitmentKey(portId, channelId, sequence), commitmentSlot));
    }

    // CapabilityPath

    function portCapabilityPath(string calldata portId) external pure returns (bytes memory) {
        return abi.encodePacked(portId);
    }

    function channelCapabilityPath(string calldata portId, string calldata channelId) external pure returns (bytes memory) {
        return abi.encodePacked(portId, "/", channelId);
    }
}
