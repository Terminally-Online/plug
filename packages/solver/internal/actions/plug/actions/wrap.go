package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/weth_address"
	"solver/internal/actions"
	"solver/internal/helpers/llama"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func handleWrap(lookup *actions.SchemaLookup[SwapRequest], wethAddress string) ([]signature.Plug, error) {
	isEthToWeth := common.HexToAddress(lookup.Inputs.TokenIn) == utils.NativeTokenAddress &&
		strings.EqualFold(lookup.Inputs.Token, wethAddress)
	isWethToEth := strings.EqualFold(lookup.Inputs.TokenIn, wethAddress) &&
		common.HexToAddress(lookup.Inputs.Token) == utils.NativeTokenAddress

	if !isEthToWeth && !isWethToEth {
		return nil, nil
	}

	amountOut, ok := new(big.Int).SetString(lookup.Inputs.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amountOut: %s", lookup.Inputs.Amount)
	}

	wethAbi, err := weth_address.WethAddressMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("WETH")
	}

	var tx signature.Plug
	if isEthToWeth {
		calldata, err := wethAbi.Pack("deposit")
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}
		tx = signature.Plug{
			To:    common.HexToAddress(wethAddress),
			Data:  calldata,
			Value: amountOut,
		}
	} else {
		calldata, err := wethAbi.Pack("withdraw", amountOut)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}
		tx = signature.Plug{
			To:   common.HexToAddress(wethAddress),
			Data: calldata,
		}
	}

	var ethPrice float64
	ethPrices, err := llama.GetPrices([]string{"ethereum:0x0000000000000000000000000000000000000000"})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ETH price: %v", err)
	}
	if _, ok := ethPrices["ethereum:0x0000000000000000000000000000000000000000"]; !ok {
		return nil, fmt.Errorf("ETH price not found in API response")
	}
	ethPrice = ethPrices["ethereum:0x0000000000000000000000000000000000000000"].Price

	buySymbol := "ETH"
	sellSymbol := "WETH"
	if isEthToWeth {
		buySymbol = "WETH"
		sellSymbol = "ETH"
	}

	return []signature.Plug{{
		To:    tx.To,
		Data:  tx.Data,
		Value: tx.Value,
		Meta: BebopTransactionMeta{
			Expiry:      0,
			Slippage:    0.0,
			PriceImpact: 0.0,
			BuyTokens: map[string]BebopQuoteResponseBuyTokens{
				common.HexToAddress(lookup.Inputs.TokenIn).Hex(): {
					BebopQuoteResponseToken: BebopQuoteResponseToken{
						Amount:         lookup.Inputs.Amount,
						Decimals:       18,
						PriceUsd:       ethPrice,
						Symbol:         buySymbol,
						Price:          ethPrice,
						PriceBeforeFee: ethPrice,
					},
					AmountBeforeFee:   lookup.Inputs.Amount,
					DeltaFromExpected: 0,
				},
			},
			SellTokens: map[string]BebopQuoteResponseToken{
				common.HexToAddress(lookup.Inputs.Token).Hex(): {
					Amount:         lookup.Inputs.Amount,
					Decimals:       18,
					PriceUsd:       ethPrice,
					Symbol:         sellSymbol,
					Price:          ethPrice,
					PriceBeforeFee: ethPrice,
				},
			},
			Warnings:           []any{},
			RequiredSignatures: []any{},
			SettlementAddress:  wethAddress,
			PartnerFeeNative:   "0",
		},
	}}, nil
}
