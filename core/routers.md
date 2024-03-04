---
head:
  - - meta
    - property: og:title
      content: Routers
  - - meta
    - name: description
      content: Routers serve as the key entrypoint for external execution of Plug.
  - - meta
    - property: og:description
      content: Routers serve as the key entrypoint for external execution of Plug.
---

# Routers

For [Plug](/), a significant portion of the execution power lies in the ability to operate through a single Router. In function, this means that regardless of the target, the intents of [Plug](/) users are resolved most often through a `Router`.

Without a `Router`, [Solvers](/core/solvers) would have to run a transaction for every unique [Socket](/core/sockets). So, if they want to settle an intent for 10 different users they would have to interface with each of those targets. Instead, they simply interact with the `Router` that coordinates transactions for each of the [Sockets](/core/sockets) through a single transaction. Notably, there is often little reason for end-users to interact with the `Router` themselves. The existence and power of the `Router` is almost entirely realized by the [Solver](/core/solvers) of each intent.

## The Abstraction

An important benefit of a `Router` model is that it means there is no need for protocols already deployed to update nor is there a need for new protocols to integrate a piece of [Plug](/) directly. Instead, every contract deployed can be interacted with easily.

With execution localization, for a protocol to support intents there would have to be a wrapper for every connection. There would need to be a mechanism that enables the protocol to recover the sender from the message forwarded by the [Solver](/core/solvers) which is not ideal due to the blocking experienced by all contracts already deployed as well as those that do not include the logic on their own. This assumption and need would improperly place the onus of development and integration on protocols.

Instead, a user of [Plug](/) simply signs a message for their [Vault](/instances/vaults) and execution is automatically routed through it enabling the instant support for ~100% of smart contracts deployed. There's no sender recovery mechanism within any of the protocols. All actions are properly attributed without abstraction without introducing new risks like [ERC-2771](https://eips.ethereum.org/EIPS/eip-2771) does.

To enable this, a `Router` contains key functions that enable execution on behalf of [Sockets](/core/sockets) with:

- [plug](/core/routers/plug): Execute [LivePlugs](/generated/base-types/LivePlugs) bundles in a single or batched method.
