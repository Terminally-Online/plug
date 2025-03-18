package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)


func createEmptyAccountOption(index int, chainId uint64, address common.Address) actions.Option {
	collateralIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "collateral")
	debtIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "debt")

	return actions.Option{
		Label: fmt.Sprintf("Account #%d", index),
		Name:  utils.FormatAddress(euler_utils.GetSubAccountAddress(address, uint8(index))),
		Value: fmt.Sprintf("%d", index),
		Info:  &actions.OptionInfo{Label: "Health Factor: -", Value: "$0.00"},
		Icon:  &actions.OptionIcon{Default: fmt.Sprintf("%s%%7C%s", collateralIcon, debtIcon)},
	}
}

func TokenOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	vaults, err := reads.GetVerifiedVaults(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	_, supplyVaultOptions, _, err := GetSupplyTokenToVaultOptions(lookup.ChainId, vaults)
	if err != nil {
		return nil, err
	}
	_, borrowVaultOptions, _, err := GetBorrowTokenToVaultOptions(lookup.ChainId, vaults)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: actions.BaseLendActionTypeFields},
		1: {Complex: map[string][]actions.Option{
			"-1": borrowVaultOptions,
			"1":  supplyVaultOptions,
		}},
		2: {Simple: actions.BaseThresholdFields},
	}, nil
}
