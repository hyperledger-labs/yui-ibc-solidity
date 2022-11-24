// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";
import "./IBCCommitment.sol";

contract IBCClient is IBCHost {
    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external {
        address clientImpl = clientRegistry[msg_.clientType];
        require(clientImpl != address(0), "unregistered client type");
        string memory clientId = generateClientIdentifier(msg_.clientType);
        clientTypes[clientId] = msg_.clientType;
        clientImpls[clientId] = clientImpl;
        (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok) =
            IClient(clientImpl).createClient(clientId, msg_.height, msg_.clientStateBytes, msg_.consensusStateBytes);
        require(ok, "failed to create client");
        updateCommitments(clientId, clientStateCommitment, updates);
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        require(commitments[IBCCommitment.clientStateCommitmentKey(msg_.clientId)] != bytes32(0));
        (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok) =
            getClient(msg_.clientId).updateClient(msg_.clientId, msg_.clientMessage);
        require(ok, "failed to update client");
        updateCommitments(msg_.clientId, clientStateCommitment, updates);
    }

    function updateCommitments(
        string memory clientId,
        bytes32 clientStateCommitment,
        ConsensusStateUpdates[] memory updates
    ) private {
        commitments[keccak256(IBCCommitment.clientStatePath(clientId))] = clientStateCommitment;
        for (uint256 i = 0; i < updates.length; i++) {
            commitments[IBCCommitment.consensusStateCommitmentKey(
                clientId, updates[i].height.revision_number, updates[i].height.revision_height
            )] = updates[i].consensusStateCommitment;
        }
    }

    function generateClientIdentifier(string calldata clientType) private returns (string memory) {
        string memory identifier = string(abi.encodePacked(clientType, "-", uint2str(nextClientSequence)));
        nextClientSequence++;
        emit GeneratedClientIdentifier(identifier);
        return identifier;
    }
}
