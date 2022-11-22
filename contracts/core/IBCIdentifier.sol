// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

library IBCIdentifier {
    // commitment path

    function clientStatePath(string calldata clientId) external pure returns (bytes memory) {
        return abi.encodePacked("clients/", clientId, "/clientState");
    }

    function consensusStatePath(string calldata clientId, uint64 revisionNumber, uint64 revisionHeight)
        external
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            "clients/", clientId, "/consensusStates/", uint2str(revisionNumber), "-", uint2str(revisionHeight)
        );
    }

    function connectionPath(string calldata connectionId) external pure returns (bytes memory) {
        return abi.encodePacked("connections/", connectionId);
    }

    function channelPath(string calldata portId, string calldata channelId) external pure returns (bytes memory) {
        return abi.encodePacked("channelEnds/ports/", portId, "/channels/", channelId);
    }

    function packetCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes memory)
    {
        return
            abi.encodePacked("commitments/ports/", portId, "/channels/", channelId, "/sequences/", uint2str(sequence));
    }

    function packetAcknowledgementCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("acks/ports/", portId, "/channels/", channelId, "/sequences/", uint2str(sequence));
    }

    function packetReceiptCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("receipts/ports/", portId, "/channels/", channelId, "/sequences/", uint2str(sequence));
    }

    function nextSequenceRecvCommitmentPath(string calldata portId, string calldata channelId)
        external
        pure
        returns (bytes memory)
    {
        return abi.encodePacked("nextSequenceRecv/ports/", portId, "/channels/", channelId);
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
