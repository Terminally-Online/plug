package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/coil"
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
}

var TransferFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "transfer",
}

// Transfer handles token transfer requests by creating transaction signatures (Plugs) for both native and ERC20 tokens.
// It processes a TransferRequest which includes the token details (in format "address:decimals:standard"), amount, and recipient.
//
// The token parameter must follow the format "address:decimals:standard" where:
// - address: The token contract address (or native token address for ETH)
// - decimals: The number of decimal places the token uses
// - standard: The token standard (currently only supports 20 for ERC20)
//
// Inputs that may come in as linked inputs <-{coil_name} include:
// - Amount
// - Recipient
//
// For native token transfers, it creates a simple value transfer.
// For ERC20 tokens, it generates the appropriate transfer function calldata.
//
// Returns a slice of Plugs containing the transaction parameters and any error encountered.
func Transfer(lookup *actions.SchemaLookup[TransferRequest]) ([]signature.Plug, error) {
	parts := strings.Split(lookup.Inputs.Token, ":")
	token := common.HexToAddress(parts[0])
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid token format: %s, expected format 'address:decimals:standard'", lookup.Inputs.Token)
	}
	decimals, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return nil, err
	}
	standard, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}
	if standard != 20 {
		return nil, utils.ErrNotImplemented("transfer support for 721 and 1155 are not yet implemented")
	}

	var updates []coil.Update
	recipient, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Recipient,
		lookup.Inputs.Recipient.GetValueWithError,
		&TransferFunc,
		"_to",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	amountFunc := func() (*big.Int, error) {
		// Check if this is a linked value - if so, we don't need to convert it here and can just return placeholder 0
		if lookup.Inputs.Amount.GetIsLinked() {
			return big.NewInt(0), nil
		}

		// Otherwise convert the string amount to a uint
		amount, err := utils.StringToUint(lookup.Inputs.Amount.GetValue(), uint8(decimals))
		if err != nil {
			return nil, fmt.Errorf("failed to convert amount: %w", err)
		}
		return amount, nil
	}

	fmt.Printf("transfer lookup.Inputs.Amount: %v\n", lookup.Inputs.Amount)
	fmt.Printf("transfer previousActionDefinition: %v\n", lookup.PreviousActionDefinition)
	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		amountFunc,
		&TransferFunc,
		"_value",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	if token == utils.NativeTokenAddress {
		return []signature.Plug{{
			To:      recipient,
			Value:   amount,
			Updates: updates,
		}}, nil
	}

	calldata, err := TransferFunc.GetCalldata(recipient, amount)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:      common.HexToAddress(lookup.Inputs.Token),
		Data:    calldata,
		Updates: updates,
	}}, nil
}
