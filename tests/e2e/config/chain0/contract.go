package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xDf2b5e65098EDCeE79E0633261F2347e3c70A682"
	IBCClientAddress = "0x94872f303a5c0f5afc20470A3De17198250A169b"
	IBCConnectionAddress = "0x57837cD742A4Ff8a3313587495F9965C0aBcC530"
	IBCChannelAddress = "0xCE76b1a5a7A6DC9cb5ddbE307c08c02bE1694638"
	IBCRoutingModuleAddress = "0x82eE5F6fDfb6216Ed8f00Db7848cb7FCccF6c20c"
	IBCIdentifierAddress = "0x6889D1E8f2269Ee96D08Fc5dF4295da21297d525"
	IBFT2ClientAddress = "0xC799723F0Bc6142b06091F193b9D04374cdc5530"
	SimpleTokenModuleAddress = "0xf157aD1Da0952ecC896485f362Edd506Bd9A567A"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetIBCClientAddress() common.Address {
	return common.HexToAddress(IBCClientAddress)
}

func (contractConfig) GetIBCConnectionAddress() common.Address {
	return common.HexToAddress(IBCConnectionAddress)
}

func (contractConfig) GetIBCChannelAddress() common.Address {
	return common.HexToAddress(IBCChannelAddress)
}

func (contractConfig) GetIBCRoutingModuleAddress() common.Address {
	return common.HexToAddress(IBCRoutingModuleAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
