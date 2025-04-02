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
[v0.4.2]:
├─ Mining Configuration:
│  ├─ Leading Zeros: 4
│  ├─ Total Zeros: 8
│  ├─ Factory: 0x0000000000ffe8b47b3e2130213b802212439497
│  └─ Quick Mode: Yes
│
└─ Contracts:
   ├─ Plug.Assert.sol [256] — "0x0000000035F357c2f503DA504B0B7dBDC534539C"
   ├─ Plug.Boolean.sol [256] — "0x000000000cAF2bfe5bbe3F7BBD5e70aCDCA6D1FE"
   ├─ Plug.Coercion.sol [256] — "0x0000000034E10d8cA2843b56453A0A373023b792"
   ├─ Plug.Database.sol [256] — "0x000000002Cea9833a9D2dc60e35846Cbb7fC1442"
   ├─ Plug.EVM.sol [4217] — "0x000000001B2147E34d7A00925B016e6bC697C9DC"
   ├─ Plug.Factory.sol [256] — "0x00000000026dF9927AE0fB3CFB5f4ce0298f6C45"
   ├─ Plug.Math.sol [65536] — "0x0000000000269af70428b90fEC44d94f56b43d21"
   ├─ Plug.Socket.sol [256] — "0x00000000906bb1a5fe6527c051A4C3b1c4595a8a"
   ├─ Plug.Ticket.sol [256] — "0x000000003525F8830Dbf2eaAdBCEC33cbFC3E79e"
   ├─ Plug.Token.sol [256] — "0x000000003525F8830Dbf2eaAdBCEC33cbFC3E79e"
   └─ Plug.sol [4217] — "0x000000004C26dFdF00334a42652d5880608647Fb"
```
