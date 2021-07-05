pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

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
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCHost host, IBCMsgs.MsgUpdateClient calldata msg_) external {
        host.onlyIBCModule();
        bytes memory clientStateBytes;
        bytes memory consensusStateBytes;
        uint64 height;
        bool found;
    
        (clientStateBytes, found) = host.getClientState(msg_.clientId);
        require(found, "clientState not found");

        (clientStateBytes, consensusStateBytes, height) = getClient(host, msg_.clientId).checkHeaderAndUpdateState(host, msg_.clientId, clientStateBytes, msg_.header);
    
        //// persist states ////
        host.setClientState(msg_.clientId, clientStateBytes);
        host.setConsensusState(msg_.clientId, height, consensusStateBytes);
    }

    // TODO implements
    function validateSelfClient(IBCHost host, bytes calldata clientStateBytes) external view returns (bool) {
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
}
