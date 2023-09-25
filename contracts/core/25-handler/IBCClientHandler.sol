// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Address.sol";
import "../24-host/IBCHost.sol";
import "../02-client/IIBCClient.sol";

/**
 * @dev IBCClientHandler is a contract that calls a contract that implements `IIBCClient` with delegatecall.
 */
abstract contract IBCClientHandler {
    using Address for address;

    // IBC Client contract address
    address immutable ibcClient;

    event GeneratedClientIdentifier(string);

    constructor(address _ibcClient) {
        require(Address.isContract(_ibcClient));
        ibcClient = _ibcClient;
    }

    /**
     * @dev registerClient registers a new client type into the client registry
     */
    function registerClient(string calldata clientType, ILightClient client) public virtual {
        ibcClient.functionDelegateCall(abi.encodeWithSelector(IIBCClient.registerClient.selector, clientType, client));
    }

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external returns (string memory clientId) {
        bytes memory res =
            ibcClient.functionDelegateCall(abi.encodeWithSelector(IIBCClient.createClient.selector, msg_));
        clientId = abi.decode(res, (string));
        emit GeneratedClientIdentifier(clientId);
        return clientId;
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        ibcClient.functionDelegateCall(abi.encodeWithSelector(IIBCClient.updateClient.selector, msg_));
    }
}
