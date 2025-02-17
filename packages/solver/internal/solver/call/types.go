package call

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type MulticallCalldata struct {
	Target     common.Address
	Method     string
	Args       []interface{}
	ABI        *abi.ABI
	OutputType interface{}
}

