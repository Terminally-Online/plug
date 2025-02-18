package plug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"solver/internal/solver/simulation"
	"solver/internal/utils"
)

func PostNext() (simulation.SimulationDefinitions, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulation.simulateNext")
	response, err := utils.MakeHTTPRequest(
		url,
		"POST",
		map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    os.Getenv("PLUG_APP_API_KEY"),
		},
		nil,
		nil,
		simulation.SimulationDefinitions{},
	)
	if err != nil {
		return response, err
	}
	return response, nil
}

func PostSimulations(simulations []simulation.SimulationResponse) error {
	response := simulation.SimulationResponses{
		Json: simulations,
	}
	body, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = utils.MakeHTTPRequest(
		fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulation.simulated"),
		"POST",
		map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    os.Getenv("PLUG_APP_API_KEY"),
		},
		nil,
		bytes.NewReader(body),
		simulation.SimulationResponses{},
	)
	if err != nil {
		return err
	}

	return nil
}
