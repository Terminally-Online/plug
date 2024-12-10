# Morpho Integration

## Overview

## Supporting Documentation

- docs
- [Addresses](https://docs.morpho.org/morpho/addresses/)
- [Morpho API](https://blue-api.morpho.org/graphql)
- [Rewards Program](rewards.morpho.org/v1/programs)
- [MetaMorpho Factory Contract](https://etherscan.io/address/0xA9c3D3a366466Fa809d1Ae982Fb2c46E5fC41101)
- [Morpho Contract](https://etherscan.io/address/0xbbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb)
- [Universal Reward Distributor Contract](https://etherscan.io/address/0x330eefa8a787552dc5cad3c3ca644844b1e61ddb)

---

## Contract Interfacing

For each vault specified and returned by the API call we have a `asset` field that reflects the asset that can be borrowed. This asset however does not reflect which tokens can be used as collateral.

## Scope

| System                | Name                    | Type       | Implemented | Notes                                               |
| :-------------------- | :---------------------- | :--------- | :---------- | :-------------------------------------------------- |
| Morpho Vault          | Earn via Vault          | Action     |             |                                                     |
| Morpho Vault          | Withdraw from Vault     | Action     |             |                                                     |
| Morpho Vault          | Withdraw Max from Vault | Action     |             |                                                     |
| :------------         | :---------------------- | :--------- | :---------- | :----                                               |
| Morpho Market         | Supply Collateral       | Action     | 12/09/2024  | Needs approval transaction to the router.           |
| Morpho Market         | Withdraw Collateral     | Action     | 12/09/2024  |                                                     |
| Morpho Market         | Withdraw All Collateral | Action     | 12/09/2024  |                                                     |
| Morpho Market         | Borrow                  | Action     | 12/10/2024  |                                                     |
| Morpho Market         | Repay                   | Action     | 12/10/2024  | Needs approval transaction to the router.           |
| Morpho Market         | Repay in Full           | Action     | 12/10/2024  | Needs approval transaction to the router.           |
| Rewards               | Claim Rewards           | Action     | 12/10/2024  |                                                     |
| Morpho Market         | Health Factor           | Constraint | 12/10/2024  |                                                     |
| Morpho Vault & Market | APY                     | Constraint | 12/09/2024  | Partially implemented. Does not yet support vaults. |

### Earn via Morpho Vault

`deposit()` is called on the Morpho Vault.

```
function deposit(uint256 assets, address receiver) public override returns (uint256 shares) {}
```

NOTE: must complete pre-approval of the Vault with the Vault’s underlying asset token.

Mints shares Vault shares to receiver by depositing exactly amount of underlying tokens.

- MUST emit the Deposit event.
- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the deposit execution, and are accounted for during deposit.
- MUST revert if all of assets cannot be deposited (due to deposit limit being reached, slippage, the user not approving enough underlying tokens to the Vault contract, etc).


"Deposit {0 amount} {1 token} to {1->2 vault}."

### Withdraw from Morpho Vault

`withdraw()` is called on the Morpho Vault.

```
function withdraw(uint256 assets, address receiver, address owner) public override returns (uint256 shares) {}
```

Burns shares from owner and sends exactly assets of underlying tokens to receiver.

- MUST emit the Withdraw event.
- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the withdraw execution, and are accounted for during withdraw.
- MUST revert if all of assets cannot be withdrawn (due to withdrawal limit being reached, slippage, the owner not having enough shares, etc).

NOTE: some implementations will require pre-requesting to the Vault before a withdrawal may be performed. Those methods should be performed separately.

We can use maxWithdraw(owner address) to determine if the amount is valid.


"Withdraw {0 amount} {1 token} from {1->2 vault}."

### Withdraw Max from Morpho Vault

`redeem()` is called on the Morpho Vault by using Redeem.

```
function redeem(uint256 shares, address receiver, address owner) public override returns (uint256 assets) {}
```

Burns exactly shares from owner and sends assets of underlying tokens to receiver. - MUST emit the Withdraw event.

- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the redeem execution, and are accounted for during redeem.
- MUST revert if all of shares cannot be redeemed (due to withdrawal limit being reached, slippage, the owner not having enough shares, etc).

NOTE: some implementations will require pre-requesting to the Vault before a withdrawal may be performed. Those methods should be performed separately.

maxRedeem(owner address) read on vault can be used to find this amount.


"Withdraw all {0 token} from {0->1 vault}."

## Markets

### Supply Collateral

Occurs on the Morpho contract and requires market parameters to direct funds to proper market.

"Supply {0<tokenIn:address>} {1<amountIn:uint256>} as collateral for {1->2 <token:address>}."

Supplies assets or shares on behalf of onBehalf, optionally calling back the caller's onMorphoSupply function with the given data.

Either assets or shares should be zero. Most usecases should rely on assets as an input so the caller is guaranteed to have assets tokens pulled from their balance, but the possibility to mint a specific amount of shares is given for full compatibility and precision.

Plug wll use assets as an input and pass 0 for for shares.

Takes a tuple to define the market to supply assets to:
loanToken(address), collateralToken(address), oracle(address), irm(address). lltv(uint256).

| Input        | Type         | Description                                                                           |
| ------------ | ------------ | ------------------------------------------------------------------------------------- |
| marketParams | MarketParams | The market to supply assets to.                                                       |
| assets       | uint256      | The amount of assets to supply.                                                       |
| shares       | uint256      | The amount of shares to mint.                                                         |
| onBehalf     | address      | The address that will own the increased supply position.                              |
| data         | bytes        | Arbitrary data to pass to the onMorphoSupply callback. Pass empty data if not needed. |

### Withdraw Collateral

"Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<target:string>}."

Occurs on the Morpho contract and requires market parameters to withdraw funds from the proper market.

Either assets or shares should be zero. To withdraw max, pass the shares's balance of onBehalf.

For this function, Plug will use assets and pass 0 for shares.

Takes a tuple to define the market to withdraw assets from:
loanToken(address), collateralToken(address), oracle(address), irm(address). lltv(uint256).

| Input        | Type         | Description                                                                           |
| ------------ | ------------ | ------------------------------------------------------------------------------------- |
| marketParams | MarketParams | The market to supply assets to.                                                       |
| assets       | uint256      | The amount of assets to withdraw.                                                     |
| shares       | uint256      | The amount of shares to burn.                                                         |
| onBehalf     | address      | The address that owns the supply position.                                            |
| receiver     | address      | Arbitrary data to pass to the onMorphoSupply callback. Pass empty data if not needed. |

### Withdraw All Collateral

"Withdraw all collateral from {0<market:string>}."

Occurs on the Morpho contract and requires market parameters to withdraw funds from the proper market.

Either assets or shares should be zero. To withdraw max, pass the shares's balance of onBehalf.

For this function, Plug will use shares and pass 0 for assets.

Takes a tuple to define the market to withdraw assets from:
loanToken(address), collateralToken(address), oracle(address), irm(address). lltv(uint256).

| Input        | Type         | Description                                                                           |
| ------------ | ------------ | ------------------------------------------------------------------------------------- |
| marketParams | MarketParams | The market to supply assets to.                                                       |
| assets       | uint256      | The amount of assets to withdraw.                                                     |
| shares       | uint256      | The amount of shares to burn.                                                         |
| onBehalf     | address      | The address that owns the supply position.                                            |
| receiver     | address      | Arbitrary data to pass to the onMorphoSupply callback. Pass empty data if not needed. |

To find the amount of shares we can use the graphQL api. Example of a users postion

### Borrow from Morpho Market

Borrows `assets` or `shares` on behalf of `onBehalf` and sends the assets to `receiver`.

Either `assets` or `shares` should be zero. Most use cases should rely on `assets` as an input so the caller is guaranteed to borrow `assets` of tokens, but the possibility to mint a specific amount of shares is given for full compatibility and precision.`msg.sender` must be authorized to manage `onBehalf`'s positions.

Borrowing a large amount can revert for overflow.

Borrowing an amount of shares may lead to borrow fewer assets than expected due to slippage. Consider using the `assets` parameter to avoid this.

"Borrow {0 amount} {1 token} against {2 collateral_amount} {2->3 collateral}."

### Repay

"Repay {0<amount:uint256>} {1<token:address>} tokens in {1<market:string>}."

On Morpho contract.

Repays assets or shares on behalf of onBehalf, optionally calling back the caller's onMorphoRepay function with the given data.

Either assets or shares should be zero. To repay max, pass the shares's balance of onBehalf.

For this function, Plug will take an input and pass 0 for shares.

Takes a tuple to define the market to withdraw assets from:
loanToken(address), collateralToken(address), oracle(address), irm(address). lltv(uint256).

| Input        | Type         | Description                                                                          |
| ------------ | ------------ | ------------------------------------------------------------------------------------ |
| marketParams | MarketParams | The market to supply assets to.                                                      |
| assets       | uint256      | The amount of assets to repay.                                                       |
| shares       | uint256      | The amount of shares to burn.                                                        |
| onBehalf     | address      | The address that owns the debt position.                                             |
| data         | bytes        | Arbitrary data to pass to the onMorphoRepay callback. Pass empty data if not needed. |

### Repay in Full

"Repay entire {0<token:address>} position in {1<market:string>}."

Repays assets or shares on behalf of onBehalf, optionally calling back the caller's onMorphoRepay function with the given data.

Either assets or shares should be zero. To repay max, pass the shares's balance of onBehalf.

For this function, Plug will find the share amount and pass 0 for assets.

Takes a tuple to define the market to withdraw assets from:
loanToken(address), collateralToken(address), oracle(address), irm(address). lltv(uint256).

| Input        | Type         | Description                                                                          |
| ------------ | ------------ | ------------------------------------------------------------------------------------ |
| marketParams | MarketParams | The market to supply assets to.                                                      |
| assets       | uint256      | The amount of assets to repay.                                                       |
| shares       | uint256      | The amount of shares to burn.                                                        |
| onBehalf     | address      | The address that owns the debt position.                                             |
| data         | bytes        | Arbitrary data to pass to the onMorphoRepay callback. Pass empty data if not needed. |

Before we are able to call the repay function, we need to know how many shares the user has. We can use the Morpho graphQL endpoint to understand how many shares a user has in a market and then pass that value as for shares in repay.

### Claim Rewards

"Claim rewards"



### Health Factor

### Get APY

"{0 action} APY in {1->2 vault/market} is {3 operator} than {4 threshold}."

