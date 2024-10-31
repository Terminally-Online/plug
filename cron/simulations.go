package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"solver/utils"

	"golang.org/x/exp/rand"
)

type ExecutionRequest struct {
	Id      string `json:"id"`
	Actions []struct {
		CategoryName string      `json:"categoryName"`
		ActionName   string      `json:"actionName"`
		Values       interface{} `json:"values"`
	} `json:"actions"`
	Workflow struct {
		Socket struct {
			Id            string `json:"id"`
			SocketAddress string `json:"socketAddress"`
		} `json:"socket"`
	} `json:"workflow"`
}

type ExecutionsRequest struct {
	Result struct {
		Data struct {
			Json []ExecutionRequest `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

type SimulationRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type SimulationsRequest struct {
	Json []SimulationRequest `json:"json"`
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

func GetSimulation(execution ExecutionRequest) (SimulationRequest, error) {
	statuses := []string{"success", "warning", "failure"}
	randomIndex := rand.Intn(len(statuses))

	response := SimulationRequest{
		Id:     execution.Id,
		Status: statuses[randomIndex],
	}
	return response, nil
}

func PostSimulations(simulations []SimulationRequest) {
	response := SimulationsRequest{
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
	executions, err := GetExecutions()
	if err != nil {
		log.Println(err.Error())
		return
	}

	var simulations []SimulationRequest
	for _, execution := range executions.Result.Data.Json {
		simulation, err := GetSimulation(execution)
		if err != nil {
			simulation = SimulationRequest{
				Id:     execution.Id,
				Status: "failure",
			}
		}
		simulations = append(simulations, simulation)
	}

	PostSimulations(simulations)
}
