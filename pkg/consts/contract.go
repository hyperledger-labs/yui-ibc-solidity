package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress = "0x129559FC0D7935a1806a9a192C5e550fDa7C0a06"
	IBCCommitmentAddress = "0x17c376FA3E355628bac0bCA410502A780A8e5930"
	IBFT2ClientAddress = "0x11bB8FbeFBA127838f1e3b3c08e8e594E3259c50"
	MockClientAddress = "0xBCf9C1CBd4a841DaEf9Eb43F57F6b239fE57CFc2"
	SimpleTokenAddress = "0x2B4e63574Ed6096BEb884bd2582Ee4eCFFB32E5b"
	ICS20TransferBankAddress = "0x65fc1A735A1Fe17b3da6a537e47708aF0e35B0A5"
	ICS20BankAddress = "0xC9A3729a922742acf62D50FAd158aEA940473B55"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCCommitmentAddress() common.Address {
	return common.HexToAddress(IBCCommitmentAddress)
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
