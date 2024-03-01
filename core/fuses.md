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

## The Abstraction

Functionally, all Fuses are built on top of the same abstract pattern that provides straight-forward logic appendage without significant manipulation of the intent being declared. As Fuses are lined up in a declared order, the have the ability to impact the state of the underlying transaction to execute without directly communicating with one another.

To power this, Fuses are built with three key concepts in mind:

- [encode](/core/fuse/encode): Used offchain when preparing the intent to sign for a user.
- [decode](/core/fuse/decode): Used onchain to recover declaration from a generic bytes declaration.
- [enforceFuse](/core/fuse/enforce-fuse): Used onchain when validating simulation and execution.

The system has been architected in this way to provide the best experience for both users that want human-readable data, while maintaining efficiency during [Solver simulation and execution](/core/solvers).

On the surface, this can sound a bit too esoteric to understand. In practice, it is really quite simple. Imagine you want to enforce a single `uint256` that is within a range. You [encode](/core/fuse/encode) the data accordingly and include it in the signed intent.

When the transaction is submit onchain, the [Socket](/core/sockets) will automatically interface with the `Fuse` and respond accordingly if all the rules (conditions) are not met. If all the condition checks pass, then the transaction will continue and be executed.

For a simple example, let's imagine we want a `Fuse` that makes sure the intent can only be executed before a certain `block.number` to enable very simple intent expiration. In this case, the `enforceFuse` will look like:

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

With just these couple lines of code, we now have a `Fuse` that makes sure no one can execute transactions after the period of allowance set. Even if someone has your signed intent, if the `block.number` is greater than the one set, it can never be run.

This concept carries out through every `Fuse` and conditional implementation that exists within the [Plug](/) ecosystem. `Fuses`, in the most simple form, are the embodiement of conditions that must pass before a transaction is executed.

Notably, we could have just encoded the data offchain without the need for the onchain function read. However, by doing so that means we'd have to store a huge record of encoding methods that diverge from the onchain implementations with every update or small change.

Without an abstracted pattern, a [Solver](/core/solvers) would have to store a huge record of encoding methods that diverge from the onchain implementations with every update or small change. This way, we can simply maintain the list of addresses and ABIs and automatically build based on the defined context.

If each [Fuse](/core/fuses) was declared with top-level arguments instead of encoding, there would be no standard pattern of interaction. Every [Solver](/core/solvers) would have to be aware of the shape (ABI) for every [Fuse](/core/fuses) defined within the intent and directly interface.

Instead, a [Fuse](/core/fuses) can remain unverified and provide conditional capabilities to the signing user without leaking the internal functions of piece of logic applied simply by surfacing the declaration of [encoded](/core/fuse/encode) and [decoded](/core/fuse/decode) values.
