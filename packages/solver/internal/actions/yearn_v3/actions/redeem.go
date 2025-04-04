package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type RedeemRequest struct {
	Amount string `json:"amount"`
	Gauge  string `json:"gauge"`
}

// TODO: I don't think this sentence makes sense? How should this actually be used?
// TODO Mason: agreed this doesn't make sense. will come back to it when I start testing everything
func Redeem(lookup *actions.SchemaLookup[RedeemRequest]) ([]signature.Plug, error) {
	// // TODO: Need to update the options to set the guage address (vault.Staking.Address)
	// //	     as the value in the options for this action. Right now it is setting
	// //	     the vault address which it should not be.
	// gauge, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Gauge)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	// }

	// // TODO: Should not run this if we are using linked inputs.
	// amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to convert redeem amount to uint: %w", err)
	// }

	// gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	// if err != nil {
	// 	return nil, utils.ErrABI("YearnV3Gauge")
	// }

	// calldata, err := gaugeAbi.Pack("redeem", amount, lookup.From, lookup.From)
	// if err != nil {
	// 	return nil, utils.ErrTransaction(err.Error())
	// }

	// return []signature.Plug{{
	// 	To:   *gauge,
	// 	Data: calldata,
	// }}, nil
	return nil, nil
}
