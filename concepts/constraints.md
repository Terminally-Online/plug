# Constraints

<span style="color: rgba(0,0,0,0.6)">Constraints are fundamental components in [Plug](https://onplug.io) that define the rules and conditions under which transactions can proceed. They serve as the backbone of the decision-making process, ensuring that all transactions meet specified criteria before, during, and after execution.</span>

## Types of Constraints

[Informally, an intent is a signed set of declarative constraints which allow a user to outsource transaction creation and/or execution to a third party without relinquishing full control to the transacting party](https://www.paradigm.xyz/2023/06/intents). When working with [Plug](https://onplug.io) constraints you have two types to work with:

- [**Counterfactual:**](#counterfactual) check whether something **is not** true or **has not** happened.
- [**Factual:**](#factual) check whether something **is** true or **has** happened.

This distinction may seem unnecessary at first glance, but by having both sides of the interaction pattern covered you have extreme control over how transactions are executed down to the most minute details.

### Counterfactual

The most common type of constraint used in Plug is the counterfactual implementation. Often during the execution of your transaction, the constraints are applied in a way that reverts when there is state that signals something has happened. Therefore, that constraint is looking to make sure that something _**has not** happened_.

In practice, when you define a constraint such as "**only allow this intent to be called 5 times**" the onchain execution is verifying that the amount of uses is below 5. That is to say, it is checking that it **has not** already been used 5 times.

### Factual

Factual constraints validate conditions based on actual states or events that have occurred on the blockchain. These constraints are essential for enforcing rules where the truth of a statement or occurrence of an event is critical to the transaction's logic.

For example, defining a factual constraint might take the shape of **if a payment has been received** or **that a particular NFT has been transferred to a specific address** before proceeding with a transaction. This type of constraint ensures that all conditions based on real events align with the required state at the time of transaction execution.

## Seperation of Concerns

Before Plug, protocols have had to bake in base requirements into the lowest level of their protocol. This means, that there is a wide range of different limitations placed on the interaction patterns throughout the ecosystem that follow no general pattern or standardized architecture. Protocol teams have to concern themselves with rebuilding the wheel. Users lose all control of opinion and influence while having to accept a transaction definition state that is exceptionally limiting.

For example, [Uniswap V2](https://docs.uniswap.org/contracts/v2/overview) has baked in a `deadline` of swap validity. If the transaction takes too long to settle, it will be rejected. But, [Uniswap](https://uniswap.org/) is an [AMM](https://docs.uniswap.org/concepts/uniswap-protocol), not a clock, especially not an effective one.

With Plug, [Uniswap](https://uniswap.org/) has the ability to remove this piece of logic while empowering users with a much deeper set of functionality such as deadlines that must meet the constraint where the `block number` is **before**, **after**, **between**, on **a recurring schedule** and more. All without having to change a single line of the primary Uniswap protocol. [Uniswap](https://uniswap.org/) gets to remove code, stay focused on what they are good at, and minimize the complexity of their protocol and concern surface area.

## The Limitations

Technically speaking there are effectively no limitations that exist in the base framework of constraints. By accepting inputs in the form of arbitrary data the design space is as wide as the entirety of the Ethereum ecosystem.

The functionality enabled is simply limited to what the core team of Plug develops. Additionally, due to the open source nature of the protocol anyone can come along and build new constraints without relying on our team. Due to this multi-party approach the limit of constraints is largely limited to your imagination.
