package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x2f15f8Ea36ebE0cA8542F5Da224C14DD26aaa095"
	IBCHandlerAddress = "0x488E307620F65927C3188FB85aA2C8f20562Bc55"
	IBCIdentifierAddress = "0x5120397d4E66c2b051a4ec3DB2ef31953Ac3F440"
	IBFT2ClientAddress = "0x2708C5A03Ca7CAA78214C79D2bD40Cc7421ED917"
	MockClientAddress = "0xc664Bbd720B2A7a79f7c8500FB041C7b003FCA5a"
	SimpleTokenAddress = "0x39a1b3B1a6B207d77F2ffa68F005aB8fE0688994"
	ICS20TransferBankAddress = "0x6DF80f703c32D150125eaA876Ec1D0c0Ebc3Bbf2"
	ICS20BankAddress = "0xC71b6B9dD5763E4a9eBF1a7A5330C6a6E94C788c"
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
