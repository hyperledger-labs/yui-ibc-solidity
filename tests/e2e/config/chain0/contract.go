package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x3a127904FFe33792457252e9586c0a4086e2E514"
	IBCHandlerAddress = "0xCDCC365Cb65C656f9E1c96958678ca124a067c57"
	IBCIdentifierAddress = "0x679AB82c48bD7186BcB5371295ab3f016045b988"
	IBFT2ClientAddress = "0xaA535180B01cc7FE8b23CA4F39F21413D4354b81"
	SimpleTokenAddress = "0x055BDA5Ea99a77402b65E04fa7dDBe40297d072f"
	ICS20TransferAddress = "0x2D2D2bf6E3b2bD926fe374B78C1a6CC15Ab271c8"
	ICS20VouchersAddress = "0x8d1508E7fCf6D8700988b168b7134977220C48F1"
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

func (contractConfig) GetICS20TransferAddress() common.Address {
	return common.HexToAddress(ICS20TransferAddress)
}

func (contractConfig) GetICS20VouchersAddress() common.Address {
	return common.HexToAddress(ICS20VouchersAddress)
}
