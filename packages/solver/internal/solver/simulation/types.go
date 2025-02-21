package simulation

import (
	"solver/internal/database/models"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimulationDomain struct {
	ChainId uint64         `json:"chainId"`
	From    common.Address `json:"from"`
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

// Use the models from the database package
type SimulationRequest = models.SimulationRequest
type SimulationResponse = models.SimulationResponse
type OutputData = models.OutputData

type SimulationResponses struct {
	Json []SimulationResponse `json:"json"`
}
