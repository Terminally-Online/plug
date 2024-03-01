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

For Plug, a significant portion of the execution power lies in the ability to operate through a single Router. In function, this means that regardless of the target, the intents of Plug users are resolved most often through a Router.

Without a Router, Executors would have to run a transaction for every unique target. So, if they want to settle an intent for 10 different users they would have to interface with each of those targets. Instead of that being the case, they simply interact with the Router that coordinates the communicate and transactions for each of the users through a single transaction. Notably, there is often little reason for users to interact with the Router themselves. The existence and power of the Router is almost entirely realized by the Executors of each intent.

An important Caveat to many of the existing Router models in this industry is that users themselves do not directly interact with it. There is no approval of assets to an unguarded arbitrary call mechanism. All transfers of value are localized to the target of the intent and not the Router.

## Execution Globalization

An important benefit of a `Router` model is that it means there is no need for protocols already deployed to update. Instead, every contract that has already been deployed can be interacted with ease. With localization, to support a protocol there would have to be a wrapper for every connection. There would need to be a mechanism that enables the protocol to recover the sender from the message forwarded by the Executor.

Instead, a user of [Plug](/) simply signs a message for their [Vault](/instances/vaults) and execution is automatically routed through it. There's no interaction proxy relationship within the action. All actions are properly attributed without abstraction. [Plug](/) is designed to make the primitives and more general protocols of the industry more powerful therefore there is no integration expectation or new overhead introduced.

If a protocol wants to be supported, they can be without making a single change to their onchain implementation. With even the deepest level of integration, the only thing required is the development of a [Fuse](/core/fuses) which is not an obligation or expectation the protocol itself carries.

## Onchain Implementation
