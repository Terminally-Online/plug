# Usage

<span style="color: rgba(0,0,0,0.6)">When using the set of developer tools provided by Plug, you can easily access the most critical pieces of the framework without having to build anything yourself. From creation and consumption to streaming and execution.</span>

::: warning

All of the base tools are in constant iteration and are subject to change. Breaking changes may occur at any time and you should always be on the lookout for updates. Today, we do not recommend using the `core` package in production as we will not due soft deprecations and instead will move towards a more stable and production ready package.

:::

## Installation

When working with Plug there are two primary packages you may need:

- [**types:**](https://www.npmjs.com/package/@nftchance/plug-types) The onchain types for the protocol and signatures.
- [**core:**](https://www.npmjs.com/package/@nftchance/plug-core) The framework to sign and submit intents through a simplified abstraction layer.

When working with `core` you'll automatically have `types` installed as a dependency. This means that in most cases you will be able to install `core` alone with:

::: code-group

```bash [pnpm]
pnpm i @nftchance/plug-core
```

```bash [npm]
npm i @nftchance/plug-core
```

```bash [bun]
bun i @nftchance/plug-core
```

:::

## Framework Instantiation

To get started with Plug, you'll need to install the `core` package so that you can instantiate a framework instance like:

```typescript [instantiate]
const plug = new Plug(PRIVATE_KEY);
```

With just one line of code, you've instantiated everything that is needed to create and consume intents on the Plug network. Because everything was built with Typescript you can browse all the functions and variables available with the autocomplete features of your IDE.

## Signing an Intent

When creating an intent, you'll create a set of constraints that define the limits of execution as well as any transactions that should be submitted in order to execute the intent.

In practice, this is as simple as:

```typescript [sign]
const plugs = await plug.sign([
  Plug.Revocation(SIGNER),
  Plug.LimitedCalls(1),
  {
    target: CONTRACT_ADDRESS,
    data: encodeFunctionData({
      abi: CONTRACT_ABI,
      functionName: "echo",
      args: ["Hello World"],
    }),
  },
]);
```

With just these few lines of code, we've signed an intent that will execute the `echo` function on the `CONTRACT_ADDRESS` contract with the `Hello World` message only 1 time. Additionally, at any time the signer can revoke the intent and prevent it from being executed.

## Submitting an Intent

Now that we have our intent signed, we can submit it to the Plug Executor pool so that will be executed without any further action:

```typescript [submit]
await plug.submit(plugs);
```

## Streaming Intents

On the other side of things, Solvers have the ability to listen for newly created intents that can be submitted onchain. Using the same framework used to sign intents, a Solver can open a connection to the distribution channel with:

::: code-group

```typescript [exclusive]
plug.stream({
  solvers: [new plug.Solver(PRIVATE_KEY)],
  onIntent: (intent) => console.log(intent),
});
```

:::

With this, you have the ability to receive a stream of all newly signed intents that are addressed to your solver.
