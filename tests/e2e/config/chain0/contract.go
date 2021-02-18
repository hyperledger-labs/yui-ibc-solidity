package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x98Bb2d64238474552ABe3181E31C9CC4eCAeBeA3"
	IBCClientAddress     = "0x180B6C325525dB54C0DA871F1fd924a0bcf06397"
	IBCConnectionAddress = "0xDf919000A9A2533Fb000866654ce06D565DDDb97"
	IBCChannelAddress = "0xc0ba8346289ec43cd3f68E5EBf0a3169B1d14a2d"
	IBCHandlerAddress = "0xA2f3403490466E33dcF0d74cAfc1DE0BeE0f47B4"
	SimpleTokenModuleAddress = "0x5987561e4396FC977AceFdB8DC2745305c53543a"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
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
