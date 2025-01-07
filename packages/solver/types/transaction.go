package types

import "math/big"

// Transaction represents an Ethereum transaction
type Transaction struct {
	To        string      `json:"to"`
	Data      string      `json:"data"`
	Value     big.Int     `json:"value"`
	Exclusive bool        `json:"exclusive,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
}
