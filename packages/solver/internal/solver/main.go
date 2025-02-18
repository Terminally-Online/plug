package solver

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3"
	"solver/internal/actions/ens"
	"solver/internal/actions/euler"
	"solver/internal/actions/morpho"
	"solver/internal/actions/nouns"
	"solver/internal/actions/plug"
	"solver/internal/actions/yearn_v3"
	"solver/internal/client"
	"solver/internal/solver/call"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"
)

type Solver struct {
	Simulator simulation.Simulator
	Caller    call.Caller
	Protocols map[string]actions.BaseProtocolHandler
	IsKilled  bool
}

func New() Solver {
	return Solver{
		Simulator: simulation.New(),
		Caller:    call.New(),
		Protocols: map[string]actions.BaseProtocolHandler{
			actions.ProtocolPlug:    plug.New(),
			actions.ProtocolAaveV3:  aave_v3.New(),
			actions.ProtocolYearnV3: yearn_v3.New(),
			actions.ProtocolENS:     ens.New(),
			actions.ProtocolNouns:   nouns.New(),
			actions.ProtocolMorpho:  morpho.New(),
			actions.ProtocolEuler:   euler.New(),
		},
		IsKilled: false,
	}
}

func (s *Solver) GetTransaction(rawInputs json.RawMessage, chainId uint64, from string) ([]signature.Plug, error) {
	var inputs struct {
		Protocol string `json:"protocol"`
		Action   string `json:"action"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base inputs: %v", err)
	}

	handler, exists := s.Protocols[inputs.Protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", inputs.Protocol)
	}

	client, err := client.New(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err)
	}

	params := actions.HandlerParams{
		Client:  client,
		ChainId: chainId,
		From:    from,
	}

	transactions, err := handler.GetTransaction(inputs.Action, rawInputs, params)
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		if transactions[i].Value == nil {
			transactions[i].Value = big.NewInt(0)
		}
		// TODO: Only include the gas amount when we can properly estimate it with the traces
		//       that are generated from the simulation.
		transactions[i].Gas = big.NewInt(600000)
	}

	return transactions, nil
}

// GetPlugs processes transaction inputs and returns a slice of Plug signatures.
// It handles both regular and exclusive transactions, where exclusive transactions
// must be executed independently before other transactions in the batch.
//
// Parameters:
//
//	head: Existing slice of Plug signatures to append to
//	chainId: The blockchain network identifier
//	from: The sender's address
//	inputs: Raw JSON byte array containing transaction details
//
// Returns:
//
//	plugs: Slice of Plug signatures (either appended to head or a single exclusive transaction)
//	exclusive: Boolean indicating if an exclusive transaction was found
//	error: Any error encountered during processing
func (s *Solver) GetExclusivePlugs(
	head []signature.Plug,
	inputs []byte,
	chainId uint64,
	from string,
) (plugs []signature.Plug, exclusive bool, error error) {
	plugs, err := s.GetTransaction(inputs, chainId, from)
	if err != nil {
		return nil, false, err
	}

	// NOTE: Some plug actions have exclusive transactions that need to be run alone
	//       before the rest of the Plug can run. For this, we will just break out
	//       of the loop and execute any solo transactions that are needed for
	//       the rest of the batch to run in sequence.
	for _, plug := range plugs {
		if plug.Exclusive {
			// NOTE: Set the field to false to avoid tarnishing the response shape.
			plug.Exclusive = false
			return []signature.Plug{plug}, true, nil
		}
	}

	return append(head, plugs...), false, nil
}

// GetPlug processes a simulation definition to generate a sequence of transaction signatures.
// It iterates through the input definitions, converting them to Plug signatures while handling
// both regular and exclusive transactions.
//
// Parameters:
//   - definition: A SimulationDefinition containing chain ID, sender address, and transaction inputs
//
// Returns:
//   - []signature.Plug: A slice of transaction signatures ready for execution
//   - error: Any error encountered during processing, including validation errors
//
// The function will return early if it encounters an exclusive transaction, as these must be
// executed independently. It also validates that at least one transaction is generated.
func (s *Solver) GetPlugs(definition simulation.SimulationDefinition) ([]signature.Plug, error) {
	var plugs []signature.Plug
	for _, input := range definition.Inputs {
		inputsMap := map[string]interface{}{
			"protocol": input["protocol"],
			"action":   input["action"],
		}
		for k, v := range input {
			inputsMap[k] = v
		}
		inputs, err := json.Marshal(inputsMap)
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}

		var exclusive bool
		plugs, exclusive, err = s.GetExclusivePlugs(plugs, inputs, definition.ChainId, definition.From)
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}
		if exclusive {
			break
		}
	}

	// NOTE: If there was no transaction to execute we will return a warning because
	//		 we will be halting the simulation of this workflow.
	if len(plugs) == 0 {
		return nil, utils.ErrBuild("no transactions to execute")
	}

	return plugs, nil
}

// GetLivePlugs converts a simulation definition into executable transaction signatures.
// It first generates Plug signatures using GetPlugs, then converts them into LivePlugs
// which contain additional metadata needed for transaction execution.
//
// Parameters:
//   - definition: A SimulationDefinition containing chain ID, sender address, and transaction inputs
//
// Returns:
//   - signature.LivePlugs: A struct containing the executable transaction data
//   - error: Any error encountered during the conversion process
func (s *Solver) GetLivePlugs(definition simulation.SimulationDefinition) (signature.LivePlugs, error) {
	plugs, err := s.GetPlugs(definition)
	if err != nil {
		return signature.LivePlugs{}, err
	}
	livePlugs, err := signature.GetLivePlugs(definition.ChainId, definition.From, plugs)
	if err != nil {
		return signature.LivePlugs{}, err
	}

	return livePlugs, nil
}

// _, simulationResponse, err := s.Simulator.GetSimulationResponse(definition.Id, definition.ChainId, livePlugs)
// if err != nil {
// 	simulationResponses = append(simulationResponses, simulation.SimulationResponse{
// 		Success: false,
// 	})
// 	continue
// }
