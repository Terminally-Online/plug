package signature

import (
	"math/big"
	"os"
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
	UPDATE_TYPEHASH        = "0x56b7ba00148b10c99ae43c2f84a4ec0ec1dfa2e5d7d5954f23e627d964b83435"
	SLICE_TYPEHASH         = "0x705d2fbc03b585d2178271bc5e779f63e494ae79418e11352a00b51e1edeeaae"
	PLUG_TYPEHASH          = "0x77f9edb2051551cf1c4102bcde8eba61d391a9e0536544fe44cc2330cc4913fa"
	PLUGS_TYPEHASH         = "0xf730b3caf995a40c1030675beebec0f819730393c3020ff6bd0295c733af216b"
	LIVE_PLUGS_TYPEHASH    = "0xb17d5f6d50f15d707601c6994c8e42d3c170be70fe81f012884129e8e5bf41ff"
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
	updateArrayHash := getUpdateArrayHash(plug.Updates)

	return crypto.Keccak256Hash(
		[]byte(PLUG_TYPEHASH),
		[]byte{plug.Selector},
		plug.To.Bytes(),
		crypto.Keccak256(plug.Data),
		common.LeftPadBytes(plug.Value.Bytes(), 32),
		updateArrayHash[:],
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

func getUpdateHash(update coil.Update) [32]byte {
	sliceHash := getSliceHash(update.Slice)
	return crypto.Keccak256Hash(
		[]byte(UPDATE_TYPEHASH),
		common.LeftPadBytes(update.Start.Bytes(), 32),
		sliceHash[:],
	)
}

func getSliceHash(slice coil.Slice) [32]byte {
	return crypto.Keccak256Hash(
		[]byte(SLICE_TYPEHASH),
		[]byte{slice.Index},
		common.LeftPadBytes(slice.Start.Bytes(), 32),
		common.LeftPadBytes(slice.Length.Bytes(), 32),
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

	return plugs, signature, nil
}
