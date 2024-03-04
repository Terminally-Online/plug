# 🔌 Plug Docs

> [!NOTE]
> This repository is a submodule of the larger [Plug monorepo](https://github.com/nftchance/plug) that contains all the required pieces to run the entire Plug stack yourself.

The `docs` package of Plug is a separate piece of the repository that contains up to date references for key concepts, architectural explanations, and are generally focused towards those that are slightly more technical outside of a few pages. With intents being an entirely new concept, documentation is purposefully kept at quite a high level. The deeper user-focused documentation lives inside of the application that users interact with as we do not want users forced to read documentation before they can level up their experience.

The core package of Plug powers the building and verification of declarative messages in both a server-context and interactive-wallet environment using the connected wallet when possible.

## Dependencies

In order to run `@nftchance/plug-core` it is necessary to install all of the following dependencies first:

```ml
├─ foundry - "Foundry is a blazing fast, portable and modular toolkit for Ethereum application."
└─ pnpm — "Efficient package manager for Node modules."
```

## Getting Started

To run an instance of `@nftchance/plug-docs` is incredibly straightforward. Open your terminal and run:

```bash
pnpm i
pnpm dev
```

## Generating Static References

Plug is built with a self documenting architecture. This means, that any time the core underlying packages are updated you automatically have access to the newest versions without having to write the documentation for them yourself.

To generate the references you simply run:

```bash
pnpm generate
```

Sometimes, you will find that you (or the documentation package in general), are using an old version of Plug. To update to the latest `types` package open your terminal and run:

```bash
pnpm i @nftchance/plug-types@latest
```

## Deploying

The documentation package is built using GitHub pages to keep things extremely simple and cheap. There is no manual deploy process, but this architecture also means that there are no `preview` deploys.

When deploying a new version, you simply need to merge into `master` and everything will automatically run. Important to note, during deployment it will run `actions` that may or may not fail. If you deploy, and there is a little red 'X' next to the commit, that means deployment failed. Click on it, navigate to the action and read the reason that it failed.

It is most likely that it will fail on 'Build Vitepress'. So, scroll down to it, and read the error message. The most common error is that you introduced a dead link (a link that can't be navigated to) and it needs to be fixed before it will deploy to a live version.
