// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, Upgrade} from "../../proto/Channel.sol";
import {IBCChannelLib} from "../04-channel/IBCChannelLib.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCQuerier} from "./IIBCQuerier.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";

contract IBCQuerier is IBCModuleManager, IIBCQuerier {
    function getCommitmentPrefix() public view override returns (bytes memory) {
        return _getCommitmentPrefix();
    }

    function getCommitment(bytes32 hashedPath) public view override returns (bytes32) {
        return commitments[hashedPath];
    }

    function getExpectedTimePerBlock() public view override returns (uint64) {
        return expectedTimePerBlock;
    }

    function getClientByType(string calldata clientType) public view override returns (address) {
        return clientRegistry[clientType];
    }

    function getClientType(string calldata clientId) public view override returns (string memory) {
        return clientTypes[clientId];
    }

    function getClient(string calldata clientId) public view override returns (address) {
        return clientImpls[clientId];
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
        ConnectionEnd.Data storage connection = connections[connectionId];
        return (connection, connection.state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getChannel(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (Channel.Data memory, bool)
    {
        Channel.Data storage channel = channels[portId][channelId];
        return (channel, channel.state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getNextSequenceSend(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (uint64)
    {
        return nextSequenceSends[portId][channelId];
    }

    function getNextSequenceRecv(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (uint64)
    {
        return nextSequenceRecvs[portId][channelId];
    }

    function getNextSequenceAck(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (uint64)
    {
        return nextSequenceAcks[portId][channelId];
    }

    function getPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence)
        public
        view
        override
        returns (IBCChannelLib.PacketReceipt)
    {
        return IBCChannelLib.receiptCommitmentToReceipt(
            commitments[IBCCommitment.packetReceiptCommitmentKeyCalldata(portId, channelId, sequence)]
        );
    }

    function getChannelUpgrade(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (Upgrade.Data memory, bool)
    {
        Upgrade.Data storage upgrade = upgrades[portId][channelId];
        return (upgrade, upgrade.fields.connection_hops.length != 0);
    }

    function getCanTransitionToFlushComplete(string calldata portId, string calldata channelId)
        public
        view
        override
        returns (bool)
    {
        Channel.Data storage channel = channels[portId][channelId];
        if (channel.state != Channel.State.STATE_FLUSHING) {
            return false;
        }
        return canTransitionToFlushComplete(channel.ordering, portId, channelId, channel.upgrade_sequence);
    }
}
