package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// DummyTransactOpts returns a TransactOpts with a dummy signer.
// This is useful for generating transaction data without actually signing the transaction.
func DummyTransactOpts() *bind.TransactOpts {
	return &bind.TransactOpts{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
		Context:   nil,
		NoSend:    true,
		GasLimit:  0,             // Let the node estimate
		GasPrice:  nil,           // Let the node estimate
		GasFeeCap: nil,           // Let the node estimate
		GasTipCap: nil,           // Let the node estimate
		Value:     big.NewInt(0), // No ETH sent by default
	}
}
