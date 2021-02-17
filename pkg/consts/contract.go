package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x76C523Cf88dE127169a872720Acdc233F46ec88F"
	IBCClientAddress     = "0x905A69d16226eC10f3Dd4939A904629360a7Ec74"
	IBCConnectionAddress = "0xf0C3c6B56CBe448F2801C20aB41445554efeB103"
	IBCChannelAddress = "0xaA5960eFA62bf00f5505470955A8E35C6DD99E6C"
	IBCRoutingModuleAddress = "0x3d534EA50c443164CA710Dcea0D75cb68E4A2E39"
	SimpleTokenModuleAddress = "0x4Fb426e0DF93CA7A3957C39E5C27B23Fb6210Ec6"
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
