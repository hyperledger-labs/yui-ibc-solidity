package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x20a169B3a59E0830B070B7C5BDb4965baF2582F5"
	IBCHandlerAddress = "0x417AF1C78f40220bCD9dc48cCF6d2aA12c573665"
	IBCIdentifierAddress = "0x18Bf5B933e16774c5183522e27234F9e4EB21Ed5"
	IBFT2ClientAddress = "0x7F87D1ac88Ba30b9270c50a37d7a88BC7679df89"
	SimpleTokenModuleAddress = "0x0BFA960015B86C238b2b4A9192fd19C080991003"
	SimpleTokenAddress = "0xd6C79A444BbCD1c44e83c07da2365184616115dF"
	ICS20TransferAddress = "0xEB7eA35a582eAC320e9C215F404CB19a5EC3BF25"
	ICS20VouchersAddress = "0x9dCb7FeCA4D09699f6B9df6F14b231b902dE0651"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20VouchersAddress() common.Address {
	return common.HexToAddress(ICS20VouchersAddress)
}
