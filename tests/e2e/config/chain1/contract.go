package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x5A8bEdF0927cb1142066fe7a544E54b962a57FE3"
	IBCClientAddress = "0xDeAE9d6a12b78333BF9ccd6A9fb4F12BdAf27176"
	IBCConnectionAddress = "0x610BDC77AD2e2c23B404da5cA630b48aC3C009f2"
	IBCChannelAddress = "0x27167E4d39A77237357c46603C860A733Dc40159"
	IBCRoutingModuleAddress = "0x17c7479d986ce5932214af17e6e5f7f0573705DC"
	IBCIdentifierAddress = "0x4C47fa7aA33E87FB5c8A3cdb4d87701b3Dbb3459"
	IBFT2ClientAddress = "0x1e67ba8FFf8D50939114Af5051F03dFdaDa36005"
	SimpleTokenModuleAddress = "0x5e39226ba60a8bACb284b5A736a2Ca302e171E55"
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

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
