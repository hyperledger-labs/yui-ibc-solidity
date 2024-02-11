// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import "../../../../contracts/clients/MockClient.sol";

/**
 * @dev ModifiedMockClient is a modified MockClient implementation for testing purposes.
 */
contract ModifiedMockClient is MockClient {
    using IBCHeight for Height.Data;

    constructor(address _ibcHandler) MockClient(_ibcHandler) {}

    /**
     * @dev setStatus sets the status of the client corresponding to `clientId`.
     */
    function setStatus(string calldata clientId, ClientStatus status) external virtual {
        statuses[clientId] = status;
    }
}
