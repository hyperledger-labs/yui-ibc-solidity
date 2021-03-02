package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xf8cC116326b1Fa45DF6Ba57c16C4f4eF4E17A502"
	IBCHandlerAddress = "0xDAE6c6F7fB263e58780604A20cE0d562c6a52390"
	IBCIdentifierAddress = "0xed823B51718C2fE62346ed318243FEDd2209062f"
	IBFT2ClientAddress = "0xF3943bd0ae0e272b1eaB70008429e704F9D4c451"
	SimpleTokenAddress = "0xaE0aB73acE678dDE54c3E8ace0Afe48B6935668E"
	ICS20TransferAddress = "0x7E13b07243771f02EFC1BF0BAA17f1Fa7760E998"
	ICS20BankAddress = "0xfe163a0F0bD6F88d920759a2f1C795eEC8B8cbe6"
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

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20BankAddress() common.Address {
	return common.HexToAddress(ICS20BankAddress)
}
