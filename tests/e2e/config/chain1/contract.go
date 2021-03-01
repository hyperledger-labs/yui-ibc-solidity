package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x905A69d16226eC10f3Dd4939A904629360a7Ec74"
	IBCHandlerAddress = "0xf0C3c6B56CBe448F2801C20aB41445554efeB103"
	IBCIdentifierAddress = "0xDf919000A9A2533Fb000866654ce06D565DDDb97"
	IBFT2ClientAddress = "0x76C523Cf88dE127169a872720Acdc233F46ec88F"
	SimpleTokenModuleAddress = "0xaA5960eFA62bf00f5505470955A8E35C6DD99E6C"
	SimpleTokenAddress = "0x3d534EA50c443164CA710Dcea0D75cb68E4A2E39"
	ICS20TransferAddress = "0x7D85CEbe703d1441428E958485AD3a4712e99A0c"
	ICS20VouchersAddress = "0x4Fb426e0DF93CA7A3957C39E5C27B23Fb6210Ec6"
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
