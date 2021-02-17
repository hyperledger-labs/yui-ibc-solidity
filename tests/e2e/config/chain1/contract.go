package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x98Bb2d64238474552ABe3181E31C9CC4eCAeBeA3"
	IBCClientAddress     = "0x180B6C325525dB54C0DA871F1fd924a0bcf06397"
	IBCConnectionAddress = "0xDf919000A9A2533Fb000866654ce06D565DDDb97"
	IBCChannelAddress = "0xc0ba8346289ec43cd3f68E5EBf0a3169B1d14a2d"
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