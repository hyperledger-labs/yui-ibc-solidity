package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x86d157c02b9cbEC7b72A5Ae53518530DE40Eff2F"
	IBCClientAddress = "0xe8B67B46Dd52592cDF1C6A55FFcC9e9CEAD22D30"
	IBCConnectionAddress = "0xE07e3EeFA7377c935Fd30edccf0cd325e210d0B6"
	IBCChannelAddress = "0xbCB810E2D5c105B401D9Db678652Fa8DaE21D13A"
	IBCRoutingModuleAddress = "0x4A9cB3b03AA5eEe055F32dd107c713A6013c08e4"
	IBCIdentifierAddress = "0xa7657bC9A370FBb70d985e340CFd4A36Ac82b922"
	IBFT2ClientAddress = "0x176A50aE5b2D75a65D7988CB0bE187e6d1D54E79"
	SimpleTokenModuleAddress = "0x4096D169bBf1928e6BAFE85F8ae36A381FCC7019"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
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

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
