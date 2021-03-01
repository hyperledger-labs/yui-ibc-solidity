package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xB0378a47A83D1c2c4191cc43570Fc41bef329800"
	IBCHandlerAddress = "0xF343D7E418bdA316873569C4cD25c9d6E47D11C0"
	IBCIdentifierAddress = "0xC46e4aeE0D2FB4B96DA1400fA7f63e24AbD1bf24"
	IBFT2ClientAddress = "0xEb25Ae0bFd59BC57eefbCd86d8Ca89853ce5139A"
	SimpleTokenModuleAddress = "0x7094e9D61d689B868794d645b397ceD8FC7c13C0"
	SimpleTokenAddress = "0x836A285aABb6747CbD23b6a2B1739976C11F0c7B"
	ICS20TransferAddress = "0x6F9cDF68192861f1aDDbee6D27C44eC50ED34bF7"
	ICS20VouchersAddress = "0xBCaD965Bc2F65D1f513532aAA04273A65DeDd710"
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
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20VouchersAddress() common.Address {
	return common.HexToAddress(ICS20VouchersAddress)
}
