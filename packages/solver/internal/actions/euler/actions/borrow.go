package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/solver/signature"
	"solver/internal/utils"
)

type BorrowRequest struct {
	Amount          string `json:"amount"`
	Token           string `json:"token"`
	Vault           string `json:"vault"`
	SubAccountIndex uint8  `json:"sub-account"`
}

func Borrow(lookup *actions.SchemaLookup[BorrowRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %w", err)
	}

	vault, err := reads.GetVault(lookup.Inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	calldata, err := vaultAbi.Pack(
		"borrow",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		calldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap borrow call: %w", err)
	}

	return []signature.Plug{call}, nil
}
