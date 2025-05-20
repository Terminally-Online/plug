package streams

import (
	"context"
	"encoding/json"
	"log"
	"solver/internal/avs/config"
	"solver/internal/avs/services"
	"solver/internal/avs/types"
	"solver/internal/solver/signature"
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
				dataStr, ok := message.Values["data"].(string)
				if !ok {
					continue
				}

				var data map[string]any
				if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
					continue
				}

				livePlugsStr, ok := data["plugs"].(string)
				if !ok {
					continue
				}

				var livePlugs signature.LivePlugs
				if err := json.Unmarshal([]byte(livePlugsStr), &livePlugs); err != nil {
					continue
				}

				transactionHash, err := livePlugs.Execute()
				if err != nil {
					log.Printf("Error executing transaction: %v", err)
					if err := AckMessage(ctx, consumerGroup, message.ID, stream.Stream); err != nil {
						log.Printf("Error acknowledging message: %v", err)
					}
					continue
				}

				taskDefinitionId := 0
				if td, ok := data["taskDefinitionId"].(float64); ok {
					taskDefinitionId = int(td)
				}

				// livePlugsHash := signature.GetPlugsHash(livePlugs.Plugs)
				// livePlugsHashBytes := hexutil.Bytes(livePlugsHash[:])
				executionResponse := types.ExecutionResponse{
					ProofOfTask: transactionHash,
					// Data:             &livePlugsHashBytes,
					TaskDefinitionId: taskDefinitionId,
				}

				if err = services.RelayTask(executionResponse); err != nil {
					log.Printf("Error relaying task: %v", err)
					continue
				}

				if err := AckMessage(ctx, consumerGroup, message.ID, stream.Stream); err != nil {
					log.Printf("Error acknowledging message: %v", err)
				}
			}
		}
	}
}
