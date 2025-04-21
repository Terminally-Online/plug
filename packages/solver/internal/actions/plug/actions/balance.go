package actions

import (
	"fmt"
	"solver/bindings/erc_1155"
	"solver/bindings/erc_20"
	"solver/bindings/erc_721"
	"solver/bindings/plug_evm"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceRequest struct {
	Token  string                                         `json:"token"`
	Holder coil.CoilInput[common.Address, common.Address] `json:"holder"`
}

var NativeBalanceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_evm.PlugEvmMetaData,
	FunctionName: "balanceOf",
}

var Erc20BalanceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "balanceOf",
}

var Erc721BalanceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_721.Erc721MetaData,
	FunctionName: "balanceOf",
}

var Erc1155BalanceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_1155.Erc1155MetaData,
	FunctionName: "balanceOf",
}

func Balance(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	token := common.HexToAddress(parts[0])

	if token == utils.NativeTokenAddress {
		return BalanceNative(lookup)
	}

	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid token format: %s, expected format 'address:decimals:standard'", lookup.Inputs.Token)
	}

	standard, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}

	switch standard {
	case 20:
		return Balance20(lookup)
	case 721:
		return Balance721(lookup)
	case 1155:
		return Balance1155(lookup)
	default:
		return nil, fmt.Errorf("unsupported token standard: %d", standard)
	}
}

func BalanceNative(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	var updates []coil.Update
	holder, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Holder,
		lookup.Inputs.Holder.GetValueWithError,
		&Erc20BalanceFunc,
		"_owner",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	balanceCalldata, err := NativeBalanceFunc.GetCalldata(holder)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["evm"]),
		Data:    balanceCalldata,
		Updates: updates,
	}}, nil
}

func Balance20(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	token := common.HexToAddress(strings.Split(lookup.Inputs.Token, ":")[0])

	var updates []coil.Update
	holder, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Holder,
		lookup.Inputs.Holder.GetValueWithError,
		&Erc20BalanceFunc,
		"_owner",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	balanceCalldata, err := Erc20BalanceFunc.GetCalldata(holder)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       token,
		Data:     balanceCalldata,
		Updates:  updates,
	}}, nil
}

func Balance721(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	token := common.HexToAddress(strings.Split(lookup.Inputs.Token, ":")[0])

	var updates []coil.Update
	holder, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Holder,
		lookup.Inputs.Holder.GetValueWithError,
		&Erc721BalanceFunc,
		"owner",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	balanceCalldata, err := Erc721BalanceFunc.GetCalldata(holder)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       token,
		Data:     balanceCalldata,
		Updates:  updates,
	}}, nil
}

func Balance1155(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	token := common.HexToAddress(parts[0])
	tokenId, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}

	var updates []coil.Update
	holder, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Holder,
		lookup.Inputs.Holder.GetValueWithError,
		&Erc1155BalanceFunc,
		"account",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	balanceCalldata, err := Erc1155BalanceFunc.GetCalldata(holder, tokenId)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       token,
		Data:     balanceCalldata,
		Updates:  updates,
	}}, nil
}
