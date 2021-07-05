package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x5a83E7Cf440d05A0e22c231465148D5b452D087F"
	IBCHandlerAddress = "0x433f4894bfB4EF6Fd8156393dE0361BfbeA8270A"
	IBCIdentifierAddress = "0x7E0807cdd3138F165d0a6c8bE093D7DA42C83899"
	IBFT2ClientAddress = "0xA59f1b02B85212bBCBEB8A060aa4C595ADDFA25e"
	MockClientAddress = "0xF5bd99Ee3fc9eCD0fd7511218AB89161dE50cba5"
	SimpleTokenAddress = "0xaFC143D6180b5E18198dE592B079698EAAd4126F"
	ICS20TransferBankAddress = "0x810dbcb6Bf103d499E96c1af0c98536EE3243e62"
	ICS20BankAddress = "0x6889D1E8f2269Ee96D08Fc5dF4295da21297d525"
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
