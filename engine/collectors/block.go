package collectors

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"solver/engine"
)

type BlockCollection struct {
	Number    *big.Int
	Hash      string
	Timestamp uint64
}

type BlockCollector struct {
	key    string
	client *ethclient.Client
}

func NewBlockCollector(key string, client *ethclient.Client) *BlockCollector {
	return &BlockCollector{key: key, client: client}
}

func (bc *BlockCollector) GetKey() string {
	return bc.key
}

func (bc *BlockCollector) GetCollectionStream(ctx context.Context, networkName string, stream chan<- engine.Collection) error {
	log.Printf("Starting block collection for network %s", networkName)

	headers := make(chan *types.Header)
	sub, err := bc.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new headers: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %v", err)
		case header := <-headers:
			block, err := bc.client.BlockByHash(ctx, header.Hash())
			if err != nil {
				log.Printf("Error fetching block: %v", err)
				continue
			}

			stream <- engine.Collection{
				NetworkName: networkName,
				Key:         bc.key,
				Data: BlockCollection{
					Number:    block.Number(),
					Hash:      block.Hash().Hex(),
					Timestamp: block.Time(),
				},
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
