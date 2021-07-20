package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xeB50cA91c99ceD8EDa25D362AdE560df9661Bc31"
	IBCHandlerAddress = "0x06e5dEB55CAffb1339fb447d860E305249EaD495"
	IBCIdentifierAddress = "0x0E47d25A069d0f44ffE273ca489606E9C5F64240"
	IBFT2ClientAddress = "0xBF346b5BC386c7C3378688286406B08E9327d312"
	MockClientAddress = "0x4DB8e6C8BdE4c9AFCEDb590C5446c965c073BED8"
	SimpleTokenAddress = "0xF938fE7482Fe4d1b3f84E28F1D6407836AA27d99"
	ICS20TransferBankAddress = "0xF251fB1Ca5445777Fed7bb760eB3c49636BA8DC3"
	ICS20BankAddress = "0x72f25b3D42e279917b4bd9284c22b99cBe521076"
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
