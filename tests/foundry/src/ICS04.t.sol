// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./ICS03.t.sol";
import "../../../contracts/apps/mock/IBCMockApp.sol";

abstract contract TestICS04Helper is TestIBCBase, TestMockClientHelper {
    struct ChannelInfo {
        string portId;
        string channelId;
        Channel.Order ordering;
        string version;
        string connectionId;
    }

    function setConnection(
        TestableIBCHandler handler,
        string memory clientId,
        string memory connectionId,
        string memory counterpartyClientId,
        string memory counterpartyConnectionId,
        Version.Data memory version
    ) internal {
        Version.Data[] memory versions = new Version.Data[](1);
        versions[0] = version;
        handler.setConnection(
            connectionId,
            ConnectionEnd.Data({
                client_id: clientId,
                state: ConnectionEnd.State.STATE_OPEN,
                delay_period: 0,
                versions: versions,
                counterparty: Counterparty.Data({
                    client_id: counterpartyClientId,
                    connection_id: counterpartyConnectionId,
                    prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
                })
            })
        );
    }

    function msgChannelOpenInit(ChannelInfo memory channelInfo, string memory counterpartyPortId)
        internal
        pure
        returns (IBCMsgs.MsgChannelOpenInit memory)
    {
        return IBCMsgs.MsgChannelOpenInit({
            portId: channelInfo.portId,
            channel: Channel.Data({
                state: Channel.State.STATE_INIT,
                ordering: channelInfo.ordering,
                counterparty: ChannelCounterparty.Data({port_id: counterpartyPortId, channel_id: ""}),
                connection_hops: newConnectionHops(channelInfo.connectionId),
                version: channelInfo.version
            })
        });
    }

    function msgChannelOpenTry(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight
    ) internal pure returns (IBCMsgs.MsgChannelOpenTry memory) {
        assert(bytes(channelInfo.channelId).length == 0 && bytes(channelInfo.version).length == 0);
        bytes memory proofInit = genMockChannelStateProof(
            proofHeight,
            counterpartyChannelInfo.portId,
            counterpartyChannelInfo.channelId,
            Channel.Data({
                state: Channel.State.STATE_INIT,
                ordering: counterpartyChannelInfo.ordering,
                counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: ""}),
                connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                version: counterpartyChannelInfo.version
            })
        );
        return IBCMsgs.MsgChannelOpenTry({
            portId: channelInfo.portId,
            channel: Channel.Data({
                state: Channel.State.STATE_TRYOPEN,
                ordering: channelInfo.ordering,
                counterparty: ChannelCounterparty.Data({
                    port_id: counterpartyChannelInfo.portId,
                    channel_id: counterpartyChannelInfo.channelId
                }),
                connection_hops: newConnectionHops(channelInfo.connectionId),
                version: ""
            }),
            counterpartyVersion: counterpartyChannelInfo.version,
            proofInit: proofInit,
            proofHeight: proofHeight
        });
    }

    function msgChannelOpenAck(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight
    ) internal pure returns (IBCMsgs.MsgChannelOpenAck memory) {
        return IBCMsgs.MsgChannelOpenAck({
            portId: channelInfo.portId,
            channelId: channelInfo.channelId,
            counterpartyVersion: counterpartyChannelInfo.version,
            counterpartyChannelId: counterpartyChannelInfo.channelId,
            proofTry: genMockChannelStateProof(
                proofHeight,
                counterpartyChannelInfo.portId,
                counterpartyChannelInfo.channelId,
                Channel.Data({
                    state: Channel.State.STATE_TRYOPEN,
                    ordering: counterpartyChannelInfo.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: channelInfo.channelId}),
                    connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                    version: counterpartyChannelInfo.version
                })
                ),
            proofHeight: proofHeight
        });
    }

    function msgChannelOpenConfirm(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight
    ) internal pure returns (IBCMsgs.MsgChannelOpenConfirm memory) {
        return IBCMsgs.MsgChannelOpenConfirm({
            portId: channelInfo.portId,
            channelId: channelInfo.channelId,
            proofAck: genMockChannelStateProof(
                proofHeight,
                counterpartyChannelInfo.portId,
                counterpartyChannelInfo.channelId,
                Channel.Data({
                    state: Channel.State.STATE_OPEN,
                    ordering: counterpartyChannelInfo.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: channelInfo.channelId}),
                    connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                    version: counterpartyChannelInfo.version
                })
                ),
            proofHeight: proofHeight
        });
    }

    function newConnectionHops(string memory connectionId) internal pure returns (string[] memory) {
        string[] memory connectionHops = new string[](1);
        connectionHops[0] = connectionId;
        return connectionHops;
    }

    function validatePostStateAfterChanOpenInit(
        IBCMsgs.MsgChannelOpenInit memory msg_,
        Channel.Data memory channel,
        string memory version
    ) internal {
        assertTrue(channel.state == Channel.State.STATE_INIT);
        assertTrue(channel.ordering == msg_.channel.ordering);
        assertEq(channel.counterparty.port_id, msg_.channel.counterparty.port_id);
        assertEq(channel.counterparty.channel_id, "");
        assertEq(channel.connection_hops.length, msg_.channel.connection_hops.length);
        for (uint256 i = 0; i < channel.connection_hops.length; i++) {
            assertEq(channel.connection_hops[i], msg_.channel.connection_hops[i]);
        }
        assertEq(channel.version, version);
    }

    function validatePostStateAfterChanOpenTry(
        IBCMsgs.MsgChannelOpenTry memory msg_,
        Channel.Data memory channel,
        string memory version
    ) internal {
        assertTrue(channel.state == Channel.State.STATE_TRYOPEN);
        assertTrue(channel.ordering == msg_.channel.ordering);
        assertEq(channel.counterparty.port_id, msg_.channel.counterparty.port_id);
        assertEq(channel.counterparty.channel_id, msg_.channel.counterparty.channel_id);
        assertEq(channel.connection_hops.length, msg_.channel.connection_hops.length);
        for (uint256 i = 0; i < channel.connection_hops.length; i++) {
            assertEq(channel.connection_hops[i], msg_.channel.connection_hops[i]);
        }
        assertEq(channel.version, version);
    }
}

