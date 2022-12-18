package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress              = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
	IBCCommitmentTestHelperAddress = "0xdD5109D05Ac357E446992a60E64764041A0E8529"
	SimpleTokenAddress             = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
	ICS20TransferBankAddress       = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
	ICS20BankAddress               = "0xa7f733a4fEA1071f58114b203F57444969b86524"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCCommitmentTestHelperAddress() common.Address {
	return common.HexToAddress(IBCCommitmentTestHelperAddress)
}

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferBankAddress() common.Address {
	return common.HexToAddress(ICS20TransferBankAddress)
}

func (contractConfig) GetICS20BankAddress() common.Address {
	return common.HexToAddress(ICS20BankAddress)
}
