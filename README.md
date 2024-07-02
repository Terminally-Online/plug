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

To run the tests of `@nftchance/plug-core` is incredibly straightforward. Open your terminal and run:

```bash
pnpm i
forge test
```

## Building The Address Libraries

In Plug we utilize a self-referencing architecture that enables the ability to mine addresses and then refer to the constants within that contract. This is done so that any time the `Factory` is updated, we do not need to worry about manually updating the address everywhere that it is used. To generate the proper files run:

```bash
pnpm build:mine:quick
```

> **NOTE**
> Due to the architecture, you will need to mine the addresses twice in order for everything to function as the bytecode of things will change once the addresses are updated in `PlugAddressesLib`. A signal to know that you need to run it again is if a test results in `unexpected-address`.
>
> This is only required when you've made changes to a base contract. If you have not made any changes to the base contracts, you can skip this step. If something is reverting with `unexpected-address` or `invalid-initcode` there is something wrong with the compiler that you are using. Issues can be varying in severity and I will not have time to troubleshoot them all.

## Building The Package For Distribution

For version management assistance, `@nftchance/plug-core` is built with the help of `@changesets/cli`. When it is time for a release simply open your terminal and run:

```bash
pnpm changeset add
```

With a changeset created all you have to do is submit your commit/PR to the repository. Everything else will be handled for you.
