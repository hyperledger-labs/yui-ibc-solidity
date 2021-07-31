package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xff77D90D6aA12db33d3Ba50A34fB25401f6e4c4F"
	IBCHandlerAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	IBCIdentifierAddress = "0xB9c99Dc02185993bdB9C48Fc29544f6cC6604F87"
	IBFT2ClientAddress = "0x702E40245797c5a2108A566b3CE2Bf14Bc6aF841"
	MockClientAddress = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
	SimpleTokenAddress = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
	ICS20TransferBankAddress = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
	ICS20BankAddress = "0xa7f733a4fEA1071f58114b203F57444969b86524"
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
