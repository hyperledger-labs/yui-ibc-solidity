// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import {IBCClient} from "../../../contracts/core/02-client/IBCClient.sol";
import {IBCConnection} from "../../../contracts/core/03-connection/IBCConnection.sol";
import {IBCChannelHandshake} from "../../../contracts/core/04-channel/IBCChannelHandshake.sol";
import {IBCPacket} from "../../../contracts/core/04-channel/IBCPacket.sol";
import {OwnableIBCHandler} from "../../../contracts/core/OwnableIBCHandler.sol";
import {MockClient} from "../../../contracts/clients/MockClient.sol";
import {IBFT2Client} from "../../../contracts/clients/IBFT2Client.sol";
import {ICS20Bank} from "../../../contracts/apps/20-transfer/ICS20Bank.sol";
import {ICS20TransferBank} from "../../../contracts/apps/20-transfer/ICS20TransferBank.sol";
import {ERC20Token} from "../../../contracts/apps/20-transfer/ERC20Token.sol";
import {IBCCommitmentTestHelper} from "./helpers/IBCCommitmentTestHelper.sol";
import {IBCMockApp} from "../../../contracts/apps/mock/IBCMockApp.sol";

contract DeployScript is Script {
    string private constant MOCK_CLIENT_TYPE = "mock-client";
    string private constant IBFT2_CLIENT_TYPE = "hyperledger-besu-ibft2";

    function run() external {
        uint256 privateKey =
            vm.deriveKey(vm.envString("TEST_MNEMONIC"), uint32(vm.envOr("TEST_MNEMONIC_INDEX", uint32(0))));
        vm.startBroadcast(privateKey);

        // deploy core contracts
        address ibcClient = address(new IBCClient());
        address ibcConnection = address(new IBCConnection());
        address ibcChannelHandshake = address(new IBCChannelHandshake());
        address ibcPacket = address(new IBCPacket());
        OwnableIBCHandler handler = new OwnableIBCHandler(ibcClient, ibcConnection, ibcChannelHandshake, ibcPacket);

        // deploy ics20 contract
        ICS20Bank bank = new ICS20Bank();
        address transferBank = address(new ICS20TransferBank(handler, bank));
        bank.setOperator(transferBank);
        handler.bindPort("transfer", transferBank);

        // deploy mock app contract
        IBCMockApp mockApp = new IBCMockApp(handler);
        handler.bindPort("mock", address(mockApp));

        // deploy client contracts
        MockClient mockClient = new MockClient(address(handler));
        IBFT2Client ibft2Client = new IBFT2Client(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        handler.registerClient(IBFT2_CLIENT_TYPE, ibft2Client);

        // deploy test helpers
        new ERC20Token("test", "test", 1000000);
        new IBCCommitmentTestHelper();

        vm.stopBroadcast();
    }
}
