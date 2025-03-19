package coil

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func HasDynamicFields(typ abi.Type) bool {
	if typ.T != abi.TupleTy {
		return false
	}

	for _, elem := range typ.TupleElems {
		if IsDynamicType(*elem) {
			return true
		}
	}
	return false
}

func IsDynamicType(typ abi.Type) bool {
	switch typ.T {
	case abi.ArrayTy:
		return typ.Size == 0 || (typ.Elem != nil && IsDynamicType(*typ.Elem))
	case abi.StringTy, abi.BytesTy:
		return typ.Size == 0
	case abi.TupleTy:
		return HasDynamicFields(typ)
	}
	return false
}
