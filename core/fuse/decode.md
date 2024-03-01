# decode

Decode the `terms` of execution declared for a single [Fuse](/core/fuses).

## Parameters

- `$terms`: The active rules relative to the [Fuse](/core/fuses).

## Returns

- `...args`: The recovered and spread value that was previously encoded [encoded](/core/fuse/encode).

## Onchain Implementation

With `decode`, callers onchain and offchain can recover the encoded and signed data that the [Fuse](/core/fuses) will use to recover and confirm the active state of the blockchain and intent by calling [enforceFuse](/core/fuse/enforce-fuse). A simple example will look as follows:

```solidity
function decode(bytes calldata $terms)
  external
  pure
  returns (uint256 $value)
{
  $value = abi.decode($terms, (uint256));
}
```

Calling this, the "unknown" bytes representation becomes a direct representation of the held values and is returned in a verbose manner of response. In the example above, a raw bytes is provided as argument and is returned as the raw `uint256` value.
