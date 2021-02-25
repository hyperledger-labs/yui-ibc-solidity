package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x935021270c0234c75f0cB297C99F21d0D5828e5b"
	IBCClientAddress = "0x3990146141b9c156fe5c01c802302989070EEc9E"
	IBCConnectionAddress = "0xed823B51718C2fE62346ed318243FEDd2209062f"
	IBCChannelAddress = "0x666C044632FBD306C82032F728594c675fce3d06"
	IBCRoutingModuleAddress = "0xF5aEe62756aE34d54b0ae28B1ddEB121D0b3B593"
	IBCIdentifierAddress = "0xbe57A5CdE6941Bd4a2dA6Cf9d30C67312552b38C"
	IBFT2ClientAddress = "0xB5B027c7e67CE7de62Df23A627Ad3B3e9DeD0443"
	SimpleTokenModuleAddress = "0x907f52F3AfA83d0845b145EfA5F444Dc56c6307C"
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
