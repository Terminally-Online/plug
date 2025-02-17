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
	"solver/internal/solver/call"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"
	"github.com/ethereum/go-ethereum/common"
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

	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %v", err)
	}

	params := actions.HandlerParams{
		Provider: provider,
		ChainId:  chainId,
		From:     from,
	}

	transactions, err := handler.GetTransaction(inputs.Action, rawInputs, params)
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		if transactions[i].Value == nil {
			transactions[i].Value = big.NewInt(0)
		}
		transactions[i].Gas = big.NewInt(200000)
	}

	return transactions, nil
}

func (s *Solver) GetTransactions(definition simulation.SimulationDefinition) ([]signature.Plug, error) {
	var breakOuter bool
	transactionsBatch := make([]signature.Plug, 0)
	errors := make([]error, len(definition.Inputs))
	for i, input := range definition.Inputs {
		inputMap := map[string]interface{}{
			"protocol": input["protocol"],
			"action":   input["action"],
		}
		for k, v := range input {
			inputMap[k] = v
		}

		inputsJson, err := json.Marshal(inputMap)
		if err != nil {
			errors[i] = err
			continue
		}

		transactions, err := s.GetTransaction(inputsJson, definition.ChainId, definition.From)
		if err != nil {
			errors[i] = err
			continue
		}

		// NOTE: Some plug actions have exclusive transactions that need to be run alone
		//       before the rest of the Plug can run. For this, we will just break out
		//       of the loop and execute any solo transactions that are needed for
		//       the rest of the batch to run in sequence.
		for _, transaction := range transactions {
			if transaction.Exclusive {
				// NOTE: Set the field to false to avoid tarnishing the response shape.
				transaction.Exclusive = false
				transactionsBatch = []signature.Plug{transaction}
				breakOuter = true
				break
			}
		}

		if breakOuter {
			break
		}

		transactionsBatch = append(transactionsBatch, transactions...)
	}

	// If there were any errors we will return a failure.
	for _, err := range errors {
		if err != nil {
			return nil, utils.ErrBuild(err.Error())
		}
	}

	// If there was no transaction to execute we will return a warning because
	// we will be halting the simulation of this workflow.
	if len(transactionsBatch) == 0 {
		return nil, utils.ErrBuild("no transactions to execute")
	}

	return transactionsBatch, nil
}

func (s *Solver) GetPlugs(chainId uint64, from string, transactions []signature.Plug) (*signature.LivePlugs, error) {
	solver, err := signature.GetSolverHash()
	if err != nil {
		return nil, err
	}
	salt, err := signature.GetSaltHash(common.HexToAddress(from))
	if err != nil {
		return nil, err
	}
	plugs, plugsSignature, err := signature.GetSignature(
		big.NewInt(int64(chainId)),
		common.HexToAddress(from),
		signature.Plugs{
			Socket: common.HexToAddress(from),
			Plugs:  transactions,
			Solver: solver,
			Salt:   salt,
		},
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to sign: " + err.Error())
	}

	return &signature.LivePlugs{
		Plugs:     plugs,
		Signature: plugsSignature,
	}, nil
}

func (s *Solver) GetRun(transactions []signature.Plug) error {
	// TODO: Run the transactions through the entrypoint with our executor account.
	return nil
}

