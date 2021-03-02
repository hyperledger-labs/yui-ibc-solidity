package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x173290F876E8DBe940a5579dA25aa03F36df977E"
	IBCHandlerAddress = "0x0368fa510dB2e4e9b8B4e344f2daA5c3251a4958"
	IBCIdentifierAddress = "0x2D1deF28042b3c7931690dC59aEB1DD4a6Bed164"
	IBFT2ClientAddress = "0xd0A4210EeFb4cF2748FdBD4D52720725EBcD0Eb1"
	SimpleTokenAddress = "0xA6afB05A9dA3d6b13f7C9B4Fb7658D8afd48481b"
	ICS20TransferBankAddress = "0xF16Fdb1FF23359633cCe37b7554394E8beA262D4"
	ICS20BankAddress = "0x69112AC6cBe42103F63f3001114aF647E56EB98d"
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

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferBankAddress() common.Address {
	return common.HexToAddress(ICS20TransferBankAddress)
}

func (contractConfig) GetICS20BankAddress() common.Address {
	return common.HexToAddress(ICS20BankAddress)
}
