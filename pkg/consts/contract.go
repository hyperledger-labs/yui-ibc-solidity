package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
	IBCClientAddress     = "0x37978908bac82F0191b674235A0fEEE31e7524a4"
	IBCConnectionAddress = "0x173FFadd93aefC4F6e2D9bB20825E68A736DE047"
	IBCChannelAddress = "0x962C7179caFBC2F526cE678FcD607078A548474a"
	IBCHandlerAddress = "0x3379866A055CF3725775e8d6876be4877e1a7D89"
	SimpleTokenModuleAddress = "0xf68230F013AbBc6c2330cdE2Fa2dF84e2417B4a0"
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

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
