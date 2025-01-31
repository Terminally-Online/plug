# Euler Integration

**At the time of writing this spec. sheet there exist insufficient documentation to ensure that all things are as they should be.**

## Overview

At the core base of the protocol Euler is utilizing is ERC-4626 so this is the same standard. Similar to Morpho, Euler utilizes distinct patterns for Markets and Vaults.

In their internal specification they refer to "view contracts" as "lens". It is not lens the platform, but a lens to look into some state of the protocol.

To get the address of the vaults you will utilize ``

## Supporting Documentation

- [Documentation](https://docs.euler.finance)
- [Addresses](https://github.com/euler-xyz/euler-interfaces/tree/master/addresses)
- [Vault Whitepaper](https://github.com/euler-xyz/euler-vault-kit/blob/master/docs/whitepaper.md)
- [EVC Whitepaper](https://evc.wtf/docs/whitepaper/)

## Initial Notes & Questions to answer

- How does a [strategy](https://app.euler.finance/strategies?network=base) work on the protocol level?
  - Looping collateral and borrow, no protocol nuance
  - Have to enable a vault as collateral

- How can we get the health factor of a user's account?
  - Account Lens
  - Health Score, Time to Liquidation [AccountLens](https://basescan.org/address/0x40c1DbD5855bFbCDd3844C4327777FD1c5E039eb#readContract)

- How can we get the APY of a vault, both intrinsic and with rewards?
  - Util Lens

- Does the virtual account aspect need any management?
  - Getting grouping of sub accounts for the user
  - The EVC has a method to get sub accounts from a top level account

- Potentially use [Euler Swap API](https://github.com/euler-xyz/euler-swap-api)?

- Getting Euler base markets // Recommended vaults
  - Governed perspective to get the vaults we show [GovernedPerspective](https://basescan.org/address/0xafc8545c49df2c8216305922d9753bf60bf8c14a#readContract)


## Scope

| Name                | Type       | Implemented | Notes |
| :----------------   | :--------- | :---------: | :---- |
| Supply              | Action     |             |       |
| Withdraw            | Action     |             |       |
| Withdraw All        | Action     |             |       |
| Borrow              | Action     |             |       |
| Repay               | Action     |             |       |
| Repay Max           | Action     |             |       |
| APY                 | Constraint |             |       |
| Health Factor       | Constraint |             |       |
| Time to Liquidation | Constraint |             |       |

### Supply

```function deposit(uint256 amount, address receiver) returns (uint256) {}```
Receiver is the address of the user who will receive the shares.

```javascript [sentence]
    Deposit {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>}
```

### Withdraw

Owner is the address of the user who has the tokens, receiver is the address of the user who will receive the tokens.

```function withdraw(uint256 amount, address receiver, address owner) returns (uint256)```

```javascript [sentence]
    Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}
```

### Withdraw All

```javascript [sentence]
    Withdraw {0<token:address:uint8>} from {0=>1<vault:address>}
```

### Borrow

```function borrow(uint256 amount, address receiver) returns (uint256) {}```

```javascript [sentence]
    Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address}
```

### Repay

```function repay(uint256 amount, address receiver) returns (uint256) {}```

```javascript [sentence]
    Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>}
```

### Repay Max (With Shares)

Utilizes collateral token shares to repay the debt.
```function repayWithShares(uint256 amount, address receiver) returns (uint256 shares, uint256 debt) {}```

```javascript [sentence]
    Repay {0<amount:float>} {1<token:address:uint8>} with shares to {1=>2<vault:address>}
```

### APY

Utilizes the Util lens to get the APY of a vault.

```computeAPYs(uint256 borrowSPY, uint256 cash, uint256 borrows, uint256 interestFee) returns (uint256 borrowAPY, uint256 supplyAPY) {}```

- borrowSPY -> ?
- cash -> LiabilityVault.cash()
- borrows -> LiabilityVault.totalBorrows()
- interestFee -> LiabilityVault.interestFee()

```javascript [sentence]
    "{0<action:int8>} APY in {1<target:string>} is {2<operator:int8>} than {3<threshold:float>}%."
```

### Health Factor

I believe this can be calculated by getting the base value of the asset being borrowed vs the CollateralValueLiquidation

```javascript [sentence]
    "Health factor in {0<vault:string>} is {1<operator:int8>} than {2<threshold:float>}."
```


### Time to Liquidation

Util lens contract read

To get the liabilityValue I believe we have to make a call to the oracle lens or oracle perspective contracts.
It's also currently unclear to me how we know what collaterals are being used in this specific instance with the way that euler uses virtual accounts instead of the base address in some cases.

```function calculateTimeToLiquidation(address liabilityVault,uint256 liabilityValue,address[] memory collaterals,uint256[] memory collateralValues) external view returns (int256)```

```javascript [sentence]
    Time to liquidation of {0<vault:address>} is {1<operator:int8>} than {2<threshold:float>}
```


TODO:

- reads
  - figure out multicall - @chance help ? - pricy api
  - Implement reads and test against my position 
  - health factor calculation 
  - APY read needs BorrowSPY 
  - time to liquidation  
    - needs LiabilityValue
    - collateral question


- options
  - dependent on reads

- handler
  - unmarshall the inputs

- plugs
  - put it all together