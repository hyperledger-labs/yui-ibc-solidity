// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "./helpers/IBCTestHelper.t.sol";
import "forge-std/Test.sol";
import {ICS20Transfer} from "../../../contracts/apps/20-transfer/ICS20Transfer.sol";
import {ICS20Lib} from "../../../contracts/apps/20-transfer/ICS20Lib.sol";
import {ERC20Token} from "../../../contracts/apps/20-transfer/ERC20Token.sol";
import {LocalhostClientLib} from "../../../contracts/clients/09-localhost/LocalhostClient.sol";
import {LocalhostHelper} from "../../../contracts/clients/09-localhost/LocalhostHelper.sol";
import {ICS04PacketEventTestHelper} from "./helpers/ICS04PacketTestHelper.t.sol";
import {IIBCChannelRecvPacket, IIBCChannelAcknowledgePacket} from "../../../contracts/core/04-channel/IIBCChannel.sol";

contract TestICS20Transfer is IBCTestHelper, ICS04PacketEventTestHelper {
    using LocalhostHelper for TestableIBCHandler;

    string internal constant ICS20_APP_PORT = "transfer";
    string internal constant ICS20_APP_VERSION = "ics20-1";

    address immutable alice = address(0x01);
    address immutable bob = address(0x02);
    address immutable charlie = address(0x03);

    TestableIBCHandler ibcHandler;
    ICS20Transfer transferApp;
    ERC20Token token;

    struct ChannelInfo {
        string connectionId;
        string portId;
        string channelId;
    }

    function setUp() public {
        token = new ERC20Token("test", "test", 1000000);

        ibcHandler = defaultIBCHandler();
        transferApp = new ICS20Transfer(ibcHandler, ICS20_APP_PORT);
        ibcHandler.bindPort(ICS20_APP_PORT, transferApp);
        ibcHandler.registerLocalhostClient();
        ibcHandler.createLocalhostClient();
    }

    function testRelay() public {
        vm.recordLogs();
        (ChannelInfo memory channel0, ChannelInfo memory channel1) = createTransferChannel();
        (ChannelInfo memory channel2, ChannelInfo memory channel3) = createTransferChannel();

        string memory denom0 = ICS20Lib.addressToHexString(address(token));
        string memory denom1 = ICS20Lib.denom(channel1.portId, channel1.channelId, address(token));
        string memory denom2 = ICS20Lib.denom(channel3.portId, channel3.channelId, denom1);

        token.transfer(alice, 1000);
        {
            vm.startPrank(alice);
            assertTrue(token.approve(address(transferApp), 1000));
            transferApp.depositSendTransfer(
                channel0.channelId,
                address(token),
                1000,
                ICS20Lib.addressToHexString(bob),
                ICS20Lib.timeout(0, 2)
            );
            vm.stopPrank();
            assertEq(token.balanceOf(address(transferApp)), 1000);
            assertEq(transferApp.balanceOf(alice, denom0), 0);
            assertEq(transferApp.balanceOf(bob, denom1), 0);
            Packet memory packet = relayLastSentPacket(channel0.portId, channel0.channelId);
            relayLastWrittenAcknowledgement(packet);
            assertEq(transferApp.balanceOf(bob, denom1), 1000);
        }
        {
            vm.prank(bob);
            transferApp.sendTransfer(
                channel1.channelId,
                denom1,
                700,
                ICS20Lib.addressToHexString(alice),
                ICS20Lib.timeout(0, 2)
            );
            assertEq(transferApp.balanceOf(bob, denom1), 300);
            Packet memory packet = relayLastSentPacket(channel1.portId, channel1.channelId);
            assertEq(token.balanceOf(address(transferApp)), 1000);
            relayLastWrittenAcknowledgement(packet);            
            assertEq(transferApp.balanceOf(alice, denom0), 700);
            vm.prank(alice);
            transferApp.withdraw(alice, address(token), 200);
            assertEq(transferApp.balanceOf(alice, denom0), 500);
            assertEq(token.balanceOf(alice), 200);
            assertEq(token.balanceOf(address(transferApp)), 800);
        }
        {
            vm.prank(bob);
            transferApp.sendTransfer(
                channel2.channelId,
                denom1,
                200,
                ICS20Lib.addressToHexString(charlie),
                ICS20Lib.timeout(0, 2)
            );
            assertEq(transferApp.balanceOf(bob, denom1), 100);
            Packet memory packet = relayLastSentPacket(channel2.portId, channel2.channelId);
            assertEq(token.balanceOf(address(transferApp)), 800);
            relayLastWrittenAcknowledgement(packet);            
            assertEq(transferApp.balanceOf(charlie, denom2), 200);
        }
        {
            vm.prank(charlie);
            transferApp.sendTransfer(
                channel3.channelId,
                denom2,
                100,
                ICS20Lib.addressToHexString(bob),
                ICS20Lib.timeout(0, 2)
            );
            assertEq(transferApp.balanceOf(charlie, denom2), 100);
            relayLastWrittenAcknowledgement(relayLastSentPacket(channel3.portId, channel3.channelId));
            assertEq(transferApp.balanceOf(bob, denom1), 200);

            vm.prank(charlie);
            transferApp.sendTransfer(
                channel1.channelId,
                denom2,
                100,
                ICS20Lib.addressToHexString(alice),
                ICS20Lib.timeout(0, 2)
            );
            assertEq(transferApp.balanceOf(charlie, denom2), 0);
            relayLastWrittenAcknowledgement(relayLastSentPacket(channel1.portId, channel1.channelId));
            assertEq(transferApp.balanceOf(alice, ICS20Lib.denom(channel0.portId, channel0.channelId, denom2)), 100);
            assertEq(transferApp.balanceOf(alice, denom0), 500);
        }
        // bob transfer 100 to alice
        {
            vm.prank(bob);
            transferApp.sendTransfer(
                channel1.channelId,
                denom1,
                100,
                ICS20Lib.addressToHexString(alice),
                ICS20Lib.timeout(0, 2)
            );
            assertEq(transferApp.balanceOf(bob, denom1), 100);
            Packet memory packet = relayLastSentPacket(channel1.portId, channel1.channelId);
            assertEq(token.balanceOf(address(transferApp)), 800);
            relayLastWrittenAcknowledgement(packet);
            assertEq(transferApp.balanceOf(alice, denom0), 600);
        }
    }

    function testDepositTransferWithdraw() public {
        token.transfer(alice, 1000);
        string memory denom0 = ICS20Lib.addressToHexString(address(token));
        vm.startPrank(alice);
        token.approve(address(transferApp), 1000);
        transferApp.deposit(alice, address(token), 1000);
        transferApp.transfer(bob, denom0, 300);
        vm.stopPrank();
        assertEq(transferApp.balanceOf(alice, denom0), 700);
        assertEq(transferApp.balanceOf(bob, denom0), 300);
        vm.prank(bob);
        transferApp.withdraw(bob, address(token), 200);
        assertEq(token.balanceOf(bob), 200);
        assertEq(transferApp.balanceOf(bob, denom0), 100);
    }

    function testDeposit() public {
        token.transfer(alice, 1000);

        vm.startPrank(alice);
        vm.expectRevert();
        transferApp.deposit(alice, address(token), 1000);

        token.approve(address(transferApp), 100);
        vm.expectRevert();
        transferApp.deposit(alice, address(token), 101);

        token.approve(address(transferApp), 500);
        transferApp.deposit(bob, address(token), 500);
        transferApp.balanceOf(bob, ICS20Lib.addressToHexString(address(token)));

        token.approve(address(transferApp), 500);
        transferApp.deposit(alice, address(token), 500);
        transferApp.balanceOf(alice, ICS20Lib.addressToHexString(address(token)));
    }

    function testWithdraw() public {
        token.transfer(alice, 1000);
        vm.startPrank(alice);
        token.approve(address(transferApp), 1000);
        transferApp.deposit(alice, address(token), 1000);
        assertEq(token.balanceOf(alice), 0);

        vm.expectRevert();
        transferApp.withdraw(alice, address(token), 1001);

        vm.expectRevert();
        transferApp.withdraw(alice, address(0x01), 1000);

        transferApp.withdraw(alice, address(token), 200);
        assertEq(token.balanceOf(alice), 200);

        vm.expectRevert();
        transferApp.withdraw(alice, address(token), 900);

        transferApp.withdraw(bob, address(token), 800);
        assertEq(token.balanceOf(bob), 800);

        string memory denom = ICS20Lib.addressToHexString(address(token));
        assertEq(transferApp.balanceOf(alice, denom), 0);
    }

    function createTransferChannel() internal returns (ChannelInfo memory, ChannelInfo memory) {
        (string memory connectionId0, string memory connectionId1) = ibcHandler.createLocalhostConnection();
        (string memory channelId0, string memory channelId1) = ibcHandler.createLocalhostChannel(
            LocalhostHelper.MsgCreateChannel({
                connectionId0: connectionId0,
                connectionId1: connectionId1,
                portId0: ICS20_APP_PORT,
                portId1: ICS20_APP_PORT,
                ordering: Channel.Order.ORDER_UNORDERED,
                version: ICS20_APP_VERSION
            })
        );
        return (
            ChannelInfo({connectionId: connectionId0, portId: ICS20_APP_PORT, channelId: channelId0}),
            ChannelInfo({connectionId: connectionId1, portId: ICS20_APP_PORT, channelId: channelId1})
        );
    }

    function relayLastSentPacket(string memory portId, string memory channelId) internal returns (Packet memory) {
        Packet memory packet = getLastSentPacket(ibcHandler, portId, channelId, vm.getRecordedLogs());
        ibcHandler.recvPacket(
            IIBCChannelRecvPacket.MsgPacketRecv({
                packet: packet,
                proof: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.nil()
            })
        );
        return packet;
    }

    function relayLastWrittenAcknowledgement(Packet memory packet) internal returns (WriteAcknolwedgement memory) {
        WriteAcknolwedgement memory ack = getLastWrittenAcknowledgement(ibcHandler, vm.getRecordedLogs());
        assertEq(ack.acknowledgement, ICS20Lib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON);
        ibcHandler.acknowledgePacket(
            IIBCChannelAcknowledgePacket.MsgPacketAcknowledgement({
                packet: packet,
                acknowledgement: ack.acknowledgement,
                proof: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.nil()
            })
        );
        return ack;
    }
}