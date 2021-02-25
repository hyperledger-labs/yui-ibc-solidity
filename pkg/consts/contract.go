package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xFE202b2D0483634FFb93b6475aE3E9936D4b193d"
	IBCClientAddress = "0x395544BB5C2D2d62bE58A7d8659b590CFa5A06ab"
	IBCConnectionAddress = "0xB26e810C030c56eeEd2c5231115597f1E210FC3F"
	IBCChannelAddress = "0x344B67A08484b437C8FeD8880E60DaAF5Fe06cb7"
	IBCRoutingModuleAddress = "0xF0B0b60523062B744618c98b47DCa9Bf95247746"
	IBCIdentifierAddress = "0xb824662d7d200c5a5e5a0fC4bBc056eDC779a4C6"
	IBFT2ClientAddress = "0x5b955f58bDe2F828239117f2210B4cb233FE3603"
	SimpleTokenModuleAddress = "0x69269150ae9D0a7798687008285a34f25EEb5cb9"
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
