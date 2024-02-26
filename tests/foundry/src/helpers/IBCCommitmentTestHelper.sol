// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IBCCommitment} from "../../../../contracts/core/24-host/IBCCommitment.sol";

library IBCCommitmentTestHelper {
    function clientStatePath(string calldata clientId) external pure returns (bytes memory) {
        return IBCCommitment.clientStatePath(clientId);
    }

    function consensusStatePath(string calldata clientId, uint64 revisionNumber, uint64 revisionHeight)
        external
        pure
        returns (bytes memory)
    {
        return IBCCommitment.consensusStatePath(clientId, revisionNumber, revisionHeight);
    }

    function connectionPath(string calldata connectionId) external pure returns (bytes memory) {
        return IBCCommitment.connectionPath(connectionId);
    }

    function channelPath(string calldata portId, string calldata channelId) external pure returns (bytes memory) {
        return IBCCommitment.channelPath(portId, channelId);
    }

    function packetCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes memory)
    {
        return IBCCommitment.packetCommitmentPathCalldata(portId, channelId, sequence);
    }

    function packetAcknowledgementCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes memory)
    {
        return IBCCommitment.packetAcknowledgementCommitmentPathCalldata(portId, channelId, sequence);
    }

    function packetReceiptCommitmentPath(string calldata portId, string calldata channelId, uint64 sequence)
        external
        pure
        returns (bytes memory)
    {
        return IBCCommitment.packetReceiptCommitmentPathCalldata(portId, channelId, sequence);
    }

    function nextSequenceRecvCommitmentPath(string calldata portId, string calldata channelId)
        external
        pure
        returns (bytes memory)
    {
        return IBCCommitment.nextSequenceRecvCommitmentPathCalldata(portId, channelId);
    }
}
