package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress = "0x702E40245797c5a2108A566b3CE2Bf14Bc6aF841"
	IBCCommitmentTestHelperAddress = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
	SimpleTokenAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	ICS20TransferBankAddress = "0xa7f733a4fEA1071f58114b203F57444969b86524"
	ICS20BankAddress = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
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
