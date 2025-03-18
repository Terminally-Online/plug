package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type HealthFactorRequest struct {
		Operator        int    `json:"operator"`
		Threshold       string `json:"threshold"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

func HealthFactor(lookup *actions.SchemaLookup[HealthFactorRequest]) ([]signature.Plug, error) {
	// accountLens, err := euler_account_lens.NewEulerAccountLens(
	// 	common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["account_lens"]),
	// 	lookup.Client,
	// )
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get account lens: %w", err)
	// }
	//
	// subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)
	// evcAddress := common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["evc"])
	//
	// vaultAccountInfos, err := accountLens.GetAccountEnabledVaultsInfo(
	// 	nil,
	// 	evcAddress,
	// 	subAccountAddress,
	// )
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get account vaults info: %w", err)
	// }
	//
	// var borrowVault *euler_account_lens.VaultAccountInfo
	// for _, vaultAccountInfo := range vaultAccountInfos.VaultAccountInfo {
	// 	if !vaultAccountInfo.LiquidityInfo.QueryFailure {
	// 		borrowVault = &vaultAccountInfo
	// 		break
	// 	}
	// }
	//
	// if borrowVault == nil {
	// 	return nil, fmt.Errorf("no vault found with valid liquidity info")
	// }
	//
	// // eurc has a lltv of .9
	// // weth has a lltv of .85
	// // I have .5 eurc and .001 weth deposited
	// // I'm borrowing 0.4 usdc
	// // my health factor is currently 6.97
	//
	// totalValueBorrowed := borrowVault.LiquidityInfo.LiabilityValue
	// lltvValueCollateral := borrowVault.LiquidityInfo.CollateralValueLiquidation
	//
	// healthFactorFloat := new(big.Float).SetInt(lltvValueCollateral)
	// healthFactorFloat.Quo(healthFactorFloat, new(big.Float).SetInt(totalValueBorrowed))
	// healthFactorF64, _ := healthFactorFloat.Float64()
	//
	return nil, nil
}
