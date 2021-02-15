package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x3379866A055CF3725775e8d6876be4877e1a7D89"
	IBCClientAddress     = "0xf68230F013AbBc6c2330cdE2Fa2dF84e2417B4a0"
	IBCConnectionAddress = "0xa954B862442936E35B579EB475CE82768389D301"
	IBCChannelAddress = "0x4efa02d32DDDf9856222331ABEFf1a09248FeB87"
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