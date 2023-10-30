---
head:
    - - meta
      - property: og:title
        content: Getting Started
    - - meta
      - name: description
        content: Build your first protocol with Emporium in minutes.
    - - meta
      - property: og:description
        content: Build your first protocol with Emporium in minutes.
---

# Getting Started

## Overview

Emporium is a `Solidity` protocol and `Typescript` interface for building and interacting with protocols that support declarative EVM transactions (intents).

You can learn about the rationale behind the project in the [Why Emporium](/introduction/why-emporium) section.

## Installation

To get up and running with `Emporium`, you'll need to install the core protocol and the interface by opening a terminal and running the following command with your package manager of your choice:

::: code-group

```bash [npm]
npm i @nftchance/emporium-core
```

```bash [pnpm]
pnpm i @nftchance/emporium-core
```

```bash [bun]
bun i @nftchance/emporium-core
```

:::

## Quickstart

All in all there are really only 2-3 steps when interacting with `Emporium`.

### 1. Setup your Protocol

Integrating `Emporium` into your protocol is as simple as inheriting from the `Framework` contract and passing in your protocol's name and version to declare the [domain](https://eips.ethereum.org/EIPS/eip-712#definition-of-domainseparator) of your protocol's intents:

::: code-group

```solidity 5,7,9,10,12 [PeerToPeerBridge.sol]
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

import { Framework } from '@nftchance/emporium-core/src/contracts/abstracts/Framework.sol'  // [!code focus:2]

contract PeerToPeerBridge is Framework { // [!code focus]
    constructor(
        string memory $name, // [!code focus:2]
        string memory $version
    )
        Framework($name, $version) // [!code focus]
    { }
}
```

:::

By inheriting from `Framework`, `invoke` and `contractInvoke` are added to your protocol enabling full support for intents in just those few lines; **there is no need to write any additional code or fiddle with the internals of the protocol.**

::: danger

Note: The contracts of Emporium have not yet undergone a complete audit. Please do not launch into production without a thorough review of the codebase.

:::

### 2. Sign the Permission

With your contract declared, it is now time to configure the conditions under which the transaction can be executed and distribute the permissions. Let's go ahead and declare the permission tree for our intent and delegate to a different account:

::: code-group

```typescript 10-16 [./example.ts]
// * Create a new instance of the Emporium framework.
const framework = new Framework(
	name,
	version,
	chainId,
	constants.types,
	contract
)

// [!code focus:7]
const signedPermission = await framework.sign(owner, 'Permission', {
	delegate: getAddress(owner.account.address),
	authority: bytes32(0),
	caveats: [],
	salt: bytes32(Date.now().toString())
})

// * Retrieve the object that will be passed onchain.
const SignedPermission = signedPermission.intent
```

:::

With just these few lines of code we have:

-   Instantied a new instance of the Framework that will hold our intent references.
-   Signed the raw permission delegation object with the `owner` account.
-   Declared the object type as `Delegation`.
-   Set the `authority` to `bytes32(0)` to give the invoker full control.
-   Left caveats empty to allow the invoker free reign over the execution.
    -   When you're building a real app you will have many caveats, but for simplicity we're leaving them out for now as the precise input shape and data is highly dependent on your protocol and implementation.
-   Set the `salt` to the current timestamp to ensure the intent is unique.
-   Retrieved the `SignedDelegation` object that will be passed and verified onchain.

That's a lot happening in just a few lines so it may take a second to wrap your head around it fully. As you're getting more familiar with the architecture, you have all the help of `Typescript` autocomplete at your fingertips. **Don't be afraid to use it.**

::: info

Note, in this example `authority` was set to `0x0` which gives the invoker full control over the delegated function. If you want to restrict the invoker to a specific function, you can set the `authority` to the address of the function you want to delegate to. We will get into the details of [Caveats](/intents/caveats) and [Enforcers](/enforcers) shortly.

:::

### 3. Sign an Intent

::: warning

TODO: Write the documentation for this piece.

:::
