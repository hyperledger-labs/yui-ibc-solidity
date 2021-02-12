# IBC-Solidity

IBC implementations in solidity.

# Features

- Implementation of ICS-compatible IBC
  - includes Proto3 serialization support
- Light client of IBFT2.0

# Getting started

Launch two Besu chains with IBC Contract deployed with the following command:

```sh
$ ./scripts/setup.sh testtwochainz
```

// TODO write a description of cli tools.

# Test

After launch the chains, execute the following command:

```
$ make e2e-test
```
