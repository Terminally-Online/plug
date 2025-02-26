package cron

import (
	// "log"
	// "solver/internal/helpers/plug"
	"solver/internal/solver"
	// "solver/internal/solver/simulation"
)

func Simulations(s solver.Solver) {
	// if s.IsKilled {
	// 	return
	// }

	// next, err := plug.PostNext()
	// if err != nil {
	// 	return
	// }
	//
	// var simulationResponses []simulation.SimulationResponse
	// for _, definition := range next.Result.Data.Json {
	// 	solution, err := s.Solve(definition)
	// 	simulationResponses = append(simulationResponses, simulation.SimulationResponse{
	// 		Success: err == nil && solution.Simulation.Success,
	// 	})
	// }
	// if len(simulationResponses) == 0 {
	// 	return
	// }
	//
	// if err := plug.PostSimulations(simulationResponses); err != nil {
	// 	log.Println(err.Error())
	// }
}
