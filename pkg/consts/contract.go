package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xF251fB1Ca5445777Fed7bb760eB3c49636BA8DC3"
	IBCHandlerAddress = "0xa37a1a9Cc31e44adFb68Da558fc1F00f77983794"
	IBCIdentifierAddress = "0xBF346b5BC386c7C3378688286406B08E9327d312"
	IBFT2ClientAddress = "0x72f25b3D42e279917b4bd9284c22b99cBe521076"
	SimpleTokenAddress = "0x0D8c7066D8808f8D2118aB2594159C92D8043383"
	ICS20TransferBankAddress = "0x161689B24999e61C470FafbA50C934Fb61179f4C"
	ICS20BankAddress = "0x39fA07db5D99c9eAA9e04337F9F1f7386e41dEEB"
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
