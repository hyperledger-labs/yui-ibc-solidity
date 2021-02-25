package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x31D3660a7325bC7Ca9a65D55aB7beF3062219473"
	IBCClientAddress = "0xB841d52E647078d3B92Dc0433830B2ca00D49A1E"
	IBCConnectionAddress = "0x86d157c02b9cbEC7b72A5Ae53518530DE40Eff2F"
	IBCChannelAddress = "0x176A50aE5b2D75a65D7988CB0bE187e6d1D54E79"
	IBCRoutingModuleAddress = "0xe8B67B46Dd52592cDF1C6A55FFcC9e9CEAD22D30"
	IBCIdentifierAddress = "0xEB7eA35a582eAC320e9C215F404CB19a5EC3BF25"
	IBFT2ClientAddress = "0xa7657bC9A370FBb70d985e340CFd4A36Ac82b922"
	SimpleTokenModuleAddress = "0xE07e3EeFA7377c935Fd30edccf0cd325e210d0B6"
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
