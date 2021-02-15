package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xB998570D81c7cE63741fCBA8eaE8FF322c54D8Bb"
	IBCClientAddress     = "0xf92055Aa5Ab99cBe8890bd22f13C34E675B3cAE8"
	IBCConnectionAddress = "0x0bC8d5e2F4AbD6cfC617532972F57fdf0a6B719a"
	IBCChannelAddress = "0x6C6E7c6804ca9064cBEFA7cA90283b115D8Cf1C7"
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