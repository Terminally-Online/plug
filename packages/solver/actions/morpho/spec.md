# Morpho Integration

## Overview

## Supporting Documentation

- docs
- [Addresses](https://docs.morpho.org/morpho/addresses/)
- [Morpho API](https://blue-api.morpho.org/graphql)
- [Rewards Program](rewards.morpho.org/v1/programs)

---

## Contract Interfacing

For each vault specified and returned by the API call we have a `asset` field that reflects the asset that can be borrowed. This asset however does not reflect which tokens can be used as collateral.

To determine which vaults we support you can make the call to their API with the following query:

```graphql
query ExampleQuery($first: Int) {
  vaults(first: $first, where: { chainId_in: [1], whitelisted: true }) {
    items {
      address
      symbol
      name
      whitelisted
      metadata {
        image
      }
      asset {
        address
        decimals
        name
        symbol
        logoURI
      }
      state {
        apy
        netApy
        allocation {
          enabled
          market {
            collateralAsset {
              address
              name
              symbol
              logoURI
              decimals
            }
          }
        }
        dailyApy
        dailyNetApy
        weeklyApy
        weeklyNetApy
        monthlyApy
        monthlyNetApy
        quarterlyApy
        quarterlyNetApy
        yearlyApy
        yearlyNetApy
        allTimeApy
        allTimeNetApy
      }
      liquidity {
        underlying
        usd
      }
    }
  }
}
```

## Scope

| System        | Name                    | Type       | Implemented | Notes |
| :------------ | :---------------------- | :--------- | :---------- | :---- |
| Morpho Vault  | Earn via Vault          | Action     |             |       |
| Morpho Vault  | Withdraw from Vault     | Action     |             |       |
| Morpho Vault  | Withdraw Max from Vault | Action     |             |       |
| Morpho Market | Supply Collateral       | Action     |             |       |
| Morpho Market | Withdraw Collateral     | Action     |             |       |
| Morpho Market | Withdraw All Collateral | Action     |             |       |
| Morpho Market | Borrow                  | Action     |             |       |
| Morpho Market | Repay                   | Action     |             |       |
| Morpho Market | Repay in Full           | Action     |             |       |
| Rewards       | Claim Rewards           | Action     |             |       |
| Morpho Market | Health Factor           | Constraint |             |       |
| Morpho Vault  | Vault APY               | Constraint |             |       |
| Morpho Market | Vault APY               | Constraint |             |       |

### Contracts

| Name                         | Address                                    | Desc                   |
| :--------------------------- | :----------------------------------------- | :--------------------- |
| MetaMorpho Factory           | 0xA9c3D3a366466Fa809d1Ae982Fb2c46E5fC41101 | Deploys Morpho Vaults  |
| Morpho Market Factory        | 0xbbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb | Deploys Morpho Markets |
| Universal Reward Distributor | 0x330eefa8a787552dc5cad3c3ca644844b1e61ddb | Claim rewards          |

### Earn via Morpho Vault

deposit() is called on the Morpho Vault.

Mints shares Vault shares to receiver by depositing exactly amount of underlying tokens.

- MUST emit the Deposit event.
- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the deposit execution, and are accounted for during deposit.
- MUST revert if all of assets cannot be deposited (due to deposit limit being reached, slippage, the user not approving enough underlying tokens to the Vault contract, etc).

NOTE: most implementations will require pre-approval of the Vault with the Vault’s underlying asset token.

"Deposit {0 amount} {1 token} to {1->2 vault}."

### Withdraw from Morpho Vault

This is called on the Morpho Vault.

Burns shares from owner and sends exactly assets of underlying tokens to receiver.

- MUST emit the Withdraw event.
- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the withdraw execution, and are accounted for during withdraw.
- MUST revert if all of assets cannot be withdrawn (due to withdrawal limit being reached, slippage, the owner not having enough shares, etc).

NOTE: some implementations will require pre-requesting to the Vault before a withdrawal may be performed. Those methods should be performed separately.

"Withdraw {0 amount} {1 token} from {1->2 vault}."

### Withdraw Max from Morpho Vault

This is called on the Morpho Vault by using Redeem.

Burns exactly shares from owner and sends assets of underlying tokens to receiver. - MUST emit the Withdraw event.

- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the redeem execution, and are accounted for during redeem.
- MUST revert if all of shares cannot be redeemed (due to withdrawal limit being reached, slippage, the owner not having enough shares, etc).

NOTE: some implementations will require pre-requesting to the Vault before a withdrawal may be performed. Those methods should be performed separately.

maxRedeem() read on vault can be used to find this amount.

"Withdraw all {0 token} from {0->1 vault}."

### Borrow from Morpho Market

"Borrow {0 amount} {1 token} against {2 collateral_amount} {2->3 collateral}."

### Get Vault APY

"Deposit APY for {0 token} in {0->1 vault} is {2 direction} than {3 threshold}."

### Get Specific Asset Borrow APY

"Borrow APY for {0 token} with {0->1 collateral} collateral is {2 direction} than {3 threshold}."
