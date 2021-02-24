package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x21AAc107da35eC30f344d26324ad9b7eF30F566A"
	IBCClientAddress = "0x965b87E87165252c64bcD8dE31b61Dc7A135a8A0"
	IBCConnectionAddress = "0xe1aDE7B3a75e1D5D4675855239b3d7Ed2e7254a8"
	IBCChannelAddress = "0xC12d9eB07C4B8fa3469Fc37b2861cBFe9e390C19"
	IBCRoutingModuleAddress = "0xC5de6fc9374BA11cF47Cc1aECF94aBE620821BAa"
	IBFT2ClientAddress = "0xe95338C0CaCfdb17b7DabDd32b87d0EcD2B1A091"
	SimpleTokenModuleAddress = "0x31DEd0C2E93BB82760DCb11ddb6AfCF5045Dd371"
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

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
