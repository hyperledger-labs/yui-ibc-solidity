const IBCHost = artifacts.require("IBCHost");
const IBFT2Client = artifacts.require("IBFT2Client");
const MockClient = artifacts.require("MockClient");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannelHandshake = artifacts.require("IBCChannelHandshake");
const IBCChannelPacket = artifacts.require("IBCChannelPacket");
const IBCHandler = artifacts.require("IBCHandler");
const IBCMsgs = artifacts.require("IBCMsgs");
const IBCIdentifier = artifacts.require("IBCIdentifier");
const SimpleToken = artifacts.require("SimpleToken");
const ICS20TransferBank = artifacts.require("ICS20TransferBank");
const ICS20Bank = artifacts.require("ICS20Bank");

module.exports = function (deployer) {
  deployer.deploy(IBCIdentifier).then(function() {
    return deployer.link(IBCIdentifier, [IBCHost, IBFT2Client, IBCHandler]);
  });
  deployer.deploy(IBCMsgs).then(function() {
    return deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannelHandshake, IBCChannelPacket, IBCHandler, IBFT2Client]);
  });
  deployer.deploy(IBCClient).then(function() {
    return deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannelHandshake, IBCChannelPacket]);
  });
  deployer.deploy(IBCConnection).then(function() {
    return deployer.link(IBCConnection, [IBCHandler, IBCChannelHandshake, IBCChannelPacket]);
  });
  deployer.deploy(IBCChannelHandshake).then(function() {
    return deployer.link(IBCChannelHandshake, [IBCHandler]);
  });
  deployer.deploy(IBCChannelPacket).then(function() {
    return deployer.link(IBCChannelPacket, [IBCHandler]);
  });
  deployer.deploy(IBFT2Client);
  deployer.deploy(MockClient);
  deployer.deploy(IBCHost).then(function() {
    return deployer.deploy(IBCHandler, IBCHost.address);
  });
  deployer.deploy(SimpleToken, "simple", "simple", 1000000);
  deployer.deploy(ICS20Bank).then(function() {
    return deployer.deploy(ICS20TransferBank, IBCHost.address, IBCHandler.address, ICS20Bank.address);
  });
};
