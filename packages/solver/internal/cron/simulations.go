package cron

import (
	"log"
	"solver/internal/solver"
)

func Simulations(s *solver.Solver) {
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

	var simulations []solver.SimulationRequest
	for _, execution := range executions.Result.Data.Json {
		transactions, err := s.GetTransactions(execution)
		if err != nil {
			log.Println(err.Error())
		}
		plugs, err := s.GetPlugs(execution.ChainId, execution.From, transactions)
		if err != nil {
			simulations = append(simulations, solver.SimulationRequest{
				Id:     execution.Id,
				Status: "failure",
				Error:  err.Error(),
			})
			continue
		}
		simulation, err := s.GetSimulation(execution.Id, plugs)
		if err != nil {
			simulations = append(simulations, solver.SimulationRequest{
				Id:     execution.Id,
				Status: "failure",
				Error:  err.Error(),
			})
			continue
		}
		simulations = append(simulations, simulation)
	}
	if len(simulations) == 0 {
		return
	}

	if err := s.PostSimulations(simulations); err != nil {
		log.Println(err.Error())
	}
}
