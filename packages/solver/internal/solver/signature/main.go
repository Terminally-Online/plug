package signature

import (
	"crypto/ecdsa"
	"fmt"
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
	fmt.Printf("solver expiration: %s\n", expiration.String())
	solver, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 48}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(expiration, common.HexToAddress(os.Getenv("SOLVER_ADDRESS")))
	if err != nil {
		return nil, utils.ErrBuild("failed to pack solver: " + err.Error())
	}
	fmt.Printf("solverData %x\n", solver)
	return solver, nil
}

func GetSaltHash(from common.Address) ([]byte, error) {
	saltHashNonce := big.NewInt(time.Now().Unix())
	fmt.Printf("saltHashNonce: %s\n", saltHashNonce.String())
	salt, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(
		saltHashNonce,
		common.HexToAddress("0x50701f4f523766bFb5C195F93333107d1cB8cD90"), // TODO MASON THIS IS THE OWNER ADDRESS
		common.HexToAddress(os.Getenv("SOLVER_ADDRESS")),
		// TODO: We need a way to know the implementation address that was used when deploying the socket.
		//       There is going to be some tricky stuff here. It will not matter as long as we have everyone
		//       on one version but this is going to have to be fixed sooner than later.
		common.HexToAddress(references.Plug["socket"]),
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to pack salt: " + err.Error())
	}
	fmt.Printf("saltData %x\n", salt)
	return salt, nil
}

func getSliceHash(slice coil.Slice) [32]byte {
	var typeIdByte byte
	if slice.TypeId != nil {
		typeIdByte = byte(*slice.TypeId)
	}

	sliceTypeHashArray := getTypeHashArray(plug.SliceTypeHash)

	return crypto.Keccak256Hash(
		sliceTypeHashArray[:],
		[]byte{slice.Index},
		common.LeftPadBytes(slice.Start.Bytes(), 32),
		common.LeftPadBytes(slice.Length.Bytes(), 32),
		[]byte{typeIdByte},
	)
}

