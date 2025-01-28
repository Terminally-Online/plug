# Euler Integration

**At the time of writing this spec. sheet there exist insufficient documentation to ensure that all things are as they should be.**

## Overview

At the core base of the protocol Euler is utilizing is ERC-4626 so this is the same standard. Similar to Morpho, Euler utilizes distinct patterns for Markets and Vaults.

In their internal specification they refer to "view contracts" as "lens". It is not lens the platform, but a lens to look into some state of the protocol.

To get the address of the vaults you will utilize ``

## Supporting Documentation

- [Documentation](https://docs.euler.finance)
- [Addresses](https://github.com/euler-xyz/euler-interfaces/tree/master/addresses)
- [Vault Whitepaper](https://github.com/euler-xyz/euler-vault-kit/blob/master/docs/whitepaper.md#composing-vaults)

## Initial Notes & Questions to answer

- How does a [strategy](https://app.euler.finance/strategies?network=base) work on the protocol level?
- Do we want to surface the collateral exposure of a vault to our users? e.g. Their main USDC vault can be borrowed against with AERO as collateral.
- How to get all vaults? Find all create proxy events on the factory?
- I'm not clear how to get the health factor or the APY of a vault
- 

## Scope

| Name              | Type       | Implemented | Notes |
| :---------------- | :--------- | :---------: | :---- |
| Boost             | Action     |             |       |
| Supply            | Action     |             |       |
| Repay             | Action     |             |       |
| Repay With Shares | Action     |             |       |
| Withdraw          | Action     |             |       |
| Borrow USDC       | Action     |             |       |
| Return USDC       | Action     |             |       |
| APY               | Constraint |             |       |
| Health Factor     | Constraint |             |       |


### Boost (?)
### Supply
### Repay
### Repay With Shares
### Withdraw
### Borrow USDC
### Return USDC
### APY

### Health Factor