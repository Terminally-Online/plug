# Yearn Ecosystem Integration 

## Overview
The Yearn ecosystem consists of three main systems:
1. V3 Vaults - Core yield-generating vaults (ERC-4626)
2. veYFI - Yield boosting and governance system
3. yCRV - CRV liquidity and yield optimization

Integrating a new contract requires two sets of action files, a generalized action and protocol specific action. Generalized Action files define and gather inputs which are then passed to Protocol Specific Action files which reshape the parameters to prepare the transaction.

## Supporting Documentation

- [Yearn V3 Developer Docs](https://docs.yearn.fi/developers/v3/overview)
- [Yearn Registry Contract](https://etherscan.io/address/0xff31A1B020c868F6eA3f61Eb953344920EeCA3af)
- [ERC-4626 Standard](https://eips.ethereum.org/EIPS/eip-4626)
- [yDaemon API](https://ydaemon.yearn.farm/docs/intro)

## Contract Interfacing - V3 Vaults

Yearn V3 uses the Registry contract (`0xff31A1B020c868F6eA3f61Eb953344920EeCA3af`) as the primary entry point for discovering vaults. Each vault follows the ERC-4626 standard for deposits, withdrawals, and share calculations. To find vaults for a specific asset, use the `getEndorsedVaults(asset)` function which returns an array of vault addresses supporting that asset.

### ERC-4626 Specifications

## Scope

| System | Name | Type | Source | Supported | Implemented | Notes |
| :--- | :--- | :--- | :--- | :---: | :--- | :--- |
| V3 | Deposit Assets | Action | Vault (ERC-4626) |  |  | Requires approval |
| V3 | Withdraw Assets | Action | Vault (ERC-4626) |  |  | Check maxWithdraw |
| V3 | Stake Shares | Action | Vault (ERC-4626) |  |  | |
| V3 | Vault Discovery | Constraint | Registry |  |  | Find endorsed vaults |
| V3 | APY Tracking | Constraint | yDaemon |  |  | Historical + current |
| V3 | Risk Assessment | Constraint | yDaemon |  |  | Strategy risk + TVL |
| veYFI | Lock YFI | Action | veYFI Contract |  |  | 1w to 4y lock |
| veYFI | Exit Lock Early | Action | veYFI Contract |  |  | With penalty |
| veYFI | Calculate Boost | Constraint | veYFI Contract |  |  | 1x to 10x multiplier |
| yCRV | Deposit CRV | Action | yCRV Contract |  |  | Get yCRV tokens |
| yCRV | Stake yCRV | Action | YearnBoostedStaker |  |  | Earn crvUSD |
| yCRV | Auto-compound yCRV | Action | yvyCRV Vault |  |  | Compound yields |

## Core Actions - V3 Vaults

### Deposit
Deposits assets into a Yearn vault in exchange for vault shares.

- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `deposit(uint256 assets, address receiver)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| assets | uint256 | Amount to deposit | Must be <= maxDeposit(receiver) |
| receiver | address | Recipient of shares | Optional, defaults to sender |

### Withdraw
Withdraws underlying assets from a vault by burning shares.

- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `withdraw(uint256 assets, address receiver, address owner)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| assets | uint256 | Amount to withdraw | Must be <= maxWithdraw(owner) |
| receiver | address | Recipient of assets | |
| owner | address | Owner of shares | Usually msg.sender |

### Stake Shares
Stake vault shares in gauge for additional rewards.

- **Contract:** Retrieved via Registry's `getGaugeForVault(vault)`
- **Function:** `deposit(uint256 amount, address receiver)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of vault shares to stake | Must be <= balance |
| receiver | address | Recipient of gauge position | Optional, defaults to sender |


## Constraints - V3 Vaults 

#### Max Deposit
- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `maxDeposit(address receiver)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| receiver | address | Account to check deposit limit | Returns max assets that can be deposited |

#### Max Withdraw
- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `maxWithdraw(address owner)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| owner | address | Account to check withdraw limit | Returns max assets that can be withdrawn |

#### Max Redeem
- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `maxRedeem(address owner)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| owner | address | Account to check redeem limit | Returns max shares that can be redeemed |

### APY Tracking
Retrieves historical and current APY data for vaults.

- **Source:** yDaemon API
- **Endpoint:** `GET /api/v1/vaults/apy/{chain}/{address}`

| Field | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| type | string | APY calculation type | "net" or "gross" |
| composite | boolean | If vault has multiple yield sources | true/false |
| points | array | Historical APY data points | timestamp + value pairs |
| weekAgo | number | Rolling 7-day APY | Percentage value |
| monthAgo | number | Rolling 30-day APY | Percentage value |
| inception | number | Since vault creation APY | Percentage value |

### Risk Assessment
Retrieves risk metrics and scoring for vaults.

- **Source:** yDaemon API
- **Endpoint:** `GET /api/v1/vaults/risk/{chain}/{address}`

| Field | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| tvlImpact | number | TVL risk score | 0-5 scale |
| auditScore | number | Smart contract audit rating | 0-5 scale |
| complexityScore | number | Protocol complexity rating | 0-5 scale |
| protocolSafety | number | Protocol safety score | 0-5 scale |
| teamKnowledge | number | Team expertise rating | 0-5 scale |
| testingScore | number | Testing coverage score | 0-5 scale |
| totalScore | number | Aggregate risk rating | 0-10 scale |
| details | object | Detailed risk explanations | Text descriptions |

### Boost Calculation
Calculates the boost multiplier for a user's vault deposits based on their veYFI balance.

- **Contract:** Retrieved via Registry's `getGaugeForVault(vault)`
- **Function:** `getBoost(address user)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| user | address | User address to check boost | Account with veYFI balance |
| returns | number | Boost multiplier | 1.0x to 2.5x range |

## Core Actions - veYFI

### Lock YFI
Creates a veYFI position by locking YFI tokens for boosted vault yields.

- **Contract:** (VeYFI) `0x90c1f9220d90d3966fbee24045edd73e1d588ad5`
- **Function:** `createLock(uint256 amount, uint256 duration)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of YFI to lock | Must approve first |
| duration | uint256 | Lock duration in seconds | Max 4 years |
| returns | uint256 | veYFI token ID | NFT representing lock |

### Exit Lock Early
Withdraw locked YFI tokens before the lock expiration by paying an early exit penalty.

- **Contract:** (VeYFI) `0x90c1f9220d90d3966fbee24045edd73e1d588ad5`
- **Function:** `withdrawEarly(uint256 tokenId)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| tokenId | uint256 | veYFI token ID to exit | NFT ID of lock position |
| returns | uint256 | YFI amount returned | After penalty deduction |

### Early Exit Penalty Calculation
Calculates the amount of YFI tokens returned when exiting a lock position early.

- **Contract:** (VeYFI) `0x90c1f9220d90d3966fbee24045edd73e1d588ad5`
- **Function:** `calculateEarlyWithdrawAmount(uint256 tokenId)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| tokenId | uint256 | veYFI token ID to check | NFT ID of lock position |
| returns | uint256 | YFI amount returned | After penalty deduction |

## Core Actions - yCRV

### Deposit CRV
Convert CRV tokens to yCRV for enhanced yields.

- **Contract:** 
- **Function:** `deposit(uint256 amount)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of CRV to deposit | Must approve first |
| returns | uint256 | yCRV amount received | Exchange rate varies |

### Stake yCRV
Stake yCRV tokens in YearnBoostedStaker to earn crvUSD rewards.

- **Contract:** (YearnBoostedStaker) `0xE9A115b77A1057C918F997c32663FdcE24FB873f`
- **Function:** `stake(uint256 amount)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of yCRV to stake | Must approve first |
| returns | uint256 | Receipt token amount | Represents staked position |

### Auto-compound yCRV
Deposit yCRV into yvyCRV vault for auto-compounding yields.

- **Contract:** (YearnCRVVault) `0x27B5739e22ad9033bcBf192059122d163b60349D`
- **Function:** `deposit(uint256 amount)`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of yCRV to deposit | Must approve first |
| returns | uint256 | yvyCRV amount received | Exchange rate varies |

### Claim yCRV Rewards
Claim accumulated crvUSD rewards from staked yCRV positions.

- **Contract:** (YearnBoostedStaker) `0xE9A115b77A1057C918F997c32663FdcE24FB873f`- **Function:** `getReward()`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| returns | uint256 | Amount of crvUSD claimed | Rewards since last claim |

