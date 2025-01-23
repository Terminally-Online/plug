package morpho

import (
	"fmt"
	"math/big"
)

var (
	OraclePriceScale = new(big.Int).Exp(big.NewInt(10), big.NewInt(36), nil)
	Wad              = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	VirtualShares    = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	VirtualAssets    = big.NewInt(1)
)

func wMulDown(x *big.Int, y *big.Int) *big.Int {
	return mulDivDown(x, y, Wad)
}

func wDivDown(x *big.Int, y *big.Int) *big.Int {
	return mulDivDown(x, Wad, y)
}

func mulDivDown(x *big.Int, y *big.Int, d *big.Int) *big.Int {
	result := new(big.Int).Mul(x, y)
	return result.Div(result, d)
}

func mulDivUp(x *big.Int, y *big.Int, d *big.Int) *big.Int {
	result := new(big.Int).Mul(x, y)
	dMinusOne := new(big.Int).Sub(d, big.NewInt(1))
	result.Add(result, dMinusOne)
	return result.Div(result, d)
}

func parseBigInt(value interface{}) (*big.Int, error) {
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
