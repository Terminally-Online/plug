# 🔌 Plug App

> [!NOTE]
> This repository is a submodule of the larger [Plug monorepo](https://github.com/terminally-online/plug) that contains all the required pieces to run the entire Plug stack yourself.

The application of Plug powers the front-end user interface that enables the capability of building complex declarative EVM transactions (intents) in a trustless node-based editor.

The key functionality is powered by the following packages (each have their own function):

```ml
services
├─ authentication — "NextAuth & SIWE"
├─ client — "TRPC"
├─ database — "Docker & PostgreSQL & Prisma"
├─ ethereum — "Viem & Wagmi & WalletConnect"
├─ server — "API backend that powers the server, client interface and sdk when needed."
├─ style — "Tailwind"
└─ web — "Next"
```

## Dependencies

In order to run `@terminally-online/plug-app` it is necessary to install all of the following dependencies first:

```ml
├─ docker — "Pipeline to run containerized code processes."
└─ pnpm — "Efficient package manager for Node modules."
```

## Getting Started

### Preparing the Environment

Built into this repository is a simple script that enables the encryptoion and descryption of the `.env` file.

You do not need to keep track of several different keys and passwords when it is handled. When a push is made to the repository, `.env.encrypted` is automatically generated and you will pull down any updates with the latest commit/PR. When you run the client or server, the env will be decrypted and placed in `.env`.

For this to work, you will need to have a `.env` file in the root of your repository with the single value like:

```bash
ENCRYPTION_KEY=<THE_TEAM_ENCRYPTION_KEY>
```

Whoever is in charge of the encryption key will be able to provide it to you. Once you've set the value, your environment is ready.

### Running the App

To run an instance of `@terminally-online/plug-app` is incredibly straightforward. Open your terminal and run:

```bash
pnpm i
pnpm dev
```

> [!TIP]
> You will need a PostgreSQL database running in order to read and save the data of your application.
>
> By default, the app will try running a database through Docker. If it fails however, the build will proceed.
>
> For local development, spin one up using your preferred method, such as Docker. Again, an attempt is made automatically.
>
> For production development, it cannot be on the edge due to the use of Prisma.
