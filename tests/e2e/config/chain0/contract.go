package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x8b7e2A0d19048D4D6B710628b840b30d99bCD022"
	IBCClientAddress = "0x5a13B112eFd5d69b860c49032e4208fe8408D602"
	IBCConnectionAddress = "0xA629b31f3F0C8B48E395014b6283b7B616dfdb8D"
	IBCChannelAddress = "0x029EAfA3eED79299F635E2A6F1607875FF7fA983"
	IBCRoutingModuleAddress = "0x56CF24731e6DaAC0499C52bdd61Fae800bB7A173"
	IBCIdentifierAddress = "0xF8352321694Ed0806aD548c90Ee528C757F6B829"
	IBFT2ClientAddress = "0x3d906FFc0D2389A5162058897cDd54C96F07c496"
	SimpleTokenModuleAddress = "0xfed5D6a133522d9616C3F5e4916B27E27F026013"
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
