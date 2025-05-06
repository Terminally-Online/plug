package actions

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/helpers/bebop"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type SwapRequest struct {
	Amount  string `json:"amount"`
	Token   string `json:"token"`
	TokenIn string `json:"tokenIn"`
}

func Swap(lookup *actions.SchemaLookup[SwapRequest]) ([]signature.Plug, error) {
	tokenOutParts := strings.Split(lookup.Inputs.Token, ":")
	lookup.Inputs.Token = tokenOutParts[0]
	if len(tokenOutParts) < 3 {
		return nil, fmt.Errorf("invalid token format: %s, expected format 'address:decimals:standard'", lookup.Inputs.Token)
	}
	decimals, err := strconv.ParseUint(tokenOutParts[1], 10, 8)
	if err != nil {
		return nil, err
	}
	if standard, err := strconv.ParseUint(tokenOutParts[2], 10, 64); err != nil || standard != 20 {
		return nil, utils.ErrNotImplemented("support for 721 and 1155 are not yet implemented")
	}

	tokenInParts := strings.Split(lookup.Inputs.TokenIn, ":")
	lookup.Inputs.TokenIn = tokenInParts[0]

	wethAddress := references.Networks[lookup.ChainId].References["weth"]["address"]
	adjustedAmount, err := utils.StringToUint(lookup.Inputs.Amount, uint8(decimals))
	if err != nil {
		return nil, fmt.Errorf("failed to convert swap amount to uint: %w", err)
	}
	lookup.Inputs.Amount = adjustedAmount.String()

	if txs, err := handleWrap(lookup, wethAddress); err != nil {
		return nil, err
	} else if txs != nil {
		return txs, nil
	}

	bebopUrl := bebop.GetBebopQuoteURL(
		lookup.ChainId,
		lookup.Inputs.TokenIn,
		lookup.Inputs.Token,
		lookup.Inputs.Amount,
		lookup.From.String(),
	)
	bebopResponse, err := utils.MakeHTTPRequest(
		bebopUrl,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
			"source-auth":  os.Getenv("BEBOP_SOURCE_AUTH"), // This was given by the Bebop team to identify us.
		},
		nil,
		nil,
		bebop.BebopQuoteResponse{},
	)
	if err != nil {
		return nil, err
	}

	if bebopResponse.Error.ErrorCode != 0 {
		return nil, fmt.Errorf("bebop api error: %s", bebopResponse.Error.Message)
	}

	if len(bebopResponse.Routes) == 0 {
		return nil, fmt.Errorf("could not find route")
	}

	var quoteResponse *bebop.BebopQuote
	var bestAmountIn *big.Int
	for _, route := range bebopResponse.Routes {
		amountIn := new(big.Int)
		amountIn, success := amountIn.SetString(route.Quote.BuyTokens[lookup.Inputs.TokenIn].Amount, 10)
		if !success {
			log.Println("not success")
			continue
		}

		if quoteResponse == nil || amountIn.Cmp(bestAmountIn) > 0 {
			quoteResponse = &route.Quote
			bestAmountIn = amountIn
		}
	}
	if quoteResponse == nil {
		return nil, fmt.Errorf("failed to find route")
	}

	value, ok := new(big.Int).SetString(strings.TrimPrefix(quoteResponse.Tx.Value, "0x"), 16)
	if !ok {
		return nil, fmt.Errorf("failed to parse value: %s", quoteResponse.Tx.Value)
	}

	transactions := []signature.Plug{{
		To:    common.HexToAddress(quoteResponse.Tx.To),
		Data:  common.FromHex(quoteResponse.Tx.Data),
		Value: value,
		Meta: bebop.BebopTransactionMeta{
			Expiry:             quoteResponse.Expiry,
			Slippage:           quoteResponse.Slippage,
			PriceImpact:        quoteResponse.PriceImpact,
			BuyTokens:          quoteResponse.BuyTokens,
			SellTokens:         quoteResponse.SellTokens,
			Warnings:           quoteResponse.Warnings,
			RequiredSignatures: quoteResponse.RequiredSignatures,
			SettlementAddress:  quoteResponse.SettlementAddress,
			PartnerFeeNative:   quoteResponse.PartnerFeeNative,
		},
	}}

	// NOTE: If value is zero it means we are transferring an ERC20 and need to have
	//       the appropriate approval appended first.
	if value.Cmp(big.NewInt(0)) == 0 {
		erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
		if err != nil {
			return nil, utils.ErrABI("ERC20")
		}

		amountOut, ok := new(big.Int).SetString(lookup.Inputs.Amount, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse amountOut: %s", lookup.Inputs.Amount)
		}

		approveCalldata, err := erc20Abi.Pack("approve",
			common.HexToAddress(quoteResponse.ApprovalTarget),
			amountOut,
		)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		transactions = append([]signature.Plug{{
			To:   common.HexToAddress(lookup.Inputs.Token),
			Data: approveCalldata,
		}}, transactions...)
	}

	return transactions, nil
}
