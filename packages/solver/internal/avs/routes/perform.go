package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"solver/internal/avs/config"
	"solver/internal/avs/services"
	"solver/internal/avs/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
)

func PerformTask(context *gin.Context) {
	if context.Request.Method != http.MethodPost {
		context.JSON(http.StatusMethodNotAllowed, gin.H{"error": "invalid method"})
	}

	electedLeader, err := services.ElectRoundRobin(uint64(1))
	if err != nil {
		log.Println(err)
		return
	}

	_, performerAddress, err := config.GetAccount()
	if err != nil {
		log.Println(err)
		return
	}

	if *electedLeader != common.HexToAddress(performerAddress) {
		return
	}

	taskDefinitionId := 0
	if context.Request.Body != http.NoBody {
		var request map[string]any
		if json.NewDecoder(context.Request.Body).Decode(&request) == nil {
			if val, ok := request["taskDefinition"].(int); ok {
				taskDefinitionId = val
			}
		}
	}

	log.Printf("taskDefinitionId: %v\n", taskDefinitionId)

	var transactionHash string
	var livePlugsHash hexutil.Bytes
	response := types.ExecutionResponse{
		ProofOfTask:      transactionHash,
		Data:             &livePlugsHash,
		TaskDefinitionId: taskDefinitionId,
	}

	if len(livePlugsHash) == 0 {
		log.Println("nothing to execute right now")
		return
	}

	services.RelayTask(response)

	context.JSON(http.StatusOK, response)
}
