package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xdD5109D05Ac357E446992a60E64764041A0E8529"
	IBCClientAddress     = "0x361552A65C96621003C62C5971b910a1fdC9ba78"
	IBCConnectionAddress = "0x9eBF3956EE45B2b9F1fC85FB8990ce6be52F47a6"
	IBCChannelAddress = "0x727A5648832D2b317925CE043eA9b7fE04B4CD55"
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