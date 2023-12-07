# 🔌 Plug Core

> [!NOTE]
> This repository is a submodule of the larger [Plug monorepo](https://github.com/nftchance/plug) that contains all the required pieces to run the entire Plug stack yourself.

The core package of Plug powers the building and verification of declarative messages in both a server-context and interactive-wallet environment using the connected wallet when possible.

## Dependencies

In order to run `@nftchance/plug-core` it is necessary to install all of the following dependencies first:

```ml
└─ pnpm — "Efficient package manager for Node modules."
```

## Getting Started

To run an instance of `@nftchance/plug-core` is incredibly straightforward. Open your terminal and run:

```bash
pnpm i
pnpm dev
```

## Generating The Framework Types

The foundation of the protocol is automatically generated so that active development can be strictly focused on the actual logic of the framework. With this in place you should never manually edit `Plug.Types.sol` and instead should always open your terminal and run:

```bash
pnpm plug generate
```

By running this command, an up to date `Plug.Types.sol` will be generated and you will have everything you need at your disposal. 

> [!TIP]
> You will only need to do this when `@nftchance/plug-types` has been updated, the most up to date version is always included as a commit in the repository.

## Building The Package For Distribution

For version management assistance, `@nftchance/plug-core` is built with the help of `@changesets/cli`. When it is time for a release simply open your terminal and run: 

```bash
pnpm changeset add
```

With a changeset created all you have to do is submit your commit/PR to the repository. Everything else will be handled for you.
