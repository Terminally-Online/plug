package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type HealthFactorRequest struct {
	Target    string `json:"target"`
	Operator  int8   `json:"operator"`
	Threshold string `json:"threshold"`
}

func HealthFactor(lookup *actions.SchemaLookup[HealthFactorRequest]) ([]signature.Plug, error) {
	// Morpho has health factor formatted as a WAD (18 decimals)
	// threshold, err := utils.StringToUint(lookup.Inputs.Threshold, 18)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to convert health factor threshold to uint: %w", err)
	// }

	// market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get market: %w", err)
	// }

	// morpho, err := morpho_router.NewMorphoRouter(
	// 	common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
	// 	lookup.Client,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// position, err := morpho.Position(
	// 	lookup.Client.ReadOptions(lookup.From),
	// 	[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
	// 	lookup.From,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// borrowAssets := morpho_utils.MulDivUp(
	// 	position.BorrowShares,
	// 	new(big.Int).Add(market.State.BorrowAssets, morpho_utils.VirtualAssets),
	// 	new(big.Int).Add(market.State.BorrowShares, morpho_utils.VirtualShares),
	// )
	// maxBorrow := morpho_utils.WMulDown(
	// 	morpho_utils.MulDivDown(position.Collateral, market.State.Price, morpho_utils.OraclePriceScale),
	// 	market.LLTV,
	// )

	// // If the user has no borrow assets, the health factor is infinite and we can't divide by zero
	// if borrowAssets.Cmp(big.NewInt(0)) != 1 {
	// 	return nil, nil
	// }

	// healthFactor := morpho_utils.WMulDown(maxBorrow, borrowAssets)

	return nil, nil
}
