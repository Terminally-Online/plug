package config

import (
	"crypto/ecdsa"
	"fmt"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ChainId           = int64(8453)
	UseExecution      = utils.GetEnvOrDefault("USE_EXECUTION", "false") == "true"
	UseAVS            = utils.GetEnvOrDefault("USE_AVS", "false") == "true"
	Production        = utils.GetEnvOrDefault("AVS_ENV", "development") == "production"
	SolverUrl         = utils.GetEnvOrDefault("SOLVER_URL", "http://localhost:8080")
	PrivateKey        = utils.GetEnvOrDefault("PRIVATE_KEY", "")
	Node_1            = utils.GetEnvOrDefault("RPC_1", "")
	Node_8453         = utils.GetEnvOrDefault("RPC_8453", "")
	Port              = utils.GetEnvOrDefault("PORT", "6473")
	AttestationCenter = common.HexToAddress("0x62180042606624f02d8a130da8a3171e9b33894d")
)

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
