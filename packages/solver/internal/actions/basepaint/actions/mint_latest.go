package actions

import (
	"math/big"
	"os"
	"solver/bindings/basepaint_referral"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type MintLatestRequest struct {
	Recipient coil.CoilInput[common.Address, common.Address] `json:"recipient"`
	Count     uint64                                         `json:"count"`
}

var MintLatestFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     basepaint_referral.BasepaintReferralMetaData,
	FunctionName: "mintLatest",
}

func MintLatest(lookup *actions.SchemaLookup[MintLatestRequest]) ([]signature.Plug, error) {
	var updates []coil.Update
	recipient, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Recipient,
		lookup.Inputs.Recipient.GetValueWithError,
		&MintLatestFunc,
		"sendMintsTo",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	mintLatestCalldata, err := MintLatestFunc.GetCalldata(recipient, lookup.Inputs.Count, references.Referral)
	if err != nil {
		return nil, err
	}

	var value *big.Int
	if os.Getenv("SOLVER_FEE_SWITCH") == "enabled" {
		value = big.NewInt(0).Mul(big.NewInt(13), big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil))
	}

	basepaintContract := common.HexToAddress(references.Base.References["basepaint"]["referral"])
	return []signature.Plug{{
		To:      basepaintContract,
		Data:    mintLatestCalldata,
		Value:   value,
		Updates: updates,
	}}, nil
}
