# Morpho Integration

## Overview

## Supporting Documentation

- docs
- [Addresses](https://docs.morpho.org/morpho/addresses/)
- [Morpho API](https://docs.morpho.org/apis/morpho/)
****


## Contract Interfacing


>Earlier conversations with the Morpho team mentioned APIs for Blue Chip Vaults. This can be accomplished with the SDK: https://docs.morpho.org/sdks/



Morpho is made up of a set of Morpho Markets and Morpho Vaults. Vaults use Markets to access liquidity. 

Each Morpho market has its own contract address and markets are named based on their individual parameters in the following format:

`CollateralAsset/LoanAsset (LLTV, ORACLE, IRM)`

Using the following parameters as an example:

- CollateralAsset: `wstETH` 
- LoanAsset: `WETH` 
- LLTV: `94.5%` 
- Oracle: `ChainlinkOracle` 
- IRM: `AdaptiveCurveIRM`

The market would be named `wstETH/WETH (94.5%, ChainlinkOracle, AdaptiveCurveIRM)`

Morpho actions require that market related parameters are entered.

The market id is a bytes32 keccak256 hash of the 5 parameters of a market. It is used to identify a market in Morpho.

Market IDs can be found by reviewing the events from the Morpho Market Factory: `0xbbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb`

Creating a market will emit an event: `CreateMarket` with the market `id`. To retrieve the market parameters from this id, one has to paste the id into the idToMarketParams function on the Morpho contract.

Users can supply and withdraw directly to markets or use ERC4626-compliant Morpho Vaults. Each vault has one loan asset and can allocate deposits to multiple Morpho markets.





Alternative data access option: https://docs.morpho.org/apis/morpho/

## Scope

| System | Name          | Type       | Implemented | Notes |
| :----- | :------------ | :--------- | :---------- | :---- | 
| Morpho Vault | Deposit       | Action     |             |       |    
| Morpho Vault | Withdraw      | Action     |             |       |
|        | Borrow        | Action     |             |       |
|        | Repay         | Action     |             |       |
|        | Health Factor | Constraint |             |       |
|        | Vault APY | Constraint |             |       |



### Contracts

| Name | Address | Desc |
| :--- | :------ |:------ |
| MetaMorpho Factory | 0xA9c3D3a366466Fa809d1Ae982Fb2c46E5fC41101 | Deploys Morpho Vaults|
| Morpho Market Factory | 0xbbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb | Deploys Morpho Markets |
| Public Allocator | 0xfd32fA2ca22c76dD6E550706Ad913FC6CE91c75D |  |


### Deposit to Morpho Vault

This is called on the Morpho Vault.

Mints shares Vault shares to receiver by depositing exactly amount of underlying tokens. 
- MUST emit the Deposit event. 
- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the deposit execution, and are accounted for during deposit. 
- MUST revert if all of assets cannot be deposited (due to deposit limit being reached, slippage, the user not approving enough underlying tokens to the Vault contract, etc). 

NOTE: most implementations will require pre-approval of the Vault with the Vault’s underlying asset token.

### Withdraw from Morpho Vault

This is called on the Morpho Vault.

Burns shares from owner and sends exactly assets of underlying tokens to receiver. 
- MUST emit the Withdraw event. 
- MAY support an additional flow in which the underlying tokens are owned by the Vault contract before the withdraw execution, and are accounted for during withdraw. 
- MUST revert if all of assets cannot be withdrawn (due to withdrawal limit being reached, slippage, the owner not having enough shares, etc). Note that some implementations will require pre-requesting to the Vault before a withdrawal may be performed. Those methods should be performed separately.

