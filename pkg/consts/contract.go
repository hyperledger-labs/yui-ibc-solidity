package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
	IBCClientAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	IBCConnectionAddress = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
	IBCChannelAddress = "0xa7f733a4fEA1071f58114b203F57444969b86524"
	IBCRoutingModuleAddress = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
	IBFT2ClientAddress = "0xff77D90D6aA12db33d3Ba50A34fB25401f6e4c4F"
	SimpleTokenModuleAddress = "0x37978908bac82F0191b674235A0fEEE31e7524a4"
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

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
