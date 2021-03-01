package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x1A50a06145c7DcDB0c0d6e347822B616727533b7"
	IBCHandlerAddress = "0x73f1431524995074c4B8a6AB2f92089f8b6A7B45"
	IBCIdentifierAddress = "0x3568d0D3FDB692821F4f8FEe3cC195C68b3398BA"
	IBFT2ClientAddress = "0xD0D78e9f77C6C61E5Dcf3a67078fD91100B49704"
	SimpleTokenModuleAddress = "0x262f3C0bD596A3B3b94aA629c86a10eCC086683C"
	SimpleTokenAddress = "0xa8156F951Bf5f78bAa2C42149f487dD96cDa71f3"
	ICS20TransferAddress = "0x2Bc07709bC5556Bed112BA66684a2E949758BEaF"
	ICS20VouchersAddress = "0xd3B00E50dfF4dbA78Ca3CC5d7cCDC6C1c0fd49ac"
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
