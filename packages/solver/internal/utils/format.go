package utils

import (
	"fmt"
	"strings"
)

// FormatNumber formats a float64 with thousand separators and specified decimal places
func FormatNumber(num float64, prefix string) string {
	if num < 0.01 && num > 0 {
		return prefix + "<0.01"
	}

	baseStr := fmt.Sprintf("%.2f", num)
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
