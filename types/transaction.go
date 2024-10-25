package types

import "math/big"

// Transaction represents an Ethereum transaction
type Transaction struct {
	To    *string `json:"to"`
	Data  []byte  `json:"data"`
	Value big.Int `json:"value"`
}
