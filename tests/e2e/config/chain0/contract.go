package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x3d534EA50c443164CA710Dcea0D75cb68E4A2E39"
	IBCHandlerAddress = "0x4Fb426e0DF93CA7A3957C39E5C27B23Fb6210Ec6"
	IBCIdentifierAddress = "0x5987561e4396FC977AceFdB8DC2745305c53543a"
	IBFT2ClientAddress = "0xaA5960eFA62bf00f5505470955A8E35C6DD99E6C"
	SimpleTokenModuleAddress = "0x7D85CEbe703d1441428E958485AD3a4712e99A0c"
	SimpleTokenAddress = "0xbe57A5CdE6941Bd4a2dA6Cf9d30C67312552b38C"
	ICS20TransferAddress = "0x935021270c0234c75f0cB297C99F21d0D5828e5b"
	ICS20VouchersAddress = "0x093ab6eB24de83B13056C77AA9692542604A2d22"
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
