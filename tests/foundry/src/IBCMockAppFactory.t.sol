// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "./helpers/IBCTestHelper.t.sol";
import {Vm} from "forge-std/Test.sol";
import {IBCMockLib} from "../../../contracts/apps/mock/IBCMockLib.sol";
import {IIBCChannelRecvPacket, IIBCChannelAcknowledgePacket} from "../../../contracts/core/04-channel/IIBCChannel.sol";
import {LocalhostClientLib} from "../../../contracts/clients/09-localhost/LocalhostClient.sol";
import {LocalhostHelper} from "../../../contracts/clients/09-localhost/LocalhostHelper.sol";
import {ICS04PacketEventTestHelper} from "./helpers/ICS04PacketTestHelper.t.sol";
import {IBCMockAppFactory} from "../../../contracts/apps/mock/IBCMockAppFactory.sol";

contract IBCMockAppFactoryTest is IBCTestHelper, ICS04PacketEventTestHelper {
    using LocalhostHelper for TestableIBCHandler;

    string internal constant MOCK_APP_PORT = "mockapp";
    string internal constant MOCK_APP_VERSION = "mockapp-1";

    TestableIBCHandler ibcHandler;
    IBCMockAppFactory mockAppFactory;

    struct ChannelInfo {
        string connectionId;
        string portId;
        string channelId;
    }

    function setUp() public {
        ibcHandler = defaultIBCHandler();
        mockAppFactory = new IBCMockAppFactory(ibcHandler);
        ibcHandler.bindPort(MOCK_APP_PORT, mockAppFactory);
        ibcHandler.registerLocalhostClient();
        ibcHandler.createLocalhostClient();
    }

    function testHandshake() public {
        Channel.Order[2] memory orders = [Channel.Order.ORDER_UNORDERED, Channel.Order.ORDER_ORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channel0, ChannelInfo memory channel1) = createMockAppChannel(orders[i]);
            assertEq(address(ibcHandler.getIBCModuleByPort(channel0.portId)), address(mockAppFactory));
            assertEq(address(ibcHandler.getIBCModuleByChannel(channel0.portId, channel0.channelId)), address(mockAppFactory.lookupApp(channel0.channelId)));
            assertEq(address(ibcHandler.getIBCModuleByPort(channel1.portId)), address(mockAppFactory));
            assertEq(address(ibcHandler.getIBCModuleByChannel(channel1.portId, channel1.channelId)), address(mockAppFactory.lookupApp(channel1.channelId)));
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

    struct RelayCase {
        bytes packetData;
        bool async;
        bytes acknowledgement;
    }

    function sendAndRelay(ChannelInfo memory ca, ChannelInfo memory cb, Channel.Order ordering, RelayCase memory rc)
        private
    {
        uint64 sequence = mockAppFactory.lookupApp(ca.channelId).sendPacket(rc.packetData, H(uint64(getBlockNumber(1))), 0);
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
            mockAppFactory.lookupApp(cb.channelId).writeAcknowledgement(sequence);
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
}