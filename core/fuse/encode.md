# encode

Encode the `terms` of execution declared for a single [Fuse](/core/fuses).

## Parameters

- `...args`: The values applied to the [Fuse](/core/fuses) rule.

## Returns

- `$terms`: The encoded bytes data applied as rules to the [Fuse](/core/fuses) during submission.

## Onchain Implementation

With `encode`, callers onchain and offchain can build the data that the [Fuse](/core/fuses) will use to recover and confirm the active state of the blockchain and intent by calling [enforceFuse](/core/fuse/enforce-fuse). A simple example will look as follows:

```solidity
function encode(uint256 $value)
  external
  pure
  returns (bytes calldata $terms)
{
  $terms = abi.encode($value);
}
```

Calling this, we get an encoded `uint256` back in bytes form. Notably, we could have just encoded it offchain without the need for the onchain function read.
