package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x361552A65C96621003C62C5971b910a1fdC9ba78"
	IBCModuleAddress = "0x727A5648832D2b317925CE043eA9b7fE04B4CD55"
	IBFT2ClientAddress = "0x9eBF3956EE45B2b9F1fC85FB8990ce6be52F47a6"
	SimpleTokenModuleAddress = "0x702E40245797c5a2108A566b3CE2Bf14Bc6aF841"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBCModuleAddress() common.Address {
	return common.HexToAddress(IBCModuleAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
