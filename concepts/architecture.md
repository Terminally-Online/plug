# Architecture

<span style="color: rgba(0,0,0,0.6)">Plug is a generalized framework that combines onchain verification and execution with offchain simulation and route optimization. With a chain agnostic architecture, Plug can be used to build a variety of protocols that can be deployed on any Ethereum-based blockchain.</span>

## What is Plug?

Plug can be used to interact with nearly all smart contracts deployed past, present and future without being opinionated of the consumption pattern. Due to the generalized nature, protocols can be built on top that offer predictable and reliable outcomes in a bespoke manner.

In the simplest summary, Plug is a protocol, a platform, and a set of developer tools that enable more effective, secure, and responsive interactions between users and blockchains.

## The Execution Flow

To enable the generalized nature of Plug, the execution flow involves four unique parties:

- **User:** you, the creator of the intent.
- **Plug:** the abstraction layer that enables the generalization and intent declaration.
- **Solver:** the offchain engine that works to deliver the best possible outcomes.
- **Executor:** the individual that executes the transaction onchain.

For users, the interaction is as simple as declaring an intent and choosing the chains to execute on — everything else is handled by the Plug backend.

There is no need to understand the underlying mechanics of Plug or all the pieces that are required to deliver the best possible outcomes and execute the transaction. However, it may help you understand the difference between the Plug execution flow and the traditional transaction flow. So, let's take a high-level look at the execution process:

![General overview of the Plug architecture](/assets/architecture.png)

### Declaring Intents

When using Plug, users only have to declare the intent they want to take. This can be done through the [official Plug Platform](https://onplug.io) or through the [Plug SDK](https://github.com/nftchance/plug-core).

The [Plug platform](https://onplug.io) can be used to build intents that compose a variety of [constraints](/concepts/constraints) and [actions](/concepts/actions) throughout the blockchain ecosystem without requiring any technical knowledge or custom code. Alternatively, the [SDK](https://github.com/nftchance/plug-core) exposes the abstraction function and components needed for other primitive interfaces to directly integrate with Plug and deliver the benefits to the users of their own application.

In either case, the end user is able to create intents that combine [constraints](/concepts/constraints) and [actions](/concepts/actions) that deliver precisely defined outcomes.

### Solving for Outcomes

With an intent declared, the Plug backend begins working to find the best possible way to deliver the outcomes. Finding a route can be a complex process that requires a varying range of resources and primitives. This ability to wade through the complexity is what makes Plug so powerful, and unlocks the potential for delivering better outcomes than when relying on traditional transaction flows.

Of course, onchain primitives like [Uniswap](https://uniswap.org/) and [Curve](https://curve.fi) are not generalized. They each serve a specific purpose and are optimized for that purpose. So, the job of the Solver is first to implement the logic that consumes each primitive and then to find the best way to combine them to deliver the desired outcomes.

#### Aggregator and Solver: The Dynamic Duo

_How does this really work?_ Well, it's actually quite simple.

Driving the Solver is an aggregator that constantly scans the blockchain for the best available trading opportunities across multiple protocols. By not being confined to a single protocol, the aggregator can leverage the liquidity and unique advantages of various platforms.

Once the aggregator provides a range of options, the Solver comes into play. The Solver’s job is to determine the most efficient transaction pathway. This includes not just selecting the platform with the best rates, but also considering transaction fees, slippage, and the historical reliability of each protocol. The Solver sequences these options into a coherent strategy that maximizes the user's benefits in terms of cost, speed, and risk mitigation.

Once a route is identified, the Solver forwards the calldata to the Executor for simulation and execution.

To provide a concrete example with numbers, let's explore a scenario where a user wants to exchange 10,000 **$USDC** for **$ETH**. We will compare the outcomes when using a single protocol versus using an aggregator combined with a Solver within Plug.

Suppose the user decides to use only [Uniswap](https://uniswap.org/) for the swap:

- [Uniswap](https://uniswap.org/): 1 **$ETH** = 2,000 **$USDC**.
- Result: The user gets 5 **$ETH** for their 10,000 **$USDC**.

However, because [Uniswap](https://uniswap.org/) is just one platform, the user might face higher slippage, especially if their order significantly impacts the liquidity pool. Let's say the slippage is 0.5% due to the size of the transaction:

- Slippage Impact: 0.5% of 5 **$ETH** = 0.025 **$ETH**.
- Final **$ETH** received: 5 **$ETH** - 0.025 **$ETH** = 4.975 **$ETH**.

Now, let's consider the same transaction using Plug's aggregator and Solver:

- [Uniswap](https://uniswap.org/): 1 **$ETH** = 2,000 **$USDC**.
- [Sushiswap](https://www.sushi.com/swap): 1 **$ETH** = 1,950 **$USDC**.
- [Curve](https://curve.fi/): 1 **$ETH** = 1,980 **$USDC**.

The aggregator compiles these options, and the Solver decides to split the transaction to minimize slippage and maximize the received $ETH resulting in:

- 5,000 **$USDC** through [Sushiswap](https://www.sushi.com/swap).
- 5,000 **$USDC** through [Curve](https://curve.fi/).

Calculation:

- [Sushiswap](https://www.sushi.com/swap): 5,000 **$USDC** / 1,950 **$USDC** per **$ETH** = 2.5641 **$ETH**.
- [Curve](https://curve.fi/): 5,000 **$USDC** / 1,980 **$USDC** per **$ETH** = 2.5253 **$ETH**.
- Total **$ETH** received: 2.5641 **$ETH** + 2.5253 **$ETH** = 5.0894 **$ETH**.

Let's assume a reduced slippage of 0.3% on both platforms due to better liquidity management:

- [Sushiswap](https://www.sushi.com/swap): 0.3% of 2.5641 **$ETH** = 0.0077 **$ETH**.
- [Curve](https://curve.fi/): 0.3% of 2.5253 **$ETH** = 0.0076 **$ETH**.
- Final **$ETH** received: 5.0894 **$ETH** - (0.0077 + 0.0076) **$ETH** = 5.0741 **$ETH**.

The user was able to get a better outcome by using a combined set of actions and protocols. This is what the Solver does, but for all transactions instead of just swaps.

### Executing Transactions

Once the Solver has found the best route, the Executor receives the final routes so that the transaction can be simulated. If the simulation is successful, the Executor submits the transaction to the blockchain(s) declared in the intent. If any of the [constraints](/concepts/constraints) are not met or the [actions](/concepts/actions) fails, the Executor will revert the transaction and continue working on the next route.
