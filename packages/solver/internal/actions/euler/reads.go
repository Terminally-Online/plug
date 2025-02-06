package euler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/euler_governed_perspective"
	"solver/bindings/euler_utils_lens"
	"solver/bindings/euler_vault_lens"
	"solver/bindings/multicall_primary"
	"solver/internal/bindings/references"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetVerifiedVaultsMulticall(chainId uint64) ([]euler_vault_lens.VaultInfoFull, error) {
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

	for i, addr := range vaultAddresses {
		fmt.Printf("Debug: Vault %d address: %s\n", i, addr.Hex())
	}

	vaultLensAbi, err := euler_vault_lens.EulerVaultLensMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get vault lens ABI: %w", err)
	}

	calls := make([]multicall_primary.Multicall3Call, len(vaultAddresses))
	for i, vaultAddr := range vaultAddresses {
		callData, err := vaultLensAbi.Pack("getVaultInfoFull", vaultAddr)
		if err != nil {
			return nil, fmt.Errorf("failed to pack getVaultInfoFull call for vault %s: %w", vaultAddr.Hex(), err)
		}

		vaultLensAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["vault_lens"])
		fmt.Printf("Debug: Creating multicall for vault %d to lens contract %s\n", i, vaultLensAddr.Hex())
		
		calls[i] = multicall_primary.Multicall3Call{
			Target:   vaultLensAddr,
			CallData: callData,
		}
	}

	multicallAbi, err := multicall_primary.MulticallPrimaryMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Multicall")
	}

	input, err := multicallAbi.Pack("aggregate", calls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack multicall aggregate: %w", err)
	}

	multicallAddress := common.HexToAddress(references.Networks[chainId].References["multicall"]["primary"])
	msg := ethereum.CallMsg{
		To:  &multicallAddress,
		Data: input,
	}

	output, err := provider.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make multicall: %w", err)
	}
	
	unpacked, err := multicallAbi.Unpack("aggregate", output)
	if err != nil {
		fmt.Printf("Debug: Error unpacking aggregate: %v\n", err)
		return nil, fmt.Errorf("failed to unpack multicall aggregate: %w", err)
	}
	
	_, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("multicall result 0 was not a big.Int")
	}

	returnData, ok := unpacked[1].([][]byte)
	if !ok {
		return nil, fmt.Errorf("multicall result 1 was not a [][]byte")
	}
	
	results := make([]euler_vault_lens.VaultInfoFull, len(returnData))
	
	for i, data := range returnData {
		vaultInfo, err := vaultLensAbi.Unpack("getVaultInfoFull", data)
		if err != nil {
			fmt.Printf("Debug: Error unpacking vault info for vault %d: %v\n", i, err)
			fmt.Printf("Debug: Failed data hex: %x\n", data)
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
	vaults, err := GetVerifiedVaultsMulticall(chainId)
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
		common.HexToAddress(references.Networks[chainId].References["euler"]["util_lens"]),
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