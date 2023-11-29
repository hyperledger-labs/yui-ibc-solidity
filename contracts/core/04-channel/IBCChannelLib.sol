// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Channel} from "../../proto/Channel.sol";

library IBCChannelLib {
    enum PacketReceipt {
        NONE,
        SUCCESSFUL
    }

    string public constant ORDER_UNORDERED = "ORDER_UNORDERED";
    string public constant ORDER_ORDERED = "ORDER_ORDERED";

    bytes32 internal constant PACKET_RECEIPT_SUCCESSFUL_KECCAK256 =
        keccak256(abi.encodePacked(PacketReceipt.SUCCESSFUL));

    function receiptCommitmentToReceipt(bytes32 commitment) internal pure returns (PacketReceipt) {
        if (commitment == bytes32(0)) {
            return PacketReceipt.NONE;
        } else if (commitment == PACKET_RECEIPT_SUCCESSFUL_KECCAK256) {
            return PacketReceipt.SUCCESSFUL;
        } else {
            revert("unknown receipt");
        }
    }

    function toString(Channel.Order order) internal pure returns (string memory) {
        if (order == Channel.Order.ORDER_UNORDERED) {
            return ORDER_UNORDERED;
        } else if (order == Channel.Order.ORDER_ORDERED) {
            return ORDER_ORDERED;
        } else {
            revert("unknown order");
        }
    }
}
