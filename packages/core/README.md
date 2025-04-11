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
[v0.4.3]:
├─ Mining Configuration:
│  ├─ Leading Zeros: 4
│  ├─ Total Zeros: 8
│  ├─ Factory: 0x0000000000ffe8b47b3e2130213b802212439497
│  └─ Quick Mode: Yes
│
└─ Contracts:
   ├─ Plug.Assert.sol [256] — "0x000000004495c796bAeC14fddc116a116abA24E4"
   ├─ Plug.Boolean.sol [256] — "0x000000003BDBD95cb81c235016C29051e72Bc07d"
   ├─ Plug.Coercion.sol [256] — "0x00000000BC31055eBF956Fcc852ee8E3B670a3dD"
   ├─ Plug.Database.sol [256] — "0x0000000027b917cAFf88faa8Bd4EFcaa79e721c4"
   ├─ Plug.EVM.sol [256] — "0x00000000036A1FDee391c0A4f91f73fEDdE23FB1"
   ├─ Plug.Factory.sol [256] — "0x00000000E8d0C003F601472E5d94FA7bCCdD5EC2"
   ├─ Plug.Math.sol [256] — "0x00000000e36eE6E4D677d84A9fCE859A4454B017"
   ├─ Plug.Rewards.sol [256] — "0x00000000521807eAa08E92E42B139d87992b0376"
   ├─ Plug.Socket.sol [256] — "0x000000004321f57fFC6649F6b03b8a4FBbFa1EEA"
   ├─ Plug.Ticket.sol [256] — "0x00000000615370671ee8cBe545585D8CFbE7c834"
   ├─ Plug.Token.sol [256] — "0x00000000CB92316643CD22fF3F782f4B75866a34"
   └─ Plug.sol [256] — "0x000000008114350aa81CaE0153A8615e2578F462"
```
