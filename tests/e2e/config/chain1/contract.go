package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x8842CAF520C5eFF5f04CafDAb143CA26510A2C3f"
	IBCClientAddress     = "0x7E0807cdd3138F165d0a6c8bE093D7DA42C83899"
	IBCConnectionAddress = "0x9B05f07FC9EF14f3b18c9612eB7e31CF12FDa068"
	IBCChannelAddress = "0x2E94C9178569870655b5a12871a7FA9Aed8Bd5ef"
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