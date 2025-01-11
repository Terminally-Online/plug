package simulate

import (
	"log"
	"solver/solver"
	"solver/utils"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetSimulation(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	simulator := New(solver.New())
	_, err = simulator.GetSimulation(ExecutionRequest{
		ChainId: 1,
		From:    "0x62180042606624f02D8A130dA8A3171e9b33894d",
		Inputs: []Input{
			{
				"protocol": "ens",
				"action":   "buy",
				"name":     "nftchance.eth",
				"maxPrice": "123456",
			},
		},
	})
	if err != nil {
		t.Errorf("GetSimulation(execution) = %v; want %v", err, nil)
	}
}
