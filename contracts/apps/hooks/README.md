Implementation of [Osmosis IBC hooks middleware](https://github.com/osmosis-labs/osmosis/tree/main/x/ibc-hooks).

Instead of calling CosmWasm contracts, calls EVM contracts.

```json
// must be minified and escaped
{
    "evm": {
        "contract": "0xFFFFFFF",
        "abi": "0xFA73123F................F1"
    }
}
```

Where: 
- `contract` is address to call after packet received.
- `abi` is EVM encoded contract call.


Consisting of 2 contracts.

One contract is `OsmosisHookICS20AppStack` which is drop in replacement for `ICS20TransferBank`. 
Second contract is `IbcOsmosisHookExecutor`.

`IbcOsmosisHookExecutor` is created per IBC packet sender. Called contract can query `IbcOsmosisHookExecutor` for original sender and channel.

Why `IbcOsmosisHookExecutor` exists? 

First, it replicates authentication model of IBC Go module, which allows to identify original sender, which is unique per sender and chanel.

Second, it allows to do `abi` calls existing contracts without modification as needed.

`IIbcOsmosisHookCallback` interface can be implemented by sender chain contract, to receive metadata to handle acknowledgments and timeouts if needed.

