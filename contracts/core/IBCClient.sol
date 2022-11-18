// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";
import "./IBCIdentifier.sol";

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
        require(ok);
        commitments[keccak256(IBCIdentifier.clientStatePath(clientId))] = clientStateCommitment;
        for (uint256 i = 0; i < updates.length; i++) {
            commitments[keccak256(
                IBCIdentifier.consensusStatePath(
                    clientId, updates[i].height.revision_number, updates[i].height.revision_height
                )
            )] = updates[i].consensusStateCommitment;
        }
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        require(commitments[keccak256(IBCIdentifier.clientStatePath(msg_.clientId))] != bytes32(0));
        getClient(msg_.clientId).verifyClientMessageAndUpdateState(msg_.clientId, msg_.clientMessage);
    }
}
