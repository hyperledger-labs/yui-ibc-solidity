package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x906083bF79339343Ba9Da6Ee6FD4c446720Daf0C"
	IBCClientAddress     = "0x98Bb2d64238474552ABe3181E31C9CC4eCAeBeA3"
	IBCConnectionAddress = "0x180B6C325525dB54C0DA871F1fd924a0bcf06397"
	IBCChannelAddress = "0xDf919000A9A2533Fb000866654ce06D565DDDb97"
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