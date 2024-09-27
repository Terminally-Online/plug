package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"os"
)

var (
	SupportedChains = []int{1}
)

func GetProvider(chainId int) (*ethclient.Client, error) {
	if !IsSupportedChain(chainId) {
		return nil, ErrInvalidChainId("chainId", chainId)
	}

	err := godotenv.Load()
	if err != nil {
		return nil, ErrEnvironmentNotInitialized(err.Error())
	}

	alchemyAPIKey := os.Getenv("ALCHEMY_API_KEY")
	if alchemyAPIKey == "" {
		return nil, ErrEnvironmentVarNotSet("ALCHEMY_API_KEY")
	}

	var rpcURL string
	switch chainId {
	case 1:
		rpcURL = "wss://eth-mainnet.g.alchemy.com/v2/%v"
	case 11155111:
		rpcURL = "wss://eth-sepolia.g.alchemy.com/v2/%v"
	case 10:
		rpcURL = "wss://opt-mainnet.g.alchemy.com/v2/%v"
	case 11155420:
		rpcURL = "wss://opt-sepolia.g.alchemy.com/v2/%v"
	case 8453:
		rpcURL = "wss://base-mainnet.g.alchemy.com/v2/%v"
	case 84532:
		rpcURL = "wss://base-sepolia.g.alchemy.com/v2/%v"
	default:
		return nil, ErrInvalidProviderId("chainId", chainId)
	}

	ethClient, err := ethclient.Dial(fmt.Sprintf(rpcURL, alchemyAPIKey))
	if err != nil {
		return nil, ErrEthClientFailed(err.Error())
	}

	return ethClient, nil
}
