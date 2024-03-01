---
head:
  - - meta
    - property: og:title
      content: Fuse | Threshold
  - - meta
    - name: description
      content: Append a threshold condition to a Plug intent enforcing the value is within range.
  - - meta
    - property: og:description
      content: Append a threshold condition to a Plug intent enforcing the value is within range.
---

# Threshold Fuse

An execution threshold is a simple check that internally carries validation to ensure that the time based threshold has not been exceeded.

With the `Threshold Fuse` it becomes remarkably simple to define a condition that requires a piece of underlying blockchain state is larger or less than the one specified by the intent signer.

## Logic

In practice, this enables the ability for users to declare intents that:

- Can only be executed before the declared time.
- Can only be executed after the declared time.

If the condition is not met, the simulation and transaction will revert.

## Abstract

As an abstract, the `Threshold Fuse` powers functionality of the `BlockNumberFuse`, `TimestampFuse` and even the `BaseFeeFuse`. Powered by the same core implementation, a user can declare an intent to be executed only when the onchain value does not exceeded (or fail to meet) the defined threshold.

Due to this abstract design, the same logic is reused across each method of threshold definition. To set the method use for value lookup you simply override the `_threshold` function such as:

```solidity
/// @dev Returns the current block number.
function _threshold() internal view override returns (uint256) {
    return block.number;
}
```

Functionally, onchain there are several key mechanisms that serve threshold-like optionality:

- `block.timestamp`
- `block.number`
- `block.basefee`

After your override is implemented, the threshold will utilize the return established above.
