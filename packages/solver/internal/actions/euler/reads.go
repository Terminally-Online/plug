package euler

import (
	"fmt"
	"solver/bindings/euler_governed_perspective"
	"solver/bindings/euler_vault_lens"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"time"

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
		fmt.Printf("Error getting vault addresses, %v\n", err)
		return nil, err
	}

	// Sleep to avoid a 429.
	time.Sleep(15 * time.Second)

	vaultLens, err := euler_vault_lens.NewEulerVaultLens(
		common.HexToAddress(references.Networks[chainId].References["euler"]["vault_lens"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	vaultInfos := make([]euler_vault_lens.VaultInfoFull, len(vaultAddresses))
	for i, vaultAddr := range vaultAddresses {
		vaultInfo, err := vaultLens.GetVaultInfoFull(
			&bind.CallOpts{},
			vaultAddr,
		)
		if err != nil {
			fmt.Printf("Error getting vault %s info on chain %d: %v\n", vaultAddr, chainId, err)
			return nil, err
		}
		vaultInfos[i] = vaultInfo

		// Have to sleep to avoid a 429 from the RPC until this is multicalled.
		time.Sleep(5 * time.Second)
	}

	return vaultInfos, nil
}

func GetVerifiedVaultsMulticall(chainId uint64) ([]euler_vault_lens.VaultInfoFull, error) {
	return nil, nil
	// provider, err := utils.GetProvider(chainId)
	// if err != nil {
	// 	return nil, err
	// }

	// governedPerspective, err := euler_governedPerspective.NewEulerGovernedPerspective(
	// 	common.HexToAddress(references.Networks[chainId].References["euler"]["governed_perspective"]),
	// 	provider,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// vaultAddresses, err := governedPerspective.EulerGovernedPerspectiveCaller.VerifiedArray(nil)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("vaultAddresses", vaultAddresses)
	
	// multicall, err := multicall_primary.NewMulticallPrimary(
	// 	common.HexToAddress(references.Networks[chainId].References["multicall"]["primary"]),
    //     provider,
    // )
    // if err != nil {
    //     return nil, fmt.Errorf("failed to create multicall instance: %w", err)
    // }

    // vaultLensAbi, err := euler_vaultLens.EulerVaultLensMetaData.GetAbi()
	// if err != nil {
	// 	return nil, utils.ErrABI("EulerVaultLens")
	// }

    // calls := make([]multicall_primary.Multicall3Call, len(vaultAddresses))
    // for i, vaultAddr := range vaultAddresses {
    //     callData, err := vaultLensAbi.Pack("getVaultInfoFull", vaultAddr)
    //     if err != nil {
    //         return nil, fmt.Errorf("failed to pack getVaultInfoFull call: %w", err)
    //     }

    //     calls[i] = multicall_primary.Multicall3Call{
    //         Target:   common.HexToAddress(references.Networks[chainId].References["euler"]["vault_lens"]),
    //         CallData: callData,
    //     }
    // }

    // // results, err := multicall.TryAggregate(nil, true, calls)
    // // if err != nil {
    // //     return nil, fmt.Errorf("multicall failed: %w", err)
    // // }
	// // resultData := results.Data()

    // // vaultInfos := make([]euler_vaultLens.VaultInfoFull, len(calls))
    // // for i, tx := range results.Data() {
    // //     var vaultInfo euler_vaultLens.VaultInfoFull
    // //     data, err = vaultLensAbi.Unpack("getVaultInfoFull", tx)
    // //     if err != nil {
    // //         return nil, fmt.Errorf("failed to unpack vault info: %w", err)
    // //     }
    // //     vaultInfos[i] = vaultInfo
    // // }

    // return vaultInfos, nil
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

