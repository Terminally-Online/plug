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
| Boost               | Action     |             |       |
| Supply              | Action     |             |       |
| Repay               | Action     |             |       |
| Repay With Shares   | Action     |             |       |
| Withdraw            | Action     |             |       |
| Borrow USDC         | Action     |             |       |
| Return USDC         | Action     |             |       |
| APY                 | Constraint |             |       |
| Health Factor       | Constraint |             |       |
| Time to Liquidation | Constraint |             |       |


### Boost (?)
### Supply
### Repay
### Repay With Shares
### Withdraw
### Borrow USDC
### Return USDC
### APY

### Health Factor