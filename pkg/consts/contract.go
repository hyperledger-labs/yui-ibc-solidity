package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xF5aEe62756aE34d54b0ae28B1ddEB121D0b3B593"
	IBCClientAddress = "0xfcD670B3F4bF2679457079887B72AEa30ee72b29"
	IBCConnectionAddress = "0xF3943bd0ae0e272b1eaB70008429e704F9D4c451"
	IBCChannelAddress = "0xf8cC116326b1Fa45DF6Ba57c16C4f4eF4E17A502"
	IBCRoutingModuleAddress = "0xDAE6c6F7fB263e58780604A20cE0d562c6a52390"
	IBCIdentifierAddress = "0xed823B51718C2fE62346ed318243FEDd2209062f"
	IBFT2ClientAddress = "0x907f52F3AfA83d0845b145EfA5F444Dc56c6307C"
	SimpleTokenModuleAddress = "0xaE0aB73acE678dDE54c3E8ace0Afe48B6935668E"
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
