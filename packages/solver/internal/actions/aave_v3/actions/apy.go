package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3/reads"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type APYRequest struct {
	Direction int    `json:"direction"`
	Token     string `json:"token"`
}

var ApyFunc = actions.ActionOnchainFunctionResponse{
	Arguments: &abi.Arguments{
		{Name: "apy", Type: abi.Type{T: abi.UintTy, Size: 256}},
	},
}

func APY(lookup *actions.SchemaLookup[APYRequest]) ([]signature.Plug, error) {
	token, _, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	reserves, err := reads.GetReserves(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	var targetReserve *aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData
	for _, reserve := range reserves {
		if reserve.UnderlyingAsset == *token {
			targetReserve = &reserve
			break
		}
	}
	if targetReserve == nil {
		return nil, fmt.Errorf("token %s not supported", lookup.Inputs.Token)
	}

	var apy *big.Int
	switch lookup.Inputs.Direction {
	case -1:
		apy = targetReserve.VariableBorrowRate
	case 1:
		apy = targetReserve.LiquidityRate
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", lookup.Inputs.Direction)
	}

	apyCalldata, err := ApyFunc.Arguments.Pack(apy)
	if err != nil {
		return nil, fmt.Errorf("failed to pack price data: %w", err)
	}

	return []signature.Plug{{
		Selector: signature.ForwardedCall,
		Data:     apyCalldata,
	}}, nil
}
