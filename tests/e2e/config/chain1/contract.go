package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xed823B51718C2fE62346ed318243FEDd2209062f"
	IBCClientAddress = "0xF5aEe62756aE34d54b0ae28B1ddEB121D0b3B593"
	IBCConnectionAddress = "0x907f52F3AfA83d0845b145EfA5F444Dc56c6307C"
	IBCChannelAddress = "0xfcD670B3F4bF2679457079887B72AEa30ee72b29"
	IBCRoutingModuleAddress = "0xF3943bd0ae0e272b1eaB70008429e704F9D4c451"
	IBCIdentifierAddress = "0xB5B027c7e67CE7de62Df23A627Ad3B3e9DeD0443"
	IBFT2ClientAddress = "0x666C044632FBD306C82032F728594c675fce3d06"
	SimpleTokenModuleAddress = "0xf8cC116326b1Fa45DF6Ba57c16C4f4eF4E17A502"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetIBCClientAddress() common.Address {
	return common.HexToAddress(IBCClientAddress)
}

func (contractConfig) GetIBCConnectionAddress() common.Address {
	return common.HexToAddress(IBCConnectionAddress)
}

func (contractConfig) GetIBCChannelAddress() common.Address {
	return common.HexToAddress(IBCChannelAddress)
}

func (contractConfig) GetIBCRoutingModuleAddress() common.Address {
	return common.HexToAddress(IBCRoutingModuleAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
