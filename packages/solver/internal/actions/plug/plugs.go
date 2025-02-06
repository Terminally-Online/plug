package plug

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/weth_address"
	"solver/internal/actions"
	"solver/internal/actions/llama"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type SwapInputs struct {
	TokenIn   string `json:"tokenIn"`
	TokenOut  string `json:"tokenOut"`
	AmountOut string `json:"amountOut"`
}

func HandleTransfer(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token     string `json:"token"`     // Address of the token to transfer.
		Recipient string `json:"recipient"` // Address of the recipient.
		Amount    string `json:"amount"`    // Raw amount of tokens to transfer.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	if common.HexToAddress(inputs.Token) == utils.NativeTokenAddress {
		amount, err := utils.StringToUint(inputs.Amount, 18)
		if err != nil {
			return nil, fmt.Errorf("failed to convert native transfer amount to uint: %w", err)
		}

		transaction := ethtypes.NewTransaction(
			0,
			common.HexToAddress(inputs.Recipient),
			amount,
			utils.NativeTransferGas,
			big.NewInt(0),
			nil,
		)

		return []signature.Plug{{
			To:    common.HexToAddress(inputs.Recipient),
			Value: transaction.Value(),
		}}, nil
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}

	decimals, err := getERC20Decimals(params.ChainId, inputs.Token)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	amount, err := utils.StringToUint(inputs.Amount, *decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transfer amount to uint: %w", err)
	}

	calldata, err := erc20Abi.Pack("transfer",
		common.HexToAddress(inputs.Recipient),
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Token),
		Data: calldata,
	}}, nil
}

func HandleTransferFrom(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token     string  `json:"token"`     // Address of the token to transfer.
		Recipient string  `json:"recipient"` // Address of the recipient.
		Amount    big.Int `json:"amount"`    // Raw amount of tokens to transfer.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	tokenType, err := getTokenType(params.ChainId, inputs.Token)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	log.Printf("%d", *tokenType)

	return []signature.Plug{}, nil
}

// handleEthWethSwap handles direct conversions between ETH and WETH
func handleEthWethSwap(inputs SwapInputs, _ actions.HandlerParams, wethAddress string) ([]signature.Plug, error) {
	isEthToWeth := common.HexToAddress(inputs.TokenIn) == utils.NativeTokenAddress &&
		strings.EqualFold(inputs.TokenOut, wethAddress)
	isWethToEth := strings.EqualFold(inputs.TokenIn, wethAddress) &&
		common.HexToAddress(inputs.TokenOut) == utils.NativeTokenAddress

	if !isEthToWeth && !isWethToEth {
		return nil, nil
	}

	amountOut, ok := new(big.Int).SetString(inputs.AmountOut, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amountOut: %s", inputs.AmountOut)
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

	// Get current ETH price
	var ethPrice float64
	ethPrices, err := llama.GetPrices([]string{"ethereum:0x0000000000000000000000000000000000000000"})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ETH price: %v", err)
	}
	if _, ok := ethPrices["ethereum:0x0000000000000000000000000000000000000000"]; !ok {
		return nil, fmt.Errorf("ETH price not found in API response")
	}
	ethPrice = ethPrices["ethereum:0x0000000000000000000000000000000000000000"].Price

	// Set symbols based on conversion direction
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
				common.HexToAddress(inputs.TokenIn).Hex(): {
					BebopQuoteResponseToken: BebopQuoteResponseToken{
						Amount:         inputs.AmountOut,
						Decimals:       18,
						PriceUsd:       ethPrice,
						Symbol:         buySymbol,
						Price:          ethPrice,
						PriceBeforeFee: ethPrice,
					},
					AmountBeforeFee:   inputs.AmountOut,
					DeltaFromExpected: 0,
				},
			},
			SellTokens: map[string]BebopQuoteResponseToken{
				common.HexToAddress(inputs.TokenOut).Hex(): {
					Amount:         inputs.AmountOut,
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

// handleBebopSwap handles swaps through the Bebop API
func handleBebopSwap(inputs SwapInputs, params actions.HandlerParams) ([]signature.Plug, error) {
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

		amountOut, ok := new(big.Int).SetString(inputs.AmountOut, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse amountOut: %s", inputs.AmountOut)
		}

		approveCalldata, err := erc20Abi.Pack("approve",
			common.HexToAddress(quoteResponse.ApprovalTarget),
			amountOut,
		)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		transactions = append([]signature.Plug{{
			To:   common.HexToAddress(inputs.TokenOut),
			Data: approveCalldata,
		}}, transactions...)
	}

	return transactions, nil
}

func HandleSwap(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs SwapInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal inputs: %v", err)
	}

	wethAddress := references.Networks[params.ChainId].References["weth"]["address"]

	decimals, err := getERC20Decimals(params.ChainId, inputs.TokenOut)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	adjustedAmount, err := utils.StringToUint(inputs.AmountOut, *decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert swap amount to uint: %w", err)
	}

	inputs.AmountOut = adjustedAmount.String()

	// Try ETH ↔ WETH conversion first
	if txs, err := handleEthWethSwap(inputs, params, wethAddress); err != nil {
		return nil, err
	} else if txs != nil {
		return txs, nil
	}

	// Fall back to Bebop swap
	return handleBebopSwap(inputs, params)
}

