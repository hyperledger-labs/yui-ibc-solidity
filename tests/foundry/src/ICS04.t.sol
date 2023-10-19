// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./ICS03.t.sol";
import "../../../contracts/apps/mock/IBCMockApp.sol";

abstract contract TestICS04Helper is TestIBCBase, TestMockClientHelper {
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

    function msgChannelOpenInit(
        string memory connectionId,
        string memory portId,
        Channel.Order ordering,
        string memory version,
        string memory counterpartyPortId
    ) internal pure returns (IBCMsgs.MsgChannelOpenInit memory) {
        return IBCMsgs.MsgChannelOpenInit({
            portId: portId,
            channel: Channel.Data({
                state: Channel.State.STATE_INIT,
                ordering: ordering,
                counterparty: ChannelCounterparty.Data({port_id: counterpartyPortId, channel_id: ""}),
                connection_hops: newConnectionHops(connectionId),
                version: version
            })
        });
    }

    function msgChannelOpenTry(
        string memory connectionId,
        string memory portId,
        Channel.Order ordering,
        string memory version,
        string memory counterpartyConnectionId,
        string memory counterpartyPortId,
        string memory counterpartyChannelId,
        string memory counterpartyVersion,
        Height.Data memory proofHeight
    ) internal pure returns (IBCMsgs.MsgChannelOpenTry memory) {
        return IBCMsgs.MsgChannelOpenTry({
            portId: portId,
            channel: Channel.Data({
                state: Channel.State.STATE_TRYOPEN,
                ordering: ordering,
                counterparty: ChannelCounterparty.Data({port_id: counterpartyPortId, channel_id: counterpartyChannelId}),
                connection_hops: newConnectionHops(connectionId),
                version: version
            }),
            counterpartyVersion: version,
            proofInit: genMockChannelStateProof(
                proofHeight,
                counterpartyPortId,
                counterpartyChannelId,
                Channel.Data({
                    state: Channel.State.STATE_INIT,
                    ordering: ordering,
                    counterparty: ChannelCounterparty.Data({port_id: counterpartyPortId, channel_id: ""}),
                    connection_hops: newConnectionHops(counterpartyConnectionId),
                    version: counterpartyVersion
                })
                ),
            proofHeight: proofHeight
        });
    }

    function msgChannelOpenAck(
        string memory portId,
        string memory channelId,
        Channel.Order ordering,
        string memory counterpartyConnectionId,
        string memory counterpartyPortId,
        string memory counterpartyChannelId,
        string memory counterpartyVersion,
        Height.Data memory proofHeight
    ) internal pure returns (IBCMsgs.MsgChannelOpenAck memory) {
        return IBCMsgs.MsgChannelOpenAck({
            portId: portId,
            channelId: channelId,
            counterpartyVersion: counterpartyVersion,
            counterpartyChannelId: counterpartyChannelId,
            proofTry: genMockChannelStateProof(
                proofHeight,
                counterpartyPortId,
                counterpartyChannelId,
                Channel.Data({
                    state: Channel.State.STATE_TRYOPEN,
                    ordering: ordering,
                    counterparty: ChannelCounterparty.Data({port_id: portId, channel_id: channelId}),
                    connection_hops: newConnectionHops(counterpartyConnectionId),
                    version: counterpartyVersion
                })
                ),
            proofHeight: proofHeight
        });
    }

    function msgChannelOpenConfirm(
        string memory portId,
        string memory channelId,
        Channel.Order ordering,
        string memory counterpartyConnectionId,
        string memory counterpartyPortId,
        string memory counterpartyChannelId,
        string memory counterpartyVersion,
        Height.Data memory proofHeight
    ) internal pure returns (IBCMsgs.MsgChannelOpenConfirm memory) {
        return IBCMsgs.MsgChannelOpenConfirm({
            portId: portId,
            channelId: channelId,
            proofAck: genMockChannelStateProof(
                proofHeight,
                counterpartyPortId,
                counterpartyChannelId,
                Channel.Data({
                    state: Channel.State.STATE_OPEN,
                    ordering: ordering,
                    counterparty: ChannelCounterparty.Data({port_id: portId, channel_id: channelId}),
                    connection_hops: newConnectionHops(counterpartyConnectionId),
                    version: counterpartyVersion
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
}

contract TestICS04Handshake is TestIBCBase, TestMockClientHelper, TestICS03Helper, TestICS04Helper {
    TestableIBCHandler handler;
    TestableIBCHandler counterpartyHandler;
    IBCMockApp mockApp;
    IBCMockApp counterpartyMockApp;

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
            (string memory channelId, string memory version) = handler.channelOpenInit(
                msgChannelOpenInit(connectionId, "portidone", Channel.Order.ORDER_ORDERED, "mockapp-1", "portidtwo")
            );
            assertEq(channelId, "channel-0");
            assertEq(version, "mockapp-1");
        }
        {
            (string memory channelId, string memory version) = handler.channelOpenInit(
                msgChannelOpenInit(connectionId, "portidone", Channel.Order.ORDER_ORDERED, "mockapp-1", "portidtwo")
            );
            assertEq(channelId, "channel-1");
            assertEq(version, "mockapp-1");
        }
        {
            (string memory channelId, string memory version) = handler.channelOpenInit(
                msgChannelOpenInit(connectionId, "portidone", Channel.Order.ORDER_UNORDERED, "mockapp-1", "portidtwo")
            );
            assertEq(channelId, "channel-2");
            assertEq(version, "mockapp-1");
        }
    }
}
