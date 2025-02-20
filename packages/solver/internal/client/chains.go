package client

import (
	"fmt"
	"os"
	"solver/internal/utils"
)

func GetChainName(chainId uint64) string {
	switch chainId {
	case 1:
		return "mainnet"
	case 8453:
		return "base"
	case 10:
		return "optimism-sepolia"
	case 31337:
		return "localhost"
	default:
		return "base"
	}
}

func GetQuicknodeUrl(chainId uint64) (string, error) {
	if chainId == 31337 {
		return "http://127.0.0.1:8545", nil
	}

	quicknodeApiName := os.Getenv("QUICKNODE_API_NAME")
	quicknodeApiKey := os.Getenv("QUICKNODE_API_KEY")

	var chain string
	switch chainId {
	case 1:
		chain = ""
	case 8453:
		chain = "base-mainnet."
	case 84532:
		chain = "base-sepolia."
	case 10:
		chain = "optimism."
	case 11155420:
		chain = "optimism-sepolia."
	default:
		return "", utils.ErrChainId("chainId", chainId)
	}

	return fmt.Sprintf("https://%v.%vquiknode.pro/%v", quicknodeApiName, chain, quicknodeApiKey), nil
}
