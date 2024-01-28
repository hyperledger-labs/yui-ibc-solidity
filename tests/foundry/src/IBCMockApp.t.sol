// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "./helpers/IBCTestHelper.t.sol";
import {Vm} from "forge-std/Test.sol";
import {IBCMockApp} from "../../../contracts/apps/mock/IBCMockApp.sol";
import {IBCMockLib} from "../../../contracts/apps/mock/IBCMockLib.sol";
import {IIBCChannelRecvPacket, IIBCChannelAcknowledgePacket} from "../../../contracts/core/04-channel/IIBCChannel.sol";
import {LocalhostClientLib} from "../../../contracts/clients/LocalhostClient.sol";
import {LocalhostHelper} from "../../../contracts/helpers/LocalhostHelper.sol";
import {ICS04PacketEventTestHelper} from "./helpers/ICS04PacketTestHelper.t.sol";

contract IBCMockAppTest is IBCTestHelper, ICS04PacketEventTestHelper {
    using LocalhostHelper for TestableIBCHandler;

    string internal constant MOCK_APP_PORT = "mockapp";
    string internal constant MOCK_APP_VERSION = "mockapp-1";

    TestableIBCHandler ibcHandler;
    IBCMockApp mockApp;

    struct ChannelInfo {
        string connectionId;
        string portId;
        string channelId;
    }

    function setUp() public {
        ibcHandler = defaultIBCHandler();
        mockApp = new IBCMockApp(ibcHandler);
        ibcHandler.bindPort(MOCK_APP_PORT, mockApp);
        ibcHandler.registerLocalhostClient();
        ibcHandler.createLocalhostClient();
    }

    function testHandshake() public {
        Channel.Order[2] memory orders = [Channel.Order.ORDER_UNORDERED, Channel.Order.ORDER_ORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            createMockAppChannel(orders[i]);
        }
    }

    function testHandshakeBetweenDifferentPorts() public {
        string memory mockAppPort2 = "mockapp2";
        ibcHandler.bindPort(mockAppPort2, mockApp);

        (string memory connectionId0, string memory connectionId1) = ibcHandler.createLocalhostConnection();
        Channel.Order[2] memory orders = [Channel.Order.ORDER_UNORDERED, Channel.Order.ORDER_ORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            ibcHandler.createLocalhostChannel(
                LocalhostHelper.MsgCreateChannel({
                    connectionId0: connectionId0,
                    connectionId1: connectionId1,
                    portId0: MOCK_APP_PORT,
                    portId1: mockAppPort2,
                    ordering: orders[i],
                    version: MOCK_APP_VERSION
                })
            );
        }
    }

    function testPacketRelay() public {
        vm.recordLogs();
        Channel.Order[2] memory orders = [Channel.Order.ORDER_UNORDERED, Channel.Order.ORDER_ORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channel0, ChannelInfo memory channel1) = createMockAppChannel(orders[i]);
            mockPacketRelay(channel0, channel1, orders[i]);
            mockPacketRelay(channel1, channel0, orders[i]);
            mockPacketRelay(channel0, channel1, orders[i]);
        }
    }

    function mockPacketRelay(ChannelInfo memory ca, ChannelInfo memory cb, Channel.Order ordering) internal {
        sendAndRelay(
            ca,
            cb,
            ordering,
            RelayCase({
                packetData: IBCMockLib.MOCK_PACKET_DATA,
                async: false,
                acknowledgement: IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON
            })
        );
        sendAndRelay(
            ca,
            cb,
            ordering,
            RelayCase({
                packetData: IBCMockLib.MOCK_FAIL_PACKET_DATA,
                async: false,
                acknowledgement: IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON
            })
        );
        sendAndRelay(
            ca,
            cb,
            ordering,
            RelayCase({
                packetData: IBCMockLib.MOCK_ASYNC_PACKET_DATA,
                async: true,
                acknowledgement: IBCMockLib.SUCCESSFUL_ASYNC_ACKNOWLEDGEMENT_JSON
            })
        );
    }

    function testPacketTimeout() public {
        vm.recordLogs();
        Channel.Order[2] memory orders = [Channel.Order.ORDER_UNORDERED, Channel.Order.ORDER_ORDERED];
        (string memory connId0, string memory connId1) = ibcHandler.createLocalhostConnection();
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channel0, ChannelInfo memory channel1) =
                createMockAppChannel(orders[i], connId0, connId1);
            Height.Data memory timeoutHeight = H(uint64(block.number + 1));
            mockApp.sendPacket(IBCMockLib.MOCK_PACKET_DATA, channel0.portId, channel0.channelId, timeoutHeight, 0);
            Packet memory packet =
                getLastSentPacket(ibcHandler, channel0.portId, channel0.channelId, vm.getRecordedLogs());

            uint64 nextSequenceRecv = orders[i] == Channel.Order.ORDER_ORDERED
                ? ibcHandler.getNextSequenceRecv(channel1.portId, channel1.channelId)
                : 0;

            // TimeoutPacket must be failed if the timeout height is not reached
            vm.expectRevert();
            ibcHandler.timeoutPacket(
                IIBCChannelPacketTimeout.MsgTimeoutPacket({
                    packet: packet,
                    proof: LocalhostClientLib.sentinelProof(),
                    proofHeight: timeoutHeight,
                    nextSequenceRecv: nextSequenceRecv
                })
            );

            // Set the latest height to the timeout height
            vm.roll(timeoutHeight.revision_height);

            // TimeoutPacket must be failed if an app does not allow channel closing
            mockApp.allowCloseChannel(false);
            vm.expectRevert();
            ibcHandler.timeoutPacket(
                IIBCChannelPacketTimeout.MsgTimeoutPacket({
                    packet: packet,
                    proof: LocalhostClientLib.sentinelProof(),
                    proofHeight: timeoutHeight,
                    nextSequenceRecv: nextSequenceRecv
                })
            );

            // TimeoutPacket must be successful if an app allows channel closing
            mockApp.allowCloseChannel(true);
            ibcHandler.timeoutPacket(
                IIBCChannelPacketTimeout.MsgTimeoutPacket({
                    packet: packet,
                    proof: LocalhostClientLib.sentinelProof(),
                    proofHeight: timeoutHeight,
                    nextSequenceRecv: nextSequenceRecv
                })
            );

            (Channel.Data memory channel, bool ok) = ibcHandler.getChannel(channel0.portId, channel0.channelId);
            assertTrue(ok);
            if (orders[i] == Channel.Order.ORDER_ORDERED) {
                assertTrue(channel.state == Channel.State.STATE_CLOSED);
            } else {
                assertTrue(channel.state == Channel.State.STATE_OPEN);
            }

            // TimeoutPacket must be failed if the previous exeuction succeededs
            vm.expectRevert();
            ibcHandler.timeoutPacket(
                IIBCChannelPacketTimeout.MsgTimeoutPacket({
                    packet: packet,
                    proof: LocalhostClientLib.sentinelProof(),
                    proofHeight: timeoutHeight,
                    nextSequenceRecv: nextSequenceRecv
                })
            );
        }
    }

    struct RelayCase {
        bytes packetData;
        bool async;
        bytes acknowledgement;
    }

    function sendAndRelay(ChannelInfo memory ca, ChannelInfo memory cb, Channel.Order ordering, RelayCase memory rc)
        private
    {
        uint64 sequence = mockApp.sendPacket(rc.packetData, ca.portId, ca.channelId, H(uint64(block.number + 1)), 0);
        Packet memory packet = getLastSentPacket(ibcHandler, ca.portId, ca.channelId, vm.getRecordedLogs());
        assertEq(packet.data, rc.packetData);
        ibcHandler.recvPacket(
            IIBCChannelRecvPacket.MsgPacketRecv({
                packet: packet,
                proof: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.nil()
            })
        );
        if (ordering == Channel.Order.ORDER_UNORDERED) {
            assertTrue(ibcHandler.hasPacketReceipt(cb.portId, cb.channelId, sequence));
        } else if (ordering == Channel.Order.ORDER_ORDERED) {
            assertTrue(ibcHandler.getNextSequenceRecv(cb.portId, cb.channelId) == sequence + 1);
        } else {
            revert("unknown ordering");
        }
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(abi.encode(packet), abi.encode(getLastRecvPacket(ibcHandler, logs)));
        if (rc.async) {
            mockApp.writeAcknowledgement(cb.portId, cb.channelId, sequence);
            logs = vm.getRecordedLogs();
        }
        WriteAcknolwedgement memory ack = getLastWrittenAcknowledgement(ibcHandler, logs);
        assertEq(ack.sequence, sequence);
        assertEq(ack.acknowledgement, rc.acknowledgement);

        ibcHandler.acknowledgePacket(
            IIBCChannelAcknowledgePacket.MsgPacketAcknowledgement({
                packet: packet,
                acknowledgement: ack.acknowledgement,
                proof: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.nil()
            })
        );
        assertTrue(ibcHandler.getPacketCommitment(ca.portId, ca.channelId, sequence) == bytes32(0));
    }

    function createMockAppChannel(Channel.Order ordering) internal returns (ChannelInfo memory, ChannelInfo memory) {
        (string memory connectionId0, string memory connectionId1) = ibcHandler.createLocalhostConnection();
        (string memory channelId0, string memory channelId1) = ibcHandler.createLocalhostChannel(
            LocalhostHelper.MsgCreateChannel({
                connectionId0: connectionId0,
                connectionId1: connectionId1,
                portId0: MOCK_APP_PORT,
                portId1: MOCK_APP_PORT,
                ordering: ordering,
                version: MOCK_APP_VERSION
            })
        );
        return (
            ChannelInfo({connectionId: connectionId0, portId: MOCK_APP_PORT, channelId: channelId0}),
            ChannelInfo({connectionId: connectionId1, portId: MOCK_APP_PORT, channelId: channelId1})
        );
    }

    function createMockAppChannel(Channel.Order ordering, string memory connectionId0, string memory connectionId1)
        internal
        returns (ChannelInfo memory, ChannelInfo memory)
    {
        (string memory channelId0, string memory channelId1) = ibcHandler.createLocalhostChannel(
            LocalhostHelper.MsgCreateChannel({
                connectionId0: connectionId0,
                connectionId1: connectionId1,
                portId0: MOCK_APP_PORT,
                portId1: MOCK_APP_PORT,
                ordering: ordering,
                version: MOCK_APP_VERSION
            })
        );
        return (
            ChannelInfo({connectionId: connectionId0, portId: MOCK_APP_PORT, channelId: channelId0}),
            ChannelInfo({connectionId: connectionId1, portId: MOCK_APP_PORT, channelId: channelId1})
        );
    }
}
