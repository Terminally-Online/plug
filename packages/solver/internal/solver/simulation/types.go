package simulation

import "github.com/ethereum/go-ethereum/core/types"

type Transaction struct {
	From       string           `json:"from"`
	To         string           `json:"to"`
	ChainId    uint64           `json:"chainId"`
	Value      string           `json:"value"`
	Data       string           `json:"data"`
	Gas        *string          `json:"gas,omitempty"`
	AccessList types.AccessList `json:"accessList,omitempty"`
}
