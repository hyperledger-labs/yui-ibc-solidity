# IBC-Solidity

![Test](https://github.com/datachainlab/ibc-solidity/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/datachainlab/ibc-solidity?status.svg)](https://pkg.go.dev/github.com/datachainlab/ibc-solidity?tab=doc)

IBC implementations in ethereum.

NOTE: This is yet pre-alpha non-production-quality software.

## Features

- Implementation of [ICS](https://github.com/cosmos/ics/tree/master/spec)
- Implementation of ICS-20 that can integrate with ERC20 token
- IBFT2.0 Light Client

## Getting started

Launch two Besu chains with IBC Contract deployed with the following command:

```sh
# If NO_GEN_CODE is empty, setup-script will generate a proto3 marshaler in solidity
$ NO_GEN_CODE=1 ./scripts/setup.sh testtwochainz
```

## Example and Testing

After launch the chains, execute the following command:

```
$ make e2e-test
```

## Maintainers

- [Jun Kimura](https://github.com/bluele)
