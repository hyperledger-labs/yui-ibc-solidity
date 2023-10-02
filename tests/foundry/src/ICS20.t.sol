// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../../contracts/apps/20-transfer/ICS20Lib.sol";

contract TestICS20 is Test {
    function setUp() public {}

    function testMarshaling() public {
        {
            ICS20Lib.PacketData memory data = ICS20Lib.PacketData({
                denom: "portidone/channel-0/portidtwo/channel-1/uatom",
                sender: "cosmos12xjp5l0x5q2rts3jkujjvxskx4z0ckfzhxchkd",
                receiver: "cosmos1w3jhxarjv43k26tkv4eq8wv34g",
                amount: 1_000_000,
                memo: "memo"
            });
            bytes memory bz = ICS20LibTestHelper.marshalUnsafeJSON(data);
            ICS20Lib.PacketData memory data2 = ICS20LibTestHelper.unmarshalJSON(bz);
            assertEq(data2.denom, data.denom);
            assertEq(data2.sender, data.sender);
            assertEq(data2.receiver, data.receiver);
            assertEq(data2.amount, data.amount);
            assertEq(data2.memo, data.memo);
        }

        {
            ICS20Lib.PacketData memory data =
                ICS20Lib.PacketData({denom: "", sender: "", receiver: "", amount: 0, memo: ""});
            bytes memory bz = ICS20LibTestHelper.marshalUnsafeJSON(data);
            ICS20Lib.PacketData memory data2 = ICS20LibTestHelper.unmarshalJSON(bz);
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
            ICS20Lib.PacketData memory data = ICS20LibTestHelper.unmarshalJSON(bz);
            assertEq(data.denom, "transfer/gaiachannel/atom");
            assertEq(data.sender, "cosmos12xjp5l0x5q2rts3jkujjvxskx4z0ckfzhxchkd");
            assertEq(data.receiver, "cosmos1w3jhxarjv43k26tkv4eq8wv34g");
            assertEq(data.amount, 100);
            assertEq(data.memo, "");
        }
    }

    function testAddressToHex(address addr) public {
        string memory hexStr = ICS20LibTestHelper.addressToHexString(addr);
        (address addr2, bool ok) = ICS20LibTestHelper.hexStringToAddress(hexStr);
        assertTrue(ok);
        assertEq(addr, addr2);
    }

    function testHexToAddress(string memory any) public {
        // This should not revert if the input is not a valid hex string.
        ICS20LibTestHelper.hexStringToAddress(any);
    }
}

library ICS20LibTestHelper {
    function addressToHexString(address addr) public pure returns (string memory) {
        return ICS20Lib.addressToHexString(addr);
    }

    function hexStringToAddress(string calldata hexStr) public pure returns (address, bool) {
        return ICS20Lib.hexStringToAddress(hexStr);
    }

    function marshalUnsafeJSON(ICS20Lib.PacketData calldata data) public pure returns (bytes memory) {
        return ICS20Lib.marshalUnsafeJSON(data);
    }

    function unmarshalJSON(bytes calldata bz) public pure returns (ICS20Lib.PacketData memory) {
        return ICS20Lib.unmarshalJSON(bz);
    }
}
