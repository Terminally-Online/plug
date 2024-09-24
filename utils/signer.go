package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// DummyTransactOpts returns a TransactOpts with a dummy signer.
// This is useful for generating transaction data without actually signing the transaction.
func DummyTransactOpts(address string, value *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: common.HexToAddress(address),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
		NoSend:    true,
		Value:     value,
	}
}
