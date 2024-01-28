// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Channel} from "../../proto/Channel.sol";
import {IIBCChannelErrors} from "./IIBCChannelErrors.sol";

library IBCChannelLib {
    enum PacketReceipt {
        NONE,
        SUCCESSFUL
    }

    string internal constant ORDER_UNORDERED = "ORDER_UNORDERED";
    string internal constant ORDER_ORDERED = "ORDER_ORDERED";

    bytes32 internal constant PACKET_RECEIPT_SUCCESSFUL_KECCAK256 =
        keccak256(abi.encodePacked(PacketReceipt.SUCCESSFUL));

    function receiptCommitmentToReceipt(bytes32 commitment) internal pure returns (PacketReceipt) {
        if (commitment == bytes32(0)) {
            return PacketReceipt.NONE;
        } else if (commitment == PACKET_RECEIPT_SUCCESSFUL_KECCAK256) {
            return PacketReceipt.SUCCESSFUL;
        } else {
            revert IIBCChannelErrors.IBCChannelUnknownPacketReceiptCommitment(commitment);
        }
    }

    function buildConnectionHops(string memory connectionId) internal pure returns (string[] memory hops) {
        hops = new string[](1);
        hops[0] = connectionId;
        return hops;
    }

    function toString(Channel.Order order) internal pure returns (string memory) {
        if (order == Channel.Order.ORDER_UNORDERED) {
            return ORDER_UNORDERED;
        } else if (order == Channel.Order.ORDER_ORDERED) {
            return ORDER_ORDERED;
        } else {
            revert IIBCChannelErrors.IBCChannelUnknownChannelOrder(order);
        }
    }
}
