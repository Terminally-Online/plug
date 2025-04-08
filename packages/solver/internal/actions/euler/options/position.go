package options

import (
	"fmt"
	"math"
	"math/big"
	"solver/bindings/euler_account_lens"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func GetAddressPositions(chainId uint64, address common.Address) ([]actions.Option, error) {
	if address == utils.ZeroAddress {
		return nil, nil
	}

	accountLensAbi, err := euler_account_lens.EulerAccountLensMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerAccountLens")
	}

	accountLensAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["account_lens"])
	evcAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["evc"])

	calls := make([]client.MulticallCalldata, 256)
	for i := range 256 {
		subAccountAddress := euler_utils.GetSubAccountAddress(address, uint8(i))
		calls[i] = client.MulticallCalldata{
			Target: accountLensAddr,
			Method: "getAccountEnabledVaultsInfo",
			Args:   []any{evcAddr, subAccountAddress},
			ABI:    accountLensAbi,
			OutputType: &struct {
				VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
			}{},
		}
	}

	client, err := client.New(chainId)
	if err != nil {
		return nil, fmt.Errorf("multicall failed: %w", err)
	}
	results, err := client.Multicall(calls)
	if err != nil {
		return nil, fmt.Errorf("multicall failed: %w", err)
	}

	vaultsBySubaccount := make(map[int]struct {
		debtVault        *euler_account_lens.VaultAccountInfo
		collateralVaults []euler_account_lens.VaultAccountInfo
	})
	consecutiveEmptyAccounts := 0
	maxIndex := -1

	for i, result := range results {
		accountInfo := result.(*struct {
			VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
		})

		if len(accountInfo.VaultAccountInfo) == 0 {
			consecutiveEmptyAccounts++
			if consecutiveEmptyAccounts >= 3 {
				break
			}
			continue
		}

		consecutiveEmptyAccounts = 0
		maxIndex = i
		for _, vaultAccountInfo := range accountInfo.VaultAccountInfo {
			temp := vaultsBySubaccount[i]
			if vaultAccountInfo.Borrowed.Sign() == 0 {
				temp.collateralVaults = append(temp.collateralVaults, vaultAccountInfo)
			} else {
				temp.debtVault = &vaultAccountInfo
			}
			vaultsBySubaccount[i] = temp
		}
	}

	options := make([]actions.Option, 0)
	optionsByIndex := make(map[int]actions.Option)

	for i := 0; i <= maxIndex; i++ {
		if accountVaults, exists := vaultsBySubaccount[i]; exists {
			var debtIcon string
			var netValue float64
			var healthFactor string
			var totalCollateralValue float64

			var largestCollateralVault *struct {
				vault euler_account_lens.VaultAccountInfo
				value float64
			}
			for _, collateralVault := range accountVaults.collateralVaults {
				vaultPrice, err := reads.GetVaultPrice(collateralVault.Vault.String(), chainId)
				if err != nil {
					fmt.Printf("error getting vault price: %v\n", err)
					continue
				}

				decimals := uint8(vaultPrice.Vault.AssetDecimals.Uint64())
				assetValue := utils.UintToFloat(collateralVault.Assets, decimals)
				netValue := assetValue * vaultPrice.Price * math.Pow10(int(decimals))
				totalCollateralValue += netValue
				if largestCollateralVault == nil || netValue > largestCollateralVault.value {
					largestCollateralVault = &struct {
						vault euler_account_lens.VaultAccountInfo
						value float64
					}{
						vault: collateralVault,
						value: netValue,
					}
				}
			}

			accountAddress := euler_utils.GetSubAccountAddress(address, uint8(i))
			if accountVaults.debtVault != nil {
				netValue = utils.UintToFloat(new(big.Int).Sub(accountVaults.debtVault.LiquidityInfo.CollateralValueRaw, accountVaults.debtVault.LiquidityInfo.LiabilityValue), 18)
				debtIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(accountVaults.debtVault.Asset.Hex()))

				totalValueBorrowed := accountVaults.debtVault.LiquidityInfo.LiabilityValue
				lltvValueCollateral := accountVaults.debtVault.LiquidityInfo.CollateralValueLiquidation

				healthFactorFloat := new(big.Float).SetInt(lltvValueCollateral)
				healthFactorFloat.Quo(healthFactorFloat, new(big.Float).SetInt(totalValueBorrowed))
				healthFactorF64, _ := healthFactorFloat.Float64()
				healthFactor = fmt.Sprintf("Health Factor: %.2f", healthFactorF64)
			} else {
				netValue = totalCollateralValue
				debtIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "debt")
				healthFactor = "Health Factor: -"
			}

			collateralIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(largestCollateralVault.vault.Asset.Hex()))
			allIcon := fmt.Sprintf("%s%%7C%s", collateralIcon, debtIcon)

			optionsByIndex[i] = actions.Option{
				Label: fmt.Sprintf("Account #%d", i),
				Name:  utils.FormatAddress(accountAddress),
				Value: fmt.Sprintf("%d", i),
				Info:  &actions.OptionInfo{Label: healthFactor, Value: fmt.Sprintf("$%.2f", netValue)},
				Icon:  &actions.OptionIcon{Default: allIcon},
			}
		} else {
			option := createEmptyAccountOption(i, chainId, address)
			optionsByIndex[i] = option
		}
	}

	for i := 0; i <= maxIndex; i++ {
		if option, exists := optionsByIndex[i]; exists {
			options = append(options, option)
		}
	}

	nextIndex := maxIndex + 1
	options = append(options, createEmptyAccountOption(nextIndex, chainId, address))

	return options, nil
}

func PositionOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	addressPositions, err := GetAddressPositions(lookup.ChainId, lookup.From)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: addressPositions},
	}, nil
}

func SupplyTokenToVaultToPositionsOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	vaults, err := reads.GetVerifiedVaults(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	supplyTokenOptions, _, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(lookup.ChainId, vaults)
	if err != nil {
		return nil, err
	}
	addressPositions, err := GetAddressPositions(lookup.ChainId, lookup.From)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyTokenOptions},
		2: {Complex: supplyTokenToVaultOptions},
		3: {Simple: addressPositions},
	}, nil
}
