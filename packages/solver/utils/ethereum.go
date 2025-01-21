package utils

import (
	"fmt"
	"math/big"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ZeroAddress        = common.HexToAddress("0x0000000000000000000000000000000000000000")
	NativeTokenAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	TokenStandards = []int{0, 20, 721, 1155}
	VaultStandards = []int{4626}

	SelectorLookup                   = "8063%s14"
	UpgradableImplementationSelector = "5c60da1b"

	// NOTE: 20 & 721 share the definition of transferFrom, but the it will hit for 20 before 721.
	TokenSelectors = []struct {
		Selector string
		Type     int
	}{
		{"a9059cbb", 20},   // 20 -- transfer(...)
		{"23b872dd", 721},  // 721 -- transferFrom(...)
		{"f242432a", 1155}, // 1155 -- safeTransferFrom(...)
	}
)

var (
	addressPattern = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	hexPattern     = regexp.MustCompile("^(0x)?[0-9a-fA-F]+$")
)

var (
	Uint8Max   = new(big.Int).SetUint64(0xFF)
	Uint16Max  = new(big.Int).SetUint64(0xFFFF)
	Uint32Max  = new(big.Int).SetUint64(0xFFFFFFFF)
	Uint64Max  = new(big.Int).SetUint64(0xFFFFFFFFFFFFFFFF)
	Uint128Max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))
	Uint256Max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
)

func IsSupportedTokenStandard(standard int) bool {
	return slices.Contains(TokenStandards, standard)
}

func IsSupportedVaultStandard(standard int) bool {
	return slices.Contains(VaultStandards, standard)
}

func IsSupportedProtocol(protocols []string, slug string) bool {
	return slices.Contains(protocols, slug)
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
		max = Uint8Max
	case 16:
		max = Uint16Max
	case 32:
		max = Uint32Max
	case 64:
		max = Uint64Max
	case 128:
		max = Uint128Max
	case 256:
		max = Uint256Max
	default:
		return false
	}

	return n.Sign() >= 0 && n.Cmp(max) <= 0 || n.Cmp(max) == 0
}

func IsBytes(value string, size int) bool {
	if !hexPattern.MatchString(value) {
		return false
	}

	value = strings.TrimPrefix(value, "0x")
	return len(value) == size*2
}

func FloatToUint(value float64, decimals uint8) (*big.Int, error) {
    // Create a copy of the input value to avoid modifying it
    result := new(big.Float).SetFloat64(value)
    
    // Calculate 10^decimals
    multiplier := new(big.Int).Exp(
        big.NewInt(10), 
        big.NewInt(int64(decimals)), 
        nil,
    )
    
    // Multiply result by 10^decimals
    result.Mul(result, new(big.Float).SetInt(multiplier))
    
    // Convert to int and check for accuracy
    intValue, accuracy := result.Int(nil)
    if accuracy != big.Exact {
        return nil, fmt.Errorf("loss of precision when converting %v to uint with %d decimals", value, decimals)
    }
    
    return intValue, nil
}

func ParseAddressAndDecimals(input string) (address string, decimals uint8, err error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid input format: %s", input)
	}
	
	if !IsAddress(parts[0]) {
		return "", 0, fmt.Errorf("invalid address: %s", parts[0])
	}
	
	decimals64, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse decimals: %v", err)
	}
	
	return parts[0], uint8(decimals64), nil
}