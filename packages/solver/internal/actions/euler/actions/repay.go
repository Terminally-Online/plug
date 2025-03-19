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

type RepayRequest struct {
	Amount          string `json:"amount"`
	Token           string `json:"token"`
	Vault           string `json:"vault"`
	SubAccountIndex uint8  `json:"sub-account"`
}

func Repay(lookup *actions.SchemaLookup[RepayRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repay amount to uint: %w", err)
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

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault,
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayCalldata, err := vaultAbi.Pack(
		"repay",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		repayCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap repay call: %w", err)
	}

	return []signature.Plug{{
		To:   vault.Vault,
		Data: approveCalldata,
	}, repayCall}, nil
}
