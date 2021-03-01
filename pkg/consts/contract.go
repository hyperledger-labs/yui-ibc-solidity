package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x0D8c7066D8808f8D2118aB2594159C92D8043383"
	IBCHandlerAddress = "0x39fA07db5D99c9eAA9e04337F9F1f7386e41dEEB"
	IBCIdentifierAddress = "0xeB50cA91c99ceD8EDa25D362AdE560df9661Bc31"
	IBFT2ClientAddress = "0xa37a1a9Cc31e44adFb68Da558fc1F00f77983794"
	SimpleTokenModuleAddress = "0x161689B24999e61C470FafbA50C934Fb61179f4C"
	SimpleTokenAddress = "0x2D1deF28042b3c7931690dC59aEB1DD4a6Bed164"
	ICS20TransferAddress = "0xAb2056B46792159E075248525056BCd612e98670"
	ICS20VouchersAddress = "0xF93199cFa3E74Ffd321F62B67d4034Ad207cD927"
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

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddresss)
}

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20VouchersAddress() common.Address {
	return common.HexToAddress(ICS20VouchersAddress)
}
