package simulation

import "github.com/ethereum/go-ethereum/core/types"

type Transaction struct {
	From       string           `json:"from"`
	To         string           `json:"to"`
	ChainId    uint64           `json:"chainId"`
	Value      string           `json:"value"`
	Data       string           `json:"data"`
	Gas        *string          `json:"gas"`
	AccessList types.AccessList `json:"accessList"`
}

// import (
// 	"solver/internal/database/models"

// 	"github.com/ethereum/go-ethereum/common"
// )

// type SimulationDomain struct {
// 	ChainId uint64         `json:"chainId"`
// 	From    common.Address `json:"from"`
// }

// type SimulationInputs struct {
// 	Inputs []map[string]any `json:"inputs"`
// }

// type SimulationOptions struct {
// 	Simulate bool `json:"simulate"` // Should the Plug be simulated.
// 	Submit   bool `json:"submit"`   // Should the Plug be run onchain.
// 	IsEOA    bool `json:"isEOA"`    // Is the Plug being run by an EOA.
// }

// type SimulationDefinition struct {
// 	Id string `json:"id"`
// 	SimulationDomain
// 	SimulationInputs
// 	Options SimulationOptions `json:"options"`
// }

// type SimulationDefinitions struct {
// 	Result struct {
// 		Data struct {
// 			Json []SimulationDefinition `json:"json"`
// 		} `json:"data"`
// 	} `json:"result"`
// }

// type SimulationRequest = models.SimulationRequest
// type SimulationResponse = models.SimulationResponse
// type SimulationOutputData = models.SimulationOutputData

// type SimulationResponses struct {
// 	Json []SimulationResponse `json:"json"`
// }
