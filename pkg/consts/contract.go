package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
	IBCHandlerAddress = "0xff77D90D6aA12db33d3Ba50A34fB25401f6e4c4F"
	IBCIdentifierAddress = "0xB9c99Dc02185993bdB9C48Fc29544f6cC6604F87"
	IBFT2ClientAddress = "0x702E40245797c5a2108A566b3CE2Bf14Bc6aF841"
	SimpleTokenAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	ICS20TransferAddress = "0xa7f733a4fEA1071f58114b203F57444969b86524"
	ICS20VouchersAddress = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
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

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20VouchersAddress() common.Address {
	return common.HexToAddress(ICS20VouchersAddress)
}
