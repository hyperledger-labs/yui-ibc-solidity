package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x4F781daD4cA110BCCFcb3823fB820Ff701d9A1ca"
	IBCHandlerAddress = "0x5a1b2f8d188E220C0ADb9d44bc6f8BA325a44bD5"
	IBCIdentifierAddress = "0x0103FE92f6cE3d309b34BD2cBfb83f994402C867"
	IBFT2ClientAddress = "0x9642005B9589A9bA382dA6f326CB225D32741207"
	SimpleTokenAddress = "0x2ecC613923741aE26B1C147827E526cf1e439131"
	ICS20TransferAddress = "0x5F068E1782e3c5afD257FFe90A99a86F5B25fcA2"
	ICS20VouchersAddress = "0x1c9B5961584c59dC4d5b3b0F23b3aCBB5ae7F94b"
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
