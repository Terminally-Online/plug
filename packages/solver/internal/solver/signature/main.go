package signature

import (
	"math/big"
	"os"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	domainName    = "Plug Socket"
	domainVersion = "0.0.1"
)

func GetPlugHash(plug Plug) [32]byte {
	return crypto.Keccak256Hash(
		[]byte(PLUG_TYPEHASH),
		plug.To.Bytes(),
		crypto.Keccak256(plug.Data),
		common.LeftPadBytes(plug.Value.Bytes(), 32),
		common.LeftPadBytes(plug.Gas.Bytes(), 32),
	)
}

func GetPlugsHash(plugs Plugs) [32]byte {
	plugArrayHash := GetPlugArrayHash(plugs.Plugs)
	return crypto.Keccak256Hash(
		[]byte(PLUGS_TYPEHASH),
		plugs.Socket.Bytes(),
		plugArrayHash[:],
		crypto.Keccak256(plugs.Solver),
		crypto.Keccak256(plugs.Salt),
	)
}

func GetPlugArrayHash(plugs []Plug) [32]byte {
	var encoded []byte
	for _, plug := range plugs {
		hash := GetPlugHash(plug)
		encoded = append(encoded, hash[:]...)
	}

	return crypto.Keccak256Hash(encoded)
}

func GetLivePlugArrayHash(livePlugs LivePlugs) [32]byte {
	plugsHash := GetPlugsHash(livePlugs.Plugs)
	return crypto.Keccak256Hash(
		[]byte(LIVE_PLUGS_TYPEHASH),
		plugsHash[:],
		crypto.Keccak256(livePlugs.Signature),
	)
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
	plugsHash := GetPlugsHash(plugs)
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
