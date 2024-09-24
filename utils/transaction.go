package utils

import (
	"math/big"
)

var (
	MinActions = 1
	MaxActions = 10
	
	NativeTransferGas = uint64(21000)
)

type Transaction struct {
	Transaction string   `json:"transaction"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       *big.Int `json:"value"`
	Gas         uint64   `json:"gas"`
}
