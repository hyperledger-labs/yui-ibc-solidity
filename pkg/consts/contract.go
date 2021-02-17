package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x8E3855F48B48F98ef36AD06B4767F946d34CDCBc"
	IBCClientAddress     = "0xe1C1a0FB7fee236A305cB1B3AB2676561d07F0D2"
	IBCConnectionAddress = "0x74569C9156ECd7897587e39111F98234177c2e67"
	IBCChannelAddress = "0x3e95052495dA4d0bF002C7c5E4bb9698cB349075"
	IBCRoutingModuleAddress = "0xe72bc97d3232Fe23f7E0bbC76f93E2a2465ff34C"
	SimpleTokenModuleAddress = "0x5adcFE2e19a63f6D389D76DB784afE130F9cF012"
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
