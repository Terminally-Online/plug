package actions

import (
	"math/big"
	"os"
	"solver/bindings/basepaint_referral"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type MintLatestRequest struct {
	Recipient common.Address `json:"recipient"`
	Count     uint64         `json:"count"`
}

func MintLatest(lookup *actions.SchemaLookup[MintLatestRequest]) ([]signature.Plug, error) {
	referralAbi, err := basepaint_referral.BasepaintReferralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	mintLatestCalldata, err := referralAbi.Pack("mintLatest", lookup.Inputs.Recipient, lookup.Inputs.Count, references.Referral)
	if err != nil {
		return nil, err
	}

	var value *big.Int
	if os.Getenv("SOLVER_FEE_SWITCH") == "enabled" {
		value = big.NewInt(0).Mul(big.NewInt(13), big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil))
	}

	return []signature.Plug{
		{
			To:       common.HexToAddress(references.Base.References["basepaint"]["referral"]),
			Data:     mintLatestCalldata,
			Value:    value,
		},
	}, nil
}
