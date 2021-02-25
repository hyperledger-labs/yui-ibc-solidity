package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x344B67A08484b437C8FeD8880E60DaAF5Fe06cb7"
	IBCClientAddress = "0x69269150ae9D0a7798687008285a34f25EEb5cb9"
	IBCConnectionAddress = "0xC96b868921B769F789593e2260b075213eCc9D3e"
	IBCChannelAddress = "0x1D268db12EB341E196C682ebB94314Af2cAC0d8d"
	IBCRoutingModuleAddress = "0x21299a307a091ab0724BAA70Ab853d449eD895C5"
	IBCIdentifierAddress = "0x395544BB5C2D2d62bE58A7d8659b590CFa5A06ab"
	IBFT2ClientAddress = "0xF0B0b60523062B744618c98b47DCa9Bf95247746"
	SimpleTokenModuleAddress = "0x61a8eCb0023d93595f9DE4715abd6b7e57212b06"
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
