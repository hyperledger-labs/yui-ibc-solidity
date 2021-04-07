# IBFT 2.0 Light Client

## Introduction

This document describes a light client that validates Hyperledger Besu using IBFT 2.0 consensus protocol.[1]

The light client is a client software that connects to full nodes to interact with the blockchain.[2][3]

In general, full nodes needs to validate every block and every transaction, which requires a lot of resources. light client can be used as intermediaries with minimal trust in full nodes, enabling blockchain validation with relatively small resources. Examples of such light client are those of bitcoin, ethereum 2.0, and tendermint. Among them, tendermint's work on light client has greatly influenced us to design the IBFT 2.0 Light Client protocol.

This document will first give an overview of the IBFT 2.0 protocol and then describe the IBFT 2.0 Light Client protocol.

## IBFT 2.0 Protocol overview

IBFT 2.0 is Proof-of-Authority (PoA) Byzantine-Fault-Tolerant (BFT) blockchain consensus protocol that enables consortium network to leverage on the capabilities of Ethereum smart contracts, ensures immediate finality, is robust in an eventually synchronous network and features a dynamic validator set.

In IBFT 2.0 protocol, the total number of nodes, `n`, and the maximum number of byzantine nodes(`f(n)`) can be expressed as `f(n) = (n-1) / 3`

There is a voting mechanism that allows nodes to add or remove validators from the valudator set. Up to one vote can be added per Block and will be applied if 1/2+ votes of the validator set are included in the block within the epoch. For more details on how voting works, please refer to [5].

An overview of the consensus state transitions in IBFT 2.0 is shown below:

![ibft2-state-machine](./ibft2-state-machine.png)

For more details on the protocol, please refer to the IBFT 2.0 protocol paper[1].

## IBFT 2.0 in Hyperledger Besu

Hyperledger Besu is an enterprise Ethereum client that implements Enterprise Ethereum Alliance (EEA) specification. Besu supports multiple consensus protocols, with IBFT 2.0 being particularly popular.

When using IBFT 2.0 in Besu, a result of consensus is stored in the extra data field of the Block. The extra data is RLP encoded. RLP encoding is a space efficient object serialization scheme used in Ethereum. The extra data field has the following RLP List:

```
[32 bytes Vanity, List<Validators>, Votes, Round number, Commit Seals]
```

The second element of the list is a list of each address in a validator set of this block, and the fifth is commit seals to this block by the validator set.

Each blockchain node verifies the commit seals to validates the block according to Algorithm 1[1].

## Light Client protocol

Similar to tendermint light client[3], we have implemented a consensus verification in IBFT 2.0 light client, which was inspired by Weak subjectivity [4].

### Trusted source

The light client is assumed to be initialized based on a trusted source. In the case of IBFT 2.0 in Besu, it is assumed that it is not difficult to provide this source since it is a permissioned chain and PoA.

### Trusting period

In the IBFT 2.0 protocol, it is implicitly assumed that the nodes can trust the valiadtor set of each finalized block height.

Therefore, asynchronous nodes somehow receive blocks from trusted nodes, validate them according to Algorithm.1 in [1], and add them to their own ledgers.

In this paper, there is an implicit assumption that the validator set for each block is always trusted. However, we thought this assumption may be difficult to satisfy in some environments. e.g. an environment where changes to the validator set occur relatively frequently.

For this purpose, the light client verification follows the IBFT 2.0 failure model with an additional assumption about a period of the validator set of height can be trusted. The assumption is follows:

- The new block must be verified by the validator set of a block generated within time duration `T` before the current time.

To implement this, the light client receives a `trusting period` parameter corresponding to `T` at its initialization. In addition, the `trusting period` is set to `T=∞` if 0 is given, which is compatible with the existing verification function for nodes.

Therefore, the light client verification is assumed to operate within the IBFT 2.0 failure model, where the maximum number of byzantine nodes is `f(n) = (n-1) / 3` within the trusting period and the validator set can change at most one for each height.

### Verification

The light client, initialized with a trusting period and a header from trusted source, validates an incoming header using the trusted header and the corresponding trusted validators.

The validating function considers a submitted Block that satisfies all of 1, 2, and 3 to be valid:
(Let `B h` be the block of height `h`, `V h` be the validator set of `B h`, `BT h` be the timestamp of `B h`, `Now()` be the current time, and `TP` be the trusting period)

Let the height of the trusted block be `n`, the height of the untrusted block be `n+m`, and if `n > 0` and `m > 0`, then

1. `BT n` < `Now()` < `BT n` + `TP` holds
2. `B n+m` has 1/3 signatures of `V n`
3. `B n+m` has 2/3+ signatures of `V n+m`

1 has already been explained in [Trusting period]. Next, let's consider 2: under the assumption that the maximum byzantine number can be expressed as `f(n) = (n-1)/3`, ensure that there is at least one honest validator. Finally, in 3 it ensure that the finalized block of height `n+m` is valid.

### Liveness analysis

Liveness of the light client depends on the change of the validator set. IBFT 2.0 supports Dynamic ValidatorSet, where the size of validator set increases or decreases by at most 1 for each block.

Therefore, we need to make sure that there is always a height at which the validation function defined in [verification] can validate the block generated by the IBFT 2.0 protocol. Note that it is obvious that there is a verifiable height when the number of validator set increases, and the discussion here will focus only on when the number of validator set decreases.

Let the validator set be `V` and the decrement be `Δ`, `Δ ⊆ V`:
```
|V ∧ V-Δ| ≥ |V| * 1/3
```
If this equation holds, there will always be a block height with a validator set that can validate the block to be validated. In IBFT 2.0, this holds because `|V - Δ| ≥ 1` and `Δ = 1`.

### Fork detection

If the failure model of IBFT 2.0 is violated and there are more than `(n-1) / 3` failed nodes, multiple valid confirmed blocks may be generated.

The light client may be able to detect such failures or attacks. This works will be done in the future.

## References

1. https://arxiv.org/pdf/1909.10194.pdf
2. https://www.parity.io/what-is-a-light-client/
3. https://medium.com/tendermint/everything-you-need-to-know-about-the-tendermint-light-client-f80d03856f98
4. https://blog.ethereum.org/2014/11/25/proof-stake-learned-love-weak-subjectivity/
5. https://besu.hyperledger.org/en/stable/HowTo/Configure/Consensus-Protocols/IBFT/#adding-and-removing-validators-by-voting
