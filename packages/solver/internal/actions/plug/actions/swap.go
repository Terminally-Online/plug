package actions

import (
	"fmt"
	"log"
	"math/big"
	"net/url"
	"os"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/client"
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

type BebopTransactionMeta struct {
	Expiry             int64                                  `json:"expiry"`
	Slippage           float64                                `json:"slippage"`
	PriceImpact        float64                                `json:"priceImpact"`
	Warnings           []any                                  `json:"warnings"`
	BuyTokens          map[string]BebopQuoteResponseBuyTokens `json:"buyTokens"`
	SellTokens         map[string]BebopQuoteResponseToken     `json:"sellTokens"`
	SettlementAddress  string                                 `json:"settlementAddress"`
	RequiredSignatures []any                                  `json:"requiredSignatures"`
	PartnerFeeNative   string                                 `json:"partnerFeeNative"`
}

type BebopQuoteResponseBuyTokens struct {
	BebopQuoteResponseToken
	AmountBeforeFee   string  `json:"amountBeforeFee"`
	DeltaFromExpected float64 `json:"deltaFromExpected"`
}

type BebopQuoteResponseToken struct {
	Amount         string  `json:"amount"`
	Decimals       int     `json:"decimals"`
	PriceUsd       float64 `json:"priceUsd"`
	Symbol         string  `json:"symbol"`
	Price          float64 `json:"price"`
	PriceBeforeFee float64 `json:"priceBeforeFee"`
}

type BebopQuoteResponseToSign struct {
	PartnerID      int    `json:"partner_id"`
	Expiry         int64  `json:"expiry"`
	TakerAddress   string `json:"taker_address"`
	MakerAddress   string `json:"maker_address"`
	MakerNonce     string `json:"maker_nonce"`
	TakerToken     string `json:"taker_token"`
	MakerToken     string `json:"maker_token"`
	TakerAmount    string `json:"taker_amount"`
	MakerAmount    string `json:"maker_amount"`
	Receiver       string `json:"receiver"`
	PackedCommands string `json:"packed_commands"`
}

type BebopQuote struct {
	Type         string  `json:"type"`
	Status       string  `json:"status"`
	QuoteId      string  `json:"quoteId"`
	ChainId      int     `json:"chainId"`
	ApprovalType string  `json:"approvalType"`
	NativeToken  string  `json:"nativeToken"`
	Taker        string  `json:"taker"`
	Receiver     string  `json:"receiver"`
	Expiry       int64   `json:"expiry"`
	Slippage     float64 `json:"slippage"`
	GasFee       struct {
		Native string  `json:"native"`
		Usd    float64 `json:"usd"`
	} `json:"gasFee"`
	BuyTokens          map[string]BebopQuoteResponseBuyTokens `json:"buyTokens"`
	SellTokens         map[string]BebopQuoteResponseToken     `json:"sellTokens"`
	SettlementAddress  string                                 `json:"settlementAddress"`
	ApprovalTarget     string                                 `json:"approvalTarget"`
	RequiredSignatures []any                                  `json:"requiredSignatures"`
	PriceImpact        float64                                `json:"priceImpact"`
	PartnerFeeNative   string                                 `json:"partnerFeeNative"`
	Warnings           []any                                  `json:"warnings"`
	Tx                 struct {
		To       string `json:"to"`
		Value    string `json:"value"`
		Data     string `json:"data"`
		From     string `json:"from"`
		Gas      int    `json:"gas"`
		GasPrice int64  `json:"gasPrice"`
	} `json:"tx"`
	Makers []string `json:"makers"`
	ToSign struct {
		PartnerID      int    `json:"partner_id"`
		Expiry         int64  `json:"expiry"`
		TakerAddress   string `json:"taker_address"`
		MakerAddress   string `json:"maker_address"`
		MakerNonce     string `json:"maker_nonce"`
		TakerToken     string `json:"taker_token"`
		MakerToken     string `json:"maker_token"`
		TakerAmount    string `json:"taker_amount"`
		MakerAmount    string `json:"maker_amount"`
		Receiver       string `json:"receiver"`
		PackedCommands string `json:"packed_commands"`
	} `json:"toSign"`
	OnchainOrderType  string `json:"onchainOrderType"`
	PartialFillOffset int    `json:"partialFillOffset"`
}

type BebopQuoteRoute struct {
	Type  string     `json:"type"`
	Quote BebopQuote `json:"quote"`
}
type BebopQuoteResponse struct {
	Error struct {
		ErrorCode int    `json:"errorCode"`
		Message   string `json:"message"`
	} `json:"error,omitempty"`
	Routes []BebopQuoteRoute `json:"routes"`
	Link   string            `json:"link"`
}

// https://api.bebop.xyz/router/ethereum/docs#/v1/get_quote_v1_quote_get
func getBebopQuoteURL(lookup *actions.SchemaLookup[SwapRequest]) string {
	chainName := client.GetChainName(lookup.ChainId)

	baseURL := fmt.Sprintf("https://api.bebop.xyz/router/%s/v1/quote", chainName)

	u, _ := url.Parse(baseURL)
	q := u.Query()

	q.Set("buy_tokens", lookup.Inputs.TokenIn)
	q.Set("sell_tokens", lookup.Inputs.Token)
	q.Set("sell_amounts", lookup.Inputs.Amount)
	q.Set("taker_address", lookup.From.String())
	q.Set("source", os.Getenv("BEBOP_SOURCE"))
	q.Set("gasless", "false") // Is equivalent of defining it as 'self-execute'.
	q.Set("approval_type", "Standard")
	q.Set("skip_validation", "true")
	q.Set("skip_taker_checks", "true")

	u.RawQuery = q.Encode()
	return u.String()
}

func handleSwap(lookup *actions.SchemaLookup[SwapRequest]) ([]signature.Plug, error) {
	bebopUrl := getBebopQuoteURL(lookup)
	log.Println(bebopUrl)
	bebopResponse, err := utils.MakeHTTPRequest(
		bebopUrl,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
			"source-auth":  os.Getenv("BEBOP_SOURCE_AUTH"), // This was given by the Bebop team to identify us.
		},
		nil,
		nil,
		BebopQuoteResponse{},
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

	var quoteResponse *BebopQuote
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
		Meta: BebopTransactionMeta{
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

	return handleSwap(lookup)
}
