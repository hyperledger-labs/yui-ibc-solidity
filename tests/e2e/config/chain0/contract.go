package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x5Ea1AA48f0C4f470F5aAa12D233aA6e7b5D348Fd"
	IBCHandlerAddress = "0x8eF354C3ED818b54346b8C7B2413Cedd0E5049D6"
	IBCIdentifierAddress = "0xA22BA4045bc50Fa615Ce9E373BeDe56782197BeE"
	IBFT2ClientAddress = "0xAcFd0b27A94B41C28d31770421A45F541CcFB291"
	SimpleTokenModuleAddress = "0x2d8aCe6cDFAe54a57763D6A9930540C420a66A2f"
	SimpleTokenAddress = "0x0dc8d0F42C8b37f56F14D0Cf0a88ECB36Da04d2C"
	ICS20TransferAddress = "0x09199A80877b8481e92293dE5635dbC0DcB73127"
	ICS20VouchersAddress = "0x9c3EDEAC9F671Cc26c6635546875862a38db3739"
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
