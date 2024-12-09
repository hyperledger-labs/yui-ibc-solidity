// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ILightClient} from "./ILightClient.sol";
import {Height} from "../../proto/Client.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IBCClientLib} from "../02-client/IBCClientLib.sol";
import {IIBCClientErrors} from "../02-client/IIBCClientErrors.sol";

/**
 * @dev IBCClient is a contract that implements [ICS-2](https://github.com/cosmos/ibc/tree/main/spec/core/ics-002-client-semantics).
 */
contract IBCClient is IBCHost, IIBCClient, IIBCClientErrors {
    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(MsgCreateClient calldata msg_) external override returns (string memory clientId) {
        address clientImpl = getHostStorage().clientRegistry[msg_.clientType];
        if (clientImpl == address(0)) {
            revert IBCClientUnregisteredClientType(msg_.clientType);
        }
        clientId = generateClientIdentifier(msg_.clientType);
        ClientStorage storage client = getClientStorage()[clientId];
        client.clientType = msg_.clientType;
        client.clientImpl = clientImpl;
        Height.Data memory height =
            ILightClient(clientImpl).initializeClient(clientId, msg_.protoClientState, msg_.protoConsensusState);
        // update commitments
        mapping(bytes32 => bytes32) storage commitments = getCommitments();
        commitments[IBCCommitment.clientStateCommitmentKey(clientId)] = keccak256(msg_.protoClientState);
        commitments[IBCCommitment.consensusStateCommitmentKey(clientId, height.revision_number, height.revision_height)]
        = keccak256(msg_.protoConsensusState);
        emit GeneratedClientIdentifier(clientId);
        return clientId;
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(MsgUpdateClient calldata msg_) external override {
        (address lc, bytes4 selector, bytes memory args) = routeUpdateClient(msg_);
        // NOTE: We assume that the client contract was correctly validated by the authority at registration via `registerClient` function.
        //       For details, see the `registerClient` function in the IBCHostConfigurator.
        (bool success, bytes memory returndata) = lc.call(abi.encodePacked(selector, args));
        if (!success) {
            if (returndata.length > 0) {
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert IBCClientFailedUpdateClient(selector, args);
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
     *      WARNING: If the caller is an EOA like a relayer, the caller must validate the return values with the allow list of the contract functions before calling the LC contract with the data.
     *               This validation is always required because even if the caller trusts the IBC contract, a malicious RPC provider can return arbitrary data to the caller.
     *      Check ADR-001 for details.
     */
    function routeUpdateClient(MsgUpdateClient calldata msg_)
        public
        view
        override
        returns (address, bytes4, bytes memory)
    {
        ILightClient lc = checkAndGetClient(msg_.clientId);
        // NOTE: The `lc.routeUpdateClient` function must be validated by the authority at registration via `registerClient` function.
        (bytes4 functionId, bytes memory args) = lc.routeUpdateClient(msg_.clientId, msg_.protoClientMessage);
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
        if (!found) {
            revert IBCClientClientNotFound(clientId);
        }
        mapping(bytes32 => bytes32) storage commitments = getCommitments();
        commitments[IBCCommitment.clientStateCommitmentKey(clientId)] = keccak256(clientState);
        for (uint256 i = 0; i < heights.length; i++) {
            (consensusState, found) = lc.getConsensusState(clientId, heights[i]);
            if (!found) {
                revert IBCClientConsensusStateNotFound(clientId, heights[i]);
            }
            bytes32 key = IBCCommitment.consensusStateCommitmentKey(
                clientId, heights[i].revision_number, heights[i].revision_height
            );
            bytes32 commitment = keccak256(consensusState);
            bytes32 prev = commitments[key];
            if (prev != bytes32(0) && commitment != prev) {
                // Revert if the new commitment is inconsistent with the previous one.
                // This case may indicate misbehavior of either the LightClient or the target chain.
                // Since the definition and specification of misbehavior are defined for each LightClient,
                // if a relayer detects this error, it is recommended to submit an evidence of misbehaviour to the LightClient accordingly.
                // (e.g., via the updateClient function).
                revert IBCClientInconsistentConsensusStateCommitment(key, commitment, prev);
            }
            commitments[key] = commitment;
        }
    }

    /**
     * @dev generateClientIdentifier generates a new client identifier for a given client type
     */
    function generateClientIdentifier(string calldata clientType) internal returns (string memory) {
        HostStorage storage hostStorage = getHostStorage();
        string memory identifier =
            string(abi.encodePacked(clientType, "-", Strings.toString(hostStorage.nextClientSequence)));
        if (!IBCClientLib.validateClientId(bytes(identifier))) {
            revert IBCClientInvalidClientId(identifier);
        }
        hostStorage.nextClientSequence++;
        return identifier;
    }
}
