package utils

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ZeroAddress        = common.HexToAddress("0x0000000000000000000000000000000000000000")
	NativeTokenAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	TokenStandards = []int{0, 20, 721, 1155}
	VaultStandards = []int{4626}

	addressPattern = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	hexPattern     = regexp.MustCompile("^(0x)?[0-9a-fA-F]+$")

	Uint8Max   = new(big.Int).SetUint64(0xFF)
	Uint16Max  = new(big.Int).SetUint64(0xFFFF)
	Uint32Max  = new(big.Int).SetUint64(0xFFFFFFFF)
	Uint64Max  = new(big.Int).SetUint64(0xFFFFFFFFFFFFFFFF)
	Uint128Max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))
	Uint256Max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

	SelectorLookup                   = "8063%s14"
	UpgradableImplementationSelector = "5c60da1b"

	NativeTransferGas = uint64(21000)

	// NOTE: 20 & 721 share the definition of transferFrom, but the it will hit for 20 before 721.
	TokenSelectors = []struct {
		Selector string
		Type     int
	}{
		{"a9059cbb", 20},   // 20 -- transfer(...)
		{"23b872dd", 721},  // 721 -- transferFrom(...)
		{"f242432a", 1155}, // 1155 -- safeTransferFrom(...)
	}
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

func BuildCallOpts(address string, value *big.Int) *bind.CallOpts {
	return &bind.CallOpts{
		From: common.HexToAddress(address),
		Pending: true,
		Context: context.Background(),
	}
}

func BuildTransactionOpts(address string, value *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: common.HexToAddress(address),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
		NoSend:    true,
		Value:     value,
	}
}
