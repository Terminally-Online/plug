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
- [EthereumV2 Bundler Contract](https://etherscan.io/address/0x4095F064B8d3c3548A3bebfd0Bbfd04750E30077)
- [Rewards API](https://docs.morpho.org/apis/rewards/)

---

## Scope

| System                | Name              | Type       | Implemented | Notes |
| :-------------------- | :---------------- | :--------- | :---------: | :---- |
| Morpho Vault          | Earn via Vault    | Action     | 12/11/2024  |       |
| Morpho Market         | Supply Collateral | Action     | 12/11/2024  |       |
| Morpho Vault & Market | Withdraw          | Action     | 12/04/2024  |       |
| Morpho Vault & Market | Withdraw Max      | Action     | 12/04/2024  |       |
| Morpho Market         | Borrow            | Action     | 12/04/2024  |       |
| Morpho Market         | Repay             | Action     | 12/04/2024  |       |
| Morpho Market         | Repay in Full     | Action     | 12/04/2024  |       |
| Rewards               | Claim Rewards     | Action     | 12/04/2024  |       |
| Morpho Market         | Health Factor     | Constraint | 12/10/2024  |       |
| Morpho Vault & Market | APY               | Constraint | 12/11/2024  |       |

## Earn

```javascript [sentence]
Earn by depositing {0<amount:uint256>} {1<token:address>} to {1=>2<vault:address>}.
```

```solidity [transaction batch]
token.approve(address vault, uint256 amount)
vault.deposit(uint256 assets, address receiver)
```

## Supply Collateral

```javascript [sentence]
Supply {0<amount:uint256>} {1<token:address>} as collateral to {1=>2<market:string>}.
```

```solidity [transaction batch]
erc20.approve(address market, uint256 amount)
market.supplyCollateral(
    tuple marketParams(
        address loanToken,
        address collateralToken,
        address oracle,
        address irm,
        uint256 lltv
    ),
    uint256 assets,
    address onBehalf,
    bytes data
)
```

## Withdraw

```javascript [sentence]
Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<target:strng>}.
```

When it is a vault:

```solidity [transaction batch]
vault.withdraw(uint256 assets, address receiver, address owner)
```

When it is a market:

```solidity [transaction batch]
market.withdraw(
    tuple marketParams(
        address loanToken,
        address collateralToken,
        address oracle,
        address irm,
        uint256 lltv
    ),
    uint256 assets,
    uint256 shares, // NOTE: This will be zero because we are defining the amount of assets to withdraw.
    address onBehalf,
    address receiver
)
```

## Withdraw All

```javascript [sentence]
Withdraw all {0<token:address>} from {0=>1<target:string>}
```

When it is a vault:

```solidity [transaction batch]
vault.redeem(uint256 shares, address receiver, address owner)
```

When it is a market:

Use the `position()` read to determine the shares owned by the address.

```solidity [transaction batch]
market.withdraw(
    tuple marketParams(
        address loanToken,
        address collateralToken,
        address oracle,
        address irm,
        uint256 lltv
    ),
    uint256 assets, // NOTE: This will be zero because we are defining the amount of shares to withdraw.
    uint256 shares,
    address onBehalf,
    address receiver
)
```

## Borrow

```javascript [sentence]
Borrow {0<amount:uint256>} {1<token:address>} from {1=>2<market:string>}
```

```solidity [transaction batch]
market.borrow(
    tuple marketParams(
        address loanToken,
        address collateralToken,
        address oracle,
        address irm,
        uint256 lltv
    ),
    uint256 assets,
    uint256 shares, // NOTE: This will be zero because we are defining the amount of assets to borrow.
    address onBehalf,
    address receiver
)
```

## Repay

```javascript [sentence]
Repay {0<amount:uint256>} {1<token:address>} in {1=>2<market:string>}
```

```solidity [transaction batch]
loanToken.approve(address market, uint256 amount)
market.repay(
    tuple marketParams(
        address loanToken,
        address collateralToken,
        address oracle,
        address irm,
        uint256 lltv
    ),
    uint256 assets,
    uint256 shares, // NOTE: This will be zero because we will repay assets
    address onBehalf,
    bytes data
)
```

# Repay All

```javascript [sentence]
Repay all {0<token:address>} in {0=>1<market:string>}
```

```solidity [transaction batch]
loanToken.approve(address market, uint256 amount)
market.repay(
    tuple marketParams(
        address loanToken,
        address collateralToken,
        address oracle,
        address irm,
        uint256 lltv
    ),
    uint256 assets,
    uint256 shares, // NOTE: This will be zero because we need to use assets so that we can approve the correct amount
    address onBehalf,
    bytes data
)
```

## Claim Rewards

```javascript [sentence]
Claim all reward distributions.
```

```solidity [transaction batch]
distributor.claim(
    address account,
    address reward,
    uint256 claimableAmount,
    bytes32[] proof
)
```

## Health Factor

A bunch of math here. Just go check the code.

## APY

Needs to be implemented for vaults.
