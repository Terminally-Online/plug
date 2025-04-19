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
[v0.5.0]:
├─ Mining Configuration:
│  ├─ Leading Zeros: 4
│  ├─ Total Zeros: 8
│  ├─ Factory: 0x0000000000ffe8b47b3e2130213b802212439497
│  └─ Quick Mode: Yes
│
└─ Contracts:
   ├─ Plug.Assert.sol [256] — "0x00000000F6BC829aB754e492c356b55afA158cE6"
   ├─ Plug.Boolean.sol [256] — "0x000000002992e33AC9B870aAF734d71638543EdB"
   ├─ Plug.Coercion.sol [256] — "0x0000000072f0b1700faEb2db53328E6343CF6F92"
   ├─ Plug.Database.sol [4217] — "0x0000000007189742A4a2ED3F90008a6B7FAedc21"
   ├─ Plug.EVM.sol [256] — "0x00000000bb7a076fd9606836A7D3394A078Bc486"
   ├─ Plug.Factory.sol [256] — "0x0000000035F767aB09BeE35323D2405290377873"
   ├─ Plug.Math.sol [256] — "0x00000000b4117A3E87156756c0C6588766DED50F"
   ├─ Plug.Rewards.sol [256] — "0x00000000b7E44f782B239B9f710D9c64aF2FD3DC"
   ├─ Plug.Socket.sol [256] — "0x00000000Be03b0b5ebBA6b138aB72AEE1097Ea20"
   ├─ Plug.Ticket.sol [256] — "0x0000000017e713811e77c627ADdd3f94Bf5218eD"
   ├─ Plug.Token.sol [256] — "0x00000000A3CFFc9825D7C095167c91FDfc1D565C"
   └─ Plug.sol [256] — "0x000000007f3fA8e1Bdca1Adecd8528f0D63a3FE9"
```
