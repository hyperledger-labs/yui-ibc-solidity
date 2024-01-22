// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ILightClient} from "./ILightClient.sol";
import {Height} from "../../proto/Client.sol";
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
        (bytes32 clientStateCommitment, ILightClient.ConsensusStateUpdate memory update, bool ok) =
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
        (address lc, bytes4 selector, bytes memory args) = routeUpdateClient(msg_);
        (bool success, bytes memory returndata) = lc.call(abi.encodePacked(selector, args));
        if (!success) {
            if (returndata.length > 0) {
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert("update client failed");
            }
        }
        Height.Data[] memory heights = abi.decode(returndata, (Height.Data[]));
        if (heights.length > 0) {
            updateClientCommitments(msg_.clientId, heights);
        }
    }

    /**
     * @dev routeUpdateClient returns the LC contract address and the calldata to the receiving function of the client message.
     *      Light client contract may encode a client message as other encoding scheme(e.g. ethereum ABI)
     */
    function routeUpdateClient(MsgUpdateClient calldata msg_)
        public
        view
        override
        returns (address, bytes4, bytes memory)
    {
        ILightClient lc = checkAndGetClient(msg_.clientId);
        (bytes4 functionId, bytes memory args) = lc.routeUpdateClient(msg_.clientId, msg_.clientMessage);
        return (address(lc), functionId, args);
    }

    /**
     * @dev updateClientCommitments updates the commitments of the light client's states corresponding to the given heights.
     */
    function updateClientCommitments(string calldata clientId, Height.Data[] memory heights) public override {
        ILightClient lc = checkAndGetClient(clientId);
        bytes memory clientState;
        bytes memory consensusState;
        bool found;
        (clientState, found) = lc.getClientState(clientId);
        require(found, "client not found");
        commitments[IBCCommitment.clientStateCommitmentKey(clientId)] = keccak256(clientState);
        for (uint256 i = 0; i < heights.length; i++) {
            (consensusState, found) = lc.getConsensusState(clientId, heights[i]);
            require(found, "consensus state not found");
            bytes32 key = IBCCommitment.consensusStateCommitmentKey(
                clientId, heights[i].revision_number, heights[i].revision_height
            );
            require(commitments[key] == bytes32(0), "consensus state already exists");
            commitments[key] = keccak256(consensusState);
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
