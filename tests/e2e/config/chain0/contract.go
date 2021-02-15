package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x5baB27B8d4e123fcE29B48B8b0307f57716a5a19"
	IBCClientAddress     = "0x0103FE92f6cE3d309b34BD2cBfb83f994402C867"
	IBCConnectionAddress = "0xEa34C547930d232E5Bf6D616Df9d1929F678f802"
	IBCChannelAddress = "0x93DdB80f7E23bFa1172B9930c45EcDEfbfFe0c5E"
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