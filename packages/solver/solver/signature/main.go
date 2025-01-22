package signature

import (
	"encoding/json"
	"log"
	"math/big"
	"os"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	domainName    = "My DApp"
	domainVersion = "1"
)

func GetHash[T any](data T) [32]byte {
	encoded, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to encode data: %v", err)
	}

	return crypto.Keccak256Hash(encoded)
}

func GetSignature(chainId *big.Int, socket common.Address, plug Plugs) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return nil, utils.ErrBuildFailed(err.Error())
	}

	domainHash := GetHash(EIP712Domain{
		Name:              domainName,
		Version:           domainVersion,
		ChainId:           chainId,
		VerifyingContract: socket,
	})
	messageHash := GetHash(Plug{
		To:    common.HexToAddress("0xRecipientAddress"),
		Data:  []byte{},
		Value: big.NewInt(0),
	})
	signatureHash := crypto.Keccak256(append(
		domainHash[:],
		messageHash[:]...,
	))
	signature, err := crypto.Sign(signatureHash, privateKey)
	if err != nil {
		log.Fatalf("Failed to sign message: %v", err)
	}

	return signature, nil
}
