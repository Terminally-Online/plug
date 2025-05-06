package actions

import (
	"fmt"
	"math/big"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type APYRequest struct {
	Direction int    `json:"direction"` // -1 for borrow, 1 for deposit
	Target    string `json:"target"`    // Underlying market or vault
}

var ApyFunc = actions.ActionOnchainFunctionResponse{
	Arguments: &abi.Arguments{
		{Name: "apy", Type: abi.Type{T: abi.UintTy, Size: 256}},
	},
}

func APY(lookup *actions.SchemaLookup[APYRequest]) ([]signature.Plug, error) {
	var apy float64
	isVault := len(lookup.Inputs.Target) == 42
	if isVault {
		vault, err := reads.GetVault(lookup.Inputs.Target, lookup.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch vault: %w", err)
		}

		if lookup.Inputs.Direction != 1 {
			return nil, fmt.Errorf("vaults only support deposit direction (1), got %d", lookup.Inputs.Direction)
		}

		apy = vault.DailyApys.NetApy
	} else {
		market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch market: %w", err)
		}

		switch lookup.Inputs.Direction {
		case -1:
			apy = market.DailyApys.BorrowApy
		case 1:
			apy = market.DailyApys.SupplyApy
		default:
			return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", lookup.Inputs.Direction)
		}
	}

	apyPercent := apy * 100
	apyDecimal := new(big.Float).SetFloat64(apyPercent)
	apyScaled := new(big.Float).Mul(apyDecimal, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	apyInt := new(big.Int)
	apyScaled.Int(apyInt)

	apyCalldata, err := ApyFunc.GetCalldata(apyInt)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		Selector: signature.ForwardedCall,
		Data:     apyCalldata,
	}}, nil
}
