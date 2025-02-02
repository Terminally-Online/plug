> [!NOTE]
> This repository is a submodule of the larger [Plug monorepo](https://github.com/terminally-online/plug) that contains all the required pieces to run the entire Plug stack yourself.

The core package of Plug powers the building and verification of declarative messages in both a server-context and interactive-wallet environment using the connected wallet when possible.

## Dependencies

In order to run `@terminallyonline/plug-core` it is necessary to install all of the following dependencies first:

```ml
├─ foundry - "Foundry is a blazing fast, portable and modular toolkit for Ethereum application."
└─ pnpm — "Efficient package manager for Node modules."
```

## Deployment

The deployment pipeline of Plug is setup to be as simple as possible and multichain by default. There is a single script to run (for the current version) that will deploy all the instances of the implementation based contracts. Because we are using `CREATE2` for deployments it does not matter who sends the transaction.

```bash
forge script \
    --chain <chain_name> \
    --rpc-url <rpc_url_matching_chain_name> \
    --broadcast \
    -vvvv \
    src/contracts/script/Plug.s.sol
```

When you want to verify the contract so you will need to add an `ETHERSCAN_API_KEY` to your `.env` and include the `--verify` tag when running the deployment script.

## Addresses

All contracts that should be deterministic are. That means `Plug.sol`, `Plug.Factory.sol` and the base `Plug.Socket.sol` are all found at constant addresses across chains. The only time the address will change when a new version is deployed. This will happen most common for the Socket, but we will always keep the list updated.

```ml
[v0.3.2]:
├─ Mining Configuration:
│  ├─ Leading Zeros: 2
│  ├─ Total Zeros: 4
│  ├─ Factory: 0x0000000000ffe8b47b3e2130213b802212439497
│  └─ Quick Mode: Yes
│
└─ Contracts:
   ├─ Plug.Factory.sol [256] — "0x0000000030c2d2825F563E2F7b78943B0Ea9D145"
   ├─ Plug.Socket.sol [256] — "0x0000000011A65597897563205669f9c46dEEE244"
   └─ Plug.sol [256] — "0x0000000021EAfaa2A0ADeec53B7E25F662920212"
```

To mine new addresses run:

```bash
pnpm build:mine:quick --version 0.3.3
```
