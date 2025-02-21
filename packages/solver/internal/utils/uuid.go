package utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateUUID generates a new UUID v4 string using crypto/rand
func GenerateUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(fmt.Sprintf("failed to generate uuid: %v", err))
	}
	// Set version (4) and variant (2) bits according to RFC 4122
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant 2
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}
