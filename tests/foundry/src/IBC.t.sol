// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../../contracts/core/IBCHost.sol";
import "../../../contracts/core/IBCHandler.sol";
import "../../../contracts/core/MockClient.sol";

contract IBCTest {
    IBCHost host;
    IBCHandler handler;
    MockClient mockClient;

    string private constant mockClientType = "mock-client";

    function setUp() public {
        mockClient = new MockClient();
        host = new IBCHost();
        handler = new IBCHandler(host);
        host.setIBCModule(address(handler));
        handler.registerClient(mockClientType, mockClient);
    }

    function testClientRegistration() public view {
        (address client, bool found) = host.getClientImpl(mockClientType);
        assert(found);
        assert(address(mockClient) == client);
    }
}
