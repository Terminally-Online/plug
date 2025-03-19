package reads

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_utils_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type VaultPriceInfo struct {
	Vault euler_vault_lens.VaultInfoFull
	Price float64
}

func GetVaultPrices(chainId uint64) (map[string]VaultPriceInfo, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	utilLensAbi, err := euler_utils_lens.EulerUtilsLensMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerUtilsLens")
	}

	utilLensAddress := common.HexToAddress(references.Networks[chainId].References["euler"]["utils_lens"])
	calls := make([]client.MulticallCalldata, len(vaults))
	for i, vault := range vaults {
		calls[i] = client.MulticallCalldata{
			Target: utilLensAddress,
			Method: "getAssetPriceInfo",
			Args:   []interface{}{vault.Asset, vault.UnitOfAccount},
			ABI:    utilLensAbi,
			OutputType: &euler_utils_lens.AssetPriceInfo{
				Asset:        common.Address{},
				AmountIn:     big.NewInt(0),
				AmountOutBid: big.NewInt(0),
			},
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

	prices := make(map[string]VaultPriceInfo)
	for i, result := range results {
		priceInfo := result.(*euler_utils_lens.AssetPriceInfo)
		if priceInfo.QueryFailure || priceInfo.AmountIn.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		vault := vaults[i]
		vaultAddr := vault.Vault.String()
		ratio := utils.UintToFloat(new(big.Int).Div(priceInfo.AmountOutBid, priceInfo.AmountIn), uint8(vault.UnitOfAccountDecimals.Uint64()))
		prices[vaultAddr] = VaultPriceInfo{
			Vault: vault,
			Price: ratio,
		}
	}

	return prices, nil

}

func GetVaultPrice(vault string, chainId uint64) (VaultPriceInfo, error) {
	prices, err := GetVaultPrices(chainId)
	if err != nil {
		return VaultPriceInfo{}, err
	}

	priceInfo, ok := prices[vault]

	if !ok {
		return VaultPriceInfo{}, fmt.Errorf("vault price not found for address: %s", vault)
	}

	return priceInfo, nil
}
