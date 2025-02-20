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
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Solver struct {
	Protocols map[string]actions.BaseProtocolHandler
	IsKilled  bool
}

func New() Solver {
	return Solver{
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

func (s *Solver) GetPlugsArray(
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
		plugs, exclusive, err = s.GetPlugsArray(plugs, inputs, definition.ChainId, definition.From)
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

func (s *Solver) GetLivePlugs(definition simulation.SimulationDefinition) (signature.LivePlugs, error) {
	plugs, err := s.GetPlugs(definition)
	if err != nil {
		return signature.LivePlugs{}, err
	}
	solver, err := signature.GetSolverHash()
	if err != nil {
		return signature.LivePlugs{}, err
	}
	salt, err := signature.GetSaltHash(common.HexToAddress(definition.From))
	if err != nil {
		return signature.LivePlugs{}, err
	}

	plugsSigned, plugsSignature, err := signature.GetSignature(
		big.NewInt(int64(definition.ChainId)),
		common.HexToAddress(definition.From),
		signature.Plugs{
			Socket: common.HexToAddress(definition.From),
			Plugs:  plugs,
			Solver: solver,
			Salt:   salt,
		},
	)
	if err != nil {
		return signature.LivePlugs{}, utils.ErrBuild("failed to sign: " + err.Error())
	}

	return signature.LivePlugs{
		Plugs:     plugsSigned,
		Signature: plugsSignature,
	}, nil
}

func (s *Solver) Solve(definition simulation.SimulationDefinition) (solution *Solution, err error) {
	livePlugs, err := s.GetLivePlugs(definition)
	if err != nil {
		return nil, err
	}

	// NOTE: For this to be accurate we need the plugs already signed since that is all
	//       we can simulate meaning only livePlugs should be returned when we cannot
	//       realize the state of a simulated plug simulation defintion.
	var simulationRequest *simulation.SimulationRequest
	var simulationResponse *simulation.SimulationResponse
	if definition.Options.Simulate {
		simulationRequest, simulationResponse, err = simulation.Simulate(definition.Id, definition.ChainId, livePlugs)
		if err != nil {
			return nil, err
		}
	}

	return &Solution{
		Transactions: livePlugs.Plugs.Plugs, // Transactions in the `livePlug`.
		LivePlugs:    livePlugs,             // The `livePlug` included in the bundle.
		Transaction:  simulationRequest,     // Transaction the solver runs.
		Simulation:   simulationResponse,    // Simulation results of solver run.
	}, nil
}

func (s *Solver) Submit(definitions []simulation.SimulationDefinition) ([]signature.Result, error) {
	if len(definitions) == 0 {
		return nil, utils.ErrBuild("no plugs generated to execute")
	}

	chainId := definitions[0].ChainId
	errors := make([]error, len(definitions))

	var livePlugs []signature.LivePlugs
	for i, definition := range definitions {
		if definition.ChainId != chainId {
			errors[i] = utils.ErrChainId("chainId", definition.ChainId)
			continue

		}

		solution, err := s.Solve(definition)
		if err != nil {
			errors[i] = err
			continue
		}

		if definition.Options.Submit && solution.Simulation.Success {
			livePlugs = append(livePlugs, solution.LivePlugs)
		}
	}

	provider, err := client.New(chainId)
	if err != nil {
		return nil, err
	}
	results, err := provider.Plug(livePlugs)
	if err != nil {
		return nil, err
	}

	return results, nil
}
