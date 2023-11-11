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

The crypto industry has been plagued by the inefficient and unfortunate realities of native EVM transactions. In just the last few months there has been a significant rise in alternative transaction settle mechanisms primarily focused around [smart accounts](https://eips.ethereum.org/EIPS/eip-4337) and alternate mempools. Yet, while there is more noise in the market, all existing solutions lack in either **composability**, **extinsibility**, or **modernity**.

Data trends reveal that wasted gas on failed transactions and inefficiencies is equivalent to millions of dollars lost every year. Plug aims to address these crucial pain points by reimagining how EVM transactions can be more logical, efficient, and user-centric.

## The Problems

Even after a decade there is still no battle-tested plug-and-play option. Instead, developers and blockchain users have been left to deal with:

### Authorization

Traditionally, permissions are handled on a contract-by-contract basis. This leads to a fragmented and inconsistent security model where each smart contract has its unique authorization mechanism, making it cumbersome to manage permissions across multiple contracts.

### Ordering

In the classic EVM, transactions are processed sequentially. This design imposes severe limitations on transaction throughput and creates bottlenecks, particularly during network congestion.

### Anticipatory Costs

Users are required to pay transaction costs upfront. This system is inefficient because costs are incurred before knowing whether a transaction will yield any value, effectively making every transaction a gamble.

### Scalability

Managing permissions often involves declaring them for each reference or interaction, leading to a bloated and inefficient system. This makes it challenging to scale the protocol or system as new features and interactions are added.

### Unbounded Transactions

Once a transaction is submitted, there is no built-in mechanism to protect against changes in contract state or network conditions that might affect the transaction before it gets confirmed. This leaves users vulnerable to smart contract vulnerabilities, front-running, and other forms of transaction manipulation.

These issues aren't just theoretical—they manifest as real bottlenecks that hamper adoption and user experience. The need for a solution has never been more pressing.

As the developer of protocols aimed at typical consumers the requirements to interact with the blockchain have been a constant hurdle. If users cannot access what you've made; you are missing out on the vast majority of users regardless the potential.

I wanted the ability to lower the barrier of entry without negatively impacting all the existing blockchains benefits such as _censorship resistance_ and _settlement reliability_, but all existing options required a significant level of integration, customized architecture and one-off solutions that couldn't be reused project to project.

A maximum of 5 minutes and I wanted to be up and running with a new idea. So, I created **Plug**: an extensible framework that provides plug-and-play utilities for your onchain `Solidity` protocol and `Typescript` based app. Inspired by all the tried solutions before, `Plug` packages as many benefits as possible together while leaving compromises, choices and opinion up to the implementing consumers.

## Traditional Blockchain Transactions: A Refresher

In the orthodox blockchain transaction model, a user's `Account` is used to sign a `Transaction`. [This `Transaction` specifies several parameters like the contract address, the method to invoke, and the amount of gas to allocate.](/intents/imperative-transactions)

The `Transaction` is then broadcast to the network, where it waits in a mempool until miners include it in a new block.

Critically, `Transactions` not do settle based upon the ordering of the value created by the execution (_even though they theoretically could in a vacuum_), but by the magnitude of value paid to miners by the sender of the transaction.

The user has minimal control over this process beyond setting the gas price. Once the `Transaction` is broadcast, it is subject to the whims of network congestion, miner priorities, and other unpredictable factors. This results in significant amounts of wasted gas (_gas bad_) as well as a swarm of general market inefficencies.

## The Declarative Difference

Plug introduces a paradigm shift. Instead of being a passive participant in the transaction process, the user gains the power to set conditions for transaction execution. This transforms the transaction model from being immediate and rigid to being flexible and condition-based.

If a transaction doesn't meet the predetermined conditions, it simply won't execute. This eliminates the risks associated with upfront costs, as you only pay for transactions that provide value. Additionally, it allows for more strategic planning around gas usage, thereby optimizing cost-efficiency.

This means that instead of crafting a transaction solely based on the contract to call and the gas to provide, an individual has the ability to explicitly declare the [conditions that must be met to allow execution](/intents/declarative-messages).

## Developer Experience

Plug was developed with one thing in mind: **time to launch.** Too much time is wasted in the crypto development industry on reinventing the wheel and solving complex problems that have not only been solved, but had their answers shared far and wide.

To accomplish this, `Plug` is designed to streamline the process of integrating [Declarative Transactions](/intents/declarative-messages) into your protocol with a `types first` approach. Unlike what you may expect, type generation and declaration for `Plug` starts with `Solidity` in the shape of [EIP-712 Type Declarations](https://eips.ethereum.org/EIPS/eip-712#definition-of-hashstruct).

By default, `Plug` ships with the base types that are needed to power declarative transactions however if you are seeking to build a more complex protocol, you can easily extend the types to meet your needs.

With your EIP-712 types defined, `Plug` will not only unlock the ability to generate the corresponding `TypeScript` logic, but the `Solidity` smart contract as well. This means that as soon as you declare the types used onchain you can immediately start using them in your application and broader protocol stack.

Plus, with a `types first` approach your types are always kept in sync ensuring that you never have to worry about type mismatches or inconsistencies. What would have taken tens to hundreds of hours before can be completed in just a couple of minutes.

To make integration as seamless as possible it was a constant priority not to be prescriptive. At every step of development I am diligent not to introduce opinion that one may disagree with or _simply not need_.

## Composability

One of Plug's main advantages is its focus on composability. In traditional blockchain platforms, the components often exist as siloed entities with limited interoperability. Plug shatters these silos by ensuring that its elements can work synergistically.

The inherent composability allows developers to build complex decentralized applications that are more than just the sum of their parts. The use of conditional transactions particularly enriches smart contract interactions, making it possible to create intricate, multi-step decentralized workflows that were previously challenging or expensive to implement.

## Extensibility

Plug's architecture is designed for extensibility. Unlike other solutions that offer a one-size-fits-all model, Plug is modular. This design enables you to tailor the framework to meet the unique demands of your specific project.

The protocol's extensible nature means it can evolve. As blockchain technology continues to advance, new modules can be added to Plug without requiring a complete overhaul of the existing infrastructure.

## Modernity

In a field where staying up-to-date is not just an advantage but a necessity, Plug excels by incorporating the latest advancements in blockchain technology and tolling. Plug offers a cutting-edge alternative to traditional transaction models by enabling conditional transactions built on top of `abitype`, `viem`, and `hardhat`.

There is no dealing with legacy code or outdated technology. Plug is built for the future while enriching as much of the past as possible.

## Minimal Integration Architecture

Plug delivers an unparalleled experience when it comes to implementing support for intents in your protocol for one simple reason:

1. Integration has been designed to happen at the lowest level possible.

Contrary to the typical past approach that led smart contract developers to implement signatures as part of their key mechanism, transaction verification and execution is pushed to the very edge of the protocol. This means, that instead of keeping an entirely separate piece of logic and conditions held in your mind during development of the core mechanisms, you achieve genuine [seperation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) and can focus on the core of your protocol.

Instead of integrating an extensive set of functionality in the underlying pieces of your protocol, simply inherit the Plug framework and you're done. **It's really that simple.** For example, [Uniswap](https://uniswap.org/), the leading decentralized exchange of Ethereum, had to integrate a swap `deadline` into the actual `swap` function of the protocol:

```solidity
function swapTokensForExactTokens(
    uint amountOut,
    uint amountInMax,
    address[] calldata path,
    address to,
    uint deadline
)
```

With Plug, a decentralized exchange (DEX) like Uniswap, could use the `ThresholdEnforcer` instead of including `blockNumber` deadlines at the application/protocol layer. Additionally `Users` could set before, after and between conditionals, instead of just before, without changing the core smart contract function logic.

Why is this important, you might ask?

First, we have to ask the question "Why does Uniswap have to enforce transaction-level access controls at all?"

It's not ideal that `UniswapV2` is required to include `blockNumber deadlines`in an`AutomatedMarketMaker` primitive. The concerns have not been separated. Not only did Uniswap developers have to design around this nuance, but all mechanisms consuming Uniswap also have to not only be aware of, but account for this nuance.

The severity of this issue increases as you survey more modern protocols such as [Seaport](https://github.com/ProjectOpenSea/seaport), which followed a very similar route:

```solidity
function validateTime(
    OrderParameters memory orderParameters,
    uint256 shortOrderDuration,
    uint256 distantOrderExpiration
) external view returns (ErrorsAndWarnings memory errorsAndWarnings);
```

Seaport was developed by a set of the best developers in the entire industry and still they were not able to build on top of a framework that enabled genuine separation of concerns.

**Everyone is relying on very similar logic while all simultaneously having to implement it themselves. This is a huge waste of time and resources.**
