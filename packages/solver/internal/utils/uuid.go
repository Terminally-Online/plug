package utils

import (
	"crypto/rand"
	"encoding/binary"
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

func GenerateApiKey() string {
	// Word lists for generating memorable combinations
	adjectives := []string{"happy", "brave", "clever", "swift", "mighty", "jolly", "vibrant",
		"fierce", "gentle", "mystic", "cosmic", "golden", "silver", "crystal"}
	nouns := []string{"panda", "rocket", "comet", "wizard", "dragon", "phoenix", "ninja",
		"falcon", "tiger", "forest", "ocean", "mountain", "shadow", "thunder"}

	// Generate 8 random bytes (we only need a few but this keeps it similar to your UUID function)
	randBytes := make([]byte, 8)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random bytes: %v", err))
	}

	// Use the random bytes to select words and generate a number
	adjIndex := int(randBytes[0]) % len(adjectives)
	nounIndex := int(randBytes[1]) % len(nouns)

	// Use remaining bytes for a 4-digit number
	randomNum := int(binary.BigEndian.Uint16(randBytes[2:4])) % 10000

	return fmt.Sprintf("%s-%s-%04d", adjectives[adjIndex], nouns[nounIndex], randomNum)
}
