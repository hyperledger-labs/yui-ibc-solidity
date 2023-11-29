// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ILightClient, ConsensusStateUpdate} from "./ILightClient.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCClient} from "../02-client/IIBCClient.sol";

/**
 * @dev IBCClient is a contract that implements [ICS-2](https://github.com/cosmos/ibc/tree/main/spec/core/ics-002-client-semantics).
 */
contract IBCClient is IBCHost, IIBCClient {
    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(MsgCreateClient calldata msg_) external override returns (string memory clientId) {
        address clientImpl = clientRegistry[msg_.clientType];
        require(clientImpl != address(0), "unregistered client type");
        clientId = generateClientIdentifier(msg_.clientType);
        clientTypes[clientId] = msg_.clientType;
        clientImpls[clientId] = clientImpl;
        (bytes32 clientStateCommitment, ConsensusStateUpdate memory update, bool ok) =
            ILightClient(clientImpl).createClient(clientId, msg_.clientStateBytes, msg_.consensusStateBytes);
        require(ok, "failed to create client");

        // update commitments
        commitments[IBCCommitment.clientStateCommitmentKey(clientId)] = clientStateCommitment;
        commitments[IBCCommitment.consensusStateCommitmentKey(
            clientId, update.height.revision_number, update.height.revision_height
        )] = update.consensusStateCommitment;
        emit GeneratedClientIdentifier(clientId);
        return clientId;
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(MsgUpdateClient calldata msg_) external override {
        require(commitments[IBCCommitment.clientStateCommitmentKey(msg_.clientId)] != bytes32(0));
        (bytes32 clientStateCommitment, ConsensusStateUpdate[] memory updates, bool ok) =
            checkAndGetClient(msg_.clientId).updateClient(msg_.clientId, msg_.clientMessage);
        require(ok, "failed to update client");

        // update commitments
        if (clientStateCommitment != 0) {
            commitments[IBCCommitment.clientStateCommitmentKey(msg_.clientId)] = clientStateCommitment;
        }
        for (uint256 i = 0; i < updates.length; i++) {
            commitments[IBCCommitment.consensusStateCommitmentKey(
                msg_.clientId, updates[i].height.revision_number, updates[i].height.revision_height
            )] = updates[i].consensusStateCommitment;
        }
    }

    /**
     * @dev generateClientIdentifier generates a new client identifier for a given client type
     */
    function generateClientIdentifier(string calldata clientType) private returns (string memory) {
        string memory identifier = string(abi.encodePacked(clientType, "-", Strings.toString(nextClientSequence)));
        nextClientSequence++;
        return identifier;
    }
}
