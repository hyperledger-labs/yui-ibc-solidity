// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Vm} from "forge-std/Test.sol";
import "./helpers/ICS03TestHelper.t.sol";
import "./helpers/ICS04HandshakeTestHelper.t.sol";
import {IIBCChannelRecvPacket} from "../../../contracts/core/04-channel/IIBCChannel.sol";
import {ICS04PacketMockClientTestHelper, ICS04PacketEventTestHelper} from "./helpers/ICS04PacketTestHelper.t.sol";
import {IIBCChannelAcknowledgePacket} from "../../../contracts/core/04-channel/IIBCChannel.sol";

contract TestICS04Packet is
    ICS03TestHelper,
    ICS04HandshakeMockClientTestHelper,
    ICS04PacketMockClientTestHelper,
    ICS04PacketEventTestHelper
{
    TestableIBCHandler handler;
    TestableIBCHandler counterpartyHandler;
    MockClient client;
    MockClient counterpartyClient;
    IBCMockApp mockApp;
    IBCMockApp counterpartyMockApp;

    string clientId;
    string counterpartyClientId;
    string connectionId;
    string counterpartyConnectionId;

    function setUp() public {
        (TestableIBCHandler _handler, MockClient _client) = ibcHandlerMockClient();
        (TestableIBCHandler _counterpartyHandler, MockClient _counterpartyClient) = ibcHandlerMockClient();
        handler = _handler;
        counterpartyHandler = _counterpartyHandler;
        client = _client;
        counterpartyClient = _counterpartyClient;

        mockApp = new IBCMockApp(handler);
        handler.bindPort("portidone", mockApp);
        counterpartyMockApp = new IBCMockApp(counterpartyHandler);
        counterpartyHandler.bindPort("portidtwo", counterpartyMockApp);

        clientId = createMockClient(handler, 1);
        counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        connectionId = genConnectionId(0);
        counterpartyConnectionId = genConnectionId(1);

        setConnection(
            handler,
            clientId,
            connectionId,
            counterpartyClientId,
            counterpartyConnectionId,
            IBCConnectionLib.defaultIBCVersion()
        );
        setConnection(
            counterpartyHandler,
            counterpartyClientId,
            counterpartyConnectionId,
            clientId,
            connectionId,
            IBCConnectionLib.defaultIBCVersion()
        );
    }

    function testSendPacket() public {
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        vm.recordLogs();
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channelInfo, ChannelInfo memory counterpartyChannelInfo) = handshakeChannel(
                handler,
                counterpartyHandler,
                ChannelInfo("portidone", "", orders[i], "", connectionId),
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                ChannelHandshakeStep.CONFIRM,
                H(1)
            );
            // src -> dst
            assertEq(
                mockAppSendPacket(
                    mockApp,
                    channelInfo,
                    counterpartyChannelInfo,
                    IBCMockLib.MOCK_PACKET_DATA,
                    getHeight(client, clientId, 1),
                    0
                ),
                1
            );
            assertEq(
                mockAppSendPacket(
                    mockApp,
                    channelInfo,
                    counterpartyChannelInfo,
                    IBCMockLib.MOCK_PACKET_DATA,
                    H(0),
                    getTimestamp(client, clientId, 1)
                ),
                2
            );
            assertEq(
                mockAppSendPacket(
                    mockApp,
                    channelInfo,
                    counterpartyChannelInfo,
                    IBCMockLib.MOCK_PACKET_DATA,
                    getHeight(client, clientId, 1),
                    getTimestamp(client, clientId, 1)
                ),
                3
            );
            // OK - timeout height's revision height is zero, but revision number is greater than the latest height's revision number of the client
            assertEq(
                mockAppSendPacket(
                    mockApp, channelInfo, counterpartyChannelInfo, IBCMockLib.MOCK_PACKET_DATA, H(1, 0), 0
                ),
                4
            );
            // dst -> src
            assertEq(
                mockAppSendPacket(
                    counterpartyMockApp,
                    counterpartyChannelInfo,
                    channelInfo,
                    IBCMockLib.MOCK_PACKET_DATA,
                    getHeight(counterpartyClient, counterpartyClientId, 1),
                    0
                ),
                1
            );
            assertEq(
                mockAppSendPacket(
                    counterpartyMockApp,
                    counterpartyChannelInfo,
                    channelInfo,
                    IBCMockLib.MOCK_PACKET_DATA,
                    H(0),
                    getTimestamp(client, clientId, 1)
                ),
                2
            );
            assertEq(
                mockAppSendPacket(
                    counterpartyMockApp,
                    counterpartyChannelInfo,
                    channelInfo,
                    IBCMockLib.MOCK_PACKET_DATA,
                    getHeight(client, clientId, 1),
                    getTimestamp(client, clientId, 1)
                ),
                3
            );
            // OK - timeout height's revision height is zero, but revision number is greater than the latest height's revision number of the client
            assertEq(
                mockAppSendPacket(
                    counterpartyMockApp, counterpartyChannelInfo, channelInfo, IBCMockLib.MOCK_PACKET_DATA, H(1, 0), 0
                ),
                4
            );
        }
    }

    function testInvalidSendPacket() public {
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channelInfo,) = handshakeChannel(
                handler,
                counterpartyHandler,
                ChannelInfo("portidone", "", orders[i], "", connectionId),
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                ChannelHandshakeStep.CONFIRM,
                H(1)
            );

            {
                Height.Data memory timeoutHeight = getHeight(client, clientId, 1);
                // invalid port to send packet
                vm.expectRevert();
                mockApp.sendPacket(IBCMockLib.MOCK_PACKET_DATA, "invalidport", channelInfo.channelId, timeoutHeight, 0);
                // invalid channel to send packet
                vm.expectRevert();
                mockApp.sendPacket(IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, genChannelId(999), timeoutHeight, 0);
            }

            {
                // cannot send expired packet
                Height.Data memory timeoutHeight = getHeight(client, clientId, 0);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );
            }

            {
                // cannot send packet with timeout height and timestamp both are zero
                Height.Data memory timeoutHeight = H(0);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );
            }

            {
                // cannot send packet directly via handler
                Height.Data memory timeoutHeight = getHeight(client, clientId, 1);
                vm.expectRevert();
                handler.sendPacket(
                    channelInfo.portId, channelInfo.channelId, timeoutHeight, 0, IBCMockLib.MOCK_PACKET_DATA
                );
            }

            {
                // cannot send expired packet
                uint64 timeoutTimestamp = getTimestamp(client, clientId, 0);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, H(0), timeoutTimestamp
                );
            }

            {
                // cannot send packet if the client is not active
                Height.Data memory timeoutHeight = getHeight(client, clientId, 1);

                client.setStatus(clientId, ILightClient.ClientStatus.Frozen);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );

                client.setStatus(clientId, ILightClient.ClientStatus.Expired);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );

                client.setStatus(clientId, ILightClient.ClientStatus.Active);
            }
        }
    }

    function testRecvPacket() public {
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        // increase the channel sequence in order to use "channel-0" as invalid channel in tests
        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_UNORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.CONFIRM,
            H(1)
        );
        vm.recordLogs();
        for (uint256 i = 0; i < orders.length; i++) {
            for (uint256 j = 0; j < 2; j++) {
                Height.Data memory timeoutHeight = j == 0 ? H(0) : getHeight(client, clientId, 1);
                uint64 timeoutTimestamp = j == 1 ? getTimestamp(client, clientId, 1) : 0;

                (ChannelInfo memory channelInfo, ChannelInfo memory counterpartyChannelInfo) = handshakeChannel(
                    handler,
                    counterpartyHandler,
                    ChannelInfo("portidone", "", orders[i], "", connectionId),
                    ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                    ChannelHandshakeStep.CONFIRM,
                    H(1)
                );
                uint64 nextSeqRecv = 1;

                // OK
                {
                    Packet memory p0 = createPacket(
                        channelInfo,
                        counterpartyChannelInfo,
                        nextSeqRecv,
                        IBCMockLib.MOCK_PACKET_DATA,
                        timeoutHeight,
                        timeoutTimestamp
                    );
                    counterpartyHandler.recvPacket(msgPacketRecv(p0, H(1)));
                    nextSeqRecv++;
                    validateRecvPacketPostState(counterpartyHandler, counterpartyChannelInfo, nextSeqRecv);
                    Vm.Log[] memory logs = vm.getRecordedLogs();
                    Packet memory p1 = getLastRecvPacket(counterpartyHandler, logs);
                    assertEq(abi.encode(p0), abi.encode(p1));
                    WriteAcknolwedgement memory ack = getLastWrittenAcknowledgement(counterpartyHandler, logs);
                    assertEq(ack.acknowledgement, IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON);

                    // same packet relay must be failed
                    IIBCChannelRecvPacket.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(1));
                    vm.expectRevert();
                    counterpartyHandler.recvPacket(msg_);
                }
                // OK - Async-ack
                {
                    Packet memory p0 = createPacket(
                        channelInfo,
                        counterpartyChannelInfo,
                        nextSeqRecv,
                        IBCMockLib.MOCK_ASYNC_PACKET_DATA,
                        timeoutHeight,
                        timeoutTimestamp
                    );
                    counterpartyHandler.recvPacket(msgPacketRecv(p0, H(1)));
                    nextSeqRecv++;
                    Vm.Log[] memory logs = vm.getRecordedLogs();
                    Packet memory p1 = getLastRecvPacket(counterpartyHandler, logs);
                    assertEq(abi.encode(p0), abi.encode(p1));
                    assertTrue(findWrittenAcknowledgement(counterpartyHandler, logs).length == 0);
                    // same packet relay must be failed
                    IIBCChannelRecvPacket.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(1));
                    vm.expectRevert();
                    counterpartyHandler.recvPacket(msg_);
                    counterpartyMockApp.writeAcknowledgement(
                        counterpartyChannelInfo.portId, counterpartyChannelInfo.channelId, 2
                    );
                    vm.expectRevert();
                    counterpartyMockApp.writeAcknowledgement(
                        counterpartyChannelInfo.portId, counterpartyChannelInfo.channelId, 2
                    );
                }

                // valid source and dest but proof is invalid
                {
                    Packet memory p0 = createPacket(
                        channelInfo,
                        counterpartyChannelInfo,
                        nextSeqRecv,
                        IBCMockLib.MOCK_PACKET_DATA,
                        timeoutHeight,
                        timeoutTimestamp
                    );
                    IIBCChannelRecvPacket.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(1));
                    msg_.proof = hex"01";
                    vm.expectRevert();
                    counterpartyHandler.recvPacket(msg_);
                }
                // valid proof but source is invalid
                {
                    Packet memory p0 = createPacket(
                        channelInfo,
                        counterpartyChannelInfo,
                        nextSeqRecv,
                        IBCMockLib.MOCK_PACKET_DATA,
                        timeoutHeight,
                        timeoutTimestamp
                    );
                    IIBCChannelRecvPacket.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(1));
                    string memory tmp = channelInfo.channelId;
                    channelInfo.channelId = genChannelId(0);
                    msg_.proof = provePacketCommitment(
                        createPacket(
                            channelInfo,
                            counterpartyChannelInfo,
                            nextSeqRecv,
                            IBCMockLib.MOCK_PACKET_DATA,
                            timeoutHeight,
                            timeoutTimestamp
                        ),
                        msg_.proofHeight
                    );
                    channelInfo.channelId = tmp;
                    vm.expectRevert();
                    counterpartyHandler.recvPacket(msg_);
                }
                // valid proof but destination is invalid
                {
                    string memory tmp = counterpartyChannelInfo.channelId;
                    counterpartyChannelInfo.channelId = genChannelId(0);
                    Packet memory p0 = createPacket(
                        channelInfo,
                        counterpartyChannelInfo,
                        nextSeqRecv,
                        IBCMockLib.MOCK_PACKET_DATA,
                        timeoutHeight,
                        timeoutTimestamp
                    );
                    counterpartyChannelInfo.channelId = tmp;
                    IIBCChannelRecvPacket.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(1));
                    vm.expectRevert();
                    counterpartyHandler.recvPacket(msg_);
                }
                {
                    IIBCChannelRecvPacket.MsgPacketRecv memory msg_ = msgPacketRecv(
                        createPacket(
                            channelInfo,
                            counterpartyChannelInfo,
                            nextSeqRecv + 1,
                            IBCMockLib.MOCK_PACKET_DATA,
                            timeoutHeight,
                            timeoutTimestamp
                        ),
                        H(1)
                    );
                    if (orders[i] == Channel.Order.ORDER_ORDERED) {
                        // unordered packet must be failed
                        vm.expectRevert();
                        counterpartyHandler.recvPacket(msg_);
                    } else if (orders[i] == Channel.Order.ORDER_UNORDERED) {
                        // unordered packet must be succeed
                        counterpartyHandler.recvPacket(msg_);
                    }
                }
            }
        }
    }

    function testRecvPacketTimeoutHeight() public {
        vm.recordLogs();
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channelInfo,) = handshakeChannel(
                handler,
                counterpartyHandler,
                ChannelInfo("portidone", "", orders[i], "", connectionId),
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                ChannelHandshakeStep.CONFIRM,
                H(1)
            );
            mockApp.sendPacket(
                IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, H(block.number + 1), 0
            );
            Packet memory p0 =
                getLastSentPacket(handler, channelInfo.portId, channelInfo.channelId, vm.getRecordedLogs());
            {
                IIBCChannelPacketTimeout.MsgTimeoutPacket memory msg1 =
                    msgTimeoutPacket(channelInfo.ordering, p0, H(block.number));
                vm.expectRevert(abi.encodeWithSelector(IIBCChannelErrors.IBCChannelTimeoutNotReached.selector));
                handler.timeoutPacket(msg1);
            }

            IIBCChannelPacketSendRecv.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(block.number));
            vm.roll(block.number + 1);
            vm.expectRevert(
                abi.encodeWithSelector(
                    IIBCChannelErrors.IBCChannelTimeoutPacketHeight.selector,
                    block.number,
                    p0.timeoutHeight.revision_height
                )
            );
            counterpartyHandler.recvPacket(msg_);
            client.updateClient(clientId, mockClientHeader(uint64(block.number)));
            // timeout on source chain
            handler.timeoutPacket(msgTimeoutPacket(channelInfo.ordering, p0, H(block.number)));
            if (orders[i] == Channel.Order.ORDER_ORDERED) {
                ensureChannelState(handler, channelInfo, Channel.State.STATE_CLOSED);
            } else if (orders[i] == Channel.Order.ORDER_UNORDERED) {
                ensureChannelState(handler, channelInfo, Channel.State.STATE_OPEN);
            }
            assertEq(handler.getPacketCommitment(channelInfo.portId, channelInfo.channelId, p0.sequence), bytes32(0));
        }
    }

    function testRecvPacketTimeoutTimestamp() public {
        vm.recordLogs();
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        assertEq(block.timestamp, 1);
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channelInfo,) = handshakeChannel(
                handler,
                counterpartyHandler,
                ChannelInfo("portidone", "", orders[i], "", connectionId),
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                ChannelHandshakeStep.CONFIRM,
                H(1)
            );
            mockApp.sendPacket(
                IBCMockLib.MOCK_PACKET_DATA,
                channelInfo.portId,
                channelInfo.channelId,
                H(0),
                uint64(block.timestamp * 1e9 + 1)
            );
            Packet memory p0 =
                getLastSentPacket(handler, channelInfo.portId, channelInfo.channelId, vm.getRecordedLogs());
            {
                IIBCChannelPacketTimeout.MsgTimeoutPacket memory msg1 =
                    msgTimeoutPacket(channelInfo.ordering, p0, H(block.number));
                vm.expectRevert(abi.encodeWithSelector(IIBCChannelErrors.IBCChannelTimeoutNotReached.selector));
                handler.timeoutPacket(msg1);
            }

            IIBCChannelPacketSendRecv.MsgPacketRecv memory msg_ = msgPacketRecv(p0, H(block.number));
            vm.warp(block.timestamp + 1);
            vm.roll(block.number + 1);
            vm.expectRevert(
                abi.encodeWithSelector(
                    IIBCChannelErrors.IBCChannelTimeoutPacketTimestamp.selector,
                    block.timestamp * 1e9,
                    p0.timeoutTimestamp
                )
            );
            counterpartyHandler.recvPacket(msg_);
            client.updateClient(clientId, mockClientHeader(uint64(block.number)));
            // timeout on source chain
            handler.timeoutPacket(msgTimeoutPacket(channelInfo.ordering, p0, H(block.number)));
            if (orders[i] == Channel.Order.ORDER_ORDERED) {
                ensureChannelState(handler, channelInfo, Channel.State.STATE_CLOSED);
            } else if (orders[i] == Channel.Order.ORDER_UNORDERED) {
                ensureChannelState(handler, channelInfo, Channel.State.STATE_OPEN);
            }
            assertEq(handler.getPacketCommitment(channelInfo.portId, channelInfo.channelId, p0.sequence), bytes32(0));
        }
    }

    function testTimeoutOnClose() public {
        mockApp.allowCloseChannel(true);
        counterpartyMockApp.allowCloseChannel(true);
        vm.recordLogs();
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channelInfo, ChannelInfo memory counterpartyChannelInfo) = handshakeChannel(
                handler,
                counterpartyHandler,
                ChannelInfo("portidone", "", orders[i], "", connectionId),
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                ChannelHandshakeStep.CONFIRM,
                H(1)
            );
            mockApp.sendPacket(IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, H(2), 0);
            Packet memory p0 =
                getLastSentPacket(handler, channelInfo.portId, channelInfo.channelId, vm.getRecordedLogs());
            {
                IIBCChannelPacketTimeout.MsgTimeoutOnClose memory msg_ =
                    msgTimeoutOnClose(counterpartyHandler, orders[i], p0, H(1));
                vm.expectRevert();
                handler.timeoutOnClose(msg_);
            }
            counterpartyHandler.channelCloseInit(msgChannelCloseInit(counterpartyChannelInfo));
            handler.timeoutOnClose(msgTimeoutOnClose(counterpartyHandler, orders[i], p0, H(1)));
        }
    }

    function testAcknowledgementPacket() public {
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        vm.recordLogs();
        for (uint256 i = 0; i < orders.length; i++) {
            (ChannelInfo memory channelInfo, ChannelInfo memory counterpartyChannelInfo) = handshakeChannel(
                handler,
                counterpartyHandler,
                ChannelInfo("portidone", "", orders[i], "", connectionId),
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId),
                ChannelHandshakeStep.CONFIRM,
                H(1)
            );

            // OK
            {
                Packet memory p0 = createPacket(
                    channelInfo,
                    counterpartyChannelInfo,
                    1,
                    IBCMockLib.MOCK_PACKET_DATA,
                    H(0),
                    getTimestamp(client, clientId, 1)
                );
                handler.setPacketCommitment(p0);
                handler.acknowledgePacket(
                    msgPacketAcknowledgement(p0, IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON, H(1))
                );
                assertEq(handler.getPacketCommitment(channelInfo.portId, channelInfo.channelId, 1), bytes32(0));
            }
            // Receiving duplicate acks must be failed
            {
                Packet memory p0 = createPacket(
                    channelInfo,
                    counterpartyChannelInfo,
                    2,
                    IBCMockLib.MOCK_FAIL_PACKET_DATA,
                    H(0),
                    getTimestamp(client, clientId, 1)
                );
                handler.setPacketCommitment(p0);
                IIBCChannelAcknowledgePacket.MsgPacketAcknowledgement memory msg_ =
                    msgPacketAcknowledgement(p0, IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON, H(1));
                handler.acknowledgePacket(msg_);
                vm.expectRevert();
                handler.acknowledgePacket(msg_);
            }
            // Receiving an ack with no corresponding packet commitment must be failed
            {
                Packet memory p0 = createPacket(
                    channelInfo,
                    counterpartyChannelInfo,
                    3,
                    IBCMockLib.MOCK_PACKET_DATA,
                    H(0),
                    getTimestamp(client, clientId, 1)
                );
                IIBCChannelAcknowledgePacket.MsgPacketAcknowledgement memory msg_ =
                    msgPacketAcknowledgement(p0, IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON, H(1));
                vm.expectRevert();
                handler.acknowledgePacket(msg_);
            }
        }
    }

    function mockAppSendPacket(
        IBCMockApp app,
        ChannelInfo memory src,
        ChannelInfo memory dst,
        bytes memory message,
        Height.Data memory timeoutHeight,
        uint64 timeoutTimestamp
    ) internal returns (uint64) {
        IIBCHandler h = IIBCHandler(app.ibcAddress());
        uint64 sequence = app.sendPacket(message, src.portId, src.channelId, timeoutHeight, timeoutTimestamp);
        assertEq(h.getNextSequenceSend(src.portId, src.channelId), sequence + 1);
        Packet memory packet = getLastSentPacket(h, src.portId, src.channelId, vm.getRecordedLogs());
        assertEq(
            abi.encode(packet),
            abi.encode(
                Packet({
                    sequence: sequence,
                    sourcePort: src.portId,
                    sourceChannel: src.channelId,
                    destinationPort: dst.portId,
                    destinationChannel: dst.channelId,
                    timeoutHeight: timeoutHeight,
                    timeoutTimestamp: timeoutTimestamp,
                    data: message
                })
            )
        );
        return sequence;
    }
}
