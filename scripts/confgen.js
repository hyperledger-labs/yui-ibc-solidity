const IBFT2Client = artifacts.require("IBFT2Client");
const MockClient = artifacts.require("MockClient");
const IBCHandler = artifacts.require("IBCHandler");
const IBCCommitment = artifacts.require("IBCCommitment");
const SimpleToken = artifacts.require("SimpleToken");
const ICS20TransferBank = artifacts.require("ICS20TransferBank");
const ICS20Bank = artifacts.require("ICS20Bank");

var fs = require("fs");
var ejs = require("ejs");

if (!process.env.CONF_TPL) {
  console.log("You must set environment variable 'CONF_TPL'");
  process.exit(1);
}

const makePairs = function(arr) {
  var pairs = [];
  for (var i=0 ; i<arr.length ; i+=2) {
      if (arr[i+1] !== undefined) {
          pairs.push ([arr[i], arr[i+1]]);
      } else {
          console.error("invalid pair found");
          process.exit(1);
      }
  }
  return pairs;
};

const targets = makePairs(process.env.CONF_TPL.split(":"));

module.exports = function(callback) {
  targets.forEach(function(item) {
    ejs.renderFile(item[1], {
      IBCHandlerAddress: IBCHandler.address,
      IBFT2ClientAddress: IBFT2Client.address,
      MockClientAddress: MockClient.address,
      IBCCommitmentAddress: IBCCommitment.address,
      SimpleTokenAddress: SimpleToken.address,
      ICS20TransferBankAddress: ICS20TransferBank.address,
      ICS20BankAddress: ICS20Bank.address
    }, null, function(err, str){
        if (err) {
          throw err;
        }
        fs.writeFileSync(item[0], str);
        console.log('generated file', item[0]);
      });
  });
  callback();
};
