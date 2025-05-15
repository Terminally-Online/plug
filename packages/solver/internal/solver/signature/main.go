package signature

import (
	"math/big"
	"os"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	plug "github.com/terminally-online/plug/packages/references/common"
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
		big.NewInt((time.Now().Unix())),
		from,
		common.HexToAddress(os.Getenv("SOLVER_ADDRESS")),
		// TODO: We need a way to know the implementation address that was used when deploying the socket.
		//       There is going to be some tricky stuff here. It will not matter as long as we have everyone
		//       on one version but this is going to have to be fixed sooner than later.
		common.HexToAddress(references.Plug["socket"]),
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to pack salt: " + err.Error())
	}
	return salt, nil
}

func getSliceHash(slice coil.Slice) [32]byte {
	return crypto.Keccak256Hash(
		[]byte(plug.SliceTypeHash),
		[]byte{slice.Index},
		common.LeftPadBytes(slice.Start.Bytes(), 32),
		common.LeftPadBytes(slice.Length.Bytes(), 32),
	)
}

func getUpdateHash(update coil.Update) [32]byte {
	sliceHash := getSliceHash(update.Slice)
	return crypto.Keccak256Hash(
		[]byte(plug.UpdateTypeHash),
		common.LeftPadBytes(update.Start.Bytes(), 32),
		sliceHash[:],
	)
}

func getUpdateArrayHash(updates []coil.Update) [32]byte {
	var encoded []byte
	for _, update := range updates {
		hash := getUpdateHash(update)
		encoded = append(encoded, hash[:]...)
	}
	return crypto.Keccak256Hash(encoded)
}

func GetPlugHash(p Plug) [32]byte {
	updateArrayHash := getUpdateArrayHash(p.Updates)

	return crypto.Keccak256Hash(
		[]byte(plug.PlugTypeHash),
		[]byte{uint8(p.Selector)},
		p.To.Bytes(),
		crypto.Keccak256(p.Data),
		common.LeftPadBytes(p.Value.Bytes(), 32),
		updateArrayHash[:],
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

func GetPlugsHash(plugs Plugs) [32]byte {
	plugArrayHash := GetPlugArrayHash(plugs.Plugs)
	return crypto.Keccak256Hash(
		[]byte(plug.PlugsTypeHash),
		plugs.Socket.Bytes(),
		plugArrayHash[:],
		crypto.Keccak256(plugs.Solver),
		crypto.Keccak256(plugs.Salt),
	)
}

func GetSignature(chainId *big.Int, socket common.Address, plugs Plugs) (Plugs, []byte, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return Plugs{}, nil, utils.ErrBuild(err.Error())
	}

	domainHash := crypto.Keccak256(
		[]byte(plug.Eip712DomainTypeHash),
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

	return plugs, signature, nil
}
