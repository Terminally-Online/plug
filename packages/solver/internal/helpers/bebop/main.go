package bebop

import (
	"fmt"
	"net/url"
	"os"
	"solver/internal/client"
)

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
func GetBebopQuoteURL(
	chainId uint64,
	tokenIn string,
	tokenOut string,
	amount string,
	from string,
) string {
	chainName := client.GetChainName(chainId)

	baseURL := fmt.Sprintf("https://api.bebop.xyz/router/%s/v1/quote", chainName)

	u, _ := url.Parse(baseURL)
	q := u.Query()

	q.Set("buy_tokens", tokenIn)
	q.Set("sell_tokens", tokenOut)
	q.Set("sell_amounts", amount)
	q.Set("taker_address", from)
	q.Set("source", os.Getenv("BEBOP_SOURCE"))
	q.Set("gasless", "false")
	q.Set("approval_type", "Standard")
	q.Set("skip_validation", "true")
	q.Set("skip_taker_checks", "true")

	u.RawQuery = q.Encode()
	return u.String()
}
