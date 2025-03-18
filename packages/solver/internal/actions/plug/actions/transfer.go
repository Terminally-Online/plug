package actions

import (
	"fmt"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/actions/options"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type TransferRequest struct {
	Amount    string         `json:"amount"`
	Token     string         `json:"token"`
	Recipient common.Address `json:"recipient"`
}

var TransferFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "transfer",
}

func TransferOptions(lookup *actions.SchemaLookup[TransferRequest]) (map[int]actions.Options, error) {
	fungiblesIndex := 1
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesIndex)
	if err != nil {
		return nil, err
	}

	recipientIndex := 2
	recipientOptions, err := options.GetAddressOptions(lookup, recipientIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		recipientIndex: {Simple: recipientOptions},
	}, nil
}

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

	amount, err := utils.StringToUint(lookup.Inputs.Amount, uint8(decimals))
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	if standard != 20 {
		return nil, utils.ErrNotImplemented("transfer support for 721 and 1155 are not yet implemented")
	}

	if token == utils.NativeTokenAddress {
		return []signature.Plug{{
			To:    lookup.Inputs.Recipient,
			Value: amount,
		}}, nil
	}

	calldata, err := TransferFunc.GetCalldata(lookup.Inputs.Recipient, amount)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:   common.HexToAddress(lookup.Inputs.Token),
		Data: calldata,
	}}, nil
}
