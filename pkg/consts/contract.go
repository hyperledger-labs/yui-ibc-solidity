package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xbabc628B9D14Ace99D3De6DBad1A3C28d832090d"
	IBCClientAddress = "0xd0A4210EeFb4cF2748FdBD4D52720725EBcD0Eb1"
	IBCConnectionAddress = "0x173290F876E8DBe940a5579dA25aa03F36df977E"
	IBCChannelAddress = "0x0368fa510dB2e4e9b8B4e344f2daA5c3251a4958"
	IBCRoutingModuleAddress = "0xA6afB05A9dA3d6b13f7C9B4Fb7658D8afd48481b"
	IBFT2ClientAddress = "0xeeB18f5AC5bb97Fc68B6d3397e22564B2a38B78b"
	SimpleTokenModuleAddress = "0x69112AC6cBe42103F63f3001114aF647E56EB98d"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
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
