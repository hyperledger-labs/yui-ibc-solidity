package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "<%= IBCStoreAddress; %>"
	IBCModuleAddress = "<%= IBCModuleAddress; %>"
	IBFT2ClientAddress = "<%= IBFT2ClientAddress; %>"
	SimpleTokenModuleAddress = "<%= SimpleTokenModuleAddress; %>"
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
