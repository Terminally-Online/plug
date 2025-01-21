package cron

import (
	"log"
	"solver/solver"
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
		simulation, err := solverHandler.GetSimulation(execution.Id, transactions)
		if err != nil {
			log.Println(err.Error())
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
