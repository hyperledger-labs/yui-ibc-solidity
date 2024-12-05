// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, ChannelCounterparty} from "../../proto/Channel.sol";
import {IBCHeight} from "../02-client/IBCHeight.sol";
import {IBCConnectionLib} from "../03-connection/IBCConnectionLib.sol";
import {IIBCChannelHandshake} from "../04-channel/IIBCChannel.sol";
import {IIBCChannelErrors} from "../04-channel/IIBCChannelErrors.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCModuleInitializer, IIBCModule} from "../26-router/IIBCModule.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IBCChannelLib} from "./IBCChannelLib.sol";

/**
 * @dev IBCChannelHandshake is a contract that implements [ICS-4](https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics).
 */
contract IBCChannelHandshake is IBCModuleManager, IIBCChannelHandshake, IIBCChannelErrors {
    using IBCHeight for Height.Data;

    // ------------- IIBCChannelHandshake implementation ------------------- //

    /**
     * @dev channelOpenInit is called by a module to initiate a channel opening handshake with a module on another chain.
     */
    function channelOpenInit(IIBCChannelHandshake.MsgChannelOpenInit calldata msg_)
        public
        override
        returns (string memory, string memory)
    {
        if (msg_.channel.connection_hops.length != 1) {
            revert IBCChannelInvalidConnectionHopsLength(msg_.channel.connection_hops.length);
        }
        // optimistic channel handshakes are allowed, so we can skip checking if the connection state is OPEN here.
        ConnectionEnd.Data storage connection = getConnectionStorage()[msg_.channel.connection_hops[0]].connection;
        if (connection.versions.length != 1) {
            revert IBCChannelConnectionMultipleVersionsFound(
                msg_.channel.connection_hops[0], connection.versions.length
            );
        }
        if (
            !IBCConnectionLib.verifySupportedFeature(
                connection.versions[0], IBCChannelLib.toString(msg_.channel.ordering)
            )
        ) {
            revert IBCChannelConnectionFeatureNotSupported(msg_.channel.ordering);
        }
        if (msg_.channel.state != Channel.State.STATE_INIT) {
            revert IBCChannelUnexpectedChannelState(msg_.channel.state);
        }
        if (bytes(msg_.channel.counterparty.channel_id).length != 0) {
            revert IBCChannelCounterpartyChannelIdNotEmpty(msg_.channel.counterparty.channel_id);
        }

        string memory channelId = generateChannelIdentifier();
        ChannelStorage storage channelStorage = getChannelStorage()[msg_.portId][channelId];
        if (channelStorage.channel.state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED) {
            revert IBCChannelAlreadyChannelExists();
        }
        initializeSequences(msg_.portId, channelId);
        emit GeneratedChannelIdentifier(channelId);

        (address module, string memory version) = lookupModuleByPort(msg_.portId).onChanOpenInit(
            IIBCModuleInitializer.MsgOnChanOpenInit({
                order: msg_.channel.ordering,
                connectionHops: msg_.channel.connection_hops,
                portId: msg_.portId,
                channelId: channelId,
                counterparty: msg_.channel.counterparty,
                version: msg_.channel.version
            })
        );
        claimChannelCapability(msg_.portId, channelId, module);
        writeChannel(
            msg_.portId,
            channelId,
            msg_.channel.state,
            msg_.channel.ordering,
            msg_.channel.counterparty,
            msg_.channel.connection_hops,
            version
        );
        return (channelId, version);
    }

    /**
     * @dev channelOpenTry is called by a module to accept the first step of a channel opening handshake initiated by a module on another chain.
     */
    function channelOpenTry(IIBCChannelHandshake.MsgChannelOpenTry calldata msg_)
        public
        override
        returns (string memory, string memory)
    {
        if (msg_.channel.connection_hops.length != 1) {
            revert IBCChannelInvalidConnectionHopsLength(msg_.channel.connection_hops.length);
        }
        ConnectionEnd.Data storage connection = getConnectionStorage()[msg_.channel.connection_hops[0]].connection;
        if (connection.state != ConnectionEnd.State.STATE_OPEN) {
            revert IBCChannelConnectionNotOpened(msg_.channel.connection_hops[0]);
        }
        if (connection.versions.length != 1) {
            revert IBCChannelConnectionMultipleVersionsFound(
                msg_.channel.connection_hops[0], connection.versions.length
            );
        }
        if (
            !IBCConnectionLib.verifySupportedFeature(
                connection.versions[0], IBCChannelLib.toString(msg_.channel.ordering)
            )
        ) {
            revert IBCChannelConnectionFeatureNotSupported(msg_.channel.ordering);
        }
        if (msg_.channel.state != Channel.State.STATE_TRYOPEN) {
            revert IBCChannelUnexpectedChannelState(msg_.channel.state);
        }

        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_INIT,
            ordering: msg_.channel.ordering,
            counterparty: ChannelCounterparty.Data({port_id: msg_.portId, channel_id: ""}),
            connection_hops: getCounterpartyHops(msg_.channel.connection_hops[0]),
            version: msg_.counterpartyVersion,
            upgrade_sequence: 0
        });
        verifyChannelState(
            connection,
            msg_.proofHeight,
            msg_.proofInit,
            msg_.channel.counterparty.port_id,
            msg_.channel.counterparty.channel_id,
            Channel.encode(expectedChannel)
        );

        string memory channelId = generateChannelIdentifier();
        ChannelStorage storage channelStorage = getChannelStorage()[msg_.portId][channelId];
        if (channelStorage.channel.state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED) {
            revert IBCChannelAlreadyChannelExists();
        }
        initializeSequences(msg_.portId, channelId);
        emit GeneratedChannelIdentifier(channelId);

        (address module, string memory version) = lookupModuleByPort(msg_.portId).onChanOpenTry(
            IIBCModuleInitializer.MsgOnChanOpenTry({
                order: msg_.channel.ordering,
                connectionHops: msg_.channel.connection_hops,
                portId: msg_.portId,
                channelId: channelId,
                counterparty: msg_.channel.counterparty,
                counterpartyVersion: msg_.counterpartyVersion
            })
        );
        claimChannelCapability(msg_.portId, channelId, module);
        writeChannel(
            msg_.portId,
            channelId,
            msg_.channel.state,
            msg_.channel.ordering,
            msg_.channel.counterparty,
            msg_.channel.connection_hops,
            version
        );
        return (channelId, version);
    }

    /**
     * @dev channelOpenAck is called by the handshake-originating module to acknowledge the acceptance of the initial request by the counterparty module on the other chain.
     */
    function channelOpenAck(IIBCChannelHandshake.MsgChannelOpenAck calldata msg_) public override {
        Channel.Data storage channel = getChannelStorage()[msg_.portId][msg_.channelId].channel;
        if (channel.state != Channel.State.STATE_INIT) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
        if (connection.state != ConnectionEnd.State.STATE_OPEN) {
            revert IBCChannelConnectionNotOpened(channel.connection_hops[0]);
        }
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_TRYOPEN,
            ordering: channel.ordering,
            counterparty: ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
            connection_hops: getCounterpartyHops(channel.connection_hops[0]),
            version: msg_.counterpartyVersion,
            upgrade_sequence: 0
        });

        verifyChannelState(
            connection,
            msg_.proofHeight,
            msg_.proofTry,
            channel.counterparty.port_id,
            msg_.counterpartyChannelId,
            Channel.encode(expectedChannel)
        );
        channel.state = Channel.State.STATE_OPEN;
        channel.version = msg_.counterpartyVersion;
        channel.counterparty.channel_id = msg_.counterpartyChannelId;
        updateChannelCommitment(msg_.portId, msg_.channelId);
        lookupModuleByChannel(msg_.portId, msg_.channelId).onChanOpenAck(
            IIBCModule.MsgOnChanOpenAck({
                portId: msg_.portId,
                channelId: msg_.channelId,
                counterpartyVersion: msg_.counterpartyVersion
            })
        );
    }

    /**
     * @dev channelOpenConfirm is called by the counterparty module to acknowledge the acknowledgement of the handshake-originating module on the other chain and finish the channel opening handshake.
     */
    function channelOpenConfirm(IIBCChannelHandshake.MsgChannelOpenConfirm calldata msg_) public override {
        Channel.Data storage channel = getChannelStorage()[msg_.portId][msg_.channelId].channel;
        if (channel.state != Channel.State.STATE_TRYOPEN) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_OPEN,
            ordering: channel.ordering,
            counterparty: ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
            connection_hops: getCounterpartyHops(channel.connection_hops[0]),
            version: channel.version,
            upgrade_sequence: 0
        });
        verifyChannelState(
            connection,
            msg_.proofHeight,
            msg_.proofAck,
            channel.counterparty.port_id,
            channel.counterparty.channel_id,
            Channel.encode(expectedChannel)
        );
        channel.state = Channel.State.STATE_OPEN;
        updateChannelCommitment(msg_.portId, msg_.channelId);
        lookupModuleByChannel(msg_.portId, msg_.channelId).onChanOpenConfirm(
            IIBCModule.MsgOnChanOpenConfirm({portId: msg_.portId, channelId: msg_.channelId})
        );
    }

    /**
     * @dev channelCloseInit is called by either module to close their end of the channel. Once closed, channels cannot be reopened.
     */
    function channelCloseInit(IIBCChannelHandshake.MsgChannelCloseInit calldata msg_) public override {
        Channel.Data storage channel = getChannelStorage()[msg_.portId][msg_.channelId].channel;
        if (channel.state == Channel.State.STATE_UNINITIALIZED_UNSPECIFIED) {
            revert IBCChannelChannelNotFound(msg_.portId, msg_.channelId);
        } else if (channel.state == Channel.State.STATE_CLOSED) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }
        ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
        if (connection.state != ConnectionEnd.State.STATE_OPEN) {
            revert IBCChannelConnectionNotOpened(channel.connection_hops[0]);
        }
        channel.state = Channel.State.STATE_CLOSED;
        updateChannelCommitment(msg_.portId, msg_.channelId);
        lookupModuleByChannel(msg_.portId, msg_.channelId).onChanCloseInit(
            IIBCModule.MsgOnChanCloseInit({portId: msg_.portId, channelId: msg_.channelId})
        );
    }

    /**
     * @dev channelCloseConfirm is called by the counterparty module to close their end of the
     * channel, since the other end has been closed.
     */
    function channelCloseConfirm(IIBCChannelHandshake.MsgChannelCloseConfirm calldata msg_) public override {
        Channel.Data storage channel = getChannelStorage()[msg_.portId][msg_.channelId].channel;
        if (channel.state == Channel.State.STATE_UNINITIALIZED_UNSPECIFIED) {
            revert IBCChannelChannelNotFound(msg_.portId, msg_.channelId);
        } else if (channel.state == Channel.State.STATE_CLOSED) {
            revert IBCChannelUnexpectedChannelState(channel.state);
        }

        ConnectionEnd.Data storage connection = getConnectionStorage()[channel.connection_hops[0]].connection;
        if (connection.state != ConnectionEnd.State.STATE_OPEN) {
            revert IBCChannelConnectionNotOpened(channel.connection_hops[0]);
        }

        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_CLOSED,
            ordering: channel.ordering,
            counterparty: ChannelCounterparty.Data({port_id: msg_.portId, channel_id: msg_.channelId}),
            connection_hops: getCounterpartyHops(channel.connection_hops[0]),
            version: channel.version,
            upgrade_sequence: 0
        });
        verifyChannelState(
            connection,
            msg_.proofHeight,
            msg_.proofInit,
            channel.counterparty.port_id,
            channel.counterparty.channel_id,
            Channel.encode(expectedChannel)
        );
        channel.state = Channel.State.STATE_CLOSED;
        updateChannelCommitment(msg_.portId, msg_.channelId);
        lookupModuleByChannel(msg_.portId, msg_.channelId).onChanCloseConfirm(
            IIBCModule.MsgOnChanCloseConfirm({portId: msg_.portId, channelId: msg_.channelId})
        );
    }

    // ------------- Private functions ------------------- //

    /**
     * @dev writeChannel writes a channel which has successfully passed the OpenInit or OpenTry handshake step.
     */
    function writeChannel(
        string calldata portId,
        string memory channelId,
        Channel.State state,
        Channel.Order order,
        ChannelCounterparty.Data calldata counterparty,
        string[] calldata connectionHops,
        string memory version
    ) private {
        Channel.Data storage channel = getChannelStorage()[portId][channelId].channel;
        channel.state = state;
        channel.ordering = order;
        channel.counterparty = counterparty;
        for (uint256 i = 0; i < connectionHops.length; i++) {
            channel.connection_hops.push(connectionHops[i]);
        }
        channel.version = version;
        channel.upgrade_sequence = 0;
        updateChannelCommitment(portId, channelId);
    }

    function initializeSequences(string memory portId, string memory channelId) internal {
        ChannelStorage storage channelStorage = getChannelStorage()[portId][channelId];

        channelStorage.nextSequenceSend = 1;
        channelStorage.nextSequenceRecv = 1;
        channelStorage.nextSequenceAck = 1;
        channelStorage.recvStartSequence.sequence = 1;

        // Differ from the ICS-004 spec, we only store the commitment of the next sequence recv.
        // This is because, in the current spec, the next sequence send and ack are not needed for the verification on the counterparty chain.
        getCommitments()[IBCCommitment.nextSequenceRecvCommitmentKey(portId, channelId)] =
            keccak256(abi.encodePacked((bytes8(uint64(1)))));
    }

    function updateChannelCommitment(string memory portId, string memory channelId) private {
        getCommitments()[IBCCommitment.channelCommitmentKey(portId, channelId)] =
            keccak256(Channel.encode(getChannelStorage()[portId][channelId].channel));
    }

    function generateChannelIdentifier() private returns (string memory) {
        HostStorage storage hostStorage = getHostStorage();
        string memory identifier =
            string(abi.encodePacked("channel-", Strings.toString(hostStorage.nextChannelSequence)));
        hostStorage.nextChannelSequence++;
        return identifier;
    }

    function verifyChannelState(
        ConnectionEnd.Data storage connection,
        Height.Data calldata height,
        bytes calldata proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes
    ) private {
        // slither-disable-start reentrancy-no-eth
        if (
            checkAndGetClient(connection.client_id).verifyMembership(
                connection.client_id,
                height,
                0,
                0,
                proof,
                connection.counterparty.prefix.key_prefix,
                IBCCommitment.channelPath(portId, channelId),
                channelBytes
            )
        ) {
            // slither-disable-end reentrancy-no-eth
            return;
        }
        revert IBCChannelFailedVerifyChannelState(
            connection.client_id, IBCCommitment.channelPath(portId, channelId), channelBytes, proof, height
        );
    }

    function getCounterpartyHops(string memory connectionId) private view returns (string[] memory hops) {
        hops = new string[](1);
        hops[0] = getConnectionStorage()[connectionId].connection.counterparty.connection_id;
        return hops;
    }
}
