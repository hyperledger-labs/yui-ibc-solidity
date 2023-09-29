// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../../contracts/apps/20-transfer/ICS20Packet.sol";

contract TestICS20 is Test {
    function setUp() public {}

    function testMarshaling() public {
        {
            ICS20Packet.PacketData memory data = ICS20Packet.PacketData({
                denom: "portidone/channel-0/portidtwo/channel-1/uatom",
                sender: "cosmos12xjp5l0x5q2rts3jkujjvxskx4z0ckfzhxchkd",
                receiver: "cosmos1w3jhxarjv43k26tkv4eq8wv34g",
                amount: 1_000_000,
                memo: "memo"
            });
            bytes memory bz = ICS20Packet.marshalUnsafeJSON(data);
            ICS20Packet.PacketData memory data2 = ICS20Packet.unmarshalJSON(bz);
            assertEq(data2.denom, data.denom);
            assertEq(data2.sender, data.sender);
            assertEq(data2.receiver, data.receiver);
            assertEq(data2.amount, data.amount);
            assertEq(data2.memo, data.memo);
        }

        {
            ICS20Packet.PacketData memory data =
                ICS20Packet.PacketData({denom: "", sender: "", receiver: "", amount: 0, memo: ""});
            bytes memory bz = ICS20Packet.marshalUnsafeJSON(data);
            ICS20Packet.PacketData memory data2 = ICS20Packet.unmarshalJSON(bz);
            assertEq(data2.denom, data.denom);
            assertEq(data2.sender, data.sender);
            assertEq(data2.receiver, data.receiver);
            assertEq(data2.amount, data.amount);
            assertEq(data2.memo, data.memo);
        }

        {
            bytes memory bz = bytes(
                '{"amount":"100","denom":"transfer/gaiachannel/atom","receiver":"cosmos1w3jhxarjv43k26tkv4eq8wv34g","sender":"cosmos12xjp5l0x5q2rts3jkujjvxskx4z0ckfzhxchkd"}'
            );
            ICS20Packet.PacketData memory data = ICS20Packet.unmarshalJSON(bz);
            assertEq(data.denom, "transfer/gaiachannel/atom");
            assertEq(data.sender, "cosmos12xjp5l0x5q2rts3jkujjvxskx4z0ckfzhxchkd");
            assertEq(data.receiver, "cosmos1w3jhxarjv43k26tkv4eq8wv34g");
            assertEq(data.amount, 100);
            assertEq(data.memo, "");
        }
    }
}
