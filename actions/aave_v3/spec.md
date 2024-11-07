# Aave V3 Integration

Integrating a new contract requires two sets of action files, a generalized action and protocol specific action. Generalized Action files define and gather inputs which are then passed to Protocol Specific Action files which reshape the parameters to prepare the transaction.

## Supporting Documentation

-   [Aave Developer Docs Contract Overview](https://docs.aave.com/developers/getting-started/contracts-overview)
-   [Etherscan Aave Contract](https://etherscan.io/address/0x87870bca3f3fd6335c3f4ce8392d69350b4fa4e2#writeProxyContract)
-   [Pool Retrieval Contract](https://aave.com/docs/developers/smart-contracts/view-contracts)
-   [Contract Address Directory](https://aave.com/docs/resources/addresses)

## Contract Interfacing

Aave uses a modular framework for interacting with its contracts. Each pool represents an underlying asset and debt mode. To retrieve all of the reserve assets supported by the Aave protocol, we can use the `UIPoolDataProvider` contract to retrieve all of the supported reserve assets as well as their associated underlying assets, (debt mode) pools, all the associated APYS, etc.

When interacting with a specific pool, you need to target the precise pool address and not a generic one. This address is provided by `getReservesData` for a specific asset/pool.

When sending assets into the pool, a leading approval transaction is required so that the contract can pull it in.

## Scope

| Name                | Type       | Supported | Implemented | Notes                                                                         |
| :------------------ | :--------- | :-------: | :---------: | :---------------------------------------------------------------------------- |
| Deposit             | Action     |    ✔︎    | 11/07/2024  | Requires an initial token approval                                            |
| Borrow              | Action     |    ✔︎    | 11/07/2024  |                                                                               |
| Repay               | Action     |    ✔︎    | 11/07/2024  | Requires an intial token approval                                             |
| Redeem (Withdraw)   | Action     |    ✔︎    | 11/07/2024  |                                                                               |
| Health Factor       | Constraint |           |             | Preventive actions before liquidation risk increases                          |
| APY                 | Constraint |           |             | Automatically enter when rates are exceptionally high or exit when low        |
| APY Differentials   | Constraint |           |             | Automatically execute yield farming strategies when spreads become profitable |
| Available Liquidity | Constraint |           |             | Automatically fill a gap to adjust rates / claim the last opportunity         |

### Deposit

Users can deposit collateral to Aave to earn interest and increase their borrowing power

-   **Contract Address:** <Pool Address>
-   **Function:** supply(address asset,uint256 amount,address onBehalfOf,uint16 referralCode)

| Input Name | Type    | Description                                 | Notes                           |
| :--------- | :------ | :------------------------------------------ | :------------------------------ |
| asset      | address | The address of the asset to supply          |
| amount     | uint256 | The amount to supply                        |
| onBehalfOf | address | The address of the user supplying the asset | Optional, default is the sender |

### Borrow

Users can borrow assets from Aave using their deposited collateral

-   **Contract Address:** <Pool Address>
-   **Function:** borrow(address asset,uint256 amount,uint256 interestRateMode,uint16 referralCode,address onBehalfOf)

| Input Name       | Type    | Description                                       | Notes                           |
| :--------------- | :------ | :------------------------------------------------ | :------------------------------ |
| asset            | address | The address of the asset to borrow                |                                 |
| amount           | uint256 | The amount to borrow                              |                                 |
| interestRateMode | uint256 | Interest rate mode (1 for Stable, 2 for Variable) |                                 |
| referralCode     | uint256 | A referral code for tracking user referrals       | Optional                        |
| onBehalfOf       | address | The address receiving the borrowed amount         | Can borrow on behalf of another |

### Repay

Users can repay debts on Aave

-   **Contract Address:** <Pool Address>
-   **Function:** repay(address asset,uint256 amount,uint256 interestRateMode,address onBehalfOf)

| Input Name       | Type    | Description                                       | Notes                                         |
| :--------------- | :------ | :------------------------------------------------ | :-------------------------------------------- |
| asset            | address | The address of the borrowed underlying asset      |                                               |
| amount           | uint256 | The amount to repay                               | Use `uint256.max` to repay the whole debt     |
| interestRateMode | uint256 | Interest rate mode (1 for Stable, 2 for Variable) |                                               |
| onBehalfOf       | address | Address of the user for whom the debt is repaid   | Can be the borrower's address or another user |

### Redeem

Users can redeem earned interest from collateral deposited on Aave

-   **Contract Address:** <Pool Address>
-   **Function:** withdraw(address asset,uint256 amount,address to)

| Input Name | Type    | Description                                                                | Notes                                 |
| :--------- | :------ | :------------------------------------------------------------------------- | :------------------------------------ |
| asset      | address | The address of the asset to withdraw                                       |                                       |
| amount     | uint256 | The amount to withdraw. Use \`uint256.max\` to withdraw the entire balance | Use `uint256.max` for full withdrawal |
| to         | address | The recipient address for the withdrawn asset                              |                                       |
