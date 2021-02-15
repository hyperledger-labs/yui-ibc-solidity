package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xF5bd99Ee3fc9eCD0fd7511218AB89161dE50cba5"
	IBCClientAddress     = "0x5a83E7Cf440d05A0e22c231465148D5b452D087F"
	IBCConnectionAddress = "0x433f4894bfB4EF6Fd8156393dE0361BfbeA8270A"
	IBCChannelAddress = "0xaFC143D6180b5E18198dE592B079698EAAd4126F"
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