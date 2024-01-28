// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "./MockClientTestHelper.t.sol";
import "../../../../contracts/apps/mock/IBCMockApp.sol";
import "../../../../contracts/proto/Channel.sol";

abstract contract ICS04TestHelper is IBCTestHelper {
    struct ChannelInfo {
        string portId;
        string channelId;
        Channel.Order ordering;
        string version;
        string connectionId;
    }

    function genChannelId(uint64 sequence) internal pure returns (string memory) {
        return string(abi.encodePacked("channel-", Strings.toString(sequence)));
    }

    function ensureChannelState(IIBCHandler handler, ChannelInfo memory channel, Channel.State state) internal {
        (Channel.Data memory channel_, bool ok) = handler.getChannel(channel.portId, channel.channelId);
        assertTrue(ok);
        assertTrue(channel_.state == state);
    }
}

abstract contract ICS04HandshakeTestHelper is ICS04TestHelper {
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
                version: channelInfo.version,
                upgrade_sequence: 0
            })
        });
    }

    function msgChannelOpenTry(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight
    ) internal view returns (IIBCChannelHandshake.MsgChannelOpenTry memory) {
        assert(bytes(channelInfo.channelId).length == 0 && bytes(channelInfo.version).length == 0);
        bytes memory proofInit = proveChannelState(
            proofHeight,
            counterpartyChannelInfo.portId,
            counterpartyChannelInfo.channelId,
            Channel.Data({
                state: Channel.State.STATE_INIT,
                ordering: counterpartyChannelInfo.ordering,
                counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: ""}),
                connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                version: counterpartyChannelInfo.version,
                upgrade_sequence: 0
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
                version: "",
                upgrade_sequence: 0
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
    ) internal view returns (IIBCChannelHandshake.MsgChannelOpenAck memory) {
        return IIBCChannelHandshake.MsgChannelOpenAck({
            portId: channelInfo.portId,
            channelId: channelInfo.channelId,
            counterpartyVersion: counterpartyChannelInfo.version,
            counterpartyChannelId: counterpartyChannelInfo.channelId,
            proofTry: proveChannelState(
                proofHeight,
                counterpartyChannelInfo.portId,
                counterpartyChannelInfo.channelId,
                Channel.Data({
                    state: Channel.State.STATE_TRYOPEN,
                    ordering: counterpartyChannelInfo.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: channelInfo.channelId}),
                    connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                    version: counterpartyChannelInfo.version,
                    upgrade_sequence: 0
                })
                ),
            proofHeight: proofHeight
        });
    }

    function msgChannelOpenConfirm(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight
    ) internal view returns (IIBCChannelHandshake.MsgChannelOpenConfirm memory) {
        return IIBCChannelHandshake.MsgChannelOpenConfirm({
            portId: channelInfo.portId,
            channelId: channelInfo.channelId,
            proofAck: proveChannelState(
                proofHeight,
                counterpartyChannelInfo.portId,
                counterpartyChannelInfo.channelId,
                Channel.Data({
                    state: Channel.State.STATE_OPEN,
                    ordering: counterpartyChannelInfo.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: channelInfo.channelId}),
                    connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                    version: counterpartyChannelInfo.version,
                    upgrade_sequence: 0
                })
                ),
            proofHeight: proofHeight
        });
    }

    function msgChannelCloseInit(ChannelInfo memory channelInfo)
        internal
        pure
        returns (IIBCChannelHandshake.MsgChannelCloseInit memory)
    {
        return IIBCChannelHandshake.MsgChannelCloseInit({portId: channelInfo.portId, channelId: channelInfo.channelId});
    }

    function msgChannelCloseConfirm(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight
    ) internal view returns (IIBCChannelHandshake.MsgChannelCloseConfirm memory) {
        return msgChannelCloseConfirm(channelInfo, counterpartyChannelInfo, proofHeight, Channel.State.STATE_CLOSED);
    }

    function msgChannelCloseConfirm(
        ChannelInfo memory channelInfo,
        ChannelInfo memory counterpartyChannelInfo,
        Height.Data memory proofHeight,
        Channel.State state // for unexpected state testing
    ) internal view returns (IIBCChannelHandshake.MsgChannelCloseConfirm memory) {
        return IIBCChannelHandshake.MsgChannelCloseConfirm({
            portId: channelInfo.portId,
            channelId: channelInfo.channelId,
            proofInit: proveChannelState(
                proofHeight,
                counterpartyChannelInfo.portId,
                counterpartyChannelInfo.channelId,
                Channel.Data({
                    state: state,
                    ordering: counterpartyChannelInfo.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: channelInfo.portId, channel_id: channelInfo.channelId}),
                    connection_hops: newConnectionHops(counterpartyChannelInfo.connectionId),
                    version: counterpartyChannelInfo.version,
                    upgrade_sequence: 0
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
        require(bytes(channelInfo.channelId).length == 0, "channelInfo.channelId must be empty");
        require(bytes(channelInfo.version).length == 0, "channelInfo.version must be empty");
        require(bytes(counterpartyChannelInfo.channelId).length == 0, "counterpartyChannelInfo.channelId must be empty");
        require(bytes(counterpartyChannelInfo.version).length == 0, "counterpartyChannelInfo.version must be empty");
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

    function proveChannelState(
        Height.Data memory proofHeight,
        string memory portId,
        string memory channelId,
        Channel.Data memory channel
    ) internal view virtual returns (bytes memory);
}

abstract contract ICS04HandshakeMockClientTestHelper is ICS04HandshakeTestHelper, MockClientTestHelper {
    function proveChannelState(
        Height.Data memory proofHeight,
        string memory portId,
        string memory channelId,
        Channel.Data memory channel
    ) internal view virtual override returns (bytes memory) {
        return genMockProof(
            proofHeight,
            DEFAULT_COMMITMENT_PREFIX,
            IBCCommitment.channelPath(portId, channelId),
            Channel.encode(channel)
        );
    }
}
