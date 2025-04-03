package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BorrowRequest struct {
	Amount          coil.CoilInput[string, *big.Int] `json:"amount"`
	Token           string                           `json:"token"`
	Vault           string                           `json:"vault"`
	SubAccountIndex uint8                            `json:"sub-account"`
}

var BorrowFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     euler_evault_implementation.EulerEvaultImplementationMetaData,
	FunctionName: "borrow",
}

func Borrow(lookup *actions.SchemaLookup[BorrowRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	var updates []coil.Update
	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&BorrowFunc,
		"amount",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	calldata, err := BorrowFunc.GetCalldata(
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		common.HexToAddress(lookup.Inputs.Vault),
		subAccountAddress,
		big.NewInt(0),
		calldata,
		updates,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap borrow call: %w", err)
	}

	return []signature.Plug{call}, nil
}
