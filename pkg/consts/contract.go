package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x173FFadd93aefC4F6e2D9bB20825E68A736DE047"
	IBCClientAddress     = "0x962C7179caFBC2F526cE678FcD607078A548474a"
	IBCConnectionAddress = "0x3379866A055CF3725775e8d6876be4877e1a7D89"
	IBCChannelAddress = "0xf68230F013AbBc6c2330cdE2Fa2dF84e2417B4a0"
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