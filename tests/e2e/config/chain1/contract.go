package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x6b6B1d020Ae984F5104768EB3d62c2F304593D81"
	IBCClientAddress     = "0xe5e2F18E98Dced653e9b5bF59ecC6A6687ceb373"
	IBCConnectionAddress = "0x59679b1385248F33042e76fC538abc32ab286DF8"
	IBCChannelAddress = "0x5baB27B8d4e123fcE29B48B8b0307f57716a5a19"
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