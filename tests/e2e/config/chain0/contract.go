package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xB841d52E647078d3B92Dc0433830B2ca00D49A1E"
	IBCClientAddress     = "0x86d157c02b9cbEC7b72A5Ae53518530DE40Eff2F"
	IBCConnectionAddress = "0x176A50aE5b2D75a65D7988CB0bE187e6d1D54E79"
	IBCChannelAddress = "0xe8B67B46Dd52592cDF1C6A55FFcC9e9CEAD22D30"
	IBCHandlerAddress = "0xE07e3EeFA7377c935Fd30edccf0cd325e210d0B6"
	SimpleTokenModuleAddress = "0xbCB810E2D5c105B401D9Db678652Fa8DaE21D13A"
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
