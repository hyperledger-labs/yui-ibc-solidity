const Migrations = artifacts.require("Migrations");
const IBCStore = artifacts.require("IBCStore");
const IBFT2Client = artifacts.require("IBFT2Client");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannel = artifacts.require("IBCChannel");
const IBCRoutingModule = artifacts.require("IBCRoutingModule");
const IBCMsgs = artifacts.require("IBCMsgs");
const SimpleTokenModule = artifacts.require("SimpleTokenModule");
const Bytes = artifacts.require("Bytes");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(Bytes).then(function() {
    return deployer.link(Bytes, [IBCClient, IBCConnection, IBCChannel, IBCRoutingModule, IBFT2Client, SimpleTokenModule]);
  });
  deployer.deploy(IBCMsgs).then(function() {
    return deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannel, IBCRoutingModule, IBFT2Client]);
  });
  deployer.deploy(IBCStore).then(function() {
    return deployer.deploy(IBFT2Client, IBCStore.address).then(function() {
      return deployer.deploy(IBCClient, IBCStore.address).then(function() {
        return deployer.deploy(IBCConnection, IBCStore.address, IBCClient.address).then(function() {
          return deployer.deploy(IBCChannel, IBCStore.address, IBCClient.address, IBCConnection.address).then(function() {
            return deployer.deploy(IBCRoutingModule, IBCStore.address, IBCChannel.address).then(function() {
              return deployer.deploy(SimpleTokenModule, IBCStore.address, IBCChannel.address, IBCRoutingModule.address);
            });
          });
        });
      });
    });
  });
};
