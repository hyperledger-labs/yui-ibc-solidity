pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import "./IHandler.sol";

abstract contract IBCClient is IHandler, IBCHost {
    // clientType => Contract
    mapping(string => IClient) clients;

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCMsgs.MsgCreateClient memory msg_) public override {
        require(!ibcStore.hasClientState(msg_.clientId), "the clientId already exists");
        require(address(clients[msg_.clientType]) != address(0), "unregistered client type");

        ibcStore.setClientType(msg_.clientId, msg_.clientType);
        ibcStore.setClientState(msg_.clientId, msg_.clientStateBytes);
        ibcStore.setConsensusState(msg_.clientId, msg_.height, msg_.consensusStateBytes);
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient memory msg_) public override {
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
    function validateSelfClient(bytes memory clientStateBytes) internal view returns (bool) {
        return true;
    }

    function registerClient(string memory clientType, IClient client) public {
        clients[clientType] = client;
    }

    function getClient(string memory clientId) internal view returns (IClient) {
        return clients[ibcStore.getClientType(clientId)];
    }
}
