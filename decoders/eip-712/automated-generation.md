# Automated Generation

The `Plug` framework is designed to be modular and flexible for two key reasons:

1. The protocol is not prescriptive with your needs. You can use it to build whatever you want without having to deal with the consequences of choices I made for you.
2. The framework is in `early alpha` and lives in highly iterative environment. At this time, forward progress and improvements are prioritized over API/interface stability.

Typically when writing a new protocol it is not abnormal to take many months from start to finish. Even to get ready for an audit a team of developers will often spend weeks or months writing tests and documentation. **This is a lot of time and effort to spend on something that you may not even be sure will work or get users.**

## The Paradox of Blockchain Development

For a framework such as `Plug`, extended periods of silent development are even more expensive and damaging to progress that could have been. Being a blockchain developer has been a paradoxical situation where:

- Your product has to be sufficiently complex to be valuable.
- Complexity is the enemy of iteration, delivery, and usability.

As something becomes more complex you are tasked to spend an increasing amount of time on fiddling with the tiniest of details. By the time you have a product that is sufficiently complex to be valuable, it is either too complex to iterate on or you've spent so much time on it that you've run out of money.

**This is not only a waste of time but also a waste of money.**

## Resolving the Tedium

To resolve this issue, `Plug` is built on top of an automated `Types` generation pipeline that not only ensures our types work on the first try across all parts of our stack, but that we do not even need to spend time making changes to the broader smart contracts as we iterate on the framework.

This is accomplished by using the `Plug` framework to generate the `Plug` protocol. It sounds funny, but it's true.

This means that the `Plug` framework is not only self-sufficient, but it is also self-generating. Due to the declarative nature of [Intents](/intents/introduction), a few simple rules can be defined to enable automatic generation for a vast majority of the implementation.

With this simple pipeline in place we can focus on the core mechanisms of our protocol and not get bogged down in the tedium of writing and maintaining the same code over and over again.

## Getting Started

The `Type` system of `Plug` is automatically generated using `@nftchance/plug-types`. While it is unlikely that you will ever need to use this package directly, allow me to provide some general context quickly.

To generate your first smart contract there are only a few steps. First, you'll need to install the `Typescript` library by opening a terminal and running the following command with your package manager of your choice:

::: code-group

```bash [npm]
npm i @nftchance/plug-types
```

```bash [pnpm]
pnpm i @nftchance/plug-types
```

```bash [bun]
bun i @nftchance/plug-types
```

:::

### Initialize the Environment

With `@nftchance/plug-types` installed, you can now initialize your configuration by running the following command in the `root` directory of your project:

::: code-group

```bash [npm]
npm plug init
```

```bash [pnpm]
pnpm plug init
```

```bash [bun]
bun plug init
```

By running this command an `plug.config.ts` file will be created in the root directory of your project. This file is used to configure the `Plug` framework and is the only file that you will need to edit directly.

::: warning

Two quick notes: If there is already an `plug.config.ts` file in your project, this command will fail. Additionally, if you do not set a configuration file, the default configuration will be used.

:::

### Configure the Generation

By default `plug-types` ships with a pre-built version of the current `Plug` [Base Types](/decoders/base-types) that are used to power 99% of `Plug` consumers.

For now we are going to use the default types. That means all we need to do is set the `out` path to our `contracts` directory:

::: code-group

```typescript [./plug.config.ts]
import { config } from "@nftchance/plug-types";

export default config({ out: "./contracts/abstracts/" });
```

:::

Once again, just a couple of lines of code and we are officially ready to generate our first smart contract.

::: tip

Remember, the library was built with `Typescript` so autocomplete will provide you additional context and options as you type. For simplicity here I have provided the most basic configuration possible.

:::

### Running the Generation

You're already done with the 'hard' part. To start the `Solidity` generation based on the configuration provided return to your terminal and run:

::: code-group

```bash [npm]
npm plug generate
```

```bash [pnpm]
pnpm plug generate
```

```bash [bun]
bun plug generate
```

**Boom!** Your contracts have been generated.

:::

### The Generated Output

We've already covered the `out` path, but let's take a look at the output of the `Plug` library. By default, the `Plug` library will generate a single `Types.sol` that contains:

::: code-group

```ml [./{out}/Types.sol]
├─ ITypes - "An interface that declares the types used by your framework."
└─ Types - "An abstract reference to be inherited by your framework."
```

:::

This single file contains everything needed to start building on top of [Signed Pairs](/decoders/eip-712/signed-pairs) and [Intents](/intents/introduction). Although it is a single file, a lot happens including:

- The static `TypeHash` of each `Type`.
- The initialization of the `Domain` for [EIP-712](/decoders/eip-712#domain-specification).
- The [hashGetter](/decoders/hash-getters) functions for each `Type`.
- The [digestGetter](/decoders/digest-getters) functions for each nested `Type` in the [SignedPairs](/decoders/eip-712/signed-pairs) declared.
- The [signerGetter](/decoders/signer-getters) functions for the top-level of `SignedPair` `Type`.

::: tip

`ITypes` does not include interface function declarations of the `Types` contract. This is because `Types.sol` is generated as an `abstract` contract and cannot be used without being inherited. However, if an interface is required, one can be generated from the ABI of your smart contract.

Alternativey, in `Solidity` you can just import the live referenece of the protocol and use it in place of an indepenently declared and managed interface.

:::

## Adding Custom Types

In some cases you will want access to more than just the base `Plug` types of `Permissions`, `Intents`, and all the supporting shapes such as `Transaction`, `ReplayProtection`, etc.

In this case, you need to extend the types and prepare your protocol to consume a framework that has already been initialized with all the confusing [EIP-712 data types and decoders](https://eips.ethereum.org/EIPS/eip-712) taken care of.

::: danger

If you are at this level, it is assumed that you are familiar with `EIP-712` -- verification of signatures can be very complex. If you are having trouble verifying your signatures I recommend taking the `types first` approach so that you can take have a clear understanding of where an issues may be arising.

:::

By default, when you providing your own types they will be loaded alongside the core `Plug` framework types so that you can make your protocol intent-based without any work beyond the normal scope (yes, really, no additional work -- it's pretty cool, right?)

To illustrate this, let's look at the same example from earlier where we are sending `Mail` messages from one party to another. We'll start by declaring the [EIP-712 data types](https://eips.ethereum.org/EIPS/eip-712) that we need to use:

::: code-group

```typescript [constants.ts]
export const types = {
  Mail: [
    { name: "from", type: "Person" },
    { name: "to", type: "Person" },
    { name: "contents", type: "string" },
  ],
  Person: [
    { name: "name", type: "string" },
    { name: "wallet", type: "address" },
  ],
  SignedMail: [
    { name: "mail", type: "Mail" },
    { name: "signature", type: "bytes" },
  ],
} as const;
```

:::

With this configuration, we have the types needed to send `Mail` from one `Wallet` to another. Now that our types are declared, we will initialize the `plug-types` config with the `EIP-712` data types as well as add ourselves to the list of authors by updating our `plug.config.ts` to:

::: code-group

```typescript [plug.config.ts]
import { config } from "@nftchance/plug-types";

import { types } from "./constants.ts";

export default config({
  contract: {
    authors: ["<your name>"], // [!code focus]
  },
  out: "./contracts/abstracts/",
  types, // [!code focus]
});
```

With this simple addition to your configuration file you are ready to go. Run the `generate` command and go focus on the core mechanisms of your protocol.
