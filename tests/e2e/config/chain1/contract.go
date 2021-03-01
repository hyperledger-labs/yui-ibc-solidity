package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x18Bf5B933e16774c5183522e27234F9e4EB21Ed5"
	IBCHandlerAddress = "0x74D2C2854820F7D5777297b3bb5073369a2e996e"
	IBCIdentifierAddress = "0x22EbbE299c54Efb345a051F852633b800ffA535a"
	IBFT2ClientAddress = "0xab6f87479F9E955cC700f7279eE5F72324f3ea64"
	SimpleTokenModuleAddress = "0xfEBFF19Ca425651Ef460ddC99892dE234EDE6874"
	SimpleTokenAddress = "0xA129fc2a5b814678Fed5cF71048d25740C8278E6"
	ICS20TransferAddress = "0x7F87D1ac88Ba30b9270c50a37d7a88BC7679df89"
	ICS20VouchersAddress = "0x340493Eac310A364e09B888F2ECc05DB11eD609b"
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
