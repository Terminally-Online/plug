package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type TimeToLiquidationRequest struct {
	SubAccountIndex uint8 `json:"sub-account"`
}

func TimeToLiquidation(lookup *actions.SchemaLookup[TimeToLiquidationRequest]) ([]signature.Plug, error) {

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
	// ttlFloat := new(big.Float).SetInt(borrowVault.LiquidityInfo.TimeToLiquidation)
	// ttlFloat.Quo(ttlFloat, new(big.Float).SetInt64(60))
	// ttlMinutes, _ := ttlFloat.Float64()
	//
	return nil, nil
}
