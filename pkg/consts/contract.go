package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xDf919000A9A2533Fb000866654ce06D565DDDb97"
	IBCHandlerAddress = "0xc0ba8346289ec43cd3f68E5EBf0a3169B1d14a2d"
	IBCIdentifierAddress = "0x1D268db12EB341E196C682ebB94314Af2cAC0d8d"
	IBFT2ClientAddress = "0x180B6C325525dB54C0DA871F1fd924a0bcf06397"
	SimpleTokenModuleAddress = "0xA2f3403490466E33dcF0d74cAfc1DE0BeE0f47B4"
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

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
