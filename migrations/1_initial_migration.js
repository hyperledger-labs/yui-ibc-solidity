const Migrations = artifacts.require("Migrations");
const IBCHost = artifacts.require("IBCHost");
const IBFT2Client = artifacts.require("IBFT2Client");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannel = artifacts.require("IBCChannel");
const IBCHandler = artifacts.require("IBCHandler");
const IBCMsgs = artifacts.require("IBCMsgs");
const SimpleTokenModule = artifacts.require("SimpleTokenModule");
const Bytes = artifacts.require("Bytes");
const IBCIdentifier = artifacts.require("IBCIdentifier");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(Bytes).then(function() {
    return deployer.link(Bytes, [IBCClient, IBCConnection, IBCChannel, IBCHandler, IBFT2Client, SimpleTokenModule]);
  });
  deployer.deploy(IBCIdentifier).then(function() {
    return deployer.link(IBCIdentifier, [IBCHost, IBFT2Client, IBCHandler, SimpleTokenModule]);
  });
  deployer.deploy(IBCMsgs).then(function() {
    return deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannel, IBCHandler, IBFT2Client]);
  });
  deployer.deploy(IBCClient).then(function() {
    return deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannel]);
  });
  deployer.deploy(IBCConnection).then(function() {
    return deployer.link(IBCConnection, [IBCHandler, IBCChannel]);
  });
  deployer.deploy(IBCChannel).then(function() {
    return deployer.link(IBCChannel, [IBCHandler, SimpleTokenModule]);
  });
  deployer.deploy(IBFT2Client);
  deployer.deploy(IBCHost).then(function() {
    return deployer.deploy(IBCHandler, IBCHost.address).then(function() {
      return deployer.deploy(SimpleTokenModule, IBCHost.address, IBCHandler.address);
    });
  });

};
