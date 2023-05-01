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
import {SimpleToken} from "../../../contracts/apps/20-transfer/SimpleToken.sol";
import {IBCCommitmentTestHelper} from "./helpers/IBCCommitmentTestHelper.t.sol";

contract DeployScript is Script {
    string private constant MOCK_CLIENT_TYPE = "mock-client";
    string private constant IBFT2_CLIENT_TYPE = "hyperledger-besu-ibft2";

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // deploy core contracts
        address ibcClient = address(new IBCClient());
        address ibcConnection = address(new IBCConnection());
        address ibcChannelHandshake = address(new IBCChannelHandshake());
        address ibcPacket = address(new IBCPacket());
        OwnableIBCHandler handler = new OwnableIBCHandler(ibcClient, ibcConnection, ibcChannelHandshake, ibcPacket);

        // deploy app contracts
        ICS20Bank bank = new ICS20Bank();
        address transferBank = address(new ICS20TransferBank(handler, bank));
        bank.setOperator(transferBank);
        handler.bindPort("transfer", transferBank);

        // deploy client contracts
        MockClient mockClient = new MockClient(address(handler));
        IBFT2Client ibft2Client = new IBFT2Client(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        handler.registerClient(IBFT2_CLIENT_TYPE, ibft2Client);

        // deploy test helpers
        new SimpleToken("simple", "simple", 1000000);
        new IBCCommitmentTestHelper();

        vm.stopBroadcast();
    }
}
