package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress = "0x57f8dD374BBe73724a123b6834f8D6e1aD5efC78"
	IBCCommitmentAddress = "0xe1B4773F25Fb5109E04CE623c71d4693777cf908"
	IBFT2ClientAddress = "0x436d8B4cbC0a27d0dc391874ade6C43feaB7D9b4"
	MockClientAddress = "0x38b1a3755e2b674cA7141c40DA97bD3A25831560"
	SimpleTokenAddress = "0xDcf93366087CB3ab5b3916da4636Af72a2551C9e"
	ICS20TransferBankAddress = "0xcf4CDA33F6b6FE27DD725182FEC834C4C583098D"
	ICS20BankAddress = "0x96dAA2633565C9B29f3Bb727aE7ce2BffC8A438C"
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
