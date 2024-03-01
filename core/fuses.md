---
head:
  - - meta
    - property: og:title
      content: Fuses
  - - meta
    - name: description
      content: Fuses power the conditional enforcement of Plug with plug-and-play smart contracts that are pre-deployed.
  - - meta
    - property: og:description
      content: Fuses power the conditional enforcement of Plug with plug-and-play smart contracts that are pre-deployed.
---

# Fuses

The Fuses of Plug bring the key piece of conditional logic enforcement packaged into highly specialized single-purpose contracts. At the most root level, Fuses are designed to offer key functionality that a majority of intents will rely upon. However, they also provide discrete optionality to the broader state of the ecosystem they are deployed on.

## Onchain Implementation

Functionally, all Fuses are built on top of the same abstract pattern that provides straight-forward logic appendage without significant manipulation of the intent being declared. As Fuses are lined up in a declared order, the have the ability to impact the state of the underlying transaction to execute without directly communicating with one another.

To power this, Fuses are built with three key concepts in mind:

- `enforceFuse`: Used onchain when validating simulation and execution.
- `encode`: Used offchain when preparing the intent to sign for a user.
- `decode`: Used onchain to recover declaration from a generic bytes declaration.

The system has been architected in this way to provide the best experience for both users that want human-readable data, but also maintaining efficiency during bundler simulation and execution.

On the surface, this can sound a bit too abstract to understand. In practice, it is really quite simple. Imagine you want to enforce a single `uint256` that is within a range. To build the data needed for `enforceFuse` we will perform an offchain `staticcall` to `encode` with the value like:

```solidity
function encode(uint256 $value)
  external
  pure
  returns (bytes calldata $terms)
{
  $terms = abi.encode($value);
}
```

Notably, we could have just encoding it offchain without the need for the function read. However, by doing so that means we'd have to store a huge record of encoding methods that diverge from the onchain implementations with every update or small change. Instead, we can simply maintain the list of addresses and the value types to provide.

Calling this, we get an encoded `uint256` back in bytes form.

Now, working in the other direction, the process is much the same for decoding. The only difference between encoding and decoding though, is that `decode` is used both onchain and when we want to recover the values of an intent that has already been signed. That is to say, when a Fuse interaction has been encoded, you can still recover the state by referencing the Fuse itself. With the common patterns of ABI-based resolution, things remain simple.

For clarity, if with a `decode` function such as:

```solidity
function decode(bytes calldata $terms)
  external
  pure
  returns (uint256 $value)
{
  $value = abi.decode($terms, (uint256));
}
```

The function returns the raw `$value` that was first provided when encoding the declaration with `encode`. Everything is two-way. Now, when it comes to validate that the state is correct, we can compare any value against the `$value` declared when the Fuse interaction was originally encoded.

Maintaining our simple context, let's imagine we want to make sure the value provided is less than the current `block.number`. To do so, the `enforceFuse` of our Fuse contract will look like:

```solidity
function enforceFuse(
  bytes calldata $terms,
  PlugTypesLib.Current calldata $current,
  bytes32
)
  public
  view
  override
  returns (bytes memory $through)
{
  /// @dev Decode the terms to get the logic operator and threshold.
  (uint256 $threshold) = decode($live);

  /// @dev Confirm the intent has not expired.
  require($threshold < block.number, "PlugBlockNumberFuse:expired");

  /// @dev Continue the pass through.
  $through = $current.data;
}
```

Just like that, we now have a Fuse that makes sure no one can execute transactions after the period of allowance that we've set. Even if someone has your signed intent, if the `block.number` is greater than the one you set it can never be run.

This concept carries out through every Fuse and conditional implementation that exists within the Plug ecosystem. Fuses, in the most simple form, are the embodiement of conditions that must pass before a transaction is executed.
