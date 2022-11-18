const IBFT2Client = artifacts.require("IBFT2Client");
const MockClient = artifacts.require("MockClient");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannel = artifacts.require("IBCChannel");
const IBCHandler = artifacts.require("IBCHandler");
const IBCMsgs = artifacts.require("IBCMsgs");
const IBCIdentifier = artifacts.require("IBCIdentifier");
const SimpleToken = artifacts.require("SimpleToken");
const ICS20TransferBank = artifacts.require("ICS20TransferBank");
const ICS20Bank = artifacts.require("ICS20Bank");

module.exports = async function (deployer) {
  await deployer.deploy(IBCIdentifier);
  await deployer.link(IBCIdentifier, [IBCHandler, IBCClient, IBCConnection, IBCChannel, IBFT2Client]);

  await deployer.deploy(IBCMsgs);
  await deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannel, IBCHandler, IBFT2Client]);

  await deployer.deploy(IBCClient);
  await deployer.deploy(IBCConnection);
  await deployer.deploy(IBCChannel);
  await deployer.deploy(IBCHandler, IBCClient.address, IBCConnection.address, IBCChannel.address);

  await deployer.deploy(MockClient);
  await deployer.deploy(IBFT2Client);
  await deployer.deploy(SimpleToken, "simple", "simple", 1000000);
  await deployer.deploy(ICS20Bank);
  await deployer.deploy(ICS20TransferBank, IBCHandler.address, ICS20Bank.address);
};
