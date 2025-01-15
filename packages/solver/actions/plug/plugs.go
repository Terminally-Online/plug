package plug

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"solver/actions"
	"solver/bindings/erc_20"
	"solver/bindings/weth_address"
	"solver/types"
	"solver/utils"
	"strconv"
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

	log.Printf("HandleSwap inputs: TokenIn=%s, TokenOut=%s, AmountOut=%s", inputs.TokenIn, inputs.TokenOut, inputs.AmountOut)

	// Get WETH address from references
	wethAddress := utils.Mainnet.References["weth"]["address"]
	log.Printf("WETH address from references: %s", wethAddress)

	// Check if this is a direct ETH ↔ WETH conversion
	isEthToWeth := common.HexToAddress(inputs.TokenIn) == utils.NativeTokenAddress &&
		strings.EqualFold(inputs.TokenOut, wethAddress)
	isWethToEth := strings.EqualFold(inputs.TokenIn, wethAddress) &&
		common.HexToAddress(inputs.TokenOut) == utils.NativeTokenAddress

	log.Printf("Conversion check: isEthToWeth=%v, isWethToEth=%v", isEthToWeth, isWethToEth)

	if isEthToWeth || isWethToEth {
		log.Printf("Processing direct ETH ↔ WETH conversion")
		amountOut, ok := new(big.Int).SetString(inputs.AmountOut, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse amountOut: %s", inputs.AmountOut)
		}
		log.Printf("Parsed amountOut: %s", amountOut.String())

		wethAbi, err := weth_address.WethAddressMetaData.GetAbi()
		if err != nil {
			log.Printf("Failed to get WETH ABI: %v", err)
			return nil, utils.ErrABIFailed("WETH")
		}
		var tx *types.Transaction
		if isEthToWeth {
			log.Printf("Processing ETH → WETH conversion")
			// For ETH → WETH, we deposit ETH into WETH contract
			// The deposit function doesn't take any parameters, it uses msg.value
			calldata, err := wethAbi.Pack("deposit")
			if err != nil {
				log.Printf("Failed to create deposit calldata: %v", err)
				return nil, utils.ErrTransactionFailed(err.Error())
			}
			log.Printf("Created deposit calldata: %s", common.Bytes2Hex(calldata))
			tx = &types.Transaction{
				To:    wethAddress,
				Value: *amountOut,
				Data:  "0x" + common.Bytes2Hex(calldata),
			}
		} else {
			log.Printf("Processing WETH → ETH conversion")
			// For WETH → ETH, we withdraw from WETH contract
			calldata, err := wethAbi.Pack("withdraw", amountOut)
			if err != nil {
				log.Printf("Failed to create withdraw calldata: %v", err)
				return nil, utils.ErrTransactionFailed(err.Error())
			}
			log.Printf("Created withdraw calldata: %s", common.Bytes2Hex(calldata))
			tx = &types.Transaction{
				To:   wethAddress,
				Data: "0x" + common.Bytes2Hex(calldata),
			}
		}

		log.Printf("Created transaction: To=%s, Value=%s, Data=%s", tx.To, tx.Value.String(), tx.Data)

		// Set the symbols based on the conversion direction
		buySymbol := "ETH"
		sellSymbol := "WETH"
		if isEthToWeth {
			buySymbol = "WETH"
			sellSymbol = "ETH"
		}

		// Get ETH price from references - both ETH and WETH have the same price
		ethPrice, err := strconv.ParseFloat(utils.Mainnet.References["eth"]["price"], 64)
		if err != nil {
			log.Printf("Failed to parse ETH price, using default: %v", err)
			ethPrice = 2000.0 // fallback price if parsing fails
		}

		// Return in the same format as a Bebop swap
		result := []*types.Transaction{{
			To:    tx.To,
			Data:  tx.Data,
			Value: tx.Value,
			Meta: BebopTransactionMeta{
				Expiry:      0,   // Instant execution
				Slippage:    0.0, // No slippage on direct conversion
				PriceImpact: 0.0, // No price impact on direct conversion
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
				PartnerFeeNative:   "0", // No fee for direct conversion
			},
		}}
		log.Printf("Returning transaction result: %+v", result)
		return result, nil
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

func HandleWrap(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Token  string  `json:"token"`
		Amount big.Int `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal wrap inputs: %v", err)
	}

	wethAddress := utils.Mainnet.References["weth"]["address"]

	if strings.EqualFold(inputs.Token, wethAddress) {
		wethContract, err := weth_address.NewWethAddress(common.HexToAddress(wethAddress), params.Provider)
		if err != nil {
			return nil, utils.ErrContractFailed(wethAddress)
		}

		calldata, err := wethContract.WethAddressTransactor.Withdraw(utils.BuildTransactionOpts(params.From, nil), &inputs.Amount)
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}

		return []*types.Transaction{{
			To:   inputs.Token,
			Data: "0x" + common.Bytes2Hex(calldata.Data()),
		}}, nil
	}

	if common.HexToAddress(inputs.Token) == utils.NativeTokenAddress {
		wethContract, err := weth_address.NewWethAddress(common.HexToAddress(wethAddress), params.Provider)
		if err != nil {
			return nil, utils.ErrContractFailed(wethAddress)
		}

		calldata, err := wethContract.WethAddressTransactor.Deposit(utils.BuildTransactionOpts(params.From, &inputs.Amount))
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}

		return []*types.Transaction{{
			To:    wethAddress,
			Value: inputs.Amount,
			Data:  "0x" + common.Bytes2Hex(calldata.Data()),
		}}, nil
	}

	return nil, fmt.Errorf("token must be either ETH or WETH for wrapping/unwrapping")
}
