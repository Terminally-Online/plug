---
head:
  - - meta
    - property: og:title
      content: Fuse | Threshold
  - - meta
    - name: description
      content: Append a threshold condition to a Plug intent enforcing the time to be before or after a declared timestamp.
  - - meta
    - property: og:description
      content: Append a threshold condition to a Plug intent enforcing the time to be before or after a declared timestamp.
---

# Threshold Fuse

An execution threshold is a simple check that internally carries validation to ensure that the time based threshold has not been exceeded.

With the [Threshold Fuse](/instances/fuse/timestamp) it becomes remarkably simple to define a condition that requires the current temporal state of the underlying blockchain is larger or less than the one specified by the intent signer.

## Logic

In practice, this enables the ability for users to declare intents that:

- Can only be executed before the declared time.
- Can only be executed after the declared time.

If the condition is not met, the simulation and transaction will revert.

## Abstract

As an abstract, the [Threshold Fuse](/instances/fuse/threshold) powers functionality of both the `BlockNumberFuse` as well as the `TimestampFuse`. Powered by the same language, a user can declare an intent to be executed when the time is appropriate.

Due to this abstract design, the same logic is reused across each method of temporal definition. To set the method use for time lookup you simply override the function:

```solidity
/// @dev Returns the current block number.
function _threshold() internal view override returns (uint256) {
    return block.number;
}
```

After your override is implemented, the threshold will utilize the return established above.

Functionally, onchain there are two key time mechanisms:

- `block.timestamp`
- `block.number`

Although one can additionally utilize a Threshold fuse to manage time related state, one also has the ability to consume and enforce validation upon other things such as `block.basefee`.
