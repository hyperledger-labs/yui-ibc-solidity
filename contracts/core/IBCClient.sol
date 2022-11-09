// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";

library IBCClient {

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCHost host, IBCMsgs.MsgCreateClient calldata msg_) external {
        host.onlyIBCModule();
        (, bool found) = getClientByType(host, msg_.clientType);
        require(found, "unregistered client type");

        string memory clientId = host.generateClientIdentifier(msg_.clientType);
        host.setClientType(clientId, msg_.clientType);
        host.setClientState(clientId, msg_.clientStateBytes);
        host.setConsensusState(clientId, msg_.height, msg_.consensusStateBytes);
        host.setProcessedTime(clientId, msg_.height, block.timestamp);
        host.setProcessedHeight(clientId, msg_.height, block.number);
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCHost host, IBCMsgs.MsgUpdateClient calldata msg_) external {
        host.onlyIBCModule();
        bytes memory clientStateBytes;
        bytes memory consensusStateBytes;
        Height.Data memory height;
        bool found;
    
        (clientStateBytes, found) = host.getClientState(msg_.clientId);
        require(found, "clientState not found");

        (clientStateBytes, consensusStateBytes, height) = checkHeaderAndUpdateState(host, msg_.clientId, clientStateBytes, msg_.header);
    
        //// persist states ////
        host.setClientState(msg_.clientId, clientStateBytes);
        host.setConsensusState(msg_.clientId, height, consensusStateBytes);
        host.setProcessedTime(msg_.clientId, height, block.timestamp);
        host.setProcessedHeight(msg_.clientId, height, block.number);
    }

    // TODO implements
    function validateSelfClient(IBCHost, bytes calldata) external view returns (bool) {
        this; // this is a trick that suppresses "Warning: Function state mutability can be restricted to pure"
        return true;
    }

    function registerClient(IBCHost host, string memory clientType, IClient client) public {
        host.onlyIBCModule();
        host.setClientImpl(clientType, address(client));
    }

    function getClient(IBCHost host, string memory clientId) public view returns (IClient) {
        (IClient clientImpl, bool found) = getClientByType(host, host.getClientType(clientId));
        require(found, "clientImpl not found");
        return clientImpl;
    }

    function getClientByType(IBCHost host, string memory clientType) internal view returns (IClient clientImpl, bool) {
        (address addr, bool found) = host.getClientImpl(clientType);
        if (!found) {
            return (clientImpl, false);
        }
        return (IClient(addr), true);
    }

    function checkHeaderAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory headerBytes
    ) public returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, Height.Data memory height) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.checkHeaderAndUpdateState.selector,
                host, clientId, clientStateBytes, headerBytes));
        assert(success);
        return abi.decode(res, (bytes, bytes, Height.Data));
    }

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        bytes memory clientStateBytes
    ) public returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyClientState.selector,
                host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes));
        assert(success);
        return abi.decode(res, (bool));
    }

    function verifyClientConsensusState(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        string memory counterpartyClientIdentifier,
        Height.Data memory consensusHeight,
        bytes memory prefix,
        bytes memory proof,
        bytes memory consensusStateBytes
    ) public returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyClientConsensusState.selector,
                host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes));
        assert(success);
        return abi.decode(res, (bool));
    }

    function verifyConnectionState(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        bytes calldata prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory counterpartyConnectionBytes
    ) public returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyConnectionState.selector,
                host, clientId, height, prefix, proof, connectionId, counterpartyConnectionBytes));
        assert(success);
        return abi.decode(res, (bool));
    }

    function verifyChannelState(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes
    ) public returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyChannelState.selector,
                host, clientId, height, prefix, proof, portId, channelId, channelBytes));
        assert(success);
        return abi.decode(res, (bool));
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes // serialized with pb
    ) public returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyPacketCommitment.selector,
                host, clientId, height, delayPeriodTime, delayPeriodBlocks, prefix, proof, portId, channelId, sequence, commitmentBytes));
        assert(success);
        return abi.decode(res, (bool));
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes memory acknowledgement // serialized with pb
    ) public returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyPacketAcknowledgement.selector,
                host, clientId, height, delayPeriodTime, delayPeriodBlocks, prefix, proof, portId, channelId, sequence, acknowledgement));
        assert(success);
        return abi.decode(res, (bool));
    }
}
