package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x8F8aa2739608FB9136707F61229344883A226B24"
	IBCHandlerAddress = "0x56e7298A90022319950e6e3778111f46778a1F83"
	IBCIdentifierAddress = "0xB0378a47A83D1c2c4191cc43570Fc41bef329800"
	IBFT2ClientAddress = "0x6F9cDF68192861f1aDDbee6D27C44eC50ED34bF7"
	MockClientAddress = "0x2183e5Ee77A8014a6203D9BAA05522786788a247"
	SimpleTokenAddress = "0x2EC41e1b61c85839E19CB9d20fc5869f736FEFF0"
	ICS20TransferBankAddress = "0x564C69AD8Fa885B524D9518AaE2319fe93910382"
	ICS20BankAddress = "0x954Fd0093088d7863588b518F8A4FBe6949fFB8c"
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
