package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x9B05f07FC9EF14f3b18c9612eB7e31CF12FDa068"
	IBCHandlerAddress = "0x2E94C9178569870655b5a12871a7FA9Aed8Bd5ef"
	IBCIdentifierAddress = "0xf8cC116326b1Fa45DF6Ba57c16C4f4eF4E17A502"
	IBFT2ClientAddress = "0x8842CAF520C5eFF5f04CafDAb143CA26510A2C3f"
	MockClientAddress = "0x7E0807cdd3138F165d0a6c8bE093D7DA42C83899"
	SimpleTokenAddress = "0xA05d3D10aB5aB40f5e7751411a4ff975e6Ecc97e"
	ICS20TransferBankAddress = "0xA59f1b02B85212bBCBEB8A060aa4C595ADDFA25e"
	ICS20BankAddress = "0x747296FC9d600e4Ce2156dE3aeE8aa75bf2E459a"
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
