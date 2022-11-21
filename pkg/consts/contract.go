package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHandlerAddress = "0x0368fa510dB2e4e9b8B4e344f2daA5c3251a4958"
	IBCCommitmentAddress = "0xAb2056B46792159E075248525056BCd612e98670"
	IBFT2ClientAddress = "0x69112AC6cBe42103F63f3001114aF647E56EB98d"
	MockClientAddress = "0xA6afB05A9dA3d6b13f7C9B4Fb7658D8afd48481b"
	SimpleTokenAddress = "0xF16Fdb1FF23359633cCe37b7554394E8beA262D4"
	ICS20TransferBankAddress = "0xFEAB95Eeb8507978bC5edD22E9BA2F52f9d377A1"
	ICS20BankAddress = "0xff67836F5cb28030F6B8bDC32736F69e2e91d3F2"
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
