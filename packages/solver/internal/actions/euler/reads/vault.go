package reads

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_governed_perspective"
	"solver/bindings/euler_utils_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/bindings/references"
	"solver/internal/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetVerifiedVaults(chainId uint64) ([]euler_vault_lens.VaultInfoFull, error) {
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
			Args:       []any{vaultAddr},
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
