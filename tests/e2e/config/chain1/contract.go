package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x2D1deF28042b3c7931690dC59aEB1DD4a6Bed164"
	IBCClientAddress     = "0xF93199cFa3E74Ffd321F62B67d4034Ad207cD927"
	IBCConnectionAddress = "0xAb2056B46792159E075248525056BCd612e98670"
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
