package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/weth_address"
	"solver/internal/actions"
	"solver/internal/actions/options"
	"solver/internal/bindings/references"
	"solver/internal/helpers/llama"
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
	Warnings           []interface{}                          `json:"warnings"`
	BuyTokens          map[string]BebopQuoteResponseBuyTokens `json:"buyTokens"`
	SellTokens         map[string]BebopQuoteResponseToken     `json:"sellTokens"`
	SettlementAddress  string                                 `json:"settlementAddress"`
	RequiredSignatures []interface{}                          `json:"requiredSignatures"`
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

type BebopQuoteResponse struct {
	Error struct {
		ErrorCode int    `json:"errorCode"`
		Message   string `json:"message"`
	} `json:"error,omitempty"`
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
	RequiredSignatures []interface{}                          `json:"requiredSignatures"`
	PriceImpact        float64                                `json:"priceImpact"`
	PartnerFeeNative   string                                 `json:"partnerFeeNative"`
	Warnings           []interface{}                          `json:"warnings"`
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

func SwapOptions(lookup *actions.SchemaLookup[SwapRequest]) (map[int]actions.Options, error) {
	fungiblesOutIndex := 1
	fungiblesOutOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesOutIndex)
	if err != nil {
		return nil, err
	}

	fungiblesInIndex := 2
	fungiblesInOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesInIndex)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		fungiblesOutIndex: {Simple: fungiblesOutOptions},
		fungiblesInIndex:  {Simple: fungiblesInOptions},
	}, nil
}

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
			Warnings:           []interface{}{},
			RequiredSignatures: []interface{}{},
			SettlementAddress:  wethAddress,
			PartnerFeeNative:   "0",
		},
	}}, nil
}

func handleSwap(lookup *actions.SchemaLookup[SwapRequest]) ([]signature.Plug, error) {
	// TODO: Right now 'base' is hard-coded as the intended chain for swapping. This should
	//       consume the chain id provided in the params.
	bebopApiUrl := fmt.Sprintf(
		"https://api.bebop.xyz/pmm/base/v3/quote?buy_tokens=%s&sell_tokens=%s&sell_amounts=%s&taker_address=%s&gasless=false&approval_type=Standard&skip_validation=true",
		lookup.Inputs.TokenIn,
		lookup.Inputs.Token,
		lookup.Inputs.Amount,
		lookup.From,
	)

	quoteResponse, err := utils.MakeHTTPRequest(
		bebopApiUrl,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
		},
		nil,
		nil,
		BebopQuoteResponse{},
	)
	if err != nil {
		return nil, err
	}

	if quoteResponse.Error.ErrorCode != 0 {
		return nil, fmt.Errorf("bebop api error: %s", quoteResponse.Error.Message)
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
