package signature

import (
	"math/big"
	"os"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	domainName    = "Plug Socket"
	domainVersion = "0.0.1"
)

func GetSolverHash() ([]byte, error) { 
	// NOTE: This sets the expiration of a Solver provided order to five minutes from now so that our Solver
	//       cannot sign a message, someone else get a hold if it and execute way in the future or us
	//       end up having the case where things are Plugs are not properly executed because they are being
	//       executed 10k blocks late after it was held from execution.
	expiration := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(300))
	solver, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 48}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(expiration, common.HexToAddress(os.Getenv("SOLVER_ADDRESS")))
	if err != nil {
		return nil, utils.ErrBuild("failed to pack solver: " + err.Error())
	}
	return solver, nil
}

func GetSaltHash(from common.Address) ([]byte, error) { 
	salt, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(
		big.NewInt(time.Now().Unix()),
		from,
		common.HexToAddress(os.Getenv("ONE_CLICKER_ADDRESS")),
		common.HexToAddress(os.Getenv("IMPLEMENTATION_ADDRESS")),
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to pack salt: " + err.Error())
	}
	return salt, nil
}

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

func GetSignature(chainId *big.Int, socket common.Address, plugs Plugs) (Plugs, []byte, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return Plugs{}, nil, utils.ErrBuild(err.Error())
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
		return Plugs{}, nil, utils.ErrBuild(err.Error())
	}

	signature[64] += 27

	return Plugs{}, signature, nil
}
