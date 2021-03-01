package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x2FB9F1D2310D520e9D116b4752Db3c582ecB972c"
	IBCHandlerAddress = "0x07392E81183f86CD2B66f427c1E8Cc21eca2fB31"
	IBCIdentifierAddress = "0xc45b69894Bd1dD735779927dd841563ECB612C46"
	IBFT2ClientAddress = "0xB40d873AEB1fb78c92D9Bc57CCCdC323543C9203"
	SimpleTokenModuleAddress = "0xbD55811C336F26d6F19F8E0cD140AB8b37bCbe83"
	SimpleTokenAddress = "0xA95251102a0D5C21642319eC70F5b84d49A69d66"
	ICS20TransferAddress = "0xDED9f28e9B788a7509d1cbAeb116f2AD7Ff38CfB"
	ICS20VouchersAddress = "0x3568d0D3FDB692821F4f8FEe3cC195C68b3398BA"
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
