// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "./IBCTestHelper.t.sol";
import {Vm} from "forge-std/Test.sol";
import {ErrorReceipt} from "../../../../contracts/proto/Channel.sol";

abstract contract ICS04UpgradeTestHelper is IBCTestHelper {
    event WriteErrorReceipt(string portId, string channelId, uint64 upgradeSequence, string message);

    function decodeWriteErrorReceiptEvent(bytes memory data) internal pure returns (ErrorReceipt.Data memory) {
        (,, uint64 upgradeSequence, string memory message) = abi.decode(data, (string, string, uint64, string));
        return ErrorReceipt.Data({sequence: upgradeSequence, message: message});
    }

    function tryDecodeWriteErrorReceiptEvent(Vm.Log memory log)
        internal
        pure
        returns (ErrorReceipt.Data memory rc, bool ok)
    {
        if (log.topics[0] != WriteErrorReceipt.selector) {
            return (rc, false);
        }
        return (decodeWriteErrorReceiptEvent(log.data), true);
    }

    function getLastWriteErrorReceiptEvent(IIBCHandler handler, Vm.Log[] memory logs)
        internal
        pure
        returns (ErrorReceipt.Data memory)
    {
        for (uint256 i = logs.length; i > 0; i--) {
            if (logs[i - 1].emitter == address(handler)) {
                (ErrorReceipt.Data memory e, bool ok) = tryDecodeWriteErrorReceiptEvent(logs[i - 1]);
                if (ok) {
                    return e;
                }
            }
        }
        revert("no write error receipt event");
    }
}
