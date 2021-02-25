package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x57837cD742A4Ff8a3313587495F9965C0aBcC530"
	IBCClientAddress = "0x82eE5F6fDfb6216Ed8f00Db7848cb7FCccF6c20c"
	IBCConnectionAddress = "0xf157aD1Da0952ecC896485f362Edd506Bd9A567A"
	IBCChannelAddress = "0x2E44F7bc78444f7C8247a0e74ba87C2522b1D157"
	IBCRoutingModuleAddress = "0xB998570D81c7cE63741fCBA8eaE8FF322c54D8Bb"
	IBCIdentifierAddress = "0xC799723F0Bc6142b06091F193b9D04374cdc5530"
	IBFT2ClientAddress = "0xCE76b1a5a7A6DC9cb5ddbE307c08c02bE1694638"
	SimpleTokenModuleAddress = "0xf92055Aa5Ab99cBe8890bd22f13C34E675B3cAE8"
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
