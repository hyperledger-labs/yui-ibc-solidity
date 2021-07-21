package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x747296FC9d600e4Ce2156dE3aeE8aa75bf2E459a"
	IBCHandlerAddress = "0xA59f1b02B85212bBCBEB8A060aa4C595ADDFA25e"
	IBCIdentifierAddress = "0xfe163a0F0bD6F88d920759a2f1C795eEC8B8cbe6"
	IBFT2ClientAddress = "0x2E94C9178569870655b5a12871a7FA9Aed8Bd5ef"
	MockClientAddress = "0xA05d3D10aB5aB40f5e7751411a4ff975e6Ecc97e"
	SimpleTokenAddress = "0xF5bd99Ee3fc9eCD0fd7511218AB89161dE50cba5"
	ICS20TransferBankAddress = "0x433f4894bfB4EF6Fd8156393dE0361BfbeA8270A"
	ICS20BankAddress = "0x5a83E7Cf440d05A0e22c231465148D5b452D087F"
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
