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
	Amount  string `json:"amount"`
	Token   string `json:"token"`
	TokenIn string `json:"tokenIn"`
}

func HandleTransfer(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token     string         `json:"token"`
		Recipient common.Address `json:"recipient"`
		Amount    string         `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	parts := strings.Split(inputs.Token, ":")
	token := common.HexToAddress(parts[0])
	decimals, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return nil, err
	}
	standard, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}

	amount, err := utils.StringToUint(inputs.Amount, uint8(decimals))
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	if standard != 20 {
		return nil, utils.ErrNotImplemented("transfer support for 721 and 1155 are not yet implemented")
	}

	if token == utils.NativeTokenAddress {
		transaction := ethtypes.NewTransaction(
			0,
			inputs.Recipient,
			amount,
			utils.NativeTransferGas,
			big.NewInt(0),
			nil,
		)

		return []signature.Plug{{
			To:    inputs.Recipient,
			Value: transaction.Value(),
		}}, nil
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}

	calldata, err := erc20Abi.Pack("transfer",
		inputs.Recipient,
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
		strings.EqualFold(inputs.Token, wethAddress)
	isWethToEth := strings.EqualFold(inputs.TokenIn, wethAddress) &&
		common.HexToAddress(inputs.Token) == utils.NativeTokenAddress

	if !isEthToWeth && !isWethToEth {
		return nil, nil
	}

	amountOut, ok := new(big.Int).SetString(inputs.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amountOut: %s", inputs.Amount)
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
						Amount:         inputs.Amount,
						Decimals:       18,
						PriceUsd:       ethPrice,
						Symbol:         buySymbol,
						Price:          ethPrice,
						PriceBeforeFee: ethPrice,
					},
					AmountBeforeFee:   inputs.Amount,
					DeltaFromExpected: 0,
				},
			},
			SellTokens: map[string]BebopQuoteResponseToken{
				common.HexToAddress(inputs.Token).Hex(): {
					Amount:         inputs.Amount,
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

func handleBebopSwap(inputs SwapInputs, params actions.HandlerParams) ([]signature.Plug, error) {
	bebopApiUrl := fmt.Sprintf("https://api.bebop.xyz/pmm/base/v3/quote?buy_tokens=%s&sell_tokens=%s&sell_amounts=%s&taker_address=%s&gasless=false&approval_type=Standard&skip_validation=true",
		inputs.TokenIn,
		inputs.Token,
		inputs.Amount,
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

		amountOut, ok := new(big.Int).SetString(inputs.Amount, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse amountOut: %s", inputs.Amount)
		}

		approveCalldata, err := erc20Abi.Pack("approve",
			common.HexToAddress(quoteResponse.ApprovalTarget),
			amountOut,
		)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		transactions = append([]signature.Plug{{
			To:   common.HexToAddress(inputs.Token),
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

	tokenOutParts := strings.Split(inputs.Token, ":")
	inputs.Token = tokenOutParts[0]
	decimals, err := strconv.ParseUint(tokenOutParts[1], 10, 8)
	if err != nil {
		return nil, err
	}
	if standard, err := strconv.ParseUint(tokenOutParts[2], 10, 64); err != nil || standard != 20 {
		return nil, utils.ErrNotImplemented("support for 721 and 1155 are not yet implemented")
	}

	tokenInParts := strings.Split(inputs.TokenIn, ":")
	inputs.TokenIn = tokenInParts[0]

	wethAddress := references.Networks[params.ChainId].References["weth"]["address"]
	adjustedAmount, err := utils.StringToUint(inputs.Amount, uint8(decimals))
	if err != nil {
		return nil, fmt.Errorf("failed to convert swap amount to uint: %w", err)
	}
	inputs.Amount = adjustedAmount.String()

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
		wethContract, err := weth_address.NewWethAddress(common.HexToAddress(wethAddress), params.Client)
		if err != nil {
			return nil, utils.ErrContract(wethAddress)
		}

		calldata, err := wethContract.WethAddressTransactor.Withdraw(params.Client.WriteOptions(params.From, big.NewInt(0)), amount)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}

		return []signature.Plug{{
			To:   common.HexToAddress(inputs.Token),
			Data: calldata.Data(),
		}}, nil
	}

	if common.HexToAddress(inputs.Token) == utils.NativeTokenAddress {
		wethContract, err := weth_address.NewWethAddress(common.HexToAddress(wethAddress), params.Client)
		if err != nil {
			return nil, utils.ErrContract(wethAddress)
		}

		calldata, err := wethContract.WethAddressTransactor.Deposit(params.Client.WriteOptions(params.From, amount))
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
		Token     string `json:"token"`
		Operator  int8   `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal price constraint inputs: %w", err)
	}

	tokenId := fmt.Sprintf("ethereum:%s", inputs.Token)
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
		Token     string `json:"token"`
		Address   string `json:"address"`
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance constraint inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, err
	}

	threshold, err := utils.StringToUint(inputs.Threshold, decimals)
	if err != nil {
		return nil, err
	}

	erc20Contract, err := erc_20.NewErc20(*token, params.Client)
	if err != nil {
		return nil, err
	}

	balance, err := erc20Contract.BalanceOf(params.Client.ReadOptions(params.From), common.HexToAddress(inputs.Address))
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
