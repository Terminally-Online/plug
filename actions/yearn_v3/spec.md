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

The full list of underlying assets supported by Yearn can be found with the yDameon API.

Vaults are discoverable based on which assets can be deposited to them.

Yearn V3 uses the Registry contract (`0xff31A1B020c868F6eA3f61Eb953344920EeCA3af`) as the primary entry point for discovering vaults. Each vault follows the ERC-4626 standard for deposits, withdrawals, and share calculations. To find vaults for a specific asset, use the `getEndorsedVaults(asset)` function which returns an array of vault addresses supporting that asset.

Users can engage with vaults by depositing assets for shares, staking shares, and withdrawing assets.

Users may determine if they want to enter a vault based on APY which can be retried via the yDameon API.

Users may may decided to enter or exit vaults when the maximum avialable space or maximum withdrawable amount reaches a certain level. These functions exist on the vault contracts.

The Yearn ecosystem enables boosted yield to YFI stakers. Users may determine to lock YFI for up to 4 years to increase their yields on Yearn vaults. Early penalties are applied for withdrawing from a lock early.

Yearn and Curve are closely related and this integration includes specs related to yCRV, a mechanism that allows users of Curve and Yearn to earn increased yield. Users can convert their CRV by depositing to the yCRV (?) contract and then can auto-compound their yield by depositing their yCRV to the yvyCRV (?) contract.

### ERC-4626 Specifications

## Scope

| System | Name | Type | Source | Supported | Implemented | Notes |
| :--- | :--- | :--- | :--- | :---: | :--- | :--- |
| V3 | Deposit Assets | Action | Vault (ERC-4626) |  |  | Requires approval |
| V3 | Withdraw Assets | Action | Vault (ERC-4626) |  |  | Check maxWithdraw |
| V3 | Stake Shares | Action | Vault (ERC-4626) |  |  | |
| V3 | APY Tracking | Constraint | yDaemon |  |  | Historical + current |
| yCRV | Deposit CRV | Action | yCRV Contract |  |  | Get yCRV tokens |
| yCRV | Stake yCRV | Action | YearnBoostedStaker |  |  | Earn crvUSD |
| yCRV | Auto-compound yCRV | Action | yvyCRV Vault |  |  | Compound yields |
| yCRV | Claim yCRV | Action | Reward Distributor |  |  | Collect yield |
| yCRV | Claimable yCRV | Constraint | Reward Distributor |  |  | Collect yield |



## Core Actions - V3 Vaults

### Deposit
Deposits assets into a Yearn vault in exchange for vault shares.

- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `deposit(uint256 assets, address receiver)`
- **Sentence:**
  - Deposit AMOUNT ASSET into VAULT
  - `Deposit {0} {1} into {1=>2}.`
  - Deposit 1000 USDC into USDC-A yVault


| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| assets | uint256 | Amount to deposit | Must be <= maxDeposit(receiver) |
| receiver | address | Recipient of shares | Optional, defaults to sender |

### Withdraw
Withdraws underlying assets from a vault by burning shares.

- **Contract:** Retrieved via Registry's `getEndorsedVaults(asset)`
- **Function:** `withdraw(uint256 assets, address receiver, address owner)`
- **Sentence:**
  - Withdraw AMOUNT ASSET into VAULT
  - `Withdraw {0} {1} from {1=>2}.`
  - Withdraw 1000 USDC from USDC-A yVault

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| assets | uint256 | Amount to withdraw | Must be <= maxWithdraw(owner) |
| receiver | address | Recipient of assets | |
| owner | address | Owner of shares | Usually msg.sender |

### Stake Shares
Stake vault shares in gauge for additional rewards.

- **Contract:** Retrieved via Registry's `getGaugeForVault(vault)`
- **Function:** `deposit(uint256 amount, address receiver)`
- - **Sentence:**
  - Stake AMOUNT shares into VAULT_GUAGE_CONTRACT.
  - `Stake {0} shares into {1}.`
  - Stake 873.4 shares into USDC-A yVault

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of vault shares to stake | Must be <= balance |
| receiver | address | Recipient of gauge position | Optional, defaults to sender |


## Constraints - V3 Vaults 

### APY Tracking
Retrieves historical and current APY data for vaults.

- **Source:** yDaemon API
- **Endpoint:** `GET /api/v1/vaults/apy/{chain}/{address}`
- - **Sentence:**
  - If VAULT apy is DIRECTION than AMOUNT.
  - `If {0} apy is {1} than {2}.`
  - If USCD-A yVault is Greater Than 4%

| Field | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| type | string | APY calculation type | `net` or `gross` |
| composite | boolean | If vault has multiple yield sources | true/false |
| points | array | Historical APY data points | timestamp + value pairs |
| weekAgo | number | Rolling 7-day APY | Percentage value |
| monthAgo | number | Rolling 30-day APY | Percentage value |
| inception | number | Since vault creation APY | Percentage value |

## Core Actions - yCRV

### Deposit CRV
Convert CRV tokens to yCRV for enhanced yields.

- **Contract:** 
- **Function:** `deposit(uint256 amount)`
- **Sentence:**
  - Deposit AMOUNT crv.
  - `Deposit {0} crv.`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of CRV to deposit | Must approve first |
| returns | uint256 | yCRV amount received | Exchange rate varies |

### Stake yCRV
Stake yCRV tokens in YearnBoostedStaker to earn crvUSD rewards.

- **Contract:** (YearnBoostedStaker) `0xE9A115b77A1057C918F997c32663FdcE24FB873f`
- **Function:** `stake(uint256 amount)`
- **Sentence:**
  - Stake AMOUNT ycrv for crvUSD.
  - `Stake {0} ycrv.`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of yCRV to stake | Must approve first |
| returns | uint256 | Receipt token amount | Represents staked position |

### Auto-compound yCRV
Deposit yCRV into yvyCRV vault for auto-compounding yields.

- **Contract:** (YearnCRVVault) `0x27B5739e22ad9033bcBf192059122d163b60349D`
- **Function:** `deposit(uint256 amount)`
- **Sentence:**
  - Deposit AMOUNT ycrv to earn ycrv.
  - `Deposit {0} ycrv to earn ycrv.`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of yCRV to deposit | Must approve first |
| returns | uint256 | yvyCRV amount received | Exchange rate varies |

### Claim yCRV Rewards
Claim accumulated crvUSD rewards from staked yCRV positions.

- **Contract:** (Reward Distibutor) `0xB226c52EB411326CdB54824a88aBaFDAAfF16D3d`
- **Function:** `claim()`
- **Sentence:**
  - Claim my yCRV rewards.
  - `Claim my yCRV rewards.`

- **Contract:** (Reward Distibutor) `0xB226c52EB411326CdB54824a88aBaFDAAfF16D3d`
- **Function:** `getClaimable(address _account)`
- **Sentence:**
  - If my claimable ycrv is DIRECTION than AMOUNT.
  - `If my claimable ycrv is {0} than {1}`

| Input Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| amount | uint256 | Amount of yCRV to deposit | Must approve first |
| returns | uint256 | yvyCRV amount received | Exchange rate varies |



