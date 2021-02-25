pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import "../lib/Ownable.sol";

contract IBCClient is IBCHost, Ownable {

    constructor(IBCStore store) IBCHost(store) public {}

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCMsgs.MsgCreateClient memory msg_) public {
        (, bool found) = ibcStore.getClientState(msg_.clientId);
        require(!found, "the clientId already exists");
        (, found) = getClientByType(msg_.clientType);
        require(found, "unregistered client type");

        ibcStore.setClientType(msg_.clientId, msg_.clientType);
        ibcStore.setClientState(msg_.clientId, msg_.clientStateBytes);
        ibcStore.setConsensusState(msg_.clientId, msg_.height, msg_.consensusStateBytes);
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient memory msg_) public {
        bytes memory clientStateBytes;
        bytes memory consensusStateBytes;
        uint64 height;
        bool found;
    
        (clientStateBytes, found) = ibcStore.getClientState(msg_.clientId);
        require(found, "clientState not found");

        (clientStateBytes, consensusStateBytes, height) = getClient(msg_.clientId).checkHeaderAndUpdateState(msg_.clientId, clientStateBytes, msg_.header);
    
        //// persist states ////
        ibcStore.setClientState(msg_.clientId, clientStateBytes);
        ibcStore.setConsensusState(msg_.clientId, height, consensusStateBytes);
    }

    // TODO implements
    function validateSelfClient(bytes calldata clientStateBytes) external view returns (bool) {
        return true;
    }

    function registerClient(string memory clientType, IClient client) onlyOwner public {
        ibcStore.setClientImpl(clientType, address(client));
    }

    function getClient(string memory clientId) public view returns (IClient) {
        (IClient clientImpl, bool found) = getClientByType(ibcStore.getClientType(clientId));
        require(found, "clientImpl not found");
        return clientImpl;
    }

    function getClientByType(string memory clientType) internal view returns (IClient clientImpl, bool) {
        (address addr, bool found) = ibcStore.getClientImpl(clientType);
        if (!found) {
            return (clientImpl, false);
        }
        return (IClient(addr), true);
    }
}
