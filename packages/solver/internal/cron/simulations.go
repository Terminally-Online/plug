package cron

import (
	"log"
	"solver/internal/solver"
	"solver/internal/solver/simulation"
)

func Simulations(s solver.Solver) {
	// NOTE: If the solver has had its kill switch toggled prevent the running any
	//       new simulation processes that would retrieve, build and execute the
	//       active executions that the application endpoint provides.
	if s.IsKilled {
		return
	}

	executions, err := s.GetExecutions()
	if err != nil {
		return
	}

	var simulationResponses []simulation.SimulationResponse
	for _, execution := range executions.Result.Data.Json {
		transactions, err := s.GetTransactions(execution)
		if err != nil {
			log.Println(err.Error())
		}
		plugs, err := s.GetPlugs(execution.ChainId, execution.From, transactions)
		if err != nil {
			simulationResponses = append(simulationResponses, simulation.SimulationResponse{
				Success: false,
			})
			continue
		}
		_, simulationResponse, err := s.GetSimulation(execution.Id, execution.ChainId, plugs)
		if err != nil {
			simulationResponses = append(simulationResponses, simulation.SimulationResponse{
				Success: false,
			})
			continue
		}

		simulationResponses = append(simulationResponses, simulationResponse)
	}
	if len(simulationResponses) == 0 {
		return
	}

	if err := s.PostSimulations(simulationResponses); err != nil {
		log.Println(err.Error())
	}
}
