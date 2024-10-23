package cron

import (
    "fmt"
    "os"
    "solver/utils"
)

type SimulationsResponse struct { 
    Result struct {
        Data struct {
            JSON []struct {
                ID      string `json:"id"`
                Actions []struct {
                    CategoryName string `json:"categoryName"`
                    ActionName  string `json:"actionName"`
                    Values     []struct {
                        Label string `json:"label"`
                        Value string `json:"value"`
                    } `json:"values"`
                } `json:"actions"`
                SocketID string `json:"socketId"`
            } `json:"json"`
        } `json:"data"`
    } `json:"result"`
}

func GetSimulations() (SimulationsResponse, error){ 
    url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulateNext")
    response, err := utils.MakeHTTPRequest(
        url,
        "POST",
        map[string]string{
            "Content-Type": "application/json",
            "X-API-Key": os.Getenv("PLUG_APP_API_KEY"),
        },
        nil,
        nil,
        SimulationsResponse{},
    )
    if err != nil {
        return response, err
    }
    return response, nil
}

func PostResults() { 
    // Submit the simulation results
}

func Simulations() { 
    // Get the simulations
    // Loop through the simulations
    // Submit the simulation results
}
