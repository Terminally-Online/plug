package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type APYRequest struct {
	Direction int    `json:"direction"`
	Token     string `json:"token"`
}

// TODO MASON: reimplement this
func APY(lookup *actions.SchemaLookup[APYRequest]) ([]signature.Plug, error) {
	// token, _, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	// }
	//
	// reserves, err := reads.GetReserves(lookup.ChainId)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// var targetReserve *aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData
	// for _, reserve := range reserves {
	// 	if reserve.UnderlyingAsset == *token {
	// 		targetReserve = &reserve
	// 		break
	// 	}
	// }
	// if targetReserve == nil {
	// 	return nil, fmt.Errorf("token %s not supported", lookup.Inputs.Token)
	// }

	return nil, nil
}