contract TestICS04Handshake is TestIBCBase, TestMockClientHelper, TestICS03Helper, TestICS04Helper {
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
        handler.bindPort("portidone", address(mockApp));
        counterpartyMockApp = new IBCMockApp(counterpartyHandler);
        counterpartyHandler.bindPort("portidtwo", address(counterpartyMockApp));
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
            IBCMsgs.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            (string memory channelId, string memory version) = handler.channelOpenInit(msg_);
            assertEq(channelId, "channel-0");
            assertEq(version, MOCK_APP_VERSION);
            (Channel.Data memory channel,) = handler.getChannel("portidone", channelId);
            validatePostStateAfterChanOpenInit(msg_, channel, MOCK_APP_VERSION);
        }
        {
            IBCMsgs.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_ORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            (string memory channelId, string memory version) = handler.channelOpenInit(msg_);
            assertEq(channelId, "channel-1");
            assertEq(version, MOCK_APP_VERSION);
            (Channel.Data memory channel,) = handler.getChannel("portidone", channelId);
            validatePostStateAfterChanOpenInit(msg_, channel, MOCK_APP_VERSION);
        }
        {
            IBCMsgs.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
                ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, MOCK_APP_VERSION, connectionId), "portidtwo"
            );
            (string memory channelId, string memory version) = handler.channelOpenInit(msg_);
            assertEq(channelId, "channel-2");
            assertEq(version, MOCK_APP_VERSION);
            (Channel.Data memory channel,) = handler.getChannel("portidone", channelId);
            validatePostStateAfterChanOpenInit(msg_, channel, MOCK_APP_VERSION);
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
            IBCMsgs.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
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
            IBCMsgs.MsgChannelOpenInit memory msg_ = msgChannelOpenInit(
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

        {
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 1)
            );
            (string memory channelId, string memory version) = handler.channelOpenTry(msg_);
            assertEq(channelId, "channel-0");
            assertEq(version, MOCK_APP_VERSION);
            (Channel.Data memory channel,) = handler.getChannel("portidone", channelId);
            validatePostStateAfterChanOpenTry(msg_, channel, MOCK_APP_VERSION);
        }
        {
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 1)
            );
            (string memory channelId, string memory version) = handler.channelOpenTry(msg_);
            assertEq(channelId, "channel-1");
            assertEq(version, MOCK_APP_VERSION);
            (Channel.Data memory channel,) = handler.getChannel("portidone", channelId);
            validatePostStateAfterChanOpenTry(msg_, channel, MOCK_APP_VERSION);
        }
    }

    function testInvalidChanOpenTry() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        string memory connectionId = genConnectionId(0);
        string memory counterpartyConnectionId = genConnectionId(1);

        Version.Data memory orderedVersion = orderedIBCVersion();

        // connection does not exist
        {
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 1)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedIBCVersion()
            );
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 1)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        // invalid proof height
        {
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 2)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        // not binded port
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedVersion
            );
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 1)
            );
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }

        // app version mismatch between proof and msg
        {
            setConnection(
                handler, clientId, connectionId, counterpartyClientId, counterpartyConnectionId, orderedVersion
            );
            IBCMsgs.MsgChannelOpenTry memory msg_ = msgChannelOpenTry(
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
                H(0, 1)
            );
            msg_.counterpartyVersion = "mockapp-2";
            vm.expectRevert();
            handler.channelOpenTry(msg_);
        }
    }
}
