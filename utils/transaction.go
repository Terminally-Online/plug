package utils

import (
	"math/big"
)

type Transaction struct {
	Transaction string   `json:"transaction"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       *big.Int `json:"value"`
}
