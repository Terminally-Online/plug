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

Plug is a `Solidity` protocol and `Typescript` interface for building and interacting with protocols that support declarative EVM transactions.

You can learn about the rationale behind the project in the [Why Plug](/introduction/why-plug) section.

## User Quickstart

[Plug](/) has been designed to serve the end-user first. If you're not a developer all you have to do is head to the [official application](https://onplug.io) and you can get off to the races. There you will find templates and guides to lead you on your journey.

## Developer Installation

To work with `Plug` at the protocol or application layer you'll need to install the core framework package by opening a terminal and running the following command with your package manager of your choice:

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

## Protocol Quickstart

Integrating `Plug` into your protocol is as simple as inheriting from the appropriate contract and passing in your protocol's name and version to declare the [domain](https://eips.ethereum.org/EIPS/eip-712#definition-of-domainseparator) of your protocol's plugs:

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

By inheriting `PlugReceiver`, your protocol now has full support for plugs in just those few lines; **there is no need to write any additional code or fiddle with the internals of the protocol.**

::: tip

Notably, this step is optional to the extent that by not including this the value of `msg.sender` will be the address of the contract that validated and executed the [Plugs](generated/base-types/Plugs).

:::

## Application Quickstart

With your target contract prepared, it is now time to configure the conditions under which the transaction can be executed and distribute the fuses. Let's go ahead and declare the fuse tree for our intent and allow execution to safely be by an account in the Executor pool:

::: code-group

```typescript [./example.ts]
// * Create a new instance of the Plug framework.
const framework = new Plug(name, version, chainId, constants.types, contract);

// * Declare the transaction that is going to be executed.
const data = encodeFunctionData({
  abi: CONTRACT_ABI,
  functionName: "echo",
  args: ["Hello World"]
});

// * Append all the conditions of execution (fuses).
const fuses = [
  // Enable revocation.
  Plug.Revocation(SIGNER_ADDRESS),
  // Only allow the transaction to be executed once.
  Plug.LimitedCalls(1)
];

const plugs = await framework.sign(owner, "Plugs", {
  plugs: [{
    current: {
      ground: CONTRACT_ADDRESS,
      voltage: 0,
      data
    },
    fuses
  }],
  salt: Math.floor(Date.now() / 1000);
});
```

:::

After signing, all there is left to do is submit the signed bundle to the Executor pool. This pool operates on an open API mechanism which means you can choose to use the first-party service provided or spin up your own instance and settle your own transactions.

When you're ready, all you have to do is run a single line of code like:

```typescript [./example.ts]
plugs.submit();
```
