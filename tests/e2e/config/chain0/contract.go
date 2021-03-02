package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x747296FC9d600e4Ce2156dE3aeE8aa75bf2E459a"
	IBCHandlerAddress = "0xA59f1b02B85212bBCBEB8A060aa4C595ADDFA25e"
	IBCIdentifierAddress = "0x7E13b07243771f02EFC1BF0BAA17f1Fa7760E998"
	IBFT2ClientAddress = "0xA05d3D10aB5aB40f5e7751411a4ff975e6Ecc97e"
	SimpleTokenAddress = "0xF5bd99Ee3fc9eCD0fd7511218AB89161dE50cba5"
	ICS20TransferAddress = "0x433f4894bfB4EF6Fd8156393dE0361BfbeA8270A"
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

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20BankAddress() common.Address {
	return common.HexToAddress(ICS20BankAddress)
}
