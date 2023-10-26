# @nftchance/emporium

> **Important**
> You are reading the implementation documentation for Emporium solely intended to serve developers that are actively building atop the framework.
>
> -   If you are just a general user, this: ([COMING SOON]) has the answers for you.
> -   If you are a dev looking for code documentation, this: ([COMING SOON]) has the answers for you.

Emporium powers generalized counterfactual and revokable EVM intent framework. Without having to change the core logic of your protocol you can be up and running with intents in just a few seconds. To use Emporium there are several key pieces that you will find benefit in using:

```ml
packages
├─ types — "Automatically generate the types and decoders of your intent framework."
└─ core — "Intent framework smart contracts and management utilities."
```

## Getting Started

### 1️⃣ [Generating your Types](https://github.com/nftchance/emporium-types)

> This is only neccessary if you are implementing additional types/decoders in your base `emporium` implementation. You do not need to do this step if you plan on using the core variant. By default, `emporium-core` was implemented to consume the base shapes and executions required for baseline functions.

How much time have you tried writing helper utilities to work with EIP-712 signatures, types and hashes. With `emporium-types` all you have to do is setup your configuration and run the singular `npm emporium generate` in your cli.

Instead of wasting hundreds of hours getting the proper signature types, declared messages and signed outputs you can have everything prepared in a matter of seconds. With the configured defaults, for the `emporium` framework to work you do not have to change anything unless you have a need for [advanced usage](https://github.com/nftchance/emporium-types#advanced-usage-adding-your-types).

Head to the [main `emporium-types` repository](https://github.com/nftchance/emporium-types) to read the full usage documentation or go ahead and open your terminal to run:

```bash
npm install @nftchance/emporium-types
```

### 2️⃣ [Consuming the Framework](https://github.com/nftchance/emporium-core)

In most cases, you can skip the need to [generate your types]() and go straight to the consumption of `emporium-core`. With everything already packaged for you, usage is as simple as importing the contract and configuring the top-level contract of your protocol to inherit from like:

```solidity
// path: ./myProtocol.sol
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {Framework} from "@nftchance/emporium-core/"

contract MyProtocol is Framework {
    // All of your logic.
}
```

Just like that, your protocol now has support for intents!

### Reference Implementations

It can be difficult to skim a README and fully understand what is going on here. Here is a small list of applications and protocols that you may find helpful to reference:

-   [nouns-bid-intent](https://github.com/nftchance/nouns-bid-intent) - "Schedule bids for Nouns to prevent under-RFV auction closes."

## Built With

`emporium` would not be possible without all the hardwork of those before me.

When building on top of Emporium it is advised to use the same dependencies to ensure the best experience. Dependencies outside of this range are not only untested, but not within the scope of support.

```ml
dependencies
├─ solady — "Gas optimized Solidity snippets."
├─ hardhat — "An Ethereum development environment for professionals."
├─ viem — "Build reliable Ethereum apps with lightweight, composable, and type-safe modules"
└─ abitype — "Strict TypeScript types for Ethereum ABIs."
```
