package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	DummyNonce = uint64(0)
)

// BuildTransactionOpts returns a TransactOpts with a dummy signer.
// This is useful for generating transaction data without actually signing the transaction.
func BuildTransactionOpts(address string, value *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: common.HexToAddress(address),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
		NoSend:    true,
		Value:     value,
	}
}

func BuildCallOpts(address string, value *big.Int) *bind.CallOpts {
	return &bind.CallOpts{
		From: common.HexToAddress(address),
		Pending: true,
		Context: context.Background(),
	}
}