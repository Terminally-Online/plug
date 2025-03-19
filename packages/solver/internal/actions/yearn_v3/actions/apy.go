package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type APYRequest struct {
	Vault     string `json:"token"`
	Operator  int    `json:"operator"`
	Threshold string `json:"threshold"`
}


func APY(lookup *actions.SchemaLookup[APYRequest]) ([]signature.Plug, error) {
	// thresholdFloat, err := strconv.ParseFloat(lookup.Inputs.Threshold, 64)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse threshold float: %w", err)
	// }

	// vaults, err := GetVaults(lookup.ChainId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get vaults: %v", err)
	// }

	// var targetVault *YearnVault
	// for _, vault := range vaults {
	// 	if strings.EqualFold(vault.Address, lookup.Inputs.Vault) {
	// 		targetVault = &vault
	// 		break
	// 	}
	// }
	// if targetVault == nil {
	// 	return nil, fmt.Errorf("cannot find data for vault: %s", lookup.Inputs.Vault)
	// }

	// rateFloat := new(big.Float).SetInt(new(big.Int).Add(
	// 	new(big.Int).SetInt64(int64(targetVault.APR.ForwardAPR.NetAPR*100)),
	// 	new(big.Int).SetInt64(int64(targetVault.Extra.StakingRewardsAPR*100)),
	// ))

	return nil, nil
}
