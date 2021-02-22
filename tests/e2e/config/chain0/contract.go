package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xDAE6c6F7fB263e58780604A20cE0d562c6a52390"
	IBCModuleAddress = "0xfe163a0F0bD6F88d920759a2f1C795eEC8B8cbe6"
	IBFT2ClientAddress = "0xaE0aB73acE678dDE54c3E8ace0Afe48B6935668E"
	SimpleTokenModuleAddress = "0x7E13b07243771f02EFC1BF0BAA17f1Fa7760E998"
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
