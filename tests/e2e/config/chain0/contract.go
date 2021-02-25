package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x95fE42C21F121B19D779b83eDafB97ee4d3c5e9f"
	IBCClientAddress = "0x7850096a72FF96Ae76aFaF84Fa6189c8B3DFf207"
	IBCConnectionAddress = "0x9828157efEF1FBe2a5bc809604BC3C3D022230F1"
	IBCChannelAddress = "0xC05C196E5c501ff3CdF62Ec4Fa13A96852BFbb7d"
	IBCRoutingModuleAddress = "0xB428766F67e0808C22F5976337bA9FfF46f24A81"
	IBCIdentifierAddress = "0x4fe18E0CfA90b38bF6DA7d8784271c3b0A5cc425"
	IBFT2ClientAddress = "0x500d08e08460E80aDb4B5C6e971C01ec10fD360a"
	SimpleTokenModuleAddress = "0xd629B8DF9B7F80A835b75D11a3BF7049247C28Cf"
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
