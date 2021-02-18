package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x395544BB5C2D2d62bE58A7d8659b590CFa5A06ab"
	IBCClientAddress     = "0xB26e810C030c56eeEd2c5231115597f1E210FC3F"
	IBCConnectionAddress = "0x344B67A08484b437C8FeD8880E60DaAF5Fe06cb7"
	IBCChannelAddress = "0xF0B0b60523062B744618c98b47DCa9Bf95247746"
	IBCHandlerAddress = "0x69269150ae9D0a7798687008285a34f25EEb5cb9"
	SimpleTokenModuleAddress = "0xC96b868921B769F789593e2260b075213eCc9D3e"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
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

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
