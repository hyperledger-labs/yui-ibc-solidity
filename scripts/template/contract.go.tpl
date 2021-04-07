package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "<%= IBCHostAddress; %>"
	IBCHandlerAddress = "<%= IBCHandlerAddress; %>"
	IBCIdentifierAddress = "<%= IBCIdentifierAddress; %>"
	IBFT2ClientAddress = "<%= IBFT2ClientAddress; %>"
	MockClientAddress = "<%= MockClientAddress; %>"
	SimpleTokenAddress = "<%= SimpleTokenAddress; %>"
	ICS20TransferBankAddress = "<%= ICS20TransferBankAddress; %>"
	ICS20BankAddress = "<%= ICS20BankAddress; %>"
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
