package coil

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
)

func TestGetTypeInfo(t *testing.T) {
	tests := []struct {
		name     string
		typeStr  string
		expLen   uint
		expTypeId uint8
	}{
		{"uint256", "uint256", 32, TypeIDStatic},
		{"address", "address", 20, TypeIDStatic},
		{"bool", "bool", 1, TypeIDStatic},
		{"string", "string", 32, TypeIDString},
		{"bytes", "bytes", 32, TypeIDString},
		{"bytes32", "bytes32", 32, TypeIDStatic},
		{"bytes4", "bytes4", 4, TypeIDStatic},
		{"uint8", "uint8", 1, TypeIDStatic},
		{"uint256[]", "uint256[]", 32, TypeIDArray},
		{"string[]", "string[]", 32, TypeIDArray},
		{"uint256[][]", "uint256[][]", 32, TypeIDNestedArray},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, err := ParseABIType(tt.typeStr)
			if err != nil {
				t.Fatalf("failed to parse type %s: %v", tt.typeStr, err)
			}

			info := GetTypeInfo(typ)
			assert.Equal(t, tt.expLen, info.Length, "length mismatch")
			assert.Equal(t, tt.expTypeId, info.TypeId, "typeId mismatch")
		})
	}
}

func TestFindCoils(t *testing.T) {
	// Test for balanceOf function in ERC20
	arguments := abi.Arguments{
		{
			Name: "balance",
			Type: abi.Type{T: abi.UintTy, Size: 256},
		},
	}

	method := abi.Method{
		Name:    "balanceOf",
		Outputs: arguments,
	}

	testAbi := &abi.ABI{
		Methods: map[string]abi.Method{
			"balanceOf": method,
		},
	}

	t.Run("basic coil for uint256", func(t *testing.T) {
		coils, err := FindCoils(testAbi, "balanceOf", nil, nil)

		assert.NoError(t, err)
		assert.Len(t, coils, 1)
		
		coil := coils[0]
		assert.Equal(t, big.NewInt(0), coil.Start)
		assert.Equal(t, "balance", *coil.Slice.Name)
		assert.Equal(t, big.NewInt(0), coil.Slice.Start)
		assert.Equal(t, big.NewInt(32), coil.Slice.Length)
		assert.Equal(t, uint8(0), *coil.Slice.TypeId)
	})

	t.Run("filter by output name", func(t *testing.T) {
		name := "balance"
		coils, err := FindCoils(testAbi, "balanceOf", &name, nil)
		
		assert.NoError(t, err)
		assert.Len(t, coils, 1)
	})

	t.Run("filter by output index", func(t *testing.T) {
		index := "0"
		coils, err := FindCoils(testAbi, "balanceOf", nil, &index)
		
		assert.NoError(t, err)
		assert.Len(t, coils, 1)
	})

	t.Run("no match with wrong function", func(t *testing.T) {
		_, err := FindCoils(testAbi, "nonExistentFunction", nil, nil)
		assert.Error(t, err)
	})
}

func TestIsDynamicType(t *testing.T) {
	tests := []struct {
		name     string
		typeStr  string
		expected bool
	}{
		{"uint256", "uint256", false},
		{"address", "address", false},
		{"bool", "bool", false},
		{"string", "string", true},
		{"bytes", "bytes", true},
		{"bytes32", "bytes32", false},
		{"uint256[]", "uint256[]", true},
		{"uint256[][]", "uint256[][]", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, err := ParseABIType(tt.typeStr)
			if err != nil {
				t.Fatalf("failed to parse type %s: %v", tt.typeStr, err)
			}

			isDynamic := IsDynamicType(typ)
			assert.Equal(t, tt.expected, isDynamic)
		})
	}
}