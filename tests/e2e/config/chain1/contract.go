package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xF938fE7482Fe4d1b3f84E28F1D6407836AA27d99"
	IBCClientAddress     = "0x72f25b3D42e279917b4bd9284c22b99cBe521076"
	IBCConnectionAddress = "0xF251fB1Ca5445777Fed7bb760eB3c49636BA8DC3"
	IBCChannelAddress = "0xa37a1a9Cc31e44adFb68Da558fc1F00f77983794"
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