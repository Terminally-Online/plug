package basepaint

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"solver/bindings/basepaint_referral"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

func MintLatest(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Recipient common.Address `json:"recipient"`
		Count     uint64         `json:"count"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mint latest inputs: %w", err)
	}

	referralAbi, err := basepaint_referral.BasepaintReferralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	mintLatestCalldata, err := referralAbi.Pack("mintLatest", inputs.Recipient, inputs.Count, references.Referral)
	if err != nil {
		return nil, err
	}

	var value *big.Int
	if os.Getenv("SOLVER_FEE_SWITCH") == "enabled" {
		value = big.NewInt(0).Mul(big.NewInt(13), big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil))
	}

	return []signature.Plug{
		{
			To:    common.HexToAddress(references.Base.References["basepaint"]["referral"]),
			Data:  mintLatestCalldata,
			Value: value,
		},
	}, nil
}
