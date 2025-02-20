# Pool Together Integration

## Overview

Pool Together allows users to deposit tokens into a prize pool to earn interest and potentially win a prize. The prize pool is a no-loss lottery that distributes interest earned from the deposited tokens to a randomly selected winner. Users can deposit tokens into the prize pool, and the deposited tokens are used to generate interest, which is then distributed as prizes to the winners.

## Supporting Documentation

- docs
- [Addresses](https://dev.pooltogether.com/protocol/deployments/)
- [V5 Dev Docs](https://dev.pooltogether.com/)
- [Claim Keeper](https://github.com/GenerationSoftware/pt-v5-autotasks-monorepo/blob/main/packages/prize-claimer/src/cli.ts)
- [Getting Claimable Balances](https://github.com/GenerationSoftware/pooltogether-client-monorepo/blob/e32dbddb4785de712822f8fd1b7ff2dc357dabfd/apps/migrations/src/hooks/useUserV3ClaimableRewards.ts#L81)
- [Calculating APY](https://github.dev/GenerationSoftware/pooltogether-client-monorepo/tree/main/apps)

---


Prize Pools - 1 per chain, disitrbute prizes based on draw
Prize Vaults - n per chain, contribute to Prize Pools. 4626 standard vaults
Prize Claimer - allows anyone to claim prizes on behalf of winners, and earn rewards in doing so
    The external `claimPrizes` function is used to claim prizes for one specific vault and tier combination

Assets are one to one with shares



## Scope
| System                | Name                 | Type       | Implemented | Notes                        |
| :-------------------- | :----------------    | :--------- | :---------: | :----------------------------|
| Prize Vault           | Deposit              | Action     |             |                              |
| Prize Vault           | Withdraw             | Action     |             |                              |
| Prize Vault           | Withdraw Max         | Action     |             |                              |
| Prize Vault           | APY                  | Constraint |             |                              |
| Prize Claimer         | Claim rewards keeper | Action     |             |  Consider later              |
| Prize Vault           | Yield keeper         | Action     |             | Not feasible? Need an indexer|
| Prize Pool            | Draw Auction         | Action     |             |  Consider later              |
 

## Deposit

```javascript [sentence]
Deposit {0<amount:uint256>} {1<token:address>} to {1=>2<vault:address>}
```

Vaults are deployed from the Vault Factory. Functions are then called on each Vault address to enable users to deposit and withdraw.

A list of Vaults can be accessed by reading this contract: 0xa55a74A457D8a24D68DdA0b5d1E0341746d444Bf

```totalVaults() returns int```

```allVaults(uint256 index) returns address```

You can read the number of deployed vaults with totalVaults() and then get the address of each vault by calling allVaults(index). Use multi-call .

Show the Prize Vault yield in the menu of options.

```function deposit(uint256 _assets, address _receiver) external returns (uint256);```

## Withdraw (0xb460af94)

```javascript [sentence]
Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<vault:address>}
```

```withdraw(uint256 _assets, address _receiver, address _owner)```


## Withdraw Max (0xb460af94)

```javascript [sentence]
Withdraw all {0<token:address>} from {0=>1<vault:address>}
```


```withdraw(uint256 _assets, address _receiver, address _owner)```



## Claim Rewards For Others
In order to earn fees on claiming prizes, you will need to:

- Attain a list of all active accounts for the Prize Pool you're checking
- Compute the prizes won (if any) for each account that holds a non-zero TWAB balance
- If the claim fees are sufficiently profitable, then execute a prize claim
- Periodically withdraw the fees you have earned from the Prize Pool


## APY
APY must be calculated from a few different read calls that have to be made.

