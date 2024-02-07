---
head:
  - - meta
    - property: og:title
      content: Why Plug
  - - meta
    - name: description
      content: A brief preamble on why Plug was built.
  - - meta
    - property: og:description
      content: A brief preamble on why Plug was built.
---

# Why Plug

The crypto industry has been plagued by the inefficient and unfortunate realities of native EVM transactions for just under a decade. In just the last few months there has been a significant rise in alternative transaction settle mechanisms primarily focused around [smart accounts](https://eips.ethereum.org/EIPS/eip-4337) and alternate mempools. Yet, while there is more noise in the market, all existing solutions lack in either **composability**, **extinsibility**, or **modernity**.

Data trends reveal that wasted gas on failed transactions and inefficiencies is equivalent to millions of dollars lost every year and value lost to inefficient settlement is even larger. Plug aims to address these crucial pain points by reimagining how EVM transactions can be more logical, efficient, and user-centric.

## The Problems

Even after a decade there is still no battle-tested plug-and-play option. Instead, developers and blockchain users have been left to deal with:

### Authorization

Traditionally, pins are handled on a contract-by-contract basis. This leads to a fragmented and inconsistent security model where each smart contract has its unique authorization mechanism, making it cumbersome to manage conditions across multiple contracts.

### Ordering

In the classical EVM approach, transactions are processed sequentially. This design imposes severe limitations on transaction throughput and creates bottlenecks, particularly during network congestion.

### Anticipatory Costs

Users are required to pay transaction costs upfront. This system is inefficient because costs are incurred before knowing whether a transaction will yield any value, effectively making every transaction a ~gamble.

### Scalability

Managing conditions often involves declaring them for each reference or interaction, leading to a bloated and inefficient system. This makes it challenging to scale the protocol or system as new features and interactions are added.

### Unbounded Transactions

Once a transaction is submitted, there is no built-in mechanism to protect against changes in contract state or network conditions that might affect the transaction before it gets confirmed. This leaves users vulnerable to smart contract vulnerabilities, front-running, and other forms of transaction manipulation.

These issues aren't just theoretical. They manifest as real bottlenecks that hamper adoption and user experience. The need for a solution has never been more pressing.

As the developer of protocols aimed at typical consumers the requirements to interact with the blockchain have been a constant hurdle. If users cannot access what you've made; you are missing out on the vast majority of users regardless the potential.

I wanted the ability to lower the barrier of entry without negatively impacting all the existing blockchains benefits such as _censorship resistance_ and _settlement reliability_, but all existing options required a significant level of integration, customized architecture and one-off solutions that couldn't be reused project to project.

With a maximum of 5 minutes and I wanted to be up and running with a new idea. So, I created **Plug**: an extensible framework that provides plug-and-play utilities for every protocol whether deployed in the past, present or future.

Inspired by all the tried solutions before, `Plug` packages as many benefits as possible together while leaving compromises, choices and opinion up to the consuming users.

## Traditional Blockchain Transactions: A Refresher

In the orthodox blockchain transaction model, a user's `Account` is used to sign a `Transaction`. This `Transaction` specifies several parameters like [the contract address, the method to call, and the amount of gas to allocate.](/introduction/transactions#imperative-transactions)

The `Transaction` is then broadcast to the network, where it waits in a mempool until miners include it in a new block.

Critically, `Transactions` not do settle based upon the ordering of the value created by the execution (_even though they theoretically could in a vacuum_), but by the magnitude of value paid to miners by the sender of the transaction.

The user has minimal control over this process beyond setting the gas price. Once the `Transaction` is broadcast, it is subject to the whims of network congestion, miner priorities, and other unpredictable factors. This results in significant amounts of wasted gas (_gas bad_) as well as a swarm of general market inefficencies.

## The Declarative Difference

Plug introduces a paradigm shift. Instead of being a passive participant in the transaction process, the user gains the power to set conditions for transaction execution. This transforms the transaction model from being immediate and rigid to being flexible and condition-based.

If a transaction doesn't meet the predetermined conditions, it simply won't execute. This eliminates the risks associated with upfront costs, as you only pay for transactions that provide value. Additionally, it allows for more strategic planning around gas usage and unexpected outcomes, thereby optimizing the value to cost ratio.

This means that instead of crafting a transaction solely based on the contract to call and the gas to provide, an individual has the ability to explicitly declare the [conditions that must be met to allow execution](/introduction/transactions#declarative-transactions).

The concept of "If This, Then That" is straightforward but holds immense potential. It's the core logic that allows you to build complex conditional statements, which is especially powerful in smart contracts and blockchain technologies.

## An Onchain Protocol that brings IFTTT Statements

At its essence, "If This, Then That" is a conditional statement that enables automation. You set a trigger (`This`), and if that trigger occurs, a particular action (`That`) follows.

- If the trigger is met, then the action happens.
- If condition `A` is true, then perform action `B`.

This logic isn't just a fancy way of saying something. It's a powerful concept used [extensively in programming, data science](https://en.wikipedia.org/wiki/Object-capability_model), and now increasingly in blockchain technology.

You might be thinking, "Hey, this sounds a lot like [Zapier](https://zapier.com/), [Apple Shortcuts](https://apps.apple.com/us/app/shortcuts/id915249334) and all those [IFTTT (If This, Then That) services](https://ifttt.com/) I use to automate my life!" Well, you're right; the underlying logic is quite similar. Services like Zapier or IFTTT let you create "Zaps" or "Applets" that link different apps and services together based on triggers and actions.

## How does `Plug` use IFTTT?

In `Plug`, you're doing something similar but in a much more powerful and secure environment—the blockchain. Here, the "_If This, Then That_" logic enables you to create complex conditions and actions related to smart contracts. Instead of linking social media accounts or automating email notifications, you're setting rules for how digital assets and data can be accessed, moved, or transformed.

Imagine automating an entire financial system, a voting mechanism, or an ownership registry with the same ease you set a reminder to water your plants. That's the potential power of employing IFTTT logic in the blockchain via `Plug`.

By understanding the simple yet powerful logic of "If This, Then That," you're not just becoming proficient in using `Plug`; you're understanding a foundational principle of modern technology.

To sum it up in the simplest way, `Plug` is like [Zapier](https://zapier.com), but for EVM smart contracts.

## Developer Experience

Plug was developed with one thing in mind: **time to launch.** Too much time is wasted in the crypto development industry on reinventing the wheel and solving complex problems that have not only been solved, but had their answers shared far and wide.

To accomplish this, `Plug` is designed to streamline the process of integrating [Declarative Transactions](/introduction/transactions#declarative-transactions) into your protocol with a `types first` approach. Unlike what you may expect, type generation and declaration for `Plug` starts with `Solidity` in the shape of [EIP-712 Type Declarations](https://eips.ethereum.org/EIPS/eip-712#definition-of-hashstruct).

By default, `Plug` ships with the base types that are needed to power declarative transactions however if you are seeking to build a more complex protocol, you can easily extend the types to meet your needs.

With your EIP-712 types defined, `Plug` will not only unlock the ability to generate the corresponding `TypeScript` logic, but the `Solidity` smart contract as well. This means that as soon as you declare the types used onchain you can immediately start using them in your application and broader protocol stack.

## Composability

One of Plug's main advantages is its focus on composability. In traditional blockchain platforms, the components often exist as siloed entities with limited interoperability. Plug shatters these silos by ensuring that its elements can work synergistically.

The inherent composability allows developers to build complex decentralized applications that are more than just the sum of their parts. The use of conditional transactions particularly enriches smart contract interactions, making it possible to create intricate, multi-step decentralized workflows that were previously challenging or expensive to implement.

## Extensibility

Plug's architecture is designed for extensibility. Unlike other solutions that offer a one-size-fits-all model, Plug is modular. This design enables you to tailor the framework to meet the unique demands of your specific project.

The protocol's extensible nature means it can evolve. As blockchain technology continues to advance, new modules can be added to Plug without requiring a complete overhaul of the existing infrastructure.

## Modernity

In a field where staying up-to-date is not just an advantage but a necessity, Plug excels by incorporating the latest advancements in blockchain technology and tooling. Plug offers a cutting-edge alternative to traditional transaction models by enabling conditional transactions built on top of `abitype`, `viem`, and `foundry`.

There is no dealing with legacy code or outdated technology. Plug is built for the future while enriching as much of the past as possible.

## Minimal Integration Architecture

Plug delivers an unparalleled experience when it comes to implementing support for plugs in your protocol for one simple reason:

1. **Integration has been designed to happen at the lowest level possible.**

Contrary to the typical past approach that led smart contract developers to implement signatures as part of their key mechanism, transaction verification and execution is pushed to the very edge of the protocol. This means, that instead of keeping an entirely separate piece of logic and conditions held in your mind during development of the core mechanisms, you achieve genuine [seperation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) and can focus on the core of your protocol.

Instead of integrating an extensive set of functionality in the underlying pieces of your protocol, simply inherit the Plug framework and you're done. **It's really that simple.** Today, everyone is relying on very similar logic while all simultaneously having to implement it themselves. This is a huge waste of time and resources and Plug solves this.
