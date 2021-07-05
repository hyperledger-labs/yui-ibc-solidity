package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x2706E591ca1f86489FcB1C4b6A8902a09B083201"
	IBCHandlerAddress = "0xB9D736791E1e0CaE284619bb6cf061A40daEeF5F"
	IBCIdentifierAddress = "0x4B80782C5616A043fAd46b2aEe6B39a86F3EdbA3"
	IBFT2ClientAddress = "0x91bff9c82426095905f242C4Fa4568448Aa94187"
	MockClientAddress = "0x7058695E7028aFdC28EFEBf3c82dC5EA82C35aC1"
	SimpleTokenAddress = "0x81f89365D89d9fF4Af8940c892612BDc72c8eb3F"
	ICS20TransferBankAddress = "0xe8e87341bCD77BC3f1D1F1De1FB5442F54077DDF"
	ICS20BankAddress = "0x151d23662D86034726c34c4139E302bB3235912a"
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
