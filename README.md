# 🔌 Plug Core

> [!NOTE]
> This repository is a submodule of the larger [Plug monorepo](https://github.com/nftchance/plug) that contains all the required pieces to run the entire Plug stack yourself.

The core package of Plug powers the building and verification of declarative messages in both a server-context and interactive-wallet environment using the connected wallet when possible.

## Dependencies

In order to run `@nftchance/plug-core` it is necessary to install all of the following dependencies first:

```ml
├─ foundry - "Foundry is a blazing fast, portable and modular toolkit for Ethereum application."
└─ pnpm — "Efficient package manager for Node modules."
```

## Getting Started

To run an instance of `@nftchance/plug-core` is incredibly straightforward. Open your terminal and run:

```bash
pnpm i
pnpm dev
```

## Building The Package For Distribution

> [!WARNING]
> Please do not use the contracts in production yet. They are extremely dangerous and will result in unexpected things happening. Do not do it under any circumstances. If you are reading this and get frustrated by this, please DM me and I will point out a few exploits that you may be able to help patch.

For version management assistance, `@nftchance/plug-core` is built with the help of `@changesets/cli`. When it is time for a release simply open your terminal and run:

```bash
pnpm changeset add
```

With a changeset created all you have to do is submit your commit/PR to the repository. Everything else will be handled for you.
