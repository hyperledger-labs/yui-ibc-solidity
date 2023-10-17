// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./ILightClient.sol";
import "../25-handler/IBCMsgs.sol";
import "../24-host/IBCStore.sol";
import "../24-host/IBCCommitment.sol";
import "../02-client/IIBCClient.sol";

/**
 * @dev IBCClient is a contract that implements [ICS-2](https://github.com/cosmos/ibc/tree/main/spec/core/ics-002-client-semantics).
 */
contract IBCClient is IBCStore, IIBCClient {
    /**
     * @dev registerClient registers a new client type into the client registry
     */
    function registerClient(string calldata clientType, ILightClient client) external override {
        require(validateClientType(bytes(clientType)), "invalid clientType");
        require(address(clientRegistry[clientType]) == address(0), "clientType already exists");
        require(address(client) != address(this) && Address.isContract(address(client)), "invalid client address");
        clientRegistry[clientType] = address(client);
    }

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external override returns (string memory clientId) {
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

        return clientId;
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external override {
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

    /**
     * @dev validateClientType validates the client type
     *   - clientType must be non-empty
     *   - clientType must be in the form of `^[a-z][a-z0-9-]*[a-z0-9]$`
     */
    function validateClientType(bytes memory clientTypeBytes) internal pure returns (bool) {
        if (clientTypeBytes.length == 0) {
            return false;
        }
        unchecked {
            for (uint256 i = 0; i < clientTypeBytes.length; i++) {
                uint256 c = uint256(uint8(clientTypeBytes[i]));
                if (0x61 <= c && c <= 0x7a) {
                    // a-z
                    continue;
                } else if (c == 0x2d) {
                    // hyphen cannot be the first or last character
                    if (i == 0 || i == clientTypeBytes.length - 1) {
                        return false;
                    }
                    continue;
                } else if (0x30 <= c && c <= 0x39) {
                    // 0-9
                    continue;
                } else {
                    return false;
                }
            }
        }
        return true;
    }
}
