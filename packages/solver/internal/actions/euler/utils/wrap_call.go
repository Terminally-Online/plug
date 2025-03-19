package utils

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_evc"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func WrapEVCCall(chainId uint64, targetContract common.Address, onBehalfOfAccount common.Address, value *big.Int, calldata []byte, updates *[]coil.Update) (signature.Plug, error) {
	evc, err := euler_evc.EulerEvcMetaData.GetAbi()
	if err != nil {
		return signature.Plug{}, utils.ErrABI("EulerEvc")
	}

	callCalldata, err := evc.Pack(
		"call",
		targetContract,
		onBehalfOfAccount,
		value,
		calldata,
	)
	if err != nil {
		fmt.Printf("WrapEVCCall pack error: %v\n", err)
		return signature.Plug{}, utils.ErrTransaction(err.Error())
	}
	var finalUpdates []coil.Update
	if updates != nil {
		finalUpdates = *updates
	}

	return signature.Plug{
		To:      common.HexToAddress(references.Networks[chainId].References["euler"]["evc"]),
		Data:    callCalldata,
		Updates: finalUpdates,
	}, nil
}
