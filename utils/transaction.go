package utils

import (
	"math/big"
)

var (
	MinActions      = 1
	MaxActions      = 10
)

type Transaction struct {
	Transaction string   `json:"transaction"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       *big.Int `json:"value"`
}
