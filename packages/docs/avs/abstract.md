# Abstract

In partnership with [Othentic](https://www.othentic.xyz/) we have designed, Circuit, an EigenLayer AVS with the key principles of:

- **Event-Driven Architecture:** To respond as quickly as possible to market changes and execute precisely on [schedule](../execution/schedules.md) our system is constantly simulating the state of intents and looking for a valid route.
- **Collaborative Solving:** With the streaming of upcoming executions the Plug Solver works to treat [each action](../concepts/actions.md) as an atomic unit to enable for the most efficient execution of [declarative intents](../concepts/actions.html#declarative) like swaps.
- **Efficient Leader Elections:** Until now most existing intent orderflow has been swap and bridge based making it exceptionally capital dependent. With the introduction of non-swap actions it is more important than ever to have a robust election process for intent solving and execution rights.
- **Massively Redundant Execution Trees:** The worst thing that could happen is having a valid route, but never using it onchain. With a dense pool of Operators everything is in place to ensure your intent is successfully fulfilled.
- **Onchain Validation:** Upon successful onchain execution of a Plug, the Attesters will verify the results onchain and come to shared consensus on the validity of the proof provided by the actively elected Performer.
