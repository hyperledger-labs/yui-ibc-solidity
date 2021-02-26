package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x9eBF3956EE45B2b9F1fC85FB8990ce6be52F47a6"
	IBCClientAddress = "0x702E40245797c5a2108A566b3CE2Bf14Bc6aF841"
	IBCConnectionAddress = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
	IBCChannelAddress = "0xff77D90D6aA12db33d3Ba50A34fB25401f6e4c4F"
	IBCRoutingModuleAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	IBCIdentifierAddress = "0xdD5109D05Ac357E446992a60E64764041A0E8529"
	IBFT2ClientAddress = "0x727A5648832D2b317925CE043eA9b7fE04B4CD55"
	SimpleTokenModuleAddress = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
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
