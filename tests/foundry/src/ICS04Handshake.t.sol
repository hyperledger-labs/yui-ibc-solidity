// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import "../../../contracts/apps/mock/IBCMockApp.sol";
import "./helpers/ICS03TestHelper.t.sol";
import "./helpers/ICS04HandshakeTestHelper.t.sol";

contract TestICS04Handshake is ICS03TestHelper, ICS04HandshakeMockClientTestHelper {
    TestableIBCHandler handler;
    TestableIBCHandler counterpartyHandler;
    IBCMockApp mockApp;
    IBCMockApp counterpartyMockApp;
    string internal constant MOCK_APP_VERSION = "mockapp-1";

    function setUp() public {
        (TestableIBCHandler _handler,) = ibcHandlerMockClient();
        (TestableIBCHandler _counterpartyHandler,) = ibcHandlerMockClient();
        handler = _handler;
        counterpartyHandler = _counterpartyHandler;

        mockApp = new IBCMockApp(handler);
        handler.bindPort("portidone", mockApp);
        counterpartyMockApp = new IBCMockApp(counterpartyHandler);
        counterpartyHandler.bindPort("portidtwo", counterpartyMockApp);
    }

    function testBindPort() public {
        vm.expectRevert();
        handler.bindPort("portidone", mockApp);
        vm.expectRevert();
        handler.bindPort("portidone", IIBCModule(address(0x01)));
        vm.expectRevert();
        handler.bindPort("", IIBCModule(address(0x01)));
        vm.expectRevert();
        handler.bindPort("portidone", IIBCModule(address(0)));
    }

    function testChanOpenInit() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        {
            IIBCChannelHandshake.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            (string memory channelId, string memory version) = handler.channelOpenInit(msg_);
            assertEq(channelId, "channel-0");
            assertEq(version, MOCK_APP_VERSION);
            validatePostStateAfterChanOpenInit(handler, msg_, channelId, MOCK_APP_VERSION);
        }
        {
            IIBCChannelHandshake.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            (string memory channelId, string memory version) = handler.channelOpenInit(msg_);
            assertEq(channelId, "channel-1");
            assertEq(version, MOCK_APP_VERSION);
            validatePostStateAfterChanOpenInit(handler, msg_, channelId, MOCK_APP_VERSION);
        }
        {
            IIBCChannelHandshake.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            (string memory channelId, string memory version) = handler.channelOpenInit(msg_);
            assertEq(channelId, "channel-2");
            assertEq(version, MOCK_APP_VERSION);
            validatePostStateAfterChanOpenInit(handler, msg_, channelId, MOCK_APP_VERSION);
        }
    }

    function testInvalidChanOpenInit() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

        Version.Data memory orderedVersion = orderedIBCVersion();
        Version.Data memory unorderedVersion = unorderedIBCVersion();

        // connection does not exist
        {
            vm.expectRevert();
            handler.channelOpenInit(
                msgChannelOpenInit(
                    ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId),
                    "portidtwo"
                )
            );
        }
        // not binded port
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedVersion
            );
            vm.expectRevert();
            handler.channelOpenInit(
                msgChannelOpenInit(
                    ChannelInfo("portidthree", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId),
                    "portidtwo"
                )
            );
        }
        // connection does not support UNORDERED channel
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedVersion
            );
            vm.expectRevert();
            handler.channelOpenInit(
                msgChannelOpenInit(
                    ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, MOCK_APP_VERSION, connectionId),
                    "portidtwo"
                )
            );
        }
        // connection does not support ORDERED channel
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, unorderedVersion
            );
            vm.expectRevert();
            handler.channelOpenInit(
                msgChannelOpenInit(
                    ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId),
                    "portidtwo"
                )
            );
        }
        // channel state must be STATE_INIT
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, unorderedVersion
            );
            IIBCChannelHandshake.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            msg_.channel.state = Channel.State.STATE_TRYOPEN;
            vm.expectRevert();
            handler.channelOpenInit(msg_);
        }
        // channel.counterparty.channelId must be empty
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, unorderedVersion
            );
            IIBCChannelHandshake.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            msg_.channel.counterparty.channel_id = "channel-0";
            vm.expectRevert();
            handler.channelOpenInit(msg_);
        }
    }

    function testChanOpenTry() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo({
                portId: "portidone",
                channelId: "",
                ordering: Channel.Order.ORDER_ORDERED,
                version: "",
                connectionId: connectionId
            }),
            ChannelInfo({
                portId: "portidtwo",
                channelId: "",
                ordering: Channel.Order.ORDER_ORDERED,
                version: "",
                connectionId: counterpartyConnectionId
            }),
            ChannelHandshakeStep.TRY,
            H(1)
        );

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo({
                portId: "portidone",
                channelId: "",
                ordering: Channel.Order.ORDER_UNORDERED,
                version: "",
                connectionId: connectionId
            }),
            ChannelInfo({
                portId: "portidtwo",
                channelId: "",
                ordering: Channel.Order.ORDER_UNORDERED,
                version: "",
                connectionId: counterpartyConnectionId
            }),
            ChannelHandshakeStep.TRY,
            H(1)
        );
    }

    function testInvalidChanOpenTry() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

        Version.Data memory orderedVersion = orderedIBCVersion();

        // connection does not exist
        {
            IIBCChannelHandshake.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
                ChannelInfo({
                    portId: "portidone",
                    channelId: "",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: "",
                    connectionId: connectionId
                }),
                ChannelInfo({
                    portId: "portidtwo",
                    channelId: "channel-0",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: MOCK_APP_VERSION,
                    connectionId: counterpartyConnectionId
                }),
                H(1)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedIBCVersion()
            );
            IIBCChannelHandshake.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
                ChannelInfo({
                    portId: "portidone",
                    channelId: "",
                    ordering: Channel.Order.ORDER_UNORDERED,
                    version: "",
                    connectionId: connectionId
                }),
                ChannelInfo({
                    portId: "portidtwo",
                    channelId: "channel-0",
                    ordering: Channel.Order.ORDER_UNORDERED,
                    version: MOCK_APP_VERSION,
                    connectionId: counterpartyConnectionId
                }),
                H(1)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        // invalid proof height
        {
            IIBCChannelHandshake.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
                ChannelInfo({
                    portId: "portidone",
                    channelId: "",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: "",
                    connectionId: connectionId
                }),
                ChannelInfo({
                    portId: "portidtwo",
                    channelId: "channel-0",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: MOCK_APP_VERSION,
                    connectionId: counterpartyConnectionId
                }),
                H(2)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        // not binded port
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedVersion
            );
            IIBCChannelHandshake.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
                ChannelInfo({
                    portId: "portidthree",
                    channelId: "",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: "",
                    connectionId: connectionId
                }),
                ChannelInfo({
                    portId: "portidtwo",
                    channelId: "channel-0",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: MOCK_APP_VERSION,
                    connectionId: counterpartyConnectionId
                }),
                H(1)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        // app version mismatch between proof and msg
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedVersion
            );
            IIBCChannelHandshake.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
                ChannelInfo({
                    portId: "portidone",
                    channelId: "",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: "",
                    connectionId: connectionId
                }),
                ChannelInfo({
                    portId: "portidtwo",
                    channelId: "channel-0",
                    ordering: Channel.Order.ORDER_ORDERED,
                    version: MOCK_APP_VERSION,
                    connectionId: counterpartyConnectionId
                }),
                H(1)
            );
            msg_.counterpartyVersion = "mockapp-2";
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }
    }

    function testChanOpenAck() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_ORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.ACK,
            H(1)
        );

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_UNORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.ACK,
            H(1)
        );
    }

    function testInvalidChanOpenAck() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        (ChannelInfo memory channelInfo, ChannelInfo memory counterpartyChannelInfo) = handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_ORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.TRY,
            H(1)
        );

        // app version mismatch between proof and msg
        {
            IIBCChannelHandshake.MsgChannelOpenAck memory msg_ =
                msgChannelOpenAck(channelInfo, counterpartyChannelInfo, H(1));
            msg_.counterpartyVersion = "mockapp-2";
            vm.expectRevert();
            handler.channelOpenAck(msg_);
        }
        // invalid proof height
        {
            IIBCChannelHandshake.MsgChannelOpenAck memory msg_ =
                msgChannelOpenAck(channelInfo, counterpartyChannelInfo, H(2));
            vm.expectRevert();
            handler.channelOpenAck(msg_);
        }
    }

    function testChanOpenConfirm() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_ORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.CONFIRM,
            H(1)
        );

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_UNORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.CONFIRM,
            H(1)
        );
    }

    function testInvalidChanOpenConfirm() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        (ChannelInfo memory channelInfo, ChannelInfo memory counterpartyChannelInfo) = handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_ORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.ACK,
            H(1)
        );

        // invalid proof height
        {
            IIBCChannelHandshake.MsgChannelOpenConfirm memory msg_ =
                msgChannelOpenConfirm(counterpartyChannelInfo, channelInfo, H(2));
            vm.expectRevert();
            counterpartyHandler.channelOpenConfirm(msg_);
        }
    }

    function testChanClose() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

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

        ChannelHandshakeStep[4] memory steps = [
            ChannelHandshakeStep.INIT,
            ChannelHandshakeStep.TRY,
            ChannelHandshakeStep.ACK,
            ChannelHandshakeStep.CONFIRM
        ];
        Channel.Order[2] memory orders = [Channel.Order.ORDER_ORDERED, Channel.Order.ORDER_UNORDERED];
        for (uint256 i = 0; i < steps.length; i++) {
            for (uint256 j = 0; j < orders.length; j++) {
                (ChannelInfo memory channel, ChannelInfo memory counterpartyChannel) = handshakeChannel(
                    handler,
                    counterpartyHandler,
                    ChannelInfo("portidone", "", orders[j], "", connectionId),
                    ChannelInfo("portidtwo", "", orders[j], "", counterpartyConnectionId),
                    steps[i],
                    H(1)
                );

                {
                    IIBCChannelHandshake.MsgChannelCloseConfirm memory msg_ =
                        msgChannelCloseConfirm(counterpartyChannel, channel, H(1), Channel.State.STATE_OPEN);
                    mockApp.allowCloseChannel(true);
                    vm.expectRevert();
                    handler.channelCloseConfirm(msg_);
                }
                {
                    IIBCChannelHandshake.MsgChannelCloseInit memory msg_ = msgChannelCloseInit(channel);
                    mockApp.allowCloseChannel(false);
                    vm.expectRevert();
                    handler.channelCloseInit(msg_);
                    mockApp.allowCloseChannel(true);
                    handler.channelCloseInit(msg_);
                    ensureChannelState(handler, channel, Channel.State.STATE_CLOSED);
                }
                if (steps[i] == ChannelHandshakeStep.INIT) {
                    continue;
                }
                {
                    IIBCChannelHandshake.MsgChannelCloseConfirm memory msg_ =
                        msgChannelCloseConfirm(counterpartyChannel, channel, H(1));
                    counterpartyMockApp.allowCloseChannel(false);
                    vm.expectRevert();
                    counterpartyHandler.channelCloseConfirm(msg_);
                    counterpartyMockApp.allowCloseChannel(true);
                    counterpartyHandler.channelCloseConfirm(msg_);
                    ensureChannelState(counterpartyHandler, counterpartyChannel, Channel.State.STATE_CLOSED);
                }
            }
        }
    }
}
