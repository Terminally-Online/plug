package utils

import (
	"math/big"
)

var (
	OraclePriceScale = new(big.Int).Exp(big.NewInt(10), big.NewInt(36), nil)
	Wad              = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	VirtualShares    = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	VirtualAssets    = big.NewInt(1)
)

func WMulDown(x *big.Int, y *big.Int) *big.Int {
	return MulDivDown(x, y, Wad)
}

func WDivDown(x *big.Int, y *big.Int) *big.Int {
	return MulDivDown(x, Wad, y)
}

func MulDivDown(x *big.Int, y *big.Int, d *big.Int) *big.Int {
	result := new(big.Int).Mul(x, y)
	return result.Div(result, d)
}

func MulDivUp(x *big.Int, y *big.Int, d *big.Int) *big.Int {
	result := new(big.Int).Mul(x, y)
	dMinusOne := new(big.Int).Sub(d, big.NewInt(1))
	result.Add(result, dMinusOne)
	return result.Div(result, d)
}
