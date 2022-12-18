const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannelHandshake = artifacts.require("IBCChannelHandshake");
const IBCPacket = artifacts.require("IBCPacket");
const IBCCommitment = artifacts.require("IBCCommitment");
const IBCHandler = artifacts.require("OwnableIBCHandler");
const SimpleToken = artifacts.require("SimpleToken");
const ICS20TransferBank = artifacts.require("ICS20TransferBank");
const ICS20Bank = artifacts.require("ICS20Bank");
const IBFT2Client = artifacts.require("IBFT2Client");
const MockClient = artifacts.require("MockClient");

module.exports = async function (deployer) {
  await deployer.deploy(IBCCommitment);
  await deployer.link(IBCCommitment, [IBCHandler, IBCClient, IBCConnection, IBCChannelHandshake, IBCPacket]);

  await deployer.deploy(IBCClient);
  await deployer.deploy(IBCConnection);
  await deployer.deploy(IBCChannelHandshake);
  await deployer.deploy(IBCPacket);
  await deployer.deploy(IBCHandler, IBCClient.address, IBCConnection.address, IBCChannelHandshake.address, IBCPacket.address);

  await deployer.deploy(MockClient, IBCHandler.address);
  await deployer.deploy(IBFT2Client, IBCHandler.address);
  await deployer.deploy(SimpleToken, "simple", "simple", 1000000);
  await deployer.deploy(ICS20Bank);
  await deployer.deploy(ICS20TransferBank, IBCHandler.address, ICS20Bank.address);
};
