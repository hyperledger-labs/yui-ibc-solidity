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
        bool found;

        (clientStateBytes, found) = host.getClientState(msg_.clientId);
        require(found, "clientState not found");

        verifyClientMessageAndUpdateState(host, msg_.clientId, clientStateBytes, msg_.clientMessage);
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

    function verifyClientMessageAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory clientMessageBytes
    ) internal returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyClientMessageAndUpdateState.selector, host, clientId, clientStateBytes, clientMessageBytes
            )
        );
        require(success);
        return abi.decode(res, (bool));
    }

    // Verification functions

    function verifyMembership(
        IBCHost host,
        string memory clientId,
        Height.Data memory height,
        uint64 delayTimePeriod,
        uint64 delayBlockPeriod,
        bytes memory proof,
        bytes memory prefix,
        bytes memory path,
        bytes memory value
    ) private returns (bool) {
        (bool success, bytes memory res) = address(getClient(host, clientId)).delegatecall(
            abi.encodeWithSelector(
                IClient.verifyMembership.selector,
                host,
                clientId,
                height,
                delayTimePeriod,
                delayBlockPeriod,
                proof,
                prefix,
                path,
                value
            )
        );
        require(success);
        return abi.decode(res, (bool));
    }

    function verifyClientState(
        IBCHost host,
        ConnectionEnd.Data memory connection,
        Height.Data memory height,
        bytes memory proof,
        bytes memory clientStateBytes
    ) internal returns (bool) {
        return verifyMembership(
            host,
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.clientStatePath(connection.counterparty.client_id),
            clientStateBytes
        );
    }

    function verifyClientConsensusState(
        IBCHost host,
        ConnectionEnd.Data memory connection,
        Height.Data memory height,
        Height.Data memory consensusHeight,
        bytes memory proof,
        bytes memory consensusStateBytes
    ) internal returns (bool) {
        return verifyMembership(
            host,
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.consensusStatePath(
                connection.counterparty.client_id, consensusHeight.revision_number, consensusHeight.revision_height
            ),
            consensusStateBytes
        );
    }

    function verifyConnectionState(
        IBCHost host,
        ConnectionEnd.Data memory connection,
        Height.Data memory height,
        bytes memory proof,
        string memory connectionId,
        ConnectionEnd.Data memory counterpartyConnection
    ) internal returns (bool) {
        return verifyMembership(
            host,
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.connectionPath(connectionId),
            ConnectionEnd.encode(counterpartyConnection)
        );
    }

    function verifyChannelState(
        IBCHost host,
        ConnectionEnd.Data memory connection,
        Height.Data memory height,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes
    ) internal returns (bool) {
        return verifyMembership(
            host,
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.channelPath(portId, channelId),
            channelBytes
        );
    }

    function verifyPacketCommitment(
        IBCHost host,
        ConnectionEnd.Data memory connection,
        Height.Data memory height,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) internal returns (bool) {
        return verifyMembership(
            host,
            connection.client_id,
            height,
            connection.delay_period,
            calcBlockDelay(host, connection.delay_period),
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.packetCommitmentPath(portId, channelId, sequence),
            abi.encodePacked(commitmentBytes)
        );
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        ConnectionEnd.Data memory connection,
        Height.Data memory height,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 acknowledgementCommitmentBytes
    ) internal returns (bool) {
        return verifyMembership(
            host,
            connection.client_id,
            height,
            connection.delay_period,
            calcBlockDelay(host, connection.delay_period),
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCIdentifier.packetAcknowledgementCommitmentPath(portId, channelId, sequence),
            abi.encodePacked(acknowledgementCommitmentBytes)
        );
    }

    function calcBlockDelay(IBCHost host, uint64 timeDelay) private view returns (uint64) {
        uint64 blockDelay = 0;
        uint64 expectedTimePerBlock = host.getExpectedTimePerBlock();
        if (expectedTimePerBlock != 0) {
            blockDelay = (timeDelay + expectedTimePerBlock - 1) / expectedTimePerBlock;
        }
        return blockDelay;
    }
}
