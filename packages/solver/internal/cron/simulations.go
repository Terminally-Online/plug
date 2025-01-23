package cron

import (
	"log"
	"solver/internal/solver"
)

func Simulations() {
	solverHandler := solver.New()

	executions, err := solverHandler.GetExecutions()
	if err != nil {
		return
	}

	var simulations []solver.SimulationRequest
	for _, execution := range executions.Result.Data.Json {
		transactions, err := solverHandler.GetTransactions(execution)
		if err != nil {
			log.Println(err.Error())
		}
		plugs, err := solverHandler.GetPlugs(execution.ChainId, execution.From, transactions)
		if err != nil {
			simulations = append(simulations, solver.SimulationRequest{
				Id:     execution.Id,
				Status: "failure",
				Error:  err.Error(),
			})
			continue
		}
		simulation, err := solverHandler.GetSimulation(execution.Id, plugs)
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

	if err := solverHandler.PostSimulations(simulations); err != nil {
		log.Println(err.Error())
	}
}
