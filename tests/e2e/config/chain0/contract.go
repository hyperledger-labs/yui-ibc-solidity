package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xF16Fdb1FF23359633cCe37b7554394E8beA262D4"
	IBCHandlerAddress = "0xff67836F5cb28030F6B8bDC32736F69e2e91d3F2"
	IBCIdentifierAddress = "0xeeB18f5AC5bb97Fc68B6d3397e22564B2a38B78b"
	IBFT2ClientAddress = "0x69112AC6cBe42103F63f3001114aF647E56EB98d"
	SimpleTokenAddress = "0xFEAB95Eeb8507978bC5edD22E9BA2F52f9d377A1"
	ICS20TransferBankAddress = "0x36C7086273a419f07B061a7E4184384270F5ec97"
	ICS20BankAddress = "0x6fdA347f2A64fd55F43B63c28619548c9B362835"
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
