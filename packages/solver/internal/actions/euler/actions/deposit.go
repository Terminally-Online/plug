package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/euler_evault_implementation"
	"solver/bindings/euler_evc"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type DepositCollateralRequest struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

func DepositCollateral(lookup *actions.SchemaLookup[DepositCollateralRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit collateral amount to uint: %w", err)
	}

	vault, err := reads.GetVault(lookup.Inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Erc20")
	}

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault,
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	evc, err := euler_evc.EulerEvcMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvc")
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		depositCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	enableCollateralCalldata, err := evc.Pack(
		"enableCollateral",
		subAccountAddress,
		vault.Vault,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	enableCollateralCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["evc"]),
		subAccountAddress,
		big.NewInt(0),
		enableCollateralCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap enable collateral call: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, depositCall, enableCollateralCall}, nil
}
