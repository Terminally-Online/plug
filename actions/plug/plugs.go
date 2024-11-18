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
