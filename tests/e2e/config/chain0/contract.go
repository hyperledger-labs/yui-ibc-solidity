package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x5b955f58bDe2F828239117f2210B4cb233FE3603"
	IBCClientAddress = "0xB26e810C030c56eeEd2c5231115597f1E210FC3F"
	IBCConnectionAddress = "0x344B67A08484b437C8FeD8880E60DaAF5Fe06cb7"
	IBCChannelAddress = "0xF0B0b60523062B744618c98b47DCa9Bf95247746"
	IBCRoutingModuleAddress = "0x69269150ae9D0a7798687008285a34f25EEb5cb9"
	IBCIdentifierAddress = "0xA9553A168a69D18a744b711ED59Cebbb3ABa491f"
	IBFT2ClientAddress = "0x395544BB5C2D2d62bE58A7d8659b590CFa5A06ab"
	SimpleTokenModuleAddress = "0xC96b868921B769F789593e2260b075213eCc9D3e"
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
