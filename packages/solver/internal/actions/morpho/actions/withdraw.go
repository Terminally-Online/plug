package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_router"
	"solver/bindings/morpho_vault"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
	Target string                           `json:"target"`
}

var WithdrawMarketFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     morpho_router.MorphoRouterMetaData,
	FunctionName: "withdraw",
}

var WithdrawVaultFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     morpho_vault.MorphoVaultMetaData,
	FunctionName: "withdraw",
}

func Withdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	// Differentiate handling for vault and market
	if len(lookup.Inputs.Target) == 42 {
		var amountUpdates []coil.Update
		amount, amountUpdates, err := actions.GetAndUpdate(
			&lookup.Inputs.Amount,
			lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
			&WithdrawVaultFunc,
			"assets",
			amountUpdates,
			lookup.PreviousActionDefinition,
		)
		if err != nil {
			return nil, err
		}

		_, err = reads.GetVault(lookup.Inputs.Target, lookup.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to get vault: %w", err)
		}

		withdrawCalldata, err := WithdrawVaultFunc.GetCalldata(
			amount,
			lookup.From,
			lookup.From,
		)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		return []signature.Plug{{
			To:      common.HexToAddress(lookup.Inputs.Target),
			Data:    withdrawCalldata,
			Updates: amountUpdates,
		}}, nil
	}

	market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	var amountUpdates []coil.Update
	amount, amountUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&WithdrawMarketFunc,
		"assets",
		amountUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	withdrawCalldata, err := WithdrawMarketFunc.GetCalldata(
		market.Params,
		amount,
		big.NewInt(0),
		lookup.From,
		lookup.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw calldata: %w", err)
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		Data:    withdrawCalldata,
		Updates: amountUpdates,
	}}, nil
}
