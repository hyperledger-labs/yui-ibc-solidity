package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x964E727Fa15f3e7d80fBc7F1599F952EA8e6033d"
	IBCClientAddress     = "0xBB927aEf683f0B1Fd2Ab651bEC6755FAc2aa640f"
	IBCConnectionAddress = "0x20889941d4ed2D08AEd24B6470eE0b38a06EC89C"
	IBCChannelAddress = "0xeeC744224f08b37E026458e4C7603577741CC93A"
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