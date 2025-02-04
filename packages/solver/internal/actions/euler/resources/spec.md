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
| Earn                | Action     |             |       |
| Deposit Collateral  | Action     |             |       |
| Withdraw            | Action     |             |       |
| Withdraw All        | Action     |             |       |
| Borrow              | Action     |             |       |
| Repay               | Action     |             |       |
| Repay Max           | Action     |             |       |
| APY                 | Constraint |             |       |
| Health Factor       | Constraint |             |       |
| Time to Liquidation | Constraint |             |       |

### Earn

```function earn(uint256 amount, address receiver) returns (uint256) {}```

```javascript [sentence]
    Earn by depositing {0<amount:float>} {1<token:address>} to {1=>2<vault:address>}.
```

### Deposit Collateral

eVault
```function deposit(uint256 amount, address receiver) returns (uint256) {}```
evc
```function enableCollateral(address account,address vault)```
Receiver is the address of the user who will receive the shares.

```javascript [sentence]
    Deposit {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} as collateral.
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

## Repay Max

Being descoped as it's not a great UX allegedly.

### Repay With Shares

Visualizing the repaying with shares by mapping all vaults a user has collateral in and letting them pick between those is a bit of a leap right now. It can be done, but we'd struggle to track balances of all collateral tokens that a user has.

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

Can be calculated by taking the value of the collateral, dividing it by the value of the debt, and multiplying by the collateral factor.

```javascript [sentence]
    "Health factor in {0<vault:string>} is {1<operator:int8>} than {2<threshold:float>}."
```

- eurc has a lltv of .9 and weth has a lltv of .85
- I have .5 eurc and .001 weth deposited, borrowing 0.4 usdc
- my health factor is currently 6.97
- EURC coll value = 0.5 * $1 * 0.9 = 0.45
- WETH coll value = 0.001 * $2740 * 0.85 = 2.3256
- (2.3256 + 0.45) / total value borrowed (0.4) = 6.97 which tracks

### Time to Liquidation

Util lens contract read

To get the liabilityValue I believe we have to make a call to the oracle lens or oracle perspective contracts.
It's also currently unclear to me how we know what collaterals are being used in this specific instance with the way that euler uses virtual accounts instead of the base address in some cases.

```function calculateTimeToLiquidation(address liabilityVault,uint256 liabilityValue,address[] memory collaterals,uint256[] memory collateralValues) external view returns (int256)```

```javascript [sentence]
    Time to liquidation of {0<vault:address>} is {1<operator:int8>} than {2<threshold:float>}
```

Questions For Euler

Last time we talked, I thought we would be able to avoid having to do any management of sub-accounts, but after running some transactions and observing logs I'm not able to figure out how exactly the sub account address is generated. Is it  done on chain or off, and how we can track that sub account 0 has a debt position and we need to move to the next sub account to take on another position. [Answered]
    - Is there a registry of sub accounts that we can query to get the sub accounts of a user? [Nope]
    - Assumption: the only way to interact with the protocol is through feeding in the onBehalfOf address that will be the sub account via the call, batch and permit functions. [Yes]
    - What is the way that you guys sequentially assign sub accounts? I'm assuming you guys cast the address to a uint160 and then XOR with the index? Would love to stay consistent here to allow for easy back and forth between your app.
    - 

Questions For us

- Implementing sub account selection for every action is going to be a bitch, but we could do it through the right type of complex options.

- It seems like they default all actions to an empty sub-account, depositing defaults to an empty sub-account, borrowing does not work at all if you have an open position.

- So, we have to either offer users the ability to select their sub accounts via a complex action, or we have to figure out a way to intelligently select them.
  - An intelligent option doesn't seem doable as a user can get into complicated states with multiple positions.
  - E.g. A user has 3 borrowing positions of USDC in 3 sub accounts. If they hit our repay action, which are they referring to? If they want to add more colateral, which borrowing position is it going to be collateral for?
  - If we show sub-accounts, we can have those be an accordion in the positions column, as well as show the net value in a subaccount when selecting a sub account in a sentence.

``` golang
// GenerateSubAccount creates a new sub-account address from owner address and accountId
func GenerateSubAccount(owner common.Address, accountId uint8) common.Address {
    // Convert owner address to big.Int for bitwise operations
    ownerInt := new(big.Int).SetBytes(owner[:])
    
    // Create big.Int for accountId
    accountIdInt := new(big.Int).SetUint64(uint64(accountId))
    
    // XOR operation
    result := new(big.Int).Xor(ownerInt, accountIdInt)
    
    // Convert back to address
    var subAccount common.Address
    resultBytes := result.Bytes()
    
    // Ensure proper padding to 20 bytes (address length)
    copy(subAccount[20-len(resultBytes):], resultBytes)
    
    return subAccount
}

// GetAccountId recovers the accountId from a sub-account address
func GetAccountId(owner, subAccount common.Address) uint8 {
    // XOR the addresses and take the last byte
    result := new(big.Int).Xor(
        new(big.Int).SetBytes(owner[:]),
        new(big.Int).SetBytes(subAccount[:]),
    )
    
    return uint8(result.Uint64())
}

// Example usage
func main() {
    // Example owner address (you would normally parse this from a hex string)
    owner := common.HexToAddress("0x1234567890123456789012345678901234567890")
    
    // Generate first few sub-accounts
    for i := uint8(0); i < 5; i++ {
        subAccount := GenerateSubAccount(owner, i)
        fmt.Printf("Sub-account %d: %s\n", i, subAccount.Hex())
        
        // Verify we can recover the account ID
        recoveredId := GetAccountId(owner, subAccount)
        fmt.Printf("Recovered ID: %d\n\n", recoveredId)
    }
}
```