func getUpdateHash(update coil.Update) [32]byte {
	sliceHash := getSliceHash(update.Slice)
	updateHashArray := getTypeHashArray(plug.UpdateTypeHash)
	return crypto.Keccak256Hash(
		updateHashArray[:],
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

func getTypeHashArray(typehash string) [32]byte {
	typeHashBytes := common.FromHex(typehash)
	var typeHashArray [32]byte
	copy(typeHashArray[:], typeHashBytes)
	return typeHashArray
}

func bytesTo32(b []byte) [32]byte {
	var arr [32]byte
	copy(arr[:], b)
	return arr
}

func GetPlugHash(p Plug) [32]byte {
	updateArrayHash := getUpdateArrayHash(p.Updates)

	plugHashArray := getTypeHashArray(plug.PlugTypeHash)

	return crypto.Keccak256Hash(
		plugHashArray[:],
		crypto.Keccak256(p.Data),
		updateArrayHash[:],
	)
}

func GetPlugArrayHash(plugs []Plug) [32]byte {
	var encoded []byte
	for _, plug := range plugs {
		hash := GetPlugHash(plug)
		fmt.Printf("plugHash: %x\n", hash)
		encoded = append(encoded, hash[:]...)
	}

	return crypto.Keccak256Hash(encoded)
}

func GetPlugsHash(plugs Plugs) [32]byte {
	plugArrayHash := GetPlugArrayHash(plugs.Plugs)
	fmt.Printf("plugArrayHash: %x\n", plugArrayHash)

	// Get the type hash
	typeHashArray := getTypeHashArray(plug.PlugsTypeHash)
	fmt.Printf("typeHash: %x\n", typeHashArray)
	fmt.Printf("typeHash hex: %s\n", common.Bytes2Hex(typeHashArray[:]))

	// First hash the solver and salt as raw bytes
	solverHash := crypto.Keccak256(plugs.Solver)
	saltHash := crypto.Keccak256(plugs.Salt)

	fmt.Printf("solverHash: %x\n", solverHash)
	fmt.Printf("saltHash: %x\n", saltHash)

	fmt.Printf("solverBytes : %x\n", plugs.Solver)
	fmt.Printf("saltBytes   : %x\n", plugs.Salt)

	// Pack everything using abi.encode rules
	packed, err := abi.Arguments{
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // PLUGS_TYPEHASH
		{Type: abi.Type{T: abi.AddressTy}},              // socket
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // plugArrayHash
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // solverHash (was abi.BytesTy)
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // saltHash   (was abi.BytesTy)
	}.Pack(
		typeHashArray,
		plugs.Socket,
		plugArrayHash,
		bytesTo32(solverHash),
		bytesTo32(saltHash),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to pack plugs: %s", err.Error()))
	}

	// Log the packed bytes
	fmt.Printf("packed bytes: %x\n", packed)

	result := crypto.Keccak256Hash(packed)
	fmt.Printf("plugsHash: %x\n", result)
	return result
}

// CONFIRMED FUNCTIONAL
func GetDomainHash(chainId *big.Int, socket common.Address) ([]byte, error) {
	nameHash := crypto.Keccak256([]byte(domainName))
	versionHash := crypto.Keccak256([]byte(domainVersion))

	// Convert all byte slices to [32]byte arrays
	var nameHashArray [32]byte
	var versionHashArray [32]byte

	typeHashSlice := getTypeHashArray(plug.Eip712DomainTypeHash)
	copy(nameHashArray[:], nameHash)
	copy(versionHashArray[:], versionHash)

	packed, err := abi.Arguments{
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}},
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}},
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(
		typeHashSlice,
		nameHashArray,
		versionHashArray,
		chainId,
		socket,
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to pack domain: " + err.Error())
	}

	domainHash := crypto.Keccak256(packed)
	domainHashHex := common.BytesToHash(domainHash)
	fmt.Printf("domainHashHex: %s\n", domainHashHex.Hex())
	return domainHash, nil
}

func GetSignature(chainId *big.Int, socket common.Address, plugs Plugs) (Plugs, []byte, error) {
	fmt.Printf("Plugs %+v\n", plugs)

	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		return Plugs{}, nil, utils.ErrBuild(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Plugs{}, nil, utils.ErrBuild("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("address: %s\n", address.Hex())

	plugsData := plugs.Plugs[0].Data
	fmt.Printf("plugsData: %x\n", plugsData)
	plugsDataHex := common.Bytes2Hex(plugsData)
	fmt.Printf("plugsDataHex: %s\n", plugsDataHex)

	domainHash, err := GetDomainHash(chainId, socket)
	if err != nil {
		return Plugs{}, nil, err
	}

	plugsHash := GetPlugsHash(plugs)

	digestInput := append([]byte{0x19, 0x01}, domainHash[:]...)
	digestInput = append(digestInput, plugsHash[:]...)
	fmt.Printf("digestInput: %x\n", digestInput)

	signatureHash := crypto.Keccak256(digestInput)
	fmt.Printf("finalDigest: %x\n", signatureHash)

	signature, err := crypto.Sign(signatureHash, privateKey)
	if err != nil {
		return Plugs{}, nil, utils.ErrBuild(err.Error())
	}

	signature[64] += 27

	fmt.Printf("signature: %x\n", signature)

	return plugs, signature, nil
}

// "rawInput": "0x52c47baa0fb919f7f2c9d53e71f458fd9edd4b2d2378a7decc87411ad45ebffe000000000000000000000000000000000000000000000000000000000000001b55380155db822987674a0128637b0aa9e0d52056a33a8058a51d28965a16861762a97bed8fb148751b7a816725e0917a7d3e46d898fd158dfce33036aa89eabc",
// "rawOutput": "0x0000000000000000000000008533eb2f10de4b528022048612a14c97d531e58e",
