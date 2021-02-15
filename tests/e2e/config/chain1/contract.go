package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xC799723F0Bc6142b06091F193b9D04374cdc5530"
	IBCClientAddress     = "0x94872f303a5c0f5afc20470A3De17198250A169b"
	IBCConnectionAddress = "0x57837cD742A4Ff8a3313587495F9965C0aBcC530"
	IBCChannelAddress = "0xCE76b1a5a7A6DC9cb5ddbE307c08c02bE1694638"
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