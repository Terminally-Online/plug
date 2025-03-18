package utils

import (
	"fmt"
	"math/big"
)

func ParseBigInt(value interface{}) (*big.Int, error) {
	switch v := value.(type) {
	case string:
		if v == "" {
			return nil, nil
		}

		val, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse value: %v", value)
		}
		return val, nil
	case int:
		return big.NewInt(int64(v)), nil
	case int64:
		return big.NewInt(v), nil
	case float64:
		return big.NewInt(int64(v)), nil
	case nil:
		return big.NewInt(0), nil
	default:
		return nil, fmt.Errorf("unsupported type for BigInt: %T", value)
	}
}
