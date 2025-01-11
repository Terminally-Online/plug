package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"solver/simulate"
	"solver/solver"
	"solver/utils"
)

type ExecutionsRequest struct {
	Result struct {
		Data struct {
			Json []simulate.ExecutionRequest `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

type SimulationsResponse struct {
	Result struct {
		Data struct {
			Json []string `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

func GetExecutions() (ExecutionsRequest, error) {
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
		ExecutionsRequest{},
	)
	if err != nil {
		return response, err
	}
	return response, nil
}

func PostSimulations(simulations []simulate.SimulationRequest) {
	response := simulate.SimulationsRequest{
		Json: simulations,
	}
	body, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		return
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
		SimulationsResponse{},
	)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func Simulations() {
	solver := solver.New()
	simulator := simulate.New(solver)

	executions, err := GetExecutions()
	if err != nil {
		log.Println(err.Error())
		return
	}

	simulations, err := simulator.GetSimulations(executions.Result.Data.Json)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(simulations) > 0 {
		PostSimulations(simulations)
	}
}
