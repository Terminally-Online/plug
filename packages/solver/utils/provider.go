package utils

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetProvider(chainId int) (*ethclient.Client, error) {
	alchemyAPIKey := os.Getenv("ALCHEMY_API_KEY")
	if alchemyAPIKey == "" {
		return nil, ErrEnvironmentVarNotSet("ALCHEMY_API_KEY")
	}

	var rpcURL string
	switch chainId {
	case 1:
		rpcURL = fmt.Sprintf("wss://eth-mainnet.g.alchemy.com/v2/%v", alchemyAPIKey)
	case 31337:
		rpcURL = "http://127.0.0.1:8545"
	case 11155111:
		rpcURL = fmt.Sprintf("wss://eth-sepolia.g.alchemy.com/v2/%v", alchemyAPIKey)
	case 10:
		rpcURL = fmt.Sprintf("wss://opt-mainnet.g.alchemy.com/v2/%v", alchemyAPIKey)
	case 11155420:
		rpcURL = fmt.Sprintf("wss://opt-sepolia.g.alchemy.com/v2/%v", alchemyAPIKey)
	case 8453:
		rpcURL = fmt.Sprintf("wss://base-mainnet.g.alchemy.com/v2/%v", alchemyAPIKey)
	case 84532:
		rpcURL = fmt.Sprintf("wss://base-sepolia.g.alchemy.com/v2/%v", alchemyAPIKey)
	default:
		return nil, ErrInvalidChainId("chainId", chainId)
	}

	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, ErrEthClientFailed(err.Error())
	}

	return ethClient, nil
}
