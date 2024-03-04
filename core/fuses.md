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
