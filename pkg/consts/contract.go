package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress = "0xc324fBF6512022865d96510bfe761c7c776ccb76"
	IBCCommitmentAddress = "0x250038C44651F7984aA7a077b88Be96b547EA4Dd"
	IBFT2ClientAddress = "0xF7028c7261949cd2a66790723d2ebA5Df3340da0"
	MockClientAddress = "0x9af1df12689Ff2886FB8e18E2e997113aBC47940"
	SimpleTokenAddress = "0x9Fe75f5d53F95eAB37FBbC48e831614dD0C88fFf"
	ICS20TransferBankAddress = "0x49F301187662eAbA6c635C5AD9Abf27D24902A12"
	ICS20BankAddress = "0xc92c375303B5FBf30bb1bee82BEDAfE1dc6B1671"
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
