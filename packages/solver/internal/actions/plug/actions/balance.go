package actions

import (
	"solver/bindings/erc_20"
	"solver/bindings/plug_evm"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

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

var BalanceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "balanceOf",
}

func Balance(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	token, _, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, err
	}

	holder, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Holder,
		lookup.Inputs.Holder.GetValueWithError,
		&BalanceFunc,
		"_owner",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	if *token == utils.NativeTokenAddress {
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

	balanceCalldata, err := BalanceFunc.GetCalldata(holder)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       *token,
		Data:     balanceCalldata,
		Updates:  updates,
	}}, nil
}
