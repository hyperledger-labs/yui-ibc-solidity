const Migrations = artifacts.require("Migrations");
const IBCStore = artifacts.require("IBCStore");
const IBFT2Client = artifacts.require("IBFT2Client");
const IBCModule = artifacts.require("IBCModule");
const IBCMsgs = artifacts.require("IBCMsgs");
const SimpleTokenModule = artifacts.require("SimpleTokenModule");
const Bytes = artifacts.require("Bytes");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(Bytes).then(function() {
    return deployer.link(Bytes, [IBCModule, IBFT2Client, SimpleTokenModule]);
  });
  deployer.deploy(IBCMsgs).then(function() {
    return deployer.link(IBCMsgs, [IBCModule, IBFT2Client]);
  });
  deployer.deploy(IBCStore).then(function() {
    return deployer.deploy(IBFT2Client, IBCStore.address).then(function() {
      return deployer.deploy(IBCModule, IBCStore.address, IBFT2Client.address).then(function() {
        return deployer.deploy(SimpleTokenModule, IBCStore.address, IBCModule.address);
      });
    });
  });
};
