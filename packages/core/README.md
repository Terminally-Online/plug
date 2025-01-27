> [!NOTE]
> This repository is a submodule of the larger [Plug monorepo](https://github.com/terminally-online/plug) that contains all the required pieces to run the entire Plug stack yourself.

The core package of Plug powers the building and verification of declarative messages in both a server-context and interactive-wallet environment using the connected wallet when possible.

## Dependencies

In order to run `@terminallyonline/plug-core` it is necessary to install all of the following dependencies first:

```ml
├─ foundry - "Foundry is a blazing fast, portable and modular toolkit for Ethereum application."
└─ pnpm — "Efficient package manager for Node modules."
```

## Getting Started

To run the tests of `@terminallyonline/plug-core` is incredibly straightforward. Open your terminal and run:

```bash
pnpm i
forge test
```

## Building The Address Libraries

In Plug we utilize a self-referencing architecture that enables the ability to mine addresses and then refer to the constants within that contract. This is done so that any time the `Factory` is updated, we do not need to worry about manually updating the address everywhere that it is used. To generate the proper files run:

```bash
pnpm build:mine:quick
```

## Building The Package For Distribution

For version management assistance, `@terminallyonline/plug-core` is built with the help of `@changesets/cli`. When it is time for a release simply open your terminal and run:

```bash
pnpm changeset add
```

With a changeset created all you have to do is submit your commit/PR to the repository. Everything else will be handled for you.
