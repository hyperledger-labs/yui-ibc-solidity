package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xa37a1a9Cc31e44adFb68Da558fc1F00f77983794"
	IBCClientAddress     = "0x0D8c7066D8808f8D2118aB2594159C92D8043383"
	IBCConnectionAddress = "0x39fA07db5D99c9eAA9e04337F9F1f7386e41dEEB"
	IBCChannelAddress = "0x161689B24999e61C470FafbA50C934Fb61179f4C"
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