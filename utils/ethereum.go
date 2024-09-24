package utils

import (
	"math/big"
	"slices"
	"regexp"
	"strings"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	NativeTokenAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
)

var (
	addressPattern = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	hexPattern     = regexp.MustCompile("^(0x)?[0-9a-fA-F]+$")
)

var (
	uint8Max   = new(big.Int).SetUint64(0xFF)
	uint16Max  = new(big.Int).SetUint64(0xFFFF)
	uint32Max  = new(big.Int).SetUint64(0xFFFFFFFF)
	uint64Max  = new(big.Int).SetUint64(0xFFFFFFFFFFFFFFFF)
	uint128Max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))
	uint256Max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
)

func IsSupportedChain(chainId int) bool {
	return slices.Contains(SupportedChains, chainId)
}

func IsAddress(s string) bool {
	return addressPattern.MatchString(s)
}

func IsHex(s string) bool {
	return hexPattern.MatchString(s)
}

func IsUint(value string, bits int) bool {
	value = strings.TrimPrefix(value, "0x")
	n, ok := new(big.Int).SetString(value, 16)
	if !ok {
		return false
	}

	var max *big.Int
	switch bits {
	case 8:
		max = uint8Max
	case 16:
		max = uint16Max
	case 32:
		max = uint32Max
	case 64:
		max = uint64Max
	case 128:
		max = uint128Max
	case 256:
		max = uint256Max
	default:
		return false
	}

	return n.Sign() >= 0 && n.Cmp(max) <= 0
}

func IsBytes(value string, size int) bool {
	if !hexPattern.MatchString(value) {
		return false
	}

	value = strings.TrimPrefix(value, "0x")
	return len(value) == size*2
}
