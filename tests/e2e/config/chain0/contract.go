package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x640B66dA25E62116445F2FfDCCC3a9Dc5cCed0B6"
	IBCHandlerAddress = "0x9F311AB62C5688e84789F1E39E9D83E8447F5E9a"
	IBCIdentifierAddress = "0x56e7298A90022319950e6e3778111f46778a1F83"
	IBFT2ClientAddress = "0xCcaf3a11748a8bB4b1597e0838fF3ab52494CD6D"
	SimpleTokenModuleAddress = "0xE780254d8350084F5102B760Bc4bbA4CDb06565f"
	SimpleTokenAddress = "0x51a96A351c71Db9d6f2f650C743Cb5289193982c"
	ICS20TransferAddress = "0xF6C5701d21Ecb3ef9eF8591B4c2eD4B1b37ed6b9"
	ICS20VouchersAddress = "0x28569e00FF86568d7E3A518a546AB336BD1d2714"
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
