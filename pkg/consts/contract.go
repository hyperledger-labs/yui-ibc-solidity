package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x2E6DeAEBF9Ee3b79A4F3Ee39e81d3672723a1A38"
	IBCHandlerAddress = "0x9332a02ed8D3Ca3CBae5e2FD98f9996825529D1C"
	IBCIdentifierAddress = "0x9f5BE5559E56ae373AB724dB8bb5354F6A537Eee"
	IBFT2ClientAddress = "0x93115d469f4324330FD647AAF72B7F27b7AA0d57"
	MockClientAddress = "0x17125b9D9Da8AA130272D2688cAe08179c737984"
	SimpleTokenAddress = "0xc84Cada16763359cCB6fe0d0d78Ccc0D067d349d"
	ICS20TransferBankAddress = "0x2af6866018aE0D6a38e3FD059f3D2a3Cd50FF08C"
	ICS20BankAddress = "0x28d3e5A53981A7D7f1BB67222A8C85D416334D7a"
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
