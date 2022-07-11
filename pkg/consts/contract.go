package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	IBCHandlerAddress = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
	IBCIdentifierAddress = "0xdD5109D05Ac357E446992a60E64764041A0E8529"
	IBFT2ClientAddress = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
	MockClientAddress = "0xff77D90D6aA12db33d3Ba50A34fB25401f6e4c4F"
	SimpleTokenAddress = "0xa7f733a4fEA1071f58114b203F57444969b86524"
	ICS20TransferBankAddress = "0x37978908bac82F0191b674235A0fEEE31e7524a4"
	ICS20BankAddress = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetMockClientAddress() common.Address {
	return common.HexToAddress(MockClientAddress)
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
