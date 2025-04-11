package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
	Vault  string                           `json:"vault"`
}

var WithdrawFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     yearn_v3_pool.YearnV3PoolMetaData,
	FunctionName: "withdraw",
}

func Withdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	var updates []coil.Update
	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&WithdrawFunc,
		"assets",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	withdrawCalldata, err := WithdrawFunc.GetCalldata(amount, lookup.From, lookup.From)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:      common.HexToAddress(lookup.Inputs.Vault),
		Data:    withdrawCalldata,
		Updates: updates,
	}}, nil
}
