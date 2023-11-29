// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCStore} from "../24-host/IBCStore.sol";

contract IBCHost is IBCStore {
    // It represents the prefix of the commitment proof(https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments#prefix).
    // In ibc-solidity, the prefix is not required, but for compatibility with ibc-go this must be a non-empty value.
    bytes internal constant DEFAULT_COMMITMENT_PREFIX = bytes("ibc");

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
        require(clientImpl != address(0));
        return ILightClient(clientImpl);
    }
}
