package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/erc_1155"
	"solver/bindings/erc_20"
	"solver/bindings/erc_721"
	"solver/internal/actions"
	"solver/internal/coil"
	"solver/internal/helpers/squid"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type TransferRequest struct {
	Amount    coil.CoilInput[string, *big.Int]               `json:"amount"`
	Token     string                                         `json:"token"`
	Recipient coil.CoilInput[common.Address, common.Address] `json:"recipient"`
	ToChainId uint64                                         `json:"toChainId"`
	ToToken   common.Address                                 `json:"toToken"`
	ToAddress common.Address                                 `json:"toAddress"`
}

var TransferERC20Func = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "transfer",
}

var TransferERC721Func = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_721.Erc721MetaData,
	FunctionName: "safeTransferFrom",
}

var Transfer1155Func = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_1155.Erc1155MetaData,
	FunctionName: "safeTransferFrom",
}

// Transfer handles token transfer requests by creating transaction signatures (Plugs) for native, ERC20, ERC1155, and ERC721 tokens.
// It processes a TransferRequest which includes the token details (in format "address:decimals:standard"), amount, and recipient.
//
// The token parameter must follow the format "address:decimals:standard" where:
// - address: The token contract address (or native token address for ETH)
// - decimals: The number of decimal places the token uses
// - standard: The token standard
//
// Inputs that may come in as linked inputs <-{coil_name} include:
// - Amount
// - Recipient
//
// For native token transfers, it creates a simple value transfer.
// For ERC tokens, it generates the appropriate transfer function calldata.
//
// Returns a slice of Plugs containing the transaction parameters and any error encountered.
func Transfer(lookup *actions.SchemaLookup[TransferRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid token format: %s, expected format 'address:decimals:standard'", lookup.Inputs.Token)
	}

	token := common.HexToAddress(parts[0])
	if token == utils.NativeTokenAddress {
		return TransferNative(lookup)
	}

	standard, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}

	switch standard {
	case 20:
		return Transfer20(lookup)
	case 721:
		return Transfer721(lookup)
	case 1155:
		return Transfer1155(lookup)
	default:
		return nil, fmt.Errorf("unsupported token standard: %d", standard)
	}
}

func TransferNative(lookup *actions.SchemaLookup[TransferRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	decimals, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return nil, err
	}

	var updates []coil.Update
	recipient, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Recipient,
		lookup.Inputs.Recipient.GetValueWithError,
		&TransferERC20Func,
		"_to",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&TransferERC20Func,
		"_value",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	isBridge := lookup.Inputs.ToChainId != 0 && lookup.Inputs.ToChainId != lookup.ChainId
	if isBridge {
		if len(updates) > 0 {
			return nil, fmt.Errorf("cannot use coils with crosschain action")
		}

		return squid.GetPlugs(squid.SquidRouteRequest{
			FromAddress: lookup.From,
			FromChain:   lookup.ChainId,
			FromToken:   utils.NativeTokenAddress,
			FromAmount:  amount,
			ToChain:     lookup.Inputs.ToChainId,
			ToToken:     lookup.Inputs.ToToken,
			ToAddress:   lookup.Inputs.ToAddress,
		})
	}

	return []signature.Plug{{
		To:      recipient,
		Value:   amount,
		Updates: updates,
	}}, nil
}

func Transfer20(lookup *actions.SchemaLookup[TransferRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	token := common.HexToAddress(parts[0])
	decimals, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return nil, err
	}

	var updates []coil.Update
	recipient, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Recipient,
		lookup.Inputs.Recipient.GetValueWithError,
		&TransferERC20Func,
		"_to",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&TransferERC20Func,
		"_value",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	isBridge := lookup.Inputs.ToChainId != 0 && lookup.Inputs.ToChainId != lookup.ChainId
	if isBridge {
		if len(updates) > 0 {
			return nil, fmt.Errorf("cannot use coils with crosschain action")
		}

		return squid.GetPlugs(squid.SquidRouteRequest{
			FromAddress: lookup.From,
			FromChain:   lookup.ChainId,
			FromToken:   token,
			FromAmount:  amount,
			ToChain:     lookup.Inputs.ToChainId,
			ToToken:     lookup.Inputs.ToToken,
			ToAddress:   lookup.Inputs.ToAddress,
		})
	}

	transferCalldata, err := TransferERC20Func.GetCalldata(recipient, amount)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:      token,
		Data:    transferCalldata,
		Updates: updates,
	}}, nil
}

func Transfer721(lookup *actions.SchemaLookup[TransferRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	token := common.HexToAddress(parts[0])
	tokenId, ok := new(big.Int).SetString(parts[1], 10)
	if !ok {
		return nil, fmt.Errorf("could not parse tokenId")
	}

	var updates []coil.Update
	recipient, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Recipient,
		lookup.Inputs.Recipient.GetValueWithError,
		&TransferERC721Func,
		"to",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	calldata, err := TransferERC721Func.GetCalldata(lookup.From, recipient, tokenId)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:      token,
		Data:    calldata,
		Updates: updates,
	}}, nil
}

func Transfer1155(lookup *actions.SchemaLookup[TransferRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	token := common.HexToAddress(parts[0])
	tokenId, ok := new(big.Int).SetString(parts[1], 10)
	if !ok {
		return nil, fmt.Errorf("could not parse tokenId")
	}

	var updates []coil.Update
	recipient, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Recipient,
		lookup.Inputs.Recipient.GetValueWithError,
		&Transfer1155Func,
		"_to",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(0),
		&Transfer1155Func,
		"value",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	calldata, err := Transfer1155Func.GetCalldata(
		lookup.From,
		recipient,
		tokenId,
		amount,
		[]byte("plug"),
	)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:      token,
		Data:    calldata,
		Updates: updates,
	}}, nil
}
