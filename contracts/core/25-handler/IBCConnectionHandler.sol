// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Address.sol";
import "../25-handler/IBCMsgs.sol";
import "../24-host/IBCHost.sol";
import "../03-connection/IIBCConnection.sol";

/**
 * @dev IBCConnectionHandler is a contract that calls a contract that implements `IIBCConnectionHandshake` with delegatecall.
 */
abstract contract IBCConnectionHandler {
    using Address for address;

    // IBC Connection contract address
    address immutable ibcConnection;

    event GeneratedConnectionIdentifier(string);

    constructor(address _ibcConnection) {
        require(Address.isContract(_ibcConnection));
        ibcConnection = _ibcConnection;
    }

    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit calldata msg_)
        external
        returns (string memory connectionId)
    {
        bytes memory res = ibcConnection.functionDelegateCall(
            abi.encodeWithSelector(IIBCConnectionHandshake.connectionOpenInit.selector, msg_)
        );
        connectionId = abi.decode(res, (string));
        emit GeneratedConnectionIdentifier(connectionId);
        return connectionId;
    }

    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry calldata msg_)
        external
        returns (string memory connectionId)
    {
        bytes memory res = ibcConnection.functionDelegateCall(
            abi.encodeWithSelector(IIBCConnectionHandshake.connectionOpenTry.selector, msg_)
        );
        connectionId = abi.decode(res, (string));
        emit GeneratedConnectionIdentifier(connectionId);
        return connectionId;
    }

    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck calldata msg_) external {
        ibcConnection.functionDelegateCall(
            abi.encodeWithSelector(IIBCConnectionHandshake.connectionOpenAck.selector, msg_)
        );
    }

    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm calldata msg_) external {
        ibcConnection.functionDelegateCall(
            abi.encodeWithSelector(IIBCConnectionHandshake.connectionOpenConfirm.selector, msg_)
        );
    }
}
