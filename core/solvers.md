---
head:
  - - meta
    - property: og:title
      content: Executors
  - - meta
    - name: description
      content: Executors of Plug handle the simulation and execution of signed intents that have been signed by end-users.
  - - meta
    - property: og:description
      content: Executors of Plug handle the simulation and execution of signed intents that have been signed by end-users.
---

# Solvers

At the heart of Plug and intents in general, is the reality that the user signing intents does not need to be one the running the transaction. With a signed intent, effectively anyone can execute the transaction if they have been given permission and all conditions have been met.

With the removed need of self-settlement users no longer need to be glued to their devices at all times. They can simply sign an intent and rest easy knowing when all the conditions have been met the transaction will be executed.

Powering this, are `Solvers`. A `Solver` will hold signed intents, constantly simulate the transaction, and when all conditions are passing, submit the transaction without any continued involvement from the user that signed the intent.

## The Abstraction

For `Solvers` there are two key processes:

- [solve](/core/solvers/solve): Simulate and/or execute the intent.

Today, [Plug](/) is built with `Exclusive Solvers` in mind. When signing an intent, if a `Solver` is needed or wanted, their address is included inside of the intent. If someone other than the allowed `Solver` tries to submit the intent onchain it will revert and have no effect on the state of the blockchain.
