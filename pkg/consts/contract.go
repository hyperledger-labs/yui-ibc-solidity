package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xBB927aEf683f0B1Fd2Ab651bEC6755FAc2aa640f"
	IBCHandlerAddress = "0x20889941d4ed2D08AEd24B6470eE0b38a06EC89C"
	IBCIdentifierAddress = "0x69112AC6cBe42103F63f3001114aF647E56EB98d"
	IBFT2ClientAddress = "0x36C7086273a419f07B061a7E4184384270F5ec97"
	MockClientAddress = "0x964E727Fa15f3e7d80fBc7F1599F952EA8e6033d"
	SimpleTokenAddress = "0xeeC744224f08b37E026458e4C7603577741CC93A"
	ICS20TransferBankAddress = "0x30BaBba8d5E322639834459991754036Ae813FE3"
	ICS20BankAddress = "0xE6ea3b07bB7a205CF9fBf219B7C37dE74Ee81490"
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

func (contractConfig) GetMockClientAddress() common.Address {
	return common.HexToAddress(MockClientAddress)
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
