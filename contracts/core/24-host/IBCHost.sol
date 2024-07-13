// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCStore} from "../24-host/IBCStore.sol";
import {IIBCHostErrors} from "./IIBCHostErrors.sol";

contract IBCHost is IIBCHostErrors, IBCStore {
    // It represents the prefix of the commitment proof(https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments#prefix).
    // In ibc-solidity, the prefix is not required, but for compatibility with ibc-go this must be a non-empty value.
    bytes internal constant DEFAULT_COMMITMENT_PREFIX = bytes("ibc");

    /**
     * @dev hostTimestamp returns the current timestamp(Unix time in nanoseconds) of the host chain.
     */
    function hostTimestamp() internal view virtual returns (uint64) {
        return uint64(block.timestamp) * 1e9;
    }

    /**
     * @dev _getCommitmentPrefix returns the prefix of the commitment proof.
     */
    function _getCommitmentPrefix() internal view virtual returns (bytes memory) {
        return DEFAULT_COMMITMENT_PREFIX;
    }

    /**
     * @dev checkAndGetClient returns the client implementation for the given client ID.
     */
    function checkAndGetClient(string memory clientId) internal view returns (ILightClient) {
        address clientImpl = clientImpls[clientId];
        if (clientImpl == address(0)) {
            revert IBCHostClientNotFound(clientId);
        }
        return ILightClient(clientImpl);
    }
}
