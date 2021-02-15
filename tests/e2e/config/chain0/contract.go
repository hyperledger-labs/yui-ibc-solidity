package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x9B0b84b8FbB7B0e03Cfd78c6AF60CB9F3EA5d8Fc"
	IBCClientAddress     = "0xD622e327b04fF3884E0730752840C67556eD326b"
	IBCConnectionAddress = "0x6677E99A3d6e90918450230B2115c955a0DFa9aA"
	IBCChannelAddress = "0xEb25Ae0bFd59BC57eefbCd86d8Ca89853ce5139A"
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