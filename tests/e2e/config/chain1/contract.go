package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x31DEd0C2E93BB82760DCb11ddb6AfCF5045Dd371"
	IBCHandlerAddress = "0x3Ac60839D86731884d28C709A1d2f63a3728a882"
	IBCIdentifierAddress = "0x21AAc107da35eC30f344d26324ad9b7eF30F566A"
	IBFT2ClientAddress = "0xC5de6fc9374BA11cF47Cc1aECF94aBE620821BAa"
	SimpleTokenModuleAddress = "0xc2351FdD53F418F788BC4801401b4C4484b49711"
	SimpleTokenAddress = "0xA22BA4045bc50Fa615Ce9E373BeDe56782197BeE"
	ICS20TransferAddress = "0x23ef151a3C438033ddAB620452bC035EA966ceaF"
	ICS20VouchersAddress = "0x9E75ED48E7AB483Cf810EF84D7DAb844c11B1644"
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
