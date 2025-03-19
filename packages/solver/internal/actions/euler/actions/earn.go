package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/solver/signature"
	"solver/internal/utils"
)

type EarnRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Vault  string `json:"vault"`
}

func Earn(lookup *actions.SchemaLookup[EarnRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert earn amount to uint: %w", err)
	}

	vault, err := reads.GetVault(lookup.Inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
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

	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}
	depositCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		lookup.From,
		big.NewInt(0),
		depositCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, depositCall}, nil
}
