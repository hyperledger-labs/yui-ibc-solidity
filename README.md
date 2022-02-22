# IBC-Solidity

![Test](https://github.com/hyperledger-labs/yui-ibc-solidity/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/hyperledger-labs/yui-ibc-solidity?status.svg)](https://pkg.go.dev/github.com/hyperledger-labs/yui-ibc-solidity?tab=doc)

IBC implementations in solidity.

This is available not only for Ethereum and Hyperledger Besu, but also for Binance Smart Chain and other blockchains that run smart contract in EVM.

NOTE: This is yet pre-alpha non-production-quality software.

## Features

- Implementation of [ICS](https://github.com/cosmos/ibc/tree/master/spec/core)
- Implementation of [ICS-20](https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer) that can integrate with ERC20 token
- [IBFT 2.0 Light Client](./docs/ibft2-light-client.md)

## Documents

Please see [here](./docs/README.md).

## Getting started

Launch two Besu chains with IBC Contract deployed with the following command:

```sh
# If NO_GEN_CODE is empty, setup-script will generate a proto3 marshaler in solidity
$ NO_GEN_CODE=1 ./scripts/setup.sh testtwochainz
```

An example of E2E working can be found here:
- https://github.com/hyperledger-labs/yui-relayer/tree/main/tests/cases/tm2eth
- https://github.com/hyperledger-labs/yui-relayer/blob/main/.github/workflows/test.yml

## Example and Testing

After launch the chains, execute the following command:

```
$ make e2e-test
```

## For Developers

To develop this project, you need the code generator [solidity-protobuf](https://github.com/datachainlab/solidity-protobuf) to generate encoders and decoders in solidity from proto files.

Currently, you need to use [this version](https://github.com/datachainlab/solidity-protobuf/tree/fce34ce0240429221105986617f64d8d4261d87d).

## Maintainers

- [Jun Kimura](https://github.com/bluele)
