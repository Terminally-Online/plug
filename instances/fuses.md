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

## Contracts

::: info

The addresses of deployed instances will be updated shortly.

:::

## Onchain Implementation

Functionally, all Fuses are built on top of the same abstract pattern that provides straight-forward logic appendage without significant manipulation of the intent being declared. As Fuses are lined up in a declared order, the have the ability to impact the state of the underlying transaction to execute without directly communicating with one another.

To power this, Fuses are built with three key concepts in mind:

- `enforceFuse`: Called onchain when being validated during simulation and execution.
- `encode`: Called offchain when preparing the intent to sign for a user.
- `decode`: Called onchain to recover the contents of the declaration from a generic bytes declaration.

The system has been architected in this way to provide the best experience for both users that want human-readable data, but also maintaining efficiency during bundler simulation and execution.
