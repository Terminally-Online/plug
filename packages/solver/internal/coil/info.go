package coil

import "github.com/ethereum/go-ethereum/accounts/abi"

type TypeInfo struct {
	Length uint
	TypeId uint8
}

func GetTypeInfo(typ abi.Type) TypeInfo {
	switch typ.T {
	case abi.ArrayTy:
		if typ.Elem != nil {
			if typ.Elem.T == abi.ArrayTy {
				return TypeInfo{WordSize, TypeIDNestedArray}
			}
			return TypeInfo{WordSize, TypeIDArray}
		}
	case abi.StringTy, abi.BytesTy:
		return TypeInfo{WordSize, TypeIDString}
	case abi.TupleTy:
		if HasDynamicFields(typ) {
			return TypeInfo{WordSize, TypeIDStruct}
		}

		var totalSize uint = 0
		for _, field := range typ.TupleElems {
			fieldInfo := GetTypeInfo(*field)
			totalSize += fieldInfo.Length
		}
		return TypeInfo{totalSize, TypeIDStatic}
	}

	switch typ.T {
	case abi.AddressTy:
		return TypeInfo{20, TypeIDStatic}
	case abi.BoolTy:
		return TypeInfo{1, TypeIDStatic}
	case abi.IntTy, abi.UintTy:
		return TypeInfo{uint(typ.Size / 8), TypeIDStatic}
	case abi.FixedBytesTy, abi.BytesTy:
		if typ.Size > 0 {
			return TypeInfo{uint(typ.Size), TypeIDStatic}
		}
	}

	return TypeInfo{WordSize, TypeIDStatic}
}
