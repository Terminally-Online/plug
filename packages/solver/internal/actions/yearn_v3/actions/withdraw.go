package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strings"

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
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
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

	targetVault, err := reads.GetVault(lookup.ChainId, lookup.Inputs.Vault)
	if !strings.EqualFold(token.Hex(), targetVault.Token.Address) {
		return nil, fmt.Errorf("asset %s cannot be used in vault: %s", token, lookup.Inputs.Vault)
	}

	withdrawCalldata, err := WithdrawFunc.GetCalldata(amount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      common.HexToAddress(lookup.Inputs.Vault),
		Data:    withdrawCalldata,
		Updates: updates,
	}}, nil
}
