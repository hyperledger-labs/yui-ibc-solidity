module github.com/datachainlab/ibc-solidity

go 1.15

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/stretchr/testify v1.7.0
	github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef
	github.com/umbracle/go-web3 v0.0.0-20201123180055-e3fe47008350
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
