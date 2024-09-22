package collectors

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"solver/engine"
)

// Block represents an Ethereum block
type Block struct {
	Number     *big.Int
	Hash       string
	Timestamp  uint64
	Transactions []string
}

// BlockCollector implements the Collector interface for Ethereum blocks
type BlockCollector struct {
	key string
	client *ethclient.Client
}

// NewBlockCollector creates a new BlockCollector
func NewBlockCollector(key string, nodeURL string) (*BlockCollector, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}
	return &BlockCollector{key: key, client: client}, nil
}

// GetKey returns the collector's key
func (bc *BlockCollector) GetKey() string {
	return bc.key
}

// GetCollectionStream starts collecting Ethereum blocks and sends them to the stream
func (bc *BlockCollector) GetCollectionStream(stream chan<- engine.Collection) error {
	headers := make(chan *types.Header)
	sub, err := bc.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new headers: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Printf("Error in block subscription: %v", err)
				return
			case header := <-headers:
				block, err := bc.client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					log.Printf("Error fetching block: %v", err)
					continue
				}
				
				txHashes := make([]string, len(block.Transactions()))
				for i, tx := range block.Transactions() {
					txHashes[i] = tx.Hash().Hex()
				}

				blockData := Block{
					Number:     block.Number(),
					Hash:       block.Hash().Hex(),
					Timestamp:  block.Time(),
				}

				stream <- engine.Collection{
					Key:  bc.key,
					Data: blockData,
				}
			}
		}
	}()

	return nil
}
