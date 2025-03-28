package streams

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"solver/internal/avs/config"
	"solver/internal/avs/services"
	"solver/internal/avs/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const IntentStream = "circuit:intents"

func HandleStream(ctx context.Context) {
	consumerGroup := "circuit-operators"
	_, operatorAddress, err := config.GetAccount()
	if err != nil {
		log.Fatalf("Failed to get operator address: %v", err)
	}

	log.Printf("Starting stream consumer for operator: %s", operatorAddress)
	streamCh := Subscribe(ctx, consumerGroup, operatorAddress)

	for {
		select {
		case <-ctx.Done():
			return
		case stream := <-streamCh:
			for _, message := range stream.Messages {
				intentId, ok := message.Values["intent_id"].(string)
				if !ok {
					log.Println("Invalid message format: missing intent_id")
					continue
				}

				dataStr, ok := message.Values["data"].(string)
				if !ok {
					log.Println("Invalid message format: missing data")
					continue
				}

				var data map[string]any
				if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
					log.Printf("Error parsing data: %v", err)
					continue
				}

				log.Printf("Processing intent: %s", intentId)

				plugsStr, ok := data["plugs"].(string)
				if !ok {
					log.Printf("Missing Plugs data for intent %s", intentId)
					continue
				}

				var plugs signature.Plugs
				if err := json.Unmarshal([]byte(plugsStr), &plugs); err != nil {
					log.Printf("Error parsing Plugs data: %v", err)
					continue
				}

				isElectedLeader := services.IsElectedLeader(operatorAddress)
				if !isElectedLeader {
					log.Printf("Not elected leader for intent %s, skipping", intentId)
					continue
				}

				privateKey, err := crypto.HexToECDSA(config.PrivateKey)
				if err != nil {
					log.Printf("Error getting private key: %v", err)
					continue
				}

				signatureMessage := fmt.Sprintf("Request signed LivePlugs for intent %s", intentId)
				messageHash := crypto.Keccak256Hash([]byte(signatureMessage))

				sig, err := crypto.Sign(messageHash.Bytes(), privateKey)
				if err != nil {
					log.Printf("Error signing message: %v", err)
					continue
				}

				sig[64] += 27
				signatureHex := hexutil.Encode(sig)

				jsonBody, err := json.Marshal(map[string]any{
					"intentId":        intentId,
					"signedMessage":   signatureHex,
					"operatorAddress": operatorAddress,
				})
				if err != nil {
					continue
				}

				url := fmt.Sprintf("%s/solver/intent/sign", config.SolverUrl)
				response, err := utils.MakeHTTPRequest(
					url,
					"POST",
					map[string]string{
						"accept": "application/json",
					},
					nil,
					bytes.NewBuffer(jsonBody),
					struct {
						LivePlugs signature.LivePlugs `json:"livePlugs"`
					}{},
				)
				if err != nil {
					log.Printf("Error requesting signed LivePlugs: %v", err)
					continue
				}

				transactionHash, err := response.LivePlugs.Execute()
				if err != nil {
					log.Printf("Error executing transaction: %v", err)
					if err := AckMessage(ctx, consumerGroup, message.ID, stream.Stream); err != nil {
						log.Printf("Error acknowledging message: %v", err)
					}
					continue
				}

				log.Printf("Transaction executed with hash: %s", transactionHash)

				livePlugsHashData, err := response.LivePlugs.GetCallData()
				if err != nil {
					log.Printf("Error getting calldata hash: %v", err)
					continue
				}

				taskDefinitionId := 0
				if td, ok := data["taskDefinitionId"].(float64); ok {
					taskDefinitionId = int(td)
				}

				hashBytes := hexutil.Bytes(livePlugsHashData)
				executionResponse := types.ExecutionResponse{
					ProofOfTask:      transactionHash,
					Data:             &hashBytes,
					TaskDefinitionId: taskDefinitionId,
				}

				err = services.RelayTask(executionResponse)
				if err != nil {
					log.Printf("Error relaying task: %v", err)
					continue
				}

				log.Printf("Successfully relayed task execution for intent %s", intentId)

				if err := AckMessage(ctx, consumerGroup, message.ID, stream.Stream); err != nil {
					log.Printf("Error acknowledging message: %v", err)
				}
			}
		}
	}
}
