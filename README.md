# IBC-Solidity

![Test](https://github.com/hyperledger-labs/yui-ibc-solidity/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/hyperledger-labs/yui-ibc-solidity?status.svg)](https://pkg.go.dev/github.com/hyperledger-labs/yui-ibc-solidity?tab=doc)

[IBC](https://github.com/cosmos/ibc) implementations in Solidity.

**IBC compatibility:** [v4.0.0](https://github.com/cosmos/ibc-go/releases/tag/v4.0.0)

This is available not only for Ethereum and Hyperledger Besu, but also for Polygon PoS and other blockchains that supports EVM-compatible.

NOTE: This is yet pre-beta non-production-quality software.

## Features

- Implementation of [ICS](https://github.com/cosmos/ibc/tree/master/spec/core)
- Implementation of [ICS-20](https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer)
- [ERC-20 Token Transfer](./contracts/app/ICS20TransferBank.sol)

## Documents

Please see [here](./docs/README.md).

In addition, a tutorial is [here](https://labs.hyperledger.org/yui-docs/yui-ibc-solidity/).

## Supported Light Client

You can deploy a Light Client that implements [the IClient interface](./contracts/core/IClient.sol) to [integrate with IBC-Solidity](./docs/architecture.md#ibcclient).

Here are some such examples:
- [IBFT 2.0 Light Client](./contracts/core/IBFT2Client.sol)
- [Tendermint Light Client](https://github.com/datachainlab/tendermint-sol/tree/use-ibc-sol-hmy)
- [Mock Client](./contracts/core/MockClient.sol)

## Related projects

- A demo of trustless bridge
    - between Harmony and Cosmos(Tendermint): https://github.com/datachainlab/harmony-cosmos-bridge-demo
    - between Celo and Cosmos: https://github.com/ChorusOne/celo-cosmos-bridge

## Development and Testing

### Unit test

```sh
$ make test
```

### E2E test

Launch two Hyperledger Besu chains(ethereum-compatible) with the contracts deployed with the following command:

```sh
$ make network-e2e
```

After launch the chains, execute the following command:

```
$ make e2e-test
```

## E2E-test with IBC-Relayer

An example of E2E with IBC-Relayer([yui-relayer](https://github.com/hyperledger-labs/yui-relayer)) can be found here:
- https://github.com/datachainlab/yui-relayer-build/tree/v0.3/tests/cases/eth2eth
- https://github.com/datachainlab/yui-relayer-build/blob/v0.3/.github/workflows/v0.3-eth2eth.yml

## For Developers

To generate the proto encoders and decoders in solidity from proto files, you need to use the code generator [solidity-protobuf](https://github.com/datachainlab/solidity-protobuf)

Currently, [v0.1.0](https://github.com/datachainlab/solidity-protobuf/tree/v0.1.0) is required.

If you edit the proto definitions, you should execute the following command:
```
$ make SOLPB_DIR=/path/to/solidity-protobuf proto-sol
```

## Maintainers

- [Jun Kimura](https://github.com/bluele)
