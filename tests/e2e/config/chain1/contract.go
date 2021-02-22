package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xF3943bd0ae0e272b1eaB70008429e704F9D4c451"
	IBCModuleAddress = "0xDAE6c6F7fB263e58780604A20cE0d562c6a52390"
	IBFT2ClientAddress = "0xf8cC116326b1Fa45DF6Ba57c16C4f4eF4E17A502"
	SimpleTokenModuleAddress = "0xaE0aB73acE678dDE54c3E8ace0Afe48B6935668E"
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
