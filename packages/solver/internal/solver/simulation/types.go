package simulation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimulationDomain struct {
	ChainId uint64 `json:"chainId"`
	From    string `json:"from"`
}

type SimulationInputs struct {
	Inputs []map[string]any `json:"inputs"`
}

type SimulationOptions struct {
	Simulate bool `json:"simulate"` // Should the Plug be simulated.
	Submit   bool `json:"submit"`   // Should the Plug be run onchain.
	IsEOA    bool `json:"isEOA"`    // Is the Plug being run by an EOA.
}

type SimulationDefinition struct {
	Id string `json:"id"`
	SimulationDomain
	SimulationInputs
	Options SimulationOptions `json:"options"`
}

type SimulationDefinitions struct {
	Result struct {
		Data struct {
			Json []SimulationDefinition `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

type SimulationRequest struct {
	Id         string           `json:"id,omitempty"`
	ChainId    uint64           `json:"chainId"`
	From       common.Address   `json:"from"`
	To         common.Address   `json:"to"`
	Data       hexutil.Bytes    `json:"data,omitempty"`
	GasLimit   *uint64          `json:"gasLimit,omitempty"`
	Value      *big.Int         `json:"value,omitempty"`
	AccessList types.AccessList `json:"accessList,omitempty"`
	ABI        string           `json:"abi,omitempty"`
}

type SimulationGas struct {
	Used uint64 `json:"used"`
}

type SimulationResponse struct {
	Id           string        `json:"id,omitempty"`
	Gas          SimulationGas `json:"gas,omitempty"`
	Success      bool          `json:"success,omitempty"`
	Data         OutputData    `json:"data,omitempty"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
}

type SimulationResponses struct {
	Json []SimulationResponse `json:"json"`
}

type OutputData struct {
	Raw     []byte      `json:"raw"`
	Decoded interface{} `json:"decoded,omitempty"`
}
