package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x59679b1385248F33042e76fC538abc32ab286DF8"
	IBCHandlerAddress = "0x5baB27B8d4e123fcE29B48B8b0307f57716a5a19"
	IBCIdentifierAddress = "0x3d61005BE64B64480449c6B62E7a6D22A97Aa7D2"
	IBFT2ClientAddress = "0x6b6B1d020Ae984F5104768EB3d62c2F304593D81"
	MockClientAddress = "0xe5e2F18E98Dced653e9b5bF59ecC6A6687ceb373"
	SimpleTokenAddress = "0x0103FE92f6cE3d309b34BD2cBfb83f994402C867"
	ICS20TransferBankAddress = "0x93DdB80f7E23bFa1172B9930c45EcDEfbfFe0c5E"
	ICS20BankAddress = "0xEa34C547930d232E5Bf6D616Df9d1929F678f802"
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
