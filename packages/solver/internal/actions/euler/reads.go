package euler

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_governed_perspective"
	"solver/bindings/euler_utils_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"time"

	"solver/internal/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type VaultPriceInfo struct {
	vault euler_vault_lens.VaultInfoFull
	price float64
}

func GetVerifiedVaults(chainId uint64) ([]euler_vault_lens.VaultInfoFull, error) {
	cacheKey := fmt.Sprintf("euler:verifiedVaults:%d", chainId)
	res, err := utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, true, func() ([]euler_vault_lens.VaultInfoFull, error) {
		provider, err := client.New(chainId)
		if err != nil {
			return nil, err
		}

		governedPerspective, err := euler_governed_perspective.NewEulerGovernedPerspective(
			common.HexToAddress(references.Networks[chainId].References["euler"]["governed_perspective"]),
			provider,
		)
		if err != nil {
			return nil, err
		}

		vaultAddresses, err := governedPerspective.EulerGovernedPerspectiveCaller.VerifiedArray(nil)
		if err != nil {
			return nil, err
		}

		vaultLensAbi, err := euler_vault_lens.EulerVaultLensMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to get vault lens ABI: %w", err)
		}

		vaultLensAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["vault_lens"])
		calls := make([]client.MulticallCalldata, len(vaultAddresses))
		for i, vaultAddr := range vaultAddresses {
			calls[i] = client.MulticallCalldata{
				Target:     vaultLensAddr,
				Method:     "getVaultInfoFull",
				Args:       []interface{}{vaultAddr},
				ABI:        vaultLensAbi,
				OutputType: &euler_vault_lens.VaultInfoFull{},
			}
		}
		results, err := provider.Multicall(calls)
		if err != nil {
			return nil, fmt.Errorf("multicall failed: %w", err)
		}

		vaultInfos := make([]euler_vault_lens.VaultInfoFull, len(results))
		for i, result := range results {
			vaultInfos[i] = *result.(*euler_vault_lens.VaultInfoFull)
		}

		return vaultInfos, nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetVault(address string, chainId uint64) (euler_vault_lens.VaultInfoFull, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return euler_vault_lens.VaultInfoFull{}, err
	}
	for _, vault := range vaults {
		if vault.Vault == common.HexToAddress(address) {
			return vault, nil
		}
	}
	return euler_vault_lens.VaultInfoFull{}, fmt.Errorf("vault not found for address: %s", address)
}

func GetVaultApy(address string, chainId uint64) (borrowApy *big.Int, supplyApy *big.Int, err error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, nil, err
	}

	vault, err := GetVault(address, chainId)
	if err != nil {
		return nil, nil, err
	}

	utilsLens, err := euler_utils_lens.NewEulerUtilsLens(
		common.HexToAddress(references.Networks[chainId].References["euler"]["utils_lens"]),
		client,
	)
	if err != nil {
		return nil, nil, err
	}

	apy, err := utilsLens.ComputeAPYs(
		&bind.CallOpts{},
		new(big.Int).Sub(vault.BorrowCap, vault.TotalBorrowed),
		vault.TotalCash,
		vault.TotalBorrowed,
		vault.InterestFee,
	)
	if err != nil {
		return nil, nil, err
	}

	return apy.BorrowAPY, apy.SupplyAPY, nil
}

func GetVaultPrices(chainId uint64) (map[string]VaultPriceInfo, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	cacheKey := fmt.Sprintf("euler:verifiedVaultPrices:%d", chainId)
	prices, err := utils.WithCache(cacheKey, []time.Duration{10 * time.Minute}, true, func() (map[string]VaultPriceInfo, error) {

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
				vault: vault,
				price: ratio,
			}
		}

		return prices, nil
	})

	if err != nil {
		return nil, err
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
