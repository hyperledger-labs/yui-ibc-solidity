package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0xF93199cFa3E74Ffd321F62B67d4034Ad207cD927"
	IBCClientAddress     = "0xAb2056B46792159E075248525056BCd612e98670"
	IBCConnectionAddress = "0xbabc628B9D14Ace99D3De6DBad1A3C28d832090d"
	IBCChannelAddress = "0xeeB18f5AC5bb97Fc68B6d3397e22564B2a38B78b"
	IBCRoutingModuleAddress = "0xd0A4210EeFb4cF2748FdBD4D52720725EBcD0Eb1"
	SimpleTokenModuleAddress = "0x173290F876E8DBe940a5579dA25aa03F36df977E"
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

func (contractConfig) GetIBCRoutingModuleAddress() common.Address {
	return common.HexToAddress(IBCRoutingModuleAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
