---
head:
  - - meta
    - property: og:title
      content: Getting Started
  - - meta
    - name: description
      content: Build your first protocol with Plug in minutes.
  - - meta
    - property: og:description
      content: Build your first protocol with Plug in minutes.
---

# Getting Started

## Overview

Plug is a `Solidity` protocol and `Typescript` interface for building and interacting with protocols that support declarative EVM transactions (plugs).

You can learn about the rationale behind the project in the [Why Plug](/introduction/why-plug) section.

## Installation

To get up and running with `Plug`, you'll need to install the core protocol and the interface by opening a terminal and running the following command with your package manager of your choice:

::: code-group

```bash [npm]
npm i @nftchance/plug-core
```

```bash [pnpm]
pnpm i @nftchance/plug-core
```

```bash [bun]
bun i @nftchance/plug-core
```

:::

## Quickstart

All in all there are really only 2-3 steps when interacting with `Plug`.

### 1. Setup your Protocol

Integrating `Plug` into your protocol is as simple as inheriting from the `Plug` contract and passing in your protocol's name and version to declare the [domain](https://eips.ethereum.org/EIPS/eip-712#definition-of-domainseparator) of your protocol's plugs:

::: code-group

```solidity 5,7,9,10,12 [PeerToPeerBridge.sol]
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

import { Plug } from '@nftchance/plug-core/src/contracts/abstracts/Plug.sol'  // [!code focus:2]

contract PeerToPeerBridge is Plug { // [!code focus]
    constructor(
        string memory $name, // [!code focus:2]
        string memory $version
    )
        Plug($name, $version) // [!code focus]
    { }
}
```

:::

By inheriting from `Plug`, `invoke` and `contractInvoke` are added to your protocol enabling full support for plugs in just those few lines; **there is no need to write any additional code or fiddle with the internals of the protocol.**

::: danger

Note: The contracts of Plug have not yet undergone a complete audit. Please do not launch into production without a thorough review of the codebase.

:::

### 2. Sign the Pin

With your contract declared, it is now time to configure the conditions under which the transaction can be executed and distribute the pins. Let's go ahead and declare the pin tree for our intent and delegate to a different account:

::: code-group

```typescript 10-16 [./example.ts]
// * Create a new instance of the Plug framework.
const framework = new Plug(name, version, chainId, constants.types, contract);

// [!code focus:7]
const signedPin = await framework.sign(owner, "Pin", {
  delegate: getAddress(owner.account.address),
  authority: bytes32(0),
  fuses: [],
  salt: bytes32(Date.now().toString()),
});

// * Retrieve the object that will be passed onchain.
const LivePin = signedPin.intent;
```

:::

With just these few lines of code we have:

- Instantied a new instance of the Plug that will hold our intent references.
- Live the raw pin delegation object with the `owner` account.
- Declared the object type as `Delegation`.
- Set the `authority` to `bytes32(0)` to give the invoker full control.
- Left fuses empty to allow the invoker free reign over the execution.
  - When you're building a real app you will have many fuses, but for simplicity we're leaving them out for now as the precise input shape and data is highly dependent on your protocol and implementation.
- Set the `salt` to the current timestamp to ensure the intent is unique.
- Retrieved the `LiveDelegation` object that will be passed and verified onchain.

That's a lot happening in just a few lines so it may take a second to wrap your head around it fully. As you're getting more familiar with the architecture, you have all the help of `Typescript` autocomplete at your fingertips. **Don't be afraid to use it.**

::: info

Note, in this example `authority` was set to `0x0` which gives the invoker full control over the delegated function. If you want to restrict the invoker to a specific function, you can set the `authority` to the address of the function you want to delegate to. We will get into the details of [Fuses](/core/fuse) shortly.

:::

### 3. Sign an Plug

::: warning

TODO: Write the documentation for this piece.

:::
