// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, Upgrade} from "../../proto/Channel.sol";
import {IBCChannelLib} from "../04-channel/IBCChannelLib.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCQuerier} from "./IIBCQuerier.sol";

contract IBCQuerier is IBCModuleManager, IIBCQuerier {
    function getCommitmentPrefix() public view override returns (bytes memory) {
        return _getCommitmentPrefix();
    }

    function getCommitmentsSlot() public pure override returns (bytes32) {
        return COMMITMENT_STORAGE_LOCATION;
    }

    function getCommitment(bytes32 hashedPath) public view override returns (bytes32) {
        return getCommitments()[hashedPath];
    }

    function getExpectedTimePerBlock() public view override returns (uint64) {
        return getHostStorage().expectedTimePerBlock;
    }

    function getIBCModuleByPort(string calldata portId) public view override returns (IIBCModule) {
        return lookupModuleByPort(portId);
    }

    function getIBCModuleByChannel(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (IIBCModule)
    {
        return lookupModuleByChannel(portId, channelId);
    }

    function getClientByType(string calldata clientType) public view override returns (address) {
        return getHostStorage().clientRegistry[clientType];
    }

    function getClientType(string calldata clientId) public view override returns (string memory) {
        return getClientStorage()[clientId].clientType;
    }

    function getClient(string calldata clientId) public view override returns (address) {
        return getClientStorage()[clientId].clientImpl;
    }

    function getClientState(string calldata clientId) public view override returns (bytes memory, bool) {
        return checkAndGetClient(clientId).getClientState(clientId);
    }

    function getConsensusState(string calldata clientId, Height.Data calldata height)
        public
        view
        override
        returns (bytes memory consensusStateBytes, bool)
    {
        return checkAndGetClient(clientId).getConsensusState(clientId, height);
    }

    function getConnection(string calldata connectionId)
        public
        view
        override
        returns (ConnectionEnd.Data memory, bool)
    {
        ConnectionEnd.Data storage connection = getConnectionStorage()[connectionId].connection;
        return (connection, connection.state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getChannel(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (Channel.Data memory, bool)
    {
        Channel.Data storage channel = getChannelStorage()[portId][channelId].channel;
        return (channel, channel.state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getNextSequenceSend(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (uint64)
    {
        return getChannelStorage()[portId][channelId].nextSequenceSend;
    }

    function getNextSequenceRecv(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (uint64)
    {
        return getChannelStorage()[portId][channelId].nextSequenceRecv;
    }

    function getNextSequenceAck(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (uint64)
    {
        return getChannelStorage()[portId][channelId].nextSequenceAck;
    }

    function getPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence)
        public
        view
        override
        returns (IBCChannelLib.PacketReceipt)
    {
        return IBCChannelLib.receiptCommitmentToReceipt(
            getCommitments()[IBCCommitment.packetReceiptCommitmentKeyCalldata(portId, channelId, sequence)]
        );
    }

    function getChannelUpgrade(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (Upgrade.Data memory, bool)
    {
        Upgrade.Data storage upgrade = getChannelStorage()[portId][channelId].upgrade;
        return (upgrade, upgrade.fields.connection_hops.length != 0);
    }

    function getCanTransitionToFlushComplete(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (bool)
    {
        Channel.Data storage channel = getChannelStorage()[portId][channelId].channel;
        if (channel.state != Channel.State.STATE_FLUSHING) {
            return false;
        }
        return canTransitionToFlushComplete(channel.ordering, portId, channelId, channel.upgrade_sequence);
    }
}
