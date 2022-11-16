// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

library IBCIdentifier {
    // Constant values

    uint256 constant commitmentSlot = 0;

    // Commitment key generator

    function clientCommitmentKey(string memory clientId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked("clients/", clientId, "/clientState"));
    }

    function consensusCommitmentKey(string memory clientId, uint64 revisionNumber, uint64 revisionHeight)
        public
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                "clients/", clientId, "/consensusStates/", uint2str(revisionNumber), "-", uint2str(revisionHeight)
            )
        );
    }

    function connectionCommitmentKey(string memory connectionId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked("connections/", connectionId));
    }

    function channelCommitmentKey(string memory portId, string memory channelId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked("channelEnds/ports/", portId, "/channels/", channelId));
    }

    function packetCommitmentKey(string memory portId, string memory channelId, uint64 sequence)
        public
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked("commitments/ports/", portId, "/channels/", channelId, "/sequences/", uint2str(sequence))
        );
    }

    function packetAcknowledgementCommitmentKey(string memory portId, string memory channelId, uint64 sequence)
        public
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked("acks/ports/", portId, "/channels/", channelId, "/sequences/", uint2str(sequence))
        );
    }

    function packetReceiptCommitmentKey(string memory portId, string memory channelId, uint64 sequence)
        public
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked("receipts/ports/", portId, "/channels/", channelId, "/sequences/", uint2str(sequence))
        );
    }

    function nextSequenceRecvCommitmentKey(string memory portId, string memory channelId)
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked("nextSequenceRecv/ports/", portId, "/channels/", channelId));
    }

    // Slot calculator

    function clientStateCommitmentSlot(string calldata clientId) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(clientCommitmentKey(clientId), commitmentSlot));
    }

    function consensusStateCommitmentSlot(string calldata clientId, uint64 revisionNumber, uint64 revisionHeight)
        external
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(consensusCommitmentKey(clientId, revisionNumber, revisionHeight), commitmentSlot)
        );
    }

    function connectionCommitmentSlot(string calldata connectionId) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(connectionCommitmentKey(connectionId), commitmentSlot));
    }

    function channelCommitmentSlot(string calldata portId, string calldata channelId) external pure returns (bytes32) {
        return keccak256(abi.encodePacked(channelCommitmentKey(portId, channelId), commitmentSlot));
    }

    function packetCommitmentSlot(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(packetCommitmentKey(portId, channelId, sequence), commitmentSlot));
    }

    function packetAcknowledgementCommitmentSlot(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes32)
    {
        return
            keccak256(abi.encodePacked(packetAcknowledgementCommitmentKey(portId, channelId, sequence), commitmentSlot));
    }

    function packetReceiptCommitmentSlot(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(packetReceiptCommitmentKey(portId, channelId, sequence), commitmentSlot));
    }

    function nextSequenceRecvCommitmentSlot(string calldata portId, string calldata channelId)
        external
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(nextSequenceRecvCommitmentKey(portId, channelId), commitmentSlot));
    }

    // CapabilityPath

    function portCapabilityPath(string calldata portId) external pure returns (bytes memory) {
        return abi.encodePacked(portId);
    }

    function channelCapabilityPath(string calldata portId, string calldata channelId)
        external
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(portId, "/", channelId);
    }

    // Utility functions

    function uint2str(uint64 _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint64 j = _i;
        uint64 len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint64 k = len;
        while (_i != 0) {
            k = k - 1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
