pragma solidity ^0.8.9;

import "../core/IBCModule.sol";

interface IICS20Transfer is IModuleCallbacks {
    function sendTransfer(
        string calldata denom,
        uint64 amount,
        address receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external;
}
