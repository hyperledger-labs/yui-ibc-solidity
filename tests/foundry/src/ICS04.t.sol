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
        returns (IIBCChannelHandshake.MsgChannelOpenInit memory)
    {
        return IIBCChannelHandshake.MsgChannelOpenInit({
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
    ) internal pure returns (IIBCChannelHandshake.MsgChannelOpenTry memory) {
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
        return IIBCChannelHandshake.MsgChannelOpenTry({
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
    ) internal pure returns (IIBCChannelHandshake.MsgChannelOpenAck memory) {
        return IIBCChannelHandshake.MsgChannelOpenAck({
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
    ) internal pure returns (IIBCChannelHandshake.MsgChannelOpenConfirm memory) {
        return IIBCChannelHandshake.MsgChannelOpenConfirm({
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

    function validateInitializedSequences(TestableIBCHandler handler, string memory portId, string memory channelId)
        internal
    {
        assertEq(handler.getNextSequenceSend(portId, channelId), 1);
        assertEq(handler.getNextSequenceRecv(portId, channelId), 1);
        assertEq(handler.getNextSequenceAck(portId, channelId), 1);
        assertEq(
            handler.getCommitment(IBCCommitment.nextSequenceRecvCommitmentKey(portId, channelId)),
            keccak256(abi.encodePacked(uint64(1)))
        );
    }

    function validatePostStateAfterChanOpenInit(
        TestableIBCHandler handler,
        IIBCChannelHandshake.MsgChannelOpenInit memory msg_,
        string memory channelId,
        string memory version
    ) internal {
        (Channel.Data memory channel, bool ok) = handler.getChannel(msg_.portId, channelId);
        assertTrue(ok);
        assertTrue(channel.state == Channel.State.STATE_INIT);
        assertTrue(channel.ordering == msg_.channel.ordering);
        assertEq(channel.counterparty.port_id, msg_.channel.counterparty.port_id);
        assertEq(channel.counterparty.channel_id, "");
        assertEq(channel.connection_hops.length, msg_.channel.connection_hops.length);
        for (uint256 i = 0; i < channel.connection_hops.length; i++) {
            assertEq(channel.connection_hops[i], msg_.channel.connection_hops[i]);
        }
        assertEq(channel.version, version);
        validateInitializedSequences(handler, msg_.portId, channelId);
    }

    function validatePostStateAfterChanOpenTry(
        TestableIBCHandler handler,
        IIBCChannelHandshake.MsgChannelOpenTry memory msg_,
        string memory channelId,
        string memory version
    ) internal {
        (Channel.Data memory channel, bool ok) = handler.getChannel(msg_.portId, channelId);
        assertTrue(ok);
        assertTrue(channel.state == Channel.State.STATE_TRYOPEN);
        assertTrue(channel.ordering == msg_.channel.ordering);
        assertEq(channel.counterparty.port_id, msg_.channel.counterparty.port_id);
        assertEq(channel.counterparty.channel_id, msg_.channel.counterparty.channel_id);
        assertEq(channel.connection_hops.length, msg_.channel.connection_hops.length);
        for (uint256 i = 0; i < channel.connection_hops.length; i++) {
            assertEq(channel.connection_hops[i], msg_.channel.connection_hops[i]);
        }
        assertEq(channel.version, version);
        validateInitializedSequences(handler, msg_.portId, channelId);
    }

    function validatePostStateAfterChanOpenAck(
        TestableIBCHandler handler,
        IIBCChannelHandshake.MsgChannelOpenAck memory msg_,
        string memory version
    ) internal {
        (Channel.Data memory channel, bool ok) = handler.getChannel(msg_.portId, msg_.channelId);
        assertTrue(ok);
        assertTrue(channel.state == Channel.State.STATE_OPEN);
        assertEq(channel.counterparty.channel_id, msg_.counterpartyChannelId);
        assertEq(channel.version, version);
    }

    function validatePostStateAfterChanOpenConfirm(
        TestableIBCHandler handler,
        IIBCChannelHandshake.MsgChannelOpenConfirm memory msg_
    ) internal {
        (Channel.Data memory channel, bool ok) = handler.getChannel(msg_.portId, msg_.channelId);
        assertTrue(ok);
        assertTrue(channel.state == Channel.State.STATE_OPEN);
    }

    enum ChannelHandshakeStep {
        INIT,
        TRY,
        ACK,
        CONFIRM
    }

    function handshakeChannel(
        TestableIBCHandler handler,
        TestableIBCHandler counterpartyHandler,
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        ChannelHandshakeStep step,
        Height.Data memory proofHeight
    ) internal returns (ChannelInfo memory, ChannelInfo memory) {
        {
            IIBCChannelHandshake.MsgChannelOpenInit memory msg_ =
                msgChannelOpenInit(channelInfo, counterpartyChannelInfo.portId);
            (channelInfo.channelId, channelInfo.version) = handler.channelOpenInit(msg_);
            validatePostStateAfterChanOpenInit(handler, msg_, channelInfo.channelId, channelInfo.version);
        }
        if (step == ChannelHandshakeStep.INIT) {
            return (channelInfo, counterpartyChannelInfo);
        }
        {
            IIBCChannelHandshake.MsgChannelOpenTry memory msg_ =
                msgChannelOpenTry(counterpartyChannelInfo, channelInfo, proofHeight);
            (counterpartyChannelInfo.channelId, counterpartyChannelInfo.version) =
                counterpartyHandler.channelOpenTry(msg_);
            validatePostStateAfterChanOpenTry(
                counterpartyHandler, msg_, counterpartyChannelInfo.channelId, counterpartyChannelInfo.version
            );
        }
        if (step == ChannelHandshakeStep.TRY) {
            return (channelInfo, counterpartyChannelInfo);
        }
        {
            IIBCChannelHandshake.MsgChannelOpenAck memory msg_ =
                msgChannelOpenAck(channelInfo, counterpartyChannelInfo, proofHeight);
            handler.channelOpenAck(msg_);
            validatePostStateAfterChanOpenAck(handler, msg_, channelInfo.version);
        }
        if (step == ChannelHandshakeStep.ACK) {
            return (channelInfo, counterpartyChannelInfo);
        }
        {
            IIBCChannelHandshake.MsgChannelOpenConfirm memory msg_ =
                msgChannelOpenConfirm(counterpartyChannelInfo, channelInfo, proofHeight);
            counterpartyHandler.channelOpenConfirm(msg_);
            validatePostStateAfterChanOpenConfirm(counterpartyHandler, msg_);
        }
        return (channelInfo, counterpartyChannelInfo);
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
        handler.bindPort("portidone", mockApp);
        counterpartyMockApp = new IBCMockApp(counterpartyHandler);
        counterpartyHandler.bindPort("portidtwo", counterpartyMockApp);
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
            H(0, 1)
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
            H(0, 1)
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
                H(0, 1)
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
                H(0, 1)
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
                H(0, 1)
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
            H(0, 1)
        );

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_UNORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.ACK,
            H(0, 1)
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
            H(0, 1)
        );

        // app version mismatch between proof and msg
        {
            IIBCChannelHandshake.MsgChannelOpenAck memory msg_ =
                msgChannelOpenAck(channelInfo, counterpartyChannelInfo, H(0, 1));
            msg_.counterpartyVersion = "mockapp-2";
            vm.expectRevert();
            handler.channelOpenAck(msg_);
        }
        // invalid proof height
        {
            IIBCChannelHandshake.MsgChannelOpenAck memory msg_ =
                msgChannelOpenAck(channelInfo, counterpartyChannelInfo, H(0, 2));
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
            H(0, 1)
        );

        handshakeChannel(
            handler,
            counterpartyHandler,
            ChannelInfo("portidone", "", Channel.Order.ORDER_UNORDERED, "", connectionId),
            ChannelInfo("portidtwo", "", Channel.Order.ORDER_UNORDERED, "", counterpartyConnectionId),
            ChannelHandshakeStep.CONFIRM,
            H(0, 1)
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
            H(0, 1)
        );

        // invalid proof height
        {
            IIBCChannelHandshake.MsgChannelOpenConfirm memory msg_ =
                msgChannelOpenConfirm(counterpartyChannelInfo, channelInfo, H(0, 2));
            vm.expectRevert();
            counterpartyHandler.channelOpenConfirm(msg_);
        }
    }
}

contract TestICS04Packet is TestIBCBase, TestMockClientHelper, TestICS03Helper, TestICS04Helper {
    string internal constant MOCK_APP_VERSION = "mockapp-1";

    TestableIBCHandler handler;
    TestableIBCHandler counterpartyHandler;
    ModifiedMockClient client;
    ModifiedMockClient counterpartyClient;
    IBCMockApp mockApp;
    IBCMockApp counterpartyMockApp;

    string clientId;
    string counterpartyClientId;
    string connectionId;
    string counterpartyConnectionId;

    function setUp() public {
        (TestableIBCHandler _handler, ModifiedMockClient _client) = ibcHandlerMockClient();
        (TestableIBCHandler _counterpartyHandler, ModifiedMockClient _counterpartyClient) = ibcHandlerMockClient();
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
        Channel.Order[] memory orders = new Channel.Order[](2);
        orders[0] = Channel.Order.ORDER_ORDERED;
        orders[1] = Channel.Order.ORDER_UNORDERED;
        for (uint256 i = 0; i < orders.length; i++) {
            ChannelInfo memory channelInfo = ChannelInfo("portidone", "", orders[i], "", connectionId);
            ChannelInfo memory counterpartyChannelInfo =
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId);

            (channelInfo, counterpartyChannelInfo) = handshakeChannel(
                handler,
                counterpartyHandler,
                channelInfo,
                counterpartyChannelInfo,
                ChannelHandshakeStep.CONFIRM,
                H(0, 1)
            );

            assertEq(
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA,
                    channelInfo.portId,
                    channelInfo.channelId,
                    getHeight(client, clientId, 1),
                    0
                ),
                1
            );
            assertEq(
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA,
                    channelInfo.portId,
                    channelInfo.channelId,
                    H(0, 0),
                    getTimestamp(client, clientId, 1)
                ),
                2
            );
            assertEq(
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA,
                    channelInfo.portId,
                    channelInfo.channelId,
                    getHeight(client, clientId, 1),
                    getTimestamp(client, clientId, 1)
                ),
                3
            );
        }
    }

    function testInvalidSendPacket() public {
        Channel.Order[] memory orders = new Channel.Order[](2);
        orders[0] = Channel.Order.ORDER_ORDERED;
        orders[1] = Channel.Order.ORDER_UNORDERED;
        for (uint256 i = 0; i < orders.length; i++) {
            ChannelInfo memory channelInfo = ChannelInfo("portidone", "", orders[i], "", connectionId);
            ChannelInfo memory counterpartyChannelInfo =
                ChannelInfo("portidtwo", "", orders[i], "", counterpartyConnectionId);

            (channelInfo, counterpartyChannelInfo) = handshakeChannel(
                handler,
                counterpartyHandler,
                channelInfo,
                counterpartyChannelInfo,
                ChannelHandshakeStep.CONFIRM,
                H(0, 1)
            );

            {
                Height.Data memory timeoutHeight = getHeight(client, clientId, 1);
                vm.expectRevert();
                mockApp.sendPacket(IBCMockLib.MOCK_PACKET_DATA, "invalidport", channelInfo.channelId, timeoutHeight, 0);

                vm.expectRevert();
                mockApp.sendPacket(IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, "channel-999", timeoutHeight, 0);
            }

            {
                Height.Data memory timeoutHeight = getHeight(client, clientId, 0);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );
            }

            {
                uint64 timeoutTimestamp = getTimestamp(client, clientId, 0);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, H(0, 0), timeoutTimestamp
                );
            }

            {
                Height.Data memory timeoutHeight = getHeight(client, clientId, 1);
                client.setStatus(clientId, ClientStatus.Frozen);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );
                client.setStatus(clientId, ClientStatus.Expired);
                vm.expectRevert();
                mockApp.sendPacket(
                    IBCMockLib.MOCK_PACKET_DATA, channelInfo.portId, channelInfo.channelId, timeoutHeight, 0
                );
                client.setStatus(clientId, ClientStatus.Active);
            }
        }
    }
}
