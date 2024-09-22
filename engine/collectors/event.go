package collectors

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"solver/engine"
)

type EventCollection struct {
	Number *big.Int
	Logs   []types.Log
}

type EventFilter struct {
	Addresses common.Address
	Topics []common.Hash
}

type EventCollector struct {
	key     string
	client  *ethclient.Client
	filters []EventFilter
}

func NewEventCollector(key string, client *ethclient.Client, filters []EventFilter) *EventCollector {
	return &EventCollector{
		key: key, 
		client: client, 
		filters: filters,
	}
}

func (ec *EventCollector) GetKey() string {
	return ec.key
}

func (ec *EventCollector) GetCollectionStream(ctx context.Context, networkName string, stream chan<- engine.Collection) error {
	log.Printf("Starting event collection for network %s", networkName)

	headers := make(chan *types.Header)
	sub, err := ec.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new headers: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %v", err)
		case header := <-headers:
			query := ethereum.FilterQuery{
				FromBlock: header.Number,
				ToBlock:   header.Number,
			}

			if len(ec.filters) > 0 {
				var addresses []common.Address
				var topics [][]common.Hash
				for _, filter := range ec.filters {
					addresses = append(addresses, filter.Addresses)
					topics = append(topics, filter.Topics)
				}
				query.Addresses = addresses
				query.Topics = topics
			}

			logs, err := ec.client.FilterLogs(ctx, query)
			if err != nil {
				log.Printf("Error fetching logs: %v", err)
				continue
			}

			stream <- engine.Collection{
				NetworkName: networkName,
				Key:         ec.key,
				Data: EventCollection{
					Number: header.Number,
					Logs:   logs,
				},
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
