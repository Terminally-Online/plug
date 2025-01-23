package signature

import (
	"math/big"
	"os"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	domainName    = "Plug"
	domainVersion = "1"
)

func getPlugHash(plug Plug) [32]byte {
	encoded := crypto.Keccak256(
		[]byte(PLUG_TYPEHASH),
		plug.To.Bytes(),
		common.LeftPadBytes(plug.Value.Bytes(), 32),
		crypto.Keccak256(plug.Data),
	)
	var hash [32]byte
	copy(hash[:], encoded)
	return hash
}

func getPlugArrayHash(plugs []Plug) [32]byte {
	var encoded []byte
	for _, plug := range plugs {
		hash := getPlugHash(plug)
		encoded = append(encoded, hash[:]...)
	}

	return crypto.Keccak256Hash(encoded)
}

func getPlugsHash(plugs Plugs) [32]byte {
	plugArrayHash := getPlugArrayHash(plugs.Plugs)
	encoded := crypto.Keccak256(
		[]byte(PLUGS_TYPEHASH),
		plugs.Socket.Bytes(),
		plugArrayHash[:],
		crypto.Keccak256(plugs.Solver),
		crypto.Keccak256(plugs.Salt),
	)
	var hash [32]byte
	copy(hash[:], encoded)
	return hash
}

func GetSignature(chainId *big.Int, socket common.Address, plugs Plugs) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return nil, utils.ErrBuild(err.Error())
	}

	domainHash := crypto.Keccak256(
		[]byte(EIP712_DOMAIN_TYPEHASH),
		crypto.Keccak256([]byte(domainName)),
		crypto.Keccak256([]byte(domainVersion)),
		common.LeftPadBytes(chainId.Bytes(), 32),
		socket.Bytes(),
	)
	plugsHash := getPlugsHash(plugs)
	signatureHash := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainHash,
		plugsHash[:],
	)

	signature, err := crypto.Sign(signatureHash, privateKey)
	if err != nil {
		return nil, utils.ErrBuild(err.Error())
	}

	signature[64] += 27

	return signature, nil
}
