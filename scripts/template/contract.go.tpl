package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress = "<%= IBCHandlerAddress; %>"
	IBCCommitmentTestHelperAddress = "<%= IBCCommitmentTestHelperAddress; %>"
	SimpleTokenAddress = "<%= SimpleTokenAddress; %>"
	ICS20TransferBankAddress = "<%= ICS20TransferBankAddress; %>"
	ICS20BankAddress = "<%= ICS20BankAddress; %>"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCCommitmentTestHelperAddress() common.Address {
	return common.HexToAddress(IBCCommitmentTestHelperAddress)
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
