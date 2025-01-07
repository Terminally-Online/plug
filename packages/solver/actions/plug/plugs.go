package plug

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"solver/actions"
	"solver/bindings/erc_20"
	"solver/types"
	"solver/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func HandleTransfer(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Token     string  `json:"token"`     // Address of the token to transfer.
		Recipient string  `json:"recipient"` // Address of the recipient.
		Amount    big.Int `json:"amount"`    // Raw amount of tokens to transfer.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	if common.HexToAddress(inputs.Token) == utils.NativeTokenAddress {
		transaction := ethtypes.NewTransaction(
			0,
			common.HexToAddress(inputs.Recipient),
			&inputs.Amount,
			utils.NativeTransferGas,
			big.NewInt(0),
			nil,
		)

		return []*types.Transaction{{
			To:    inputs.Recipient,
			Value: *transaction.Value(),
		}}, nil
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}

	calldata, err := erc20Abi.Pack("transfer",
		common.HexToAddress(inputs.Recipient),
		&inputs.Amount,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   inputs.Token,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}

func HandleTransferFrom(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Token     string  `json:"token"`     // Address of the token to transfer.
		Recipient string  `json:"recipient"` // Address of the recipient.
		Amount    big.Int `json:"amount"`    // Raw amount of tokens to transfer.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	tokenType, err := getTokenType(inputs.Token)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	log.Printf("%d", *tokenType)

	return []*types.Transaction{}, nil
}

/*
Swap currently only supports ERC20 tokens however programatic support for NFT swaps is possible, just
not high priority at this time.
*/
func HandleSwap(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		TokenIn   string `json:"tokenIn"`   // Address of the token to transfer.
		TokenOut  string `json:"tokenOut"`  // Address of the recipient.
		AmountOut string `json:"amountOut"` // Raw amount of tokens to transfer.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	bebopApiUrl := fmt.Sprintf("https://api.bebop.xyz/pmm/ethereum/v3/quote?buy_tokens=%s&sell_tokens=%s&sell_amounts=%s&taker_address=%s&gasless=false&approval_type=Standard&skip_validation=true",
		inputs.TokenIn,
		inputs.TokenOut,
		inputs.AmountOut,
		params.From,
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

	transactions := []*types.Transaction{{
		To:    quoteResponse.Tx.To,
		Data:  quoteResponse.Tx.Data,
		Value: *value,
		Meta: BebopTransactionMeta{
			Expiry:             quoteResponse.Expiry,
			Slippage:           quoteResponse.Slippage,
			PriceImpact:        quoteResponse.PriceImpact,
			Warnings:           quoteResponse.Warnings,
			BuyTokens:          quoteResponse.BuyTokens,
			SellTokens:         quoteResponse.SellTokens,
			SettlementAddress:  quoteResponse.SettlementAddress,
			RequiredSignatures: quoteResponse.RequiredSignatures,
			PartnerFeeNative:   quoteResponse.PartnerFeeNative,
		},
	}}

	if value.Cmp(big.NewInt(0)) == 0 {
		erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
		if err != nil {
			return nil, utils.ErrABIFailed("ERC20")
		}

		amountOut, ok := new(big.Int).SetString(inputs.AmountOut, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse amountOut: %s", inputs.AmountOut)
		}

		approveCalldata, err := erc20Abi.Pack("approve",
			common.HexToAddress(quoteResponse.ApprovalTarget),
			amountOut,
		)
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}

		transactions = append([]*types.Transaction{{
			To:   inputs.TokenOut,
			Data: "0x" + common.Bytes2Hex(approveCalldata),
		}}, transactions...)
	}

	return transactions, nil
}
