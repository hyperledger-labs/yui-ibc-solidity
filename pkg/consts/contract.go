package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x74569C9156ECd7897587e39111F98234177c2e67"
	IBCClientAddress = "0xe72bc97d3232Fe23f7E0bbC76f93E2a2465ff34C"
	IBCConnectionAddress = "0x5adcFE2e19a63f6D389D76DB784afE130F9cF012"
	IBCChannelAddress = "0xb88b5D8A95653f41BdF56cD8dD5271A557Db6117"
	IBCRoutingModuleAddress = "0xe3761c2316048503deEe5F1c9d7ecd3AE1FA166B"
	IBCIdentifierAddress = "0x8E3855F48B48F98ef36AD06B4767F946d34CDCBc"
	IBFT2ClientAddress = "0x3e95052495dA4d0bF002C7c5E4bb9698cB349075"
	SimpleTokenModuleAddress = "0x7B3440bA1Bb28121b5F482c4cD020Be325b19e7b"
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
