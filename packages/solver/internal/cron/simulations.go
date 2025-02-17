package cron

import (
	"log"
	"solver/internal/solver"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
)

func Simulations(s solver.Solver) {
	// NOTE: If the solver has had its kill switch toggled prevent the running any
	//       new simulation processes that would retrieve, build and execute the
	//       active executions that the application endpoint provides.
	if s.IsKilled {
		return
	}

	next, err := s.Simulator.GetNext()
	if err != nil {
		return
	}

	var simulationResponses []simulation.SimulationResponse
	for _, definition := range next.Result.Data.Json {
		plugs, err := s.GetPlugs(definition)
		if err != nil {
			log.Println(err.Error())
		}
		livePlugs, err := signature.GetLivePlugs(definition.ChainId, definition.From, plugs)
		if err != nil {
			simulationResponses = append(simulationResponses, simulation.SimulationResponse{
				Success: false,
			})
			continue
		}
		_, simulationResponse, err := s.Simulator.GetSimulationResponse(definition.Id, definition.ChainId, livePlugs)
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

	if err := s.Simulator.PostSimulations(simulationResponses); err != nil {
		log.Println(err.Error())
	}
}
