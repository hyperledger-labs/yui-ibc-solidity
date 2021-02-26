package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x2d8aCe6cDFAe54a57763D6A9930540C420a66A2f"
	IBCHandlerAddress = "0x0dc8d0F42C8b37f56F14D0Cf0a88ECB36Da04d2C"
	IBCIdentifierAddress = "0x23ef151a3C438033ddAB620452bC035EA966ceaF"
	IBFT2ClientAddress = "0x8eF354C3ED818b54346b8C7B2413Cedd0E5049D6"
	SimpleTokenModuleAddress = "0x9c3EDEAC9F671Cc26c6635546875862a38db3739"
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
