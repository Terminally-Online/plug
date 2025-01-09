package simulate

import (
	"encoding/json"
	"solver/solver"
	"solver/types"
)

type Input map[string]any

type ExecutionRequest struct {
	Id      string  `json:"id"`
	ChainId int     `json:"chainId"`
	From    string  `json:"from"`
	Inputs  []Input `json:"inputs"`
}

type TransactionRequest struct{}

type SimulationRequest struct {
	Id          string `json:"id"`
	Status      string `json:"status"`
	Error       string `json:"error,omitempty"`
	GasEstimate int    `json:"gasEstimate,omitempty"`
}

type SimulationsRequest struct {
	Json []SimulationRequest `json:"json"`
}
type Simulator struct {
	solver *solver.Solver
}

func New(solver *solver.Solver) *Simulator {
	return &Simulator{
		solver: solver,
	}
}

func (h *Simulator) GetSimulations(executions []ExecutionRequest) ([]SimulationRequest, error) {
	var simulations []SimulationRequest
	for _, execution := range executions {
		simulation, err := h.GetSimulation(execution)
		if err != nil {
			simulation = &SimulationRequest{
				Id:     execution.Id,
				Status: "failure",
				Error:  err.Error(),
			}
		}
		simulations = append(simulations, *simulation)
	}

	return simulations, nil
}

func (h *Simulator) GetSimulation(execution ExecutionRequest) (*SimulationRequest, error) {
	transactionsBatch := make([]*types.Transaction, 0)
	var breakOuter bool
	for _, input := range execution.Inputs {
		inputMap := map[string]interface{}{
			"protocol": input["protocol"],
			"action":   input["action"],
		}
		for k, v := range input {
			inputMap[k] = v
		}

		inputsJson, err := json.Marshal(inputMap)
		if err != nil {
			return &SimulationRequest{
				Id:     execution.Id,
				Status: "failure",
				Error:  err.Error(),
			}, nil
		}

		transactions, err := h.solver.GetTransaction(inputsJson, execution.ChainId, execution.From)
		if err != nil {
			return &SimulationRequest{
				Id:     execution.Id,
				Status: "warning",
				Error:  err.Error(),
			}, nil
		}

		// NOTE: Some plug actions have exclusive transactions that need to be run alone
		//       before the rest of the Plug can run. For this, we will just break out
		//       of the loop and execute any solo transactions that are needed for
		//       the rest of the batch to run in sequence.
		for _, transaction := range transactions {
			if transaction.Exclusive {
				// NOTE: Set the field to false to avoid tarnishing the response shape.
				transaction.Exclusive = false
				transactionsBatch = []*types.Transaction{transaction}
				breakOuter = true
				break
			}
		}

		if breakOuter {
			break
		}

		transactionsBatch = append(transactionsBatch, transactions...)
	}

	if len(transactionsBatch) == 0 {
		return &SimulationRequest{
			Id:     execution.Id,
			Status: "warning",
			Error:  "has no transactions to execute",
		}, nil
	}

	// TODO: This is where we would simulate the transaction and return a real
	//       simulation response however for now we are just returning everything
	//       as a success if it made it this far.

	return &SimulationRequest{
		Id:          execution.Id,
		Status:      "success",
		GasEstimate: 100000,
	}, nil
}
