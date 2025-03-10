package coil

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	TypeIDStatic      uint8 = 0 // Static types (uint256, address, etc.)
	TypeIDArray       uint8 = 1 // Dynamic arrays (uint256[], address[], etc.)
	TypeIDString      uint8 = 2 // String and bytes (identical in typeId)
	TypeIDStruct      uint8 = 3 // Structs with dynamic fields
	TypeIDNestedArray uint8 = 4 // Nested arrays (uint256[][], etc.)
	TypeIDMapping     uint8 = 5 // Mapping-like structures

	WordSize uint = 32 // Standard EVM word size (32 bytes)
)

// FindCoils analyzes a function's outputs and creates the appropriate slices for data extraction
// from contract call results. This enables reactive transaction building by making outputs
// available as inputs to subsequent transactions.
//
// Parameters:
//   - abi: The parsed ABI containing the function definition
//   - functionName: The name of the function whose outputs we want to analyze
//   - outputName: Optional filter to only return a specific named output
//   - outputIndex: Optional filter to only return a specific output by index
//
// Returns:
//   - []Update: A list of updates that can be used to extract data from contract call results
//   - error: Any error encountered during analysis
func FindCoils(abi *abi.ABI, functionName string, outputName *string, outputIndex *string) ([]Update, error) {
	fn, exists := abi.Methods[functionName]
	if !exists {
		return nil, fmt.Errorf("function %s not found in ABI", functionName)
	}

	coils := make([]Update, 0, len(fn.Outputs))
	var currentPos uint = 0

	for i, output := range fn.Outputs {
		typeInfo := GetTypeInfo(output.Type)

		if (outputName != nil && output.Name != *outputName) ||
			(outputIndex != nil && fmt.Sprint(i) != *outputIndex) {
			currentPos += typeInfo.Length
			continue
		}

		name := output.Name
		if name == "" {
			name = fmt.Sprintf("output%d", i)
		}

		typeId := typeInfo.TypeId
		coils = append(coils, Update{
			Start: big.NewInt(int64(currentPos)),
			Slice: Slice{
				Name:   &name,
				Index:  uint8(i),
				Start:  big.NewInt(int64(currentPos)),
				Length: big.NewInt(int64(typeInfo.Length)),
				Type:   output.Type.String(),
				TypeId: &typeId,
			},
		})

		currentPos += typeInfo.Length
	}

	return coils, nil
}

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
		if hasDynamicFields(typ) {
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

func hasDynamicFields(typ abi.Type) bool {
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
		return hasDynamicFields(typ)
	}
	return false
}

func ParseABIType(typeStr string) (abi.Type, error) {
	switch {
	case typeStr == "uint256":
		return abi.Type{T: abi.UintTy, Size: 256}, nil
	case typeStr == "address":
		return abi.Type{T: abi.AddressTy, Size: 160}, nil
	case typeStr == "bool":
		return abi.Type{T: abi.BoolTy, Size: 8}, nil
	case typeStr == "string":
		return abi.Type{T: abi.StringTy}, nil
	case typeStr == "bytes":
		return abi.Type{T: abi.BytesTy}, nil
	case strings.HasPrefix(typeStr, "bytes") && len(typeStr) > 5:
		sizeStr := typeStr[5:]
		var size uint64
		_, err := fmt.Sscan(sizeStr, &size)
		if err != nil || size > 32 {
			return abi.Type{}, fmt.Errorf("invalid fixed bytes size: %s", typeStr)
		}
		return abi.Type{T: abi.FixedBytesTy, Size: int(size)}, nil
	case strings.HasSuffix(typeStr, "[]"):
		elemType, err := ParseABIType(strings.TrimSuffix(typeStr, "[]"))
		if err != nil {
			return abi.Type{}, err
		}
		return abi.Type{T: abi.ArrayTy, Size: 0, Elem: &elemType}, nil
	}

	return abi.Type{}, fmt.Errorf("unsupported type: %s", typeStr)
}
