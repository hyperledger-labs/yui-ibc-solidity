const Migrations = artifacts.require("Migrations");
const ProvableStore = artifacts.require("ProvableStore");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const Bytes = artifacts.require("Bytes");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(Bytes).then(function() {
    return deployer.link(Bytes, [IBCClient, IBCConnection]);
  });
  deployer.deploy(ProvableStore).then(function() {
    return deployer.deploy(IBCClient, ProvableStore.address).then(function() {
      return deployer.deploy(IBCConnection, ProvableStore.address, IBCClient.address);
    });
  });
};
