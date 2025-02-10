package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func FormatNumber(number float64, prefix string) string {
	if number < 0.01 && number > 0 {
		return "<" + prefix + "0.01"
	}

	baseStr := fmt.Sprintf("%.2f", number)
	parts := strings.Split(baseStr, ".")
	
	whole := parts[0]
	var result []string
	for i := len(whole); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{whole[start:i]}, result...)
	}
	
	formatted := prefix + strings.Join(result, ",")
	if len(parts) > 1 {
		formatted += "." + parts[1]
	}
	
	return formatted
}

func FormatAddress(address common.Address) string {
    return address.Hex()[:6] + "..." + address.Hex()[len(address.Hex())-4:]
}
