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
    --rpc-url <rpc_url> \
    --broadcast \
    -vvvv \
    src/contracts/script/Plug.s.sol
```

If you don't care about verifying the contract right now you can skip this step. Please don't skip it unless you're doing something for development. When you want to verify the contract so you will need to add an `ETHERSCAN_API_KEY` to your `.env` and include the `--verify` tag when running the deployment script.

When you are setting everything up you will want to remove the `--broadcast` so that it just simulates everything and you can confirm you have everything setup right before running a transaction.

Built into the scripts are redeploy protection so if you are running a batch deploy and one succeeds, but one fails, you can just rerun the same script.

## Addresses

All contracts that should be deterministic are. That means `Plug.sol`, `Plug.Factory.sol` and the base `Plug.Socket.sol` are all found at constant addresses across chains. The only time the address will change when a new version is deployed. This will happen most common for the Socket, but we will always keep the list updated.

```ml
[v0.4.1]:
├─ Mining Configuration:
│  ├─ Leading Zeros: 4
│  ├─ Total Zeros: 8
│  ├─ Factory: 0x0000000000ffe8b47b3e2130213b802212439497
│  └─ Quick Mode: Yes
│
└─ Contracts:
   ├─ Plug.Assert.sol [256] — "0x0000000005d8F29675fC43df88588bD0D5c0DeC5"
   ├─ Plug.Boolean.sol [256] — "0x00000000410986831F18E06d908bE25e5Fb949A0"
   ├─ Plug.Coercion.sol [256] — "0x0000000011fEb9342943e0029Ad2717f5a85F118"
   ├─ Plug.Database.sol [256] — "0x0000000006d777c8390a5E84Ecb88A6556A1d3B5"
   ├─ Plug.EVM.sol [256] — "0x0000000013950F5C3d277e1754d52791ceBb4091"
   ├─ Plug.Factory.sol [256] — "0x000000002D0BacD773C6055c22650A6B85a2990B"
   ├─ Plug.Math.sol [256] — "0x000000000c0352950e3aa28973824f4d01ccec4f"
   ├─ Plug.Socket.sol [256] — "0x000000000B830571FC0D6456A6Ca2b6ddAa18F6F"
   ├─ Plug.Ticket.sol [256] — "0x000000006580cEe0D3b8ea5a196F1A038FfD3604"
   └─ Plug.sol [4217] — "0x000000004C26dFdF00334a42652d5880608647Fb"
```