func HandleWrap(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`
		Amount string `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal wrap inputs: %v", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert wrap amount to uint: %w", err)
	}

	wethAddress := references.Networks[params.ChainId].References["weth"]["address"]

	if strings.EqualFold(inputs.Token, wethAddress) {
		wethContract, err := weth_address.NewWethAddress(common.HexToAddress(wethAddress), params.Provider)
		if err != nil {
			return nil, utils.ErrContract(wethAddress)
		}

		calldata, err := wethContract.WethAddressTransactor.Withdraw(utils.BuildTransactionOpts(params.From, nil), amount)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		return []signature.Plug{{
			To:   common.HexToAddress(inputs.Token),
			Data: calldata.Data(),
		}}, nil
	}

	if common.HexToAddress(inputs.Token) == utils.NativeTokenAddress {
		wethContract, err := weth_address.NewWethAddress(common.HexToAddress(wethAddress), params.Provider)
		if err != nil {
			return nil, utils.ErrContract(wethAddress)
		}

		calldata, err := wethContract.WethAddressTransactor.Deposit(utils.BuildTransactionOpts(params.From, amount))
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		return []signature.Plug{{
			To:    common.HexToAddress(wethAddress),
			Value: amount,
			Data:  calldata.Data(),
		}}, nil
	}

	return nil, fmt.Errorf("token must be either ETH or WETH for wrapping/unwrapping")
}

func HandleConstraintPrice(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token     string `json:"token"`     // Address of token to check price
		Operator  int8   `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold string `json:"threshold"` // Price threshold in USD
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal price constraint inputs: %w", err)
	}

	// Convert token address to Llama API format
	tokenId := fmt.Sprintf("ethereum:%s", inputs.Token)

	// Get current price from Llama
	prices, err := llama.GetPrices([]string{tokenId})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token price: %w", err)
	}

	price, exists := prices[tokenId]
	if !exists {
		return nil, fmt.Errorf("price not found for token: %s", inputs.Token)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold float: %w", err)
	}

	switch inputs.Operator {
	case -1:
		if price.Price >= thresholdFloat {
			return nil, fmt.Errorf("current price $%.2f is not less than threshold $%.2f", price.Price, thresholdFloat)
		}
	case 1:
		if price.Price <= thresholdFloat {
			return nil, fmt.Errorf("current price $%.2f is not greater than threshold $%.2f", price.Price, thresholdFloat)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintBalance(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token     string `json:"token"`     // Address of token to check balance
		Address   string `json:"address"`   // Address to check balance for
		Operator  int    `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold string `json:"threshold"` // Balance threshold
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance constraint inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	threshold, err := utils.StringToUint(inputs.Threshold, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert threshold to uint: %w", err)
	}

	// Get token balance using contract call
	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}

	erc20Contract, err := erc_20.NewErc20(*token, provider)
	if err != nil {
		return nil, fmt.Errorf("failed to create ERC20 contract instance: %v", err)
	}

	callOpts := utils.BuildCallOpts(params.From, big.NewInt(0))
	balance, err := erc20Contract.BalanceOf(callOpts, common.HexToAddress(inputs.Address))
	if err != nil {
		return nil, fmt.Errorf("failed to get token balance: %w", err)
	}

	switch inputs.Operator {
	case -1:
		if balance.Cmp(threshold) >= 0 {
			return nil, fmt.Errorf("current balance %.2f is not less than threshold %.2f", balance, threshold)
		}
	case 1:
		if balance.Cmp(threshold) <= 0 {
			return nil, fmt.Errorf("current balance %.2f is not greater than threshold %.2f", balance, threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
