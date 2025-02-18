package simulation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimulationDefinition struct {
	Id      string           `json:"id"`
	ChainId uint64           `json:"chainId"`
	From    string           `json:"from"`
	Inputs  []map[string]any `json:"inputs"`
}

type SimulationDefinitions struct {
	Result struct {
		Data struct {
			Json []SimulationDefinition `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

type SimulationRequest struct {
	ExecutionId string           `json:"id,omitempty"`
	ChainId     uint64           `json:"chainId"`
	From        common.Address   `json:"from"`
	To          common.Address   `json:"to"`
	Data        hexutil.Bytes    `json:"data,omitempty"`
	GasLimit    *uint64          `json:"gasLimit,omitempty"`
	Value       *big.Int         `json:"value,omitempty"`
	AccessList  types.AccessList `json:"accessList,omitempty"`
	ABI         string           `json:"abi,omitempty"`
}

type SimulationResponse struct {
	ExecutionId  string     `json:"id"`
	GasUsed      uint64     `json:"gasUsed"`
	Success      bool       `json:"success"`
	Data         OutputData `json:"data"`
	ErrorMessage string     `json:"errorMessage,omitempty"`
}

type SimulationResponses struct {
	Json []SimulationResponse `json:"json"`
}

type OutputData struct {
	Raw     []byte      `json:"raw"`
	Decoded interface{} `json:"decoded,omitempty"`
}
