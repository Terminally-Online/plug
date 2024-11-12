ycrv_spec_sheet.md

Yearn and Curve are closely related and this integration includes specs related to yCRV, a mechanism that allows users of Curve and Yearn to earn increased yield. Users can convert their CRV by depositing to the yCRV contract and then can auto-compound their yield by depositing their yCRV to the yvyCRV contract.



| System | Name | Type | Source | Supported | Implemented | Notes |
| yCRV | Deposit CRV | Action | yCRV Contract |  |  | Get yCRV tokens |
| yCRV | Stake yCRV | Action | YearnBoostedStaker |  |  | Earn crvUSD |
| yCRV | Auto-compound yCRV | Action | yvyCRV Vault |  |  | Compound yields |
| yCRV | Claim yCRV | Action | Reward Distributor |  |  | Collect yield |
| yCRV | Claimable yCRV | Constraint | Reward Distributor |  |  | Collect yield |

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



