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

# Executors

At the heart of Plug and intents in general, is the reality that the user signing intents does not need to be one the executing the transaction. With a signed intent, effectively anyone can execute the transaction if they have been given permission and all conditions have been met.

Critically, this means that instead of users needing to be glued to their devices at all times they can simply sign an intent and rest easy knowing when all the conditions have been met the transaction will be executed.

Users do not have to run their own `Executor`, though they may if preferred.

## Onchain Implementation

To provide this functionality, an `Executor` takes a pre-signed intent and executes it through the [Router](
