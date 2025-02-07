package euler

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/euler_governed_perspective"
	"solver/bindings/euler_utils_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/bindings/references"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetVerifiedVaults(chainId uint64) ([]euler_vault_lens.VaultInfoFull, error) {
	provider, err := utils.GetProvider(chainId)
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

	// Prepare multicall inputs
	calls := make([]utils.MulticallCalldata, len(vaultAddresses))
	for i, vaultAddr := range vaultAddresses {
		calls[i] = utils.MulticallCalldata{
			Target: vaultLensAddr,
			Method: "getVaultInfoFull",
			Args:   []interface{}{vaultAddr},
			ABI:    vaultLensAbi,
		}
	}

	multicallAddress := common.HexToAddress(references.Networks[chainId].References["multicall"]["primary"])
	returnData, err := utils.ExecuteMulticall(chainId, multicallAddress, calls)
	if err != nil {
		return nil, fmt.Errorf("multicall failed: %w", err)
	}

	results := make([]euler_vault_lens.VaultInfoFull, len(returnData))
	for i, data := range returnData {
		vaultInfo, err := vaultLensAbi.Unpack("getVaultInfoFull", data)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack vault info: %w", err)
		}

		if len(vaultInfo) == 0 {
			return nil, fmt.Errorf("empty vault info returned for vault %d", i)
		}

		jsonData, err := json.Marshal(vaultInfo[0])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal vault info: %w", err)
		}

		var result euler_vault_lens.VaultInfoFull
		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal vault info: %w", err)
		}

		results[i] = result
	}

	return results, nil
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
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, nil, err
	}

	vault, err := GetVault(address, chainId)
	if err != nil {
		return nil, nil, err
	}

	utilsLens, err := euler_utils_lens.NewEulerUtilsLens(
		common.HexToAddress(references.Networks[chainId].References["euler"]["utils_lens"]),
		provider,
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
