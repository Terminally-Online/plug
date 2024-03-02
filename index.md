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

Plug is a `Solidity` protocol and `Typescript` interface for building and interacting with protocols through intents. You can learn about the rationale behind the protocol in the [Why Plug](/introduction/why-plug) section.

## User Quickstart

[Plug](/) has been designed to serve the end-user first. If you're not a developer all you have to do is head to the [official application](https://onplug.io) and you can get off to the races. There you will find templates and guides to lead you on your journey.

## Developer Quickstart

To work with `Plug` at the protocol layer you'll need to install the core framework package by opening a terminal and running the following command with your package manager of your choice:

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

### Signing an Intent

With your target contract prepared, it is now time to configure the conditions under which the transaction can be executed and distribute the fuses. Let's go ahead and declare the fuse tree for our intent and allow execution to safely be by an account in the Executor pool:

::: code-group

```typescript [./signer.ts]
// * Create a new instance of the Plug framework.
const framework = new Plug(name, version, chainId, constants.types, contract);

const plugs = await framework.sign(owner, "Plugs", {
  plugs: [{
    current: {
      target: CONTRACT_ADDRESS,
      value: 0,
      data: encodeFunctionData({
        abi: CONTRACT_ABI,
        functionName: "echo",
        args: ["Hello World"]
      })
    },
    fuses: [
      Plug.Revocation(SIGNER_ADDRESS),
      Plug.LimitedCalls(1)
    ],
    fee: 0,
    maxFeePerGas: 0,
    maxPriorityFeePerGas: 0,
    solver: Plug.Solver
  }],
  salt: Math.floor(Date.now() / 1000);
});
```

:::

After signing, all there is left to do is submit the signed bundle to the Executor pool. This pool operates on an open API mechanism which means you can choose to use the first-party service provided or spin up your own instance and settle your own transactions.

When you're ready, all you have to do is run a single line of code like:

```typescript [./signer.ts]
plugs.submit();
```

### Streaming Intents

On the other side of things, [Solvers](/core/solvers) have the ability to listen for newly created intents that can be submit onchain. Using the same framework used to sign intents, a [Solver](/core/solvers) can open a connection to the distribution WebSocket with:

```typescript [./solver.ts]
const framework = new Plug(name, version, chainId, constants.types, contract);

framework.stream();
```

With this, you will receive all newly signed intents.

If you would like to only receive orders that your [Solver](/core/solvers) has permission to manage and run you can do so easily by opening your stream with:

```typescript [./solver.ts]
const client = WebSocketProvider(RPC_URL, process.env.PRIVATE_KEY);
const signature = await framework.sign("Solver", client);

framework.stream({ solvers: [SOLVER_SIGNATURE] });
```
