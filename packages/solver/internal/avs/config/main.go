package config

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ChainId           = int64(8453)
	PrivateKey        = GetEnvOrDefault("PRIVATE_KEY", "")
	Node_1            = GetEnvOrDefault("RPC_1", "")
	Node_8453         = GetEnvOrDefault("RPC_8453", "")
	Port              = GetEnvOrDefault("PORT", "8081")
	AttestationCenter = common.HexToAddress("0x62180042606624f02d8a130da8a3171e9b33894d")
)

func GetEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetNodeUrl(chainId int64) (string, error) {
	switch chainId {
	case 1:
		return Node_1, nil
	case 8453:
		return Node_8453, nil
	default:
		return "", fmt.Errorf("unsupported chain ID: %d", chainId)
	}
}

func GetAccount() (*ecdsa.PrivateKey, string, error) {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return nil, "", err
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, "", err
	}
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	return privateKey, address, nil
}
