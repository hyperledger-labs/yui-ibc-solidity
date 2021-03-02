package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x7b1907056D9b82424f54a42707602e014106Ba62"
	IBCHandlerAddress = "0xC2D142Fda35551b0df9a52Cb0e5340B8Bec89166"
	IBCIdentifierAddress = "0x173FFadd93aefC4F6e2D9bB20825E68A736DE047"
	IBFT2ClientAddress = "0x4efa02d32DDDf9856222331ABEFf1a09248FeB87"
	SimpleTokenAddress = "0x0E47d25A069d0f44ffE273ca489606E9C5F64240"
	ICS20TransferBankAddress = "0xB9F2435339644FDF60E35bF9d02F6C190a7a1930"
	ICS20BankAddress = "0x04437cDC1d0126dEB0e762e037dcEF3d5510A1f3"
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
