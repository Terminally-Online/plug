package plug

type BebopTransactionMeta struct {
	ApprovalType   string        `json:"approvalType"`
	ApprovalTarget string        `json:"approvalTarget"`
	NativeToken    string        `json:"nativeToken"`
	Expiry         int64         `json:"expiry"`
	Slippage       float64       `json:"slippage"`
	PriceImpact    float64       `json:"priceImpact"`
	Warnings       []interface{} `json:"warnings"`

	BuyTokens          map[string]BebopQuoteResponseBuyTokens `json:"buyTokens"`
	SellTokens         map[string]BebopQuoteResponseToken     `json:"sellTokens"`
	SettlementAddress  string                                 `json:"settlementAddress"`
	RequiredSignatures []interface{}                          `json:"requiredSignatures"`
	PartnerFeeNative   string                                 `json:"partnerFeeNative"`
	Makers             []string                               `json:"makers"`
	ToSign             BebopQuoteResponseToSign               `json:"toSign"`
	OnchainOrderType   string                                 `json:"onchainOrderType"`
	PartialFillOffset  int                                    `json:"partialFillOffset"`
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
