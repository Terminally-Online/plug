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

All in all there are really only 2 steps when interacting with `Plug`.

### 1. Setup your Protocol

Integrating `Plug` into your protocol is as simple as inheriting from the `Plug` contract and passing in your protocol's name and version to declare the [domain](https://eips.ethereum.org/EIPS/eip-712#definition-of-domainseparator) of your protocol's plugs:

::: code-group

```solidity 5,7,9,10,12 [PeerToPeerBridge.sol]
// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugReceiver } from '@nftchance/plug-core/src/contracts/abstracts/Plug.Receiver.sol'

contract PeerToPeerBridge is PlugReceiver {
    /// @dev Include the code of your primitive.
}
```

:::

By inheriting from `Plug`, `plug` and `plugContract` are added to your protocol enabling full support for plugs in just those few lines; **there is no need to write any additional code or fiddle with the internals of the protocol.**

If you're not a developer, this step will have already been completed for you.

### 2. Sign the Plug

With your target contract prepared, it is now time to configure the conditions under which the transaction can be executed and distribute the fuses. Let's go ahead and declare the fuse tree for our intent and allow execution to safely be by an account in the Executor pool:

::: code-group

```typescript 10-16 [./example.ts]
// * Create a new instance of the Plug framework.
const framework = new Plug(name, version, chainId, constants.types, contract);

// [!code focus:7]
const plugs = await framework.sign(owner, "Plugs", {
  delegate: getAddress(owner.account.address),
  authority: bytes32(0),
  fuses: [],
  salt: bytes32(Date.now().toString()),
});
```

:::

Behind the scenes a lot happens so it may take a minute to wrap your head around it fully. As you're getting more familiar with the architecture, you have all the help of `Typescript` autocomplete at your fingertips. **Don't be afraid to use it.**

### 3. Submit the Plug

After signing, all there is left to do is submit the Plug to the Executor pool. This pool operates on an open API mechanism which means you can choose to use the first-party service provided or spin up your own instance and settle your own transactions.

When you're ready, all you have to do is run a single line of code like:

```typescript [./example.ts]
plugs.submit();
```
