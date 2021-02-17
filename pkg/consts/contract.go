package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x2F5703804E29F4252FA9405B8D357220d11b3bd9"
	IBCClientAddress     = "0xaE1C9125BbcF63bf51294C4D15CBD472782E330D"
	IBCConnectionAddress = "0xa7f733a4fEA1071f58114b203F57444969b86524"
	IBCChannelAddress = "0x87d7778dbc81251D5A0D78DFD8a0C359887E98C9"
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