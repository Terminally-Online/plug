package signature

import (
	"math/big"
	"os"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
	saltHashNonce := big.NewInt(time.Now().Unix())
	salt, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(
		saltHashNonce,
		from,
		common.HexToAddress(os.Getenv("SOLVER_ADDRESS")),
		common.HexToAddress(references.Plug["socket"]),
	)
	if err != nil {
		return nil, utils.ErrBuild("failed to pack salt: " + err.Error())
	}
	return salt, nil
}

func GetSignature(chainId *big.Int, socket common.Address, plugs Plugs) (Plugs, []byte, error) {
	return Plugs{}, nil, nil

	// privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	// if err != nil {
	// 	return Plugs{}, nil, utils.ErrBuild(err.Error())
	// }

	// plugs.Salt, _ = GetSaltHash(socket)
	// plugs.Solver, _ = GetSolverHash()

	// wrapped, _ := plugs.Wrap()
	// digest, err := digest.GetDigest(chainId, socket, wrapped)
	// signature, err := crypto.Sign(digest[:], privateKey)
	// if err != nil {
	// 	return Plugs{}, nil, utils.ErrBuild(err.Error())
	// }
	//
	// signature[crypto.RecoveryIDOffset] += 27

	// return plugs, signature, nil
}
