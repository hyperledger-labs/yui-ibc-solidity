package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "<%= ProvableStoreAddress; %>"
	IBCClientAddress     = "<%= IBCClientAddress; %>"
	IBCConnectionAddress = "<%= IBCConnectionAddress; %>"
	IBCChannelAddress = "<%= IBCChannelAddress; %>"
	IBCRoutingModuleAddress = "<%= IBCRoutingModuleAddress; %>"
	SimpleTokenModuleAddress = "<%= SimpleTokenModuleAddress; %>"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetProvableStoreAddress() common.Address {
	return common.HexToAddress(ProvableStoreAddress)
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
