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
)

var (
	domainName    = "Plug Socket"
	domainVersion = "0.0.1"
)

const (
	EIP712_DOMAIN_TYPEHASH = "0x8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f"
	SLICE_TYPEHASH         = "0xf8939514938e0a800705081290e2e4c7efcf49061b28bf5b38f457c851eb82ac"
	UPDATE_TYPEHASH        = "0x85c9aec0e14ad33e63489c03355fa65515340a998cc26cd360d11267b451b6fd"
	PLUG_TYPEHASH          = "0x7cae6e9d732b3307b20040708ed6876bf34aeb91eb6bcfbfd18581cb0376b60b"
	PLUGS_TYPEHASH         = "0x05b2ab8b8c7ceee9902f5288470f7189883657d476121976b1079d47722718a2"
	LIVE_PLUGS_TYPEHASH    = "0x858fa8b1b482c729dcb5ae30adab7db7ed354ebaba182da4ff91412001f7fd45"
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
		[]byte(SLICE_TYPEHASH),
		[]byte{slice.Index},
		common.LeftPadBytes(slice.Start.Bytes(), 32),
		common.LeftPadBytes(slice.Length.Bytes(), 32),
	)
}

func getUpdateHash(update coil.Update) [32]byte {
	sliceHash := getSliceHash(update.Slice)
	return crypto.Keccak256Hash(
		[]byte(UPDATE_TYPEHASH),
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

func GetPlugHash(plug Plug) [32]byte {
	updateArrayHash := getUpdateArrayHash(plug.Updates)

	return crypto.Keccak256Hash(
		[]byte(PLUG_TYPEHASH),
		[]byte{uint8(plug.Selector)},
		plug.To.Bytes(),
		crypto.Keccak256(plug.Data),
		common.LeftPadBytes(plug.Value.Bytes(), 32),
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
		[]byte(PLUGS_TYPEHASH),
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

	return plugs, signature, nil
}
