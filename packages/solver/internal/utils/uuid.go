package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
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
	// Possible words for the key
	words := []string{"bing", "boop", "bop", "bam", "blap"}

	// Generate 12 random bytes to select words
	randBytes := make([]byte, 12)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random bytes: %v", err))
	}

	// Create a slice to hold the 12 selected words
	selectedWords := make([]string, 12)
	for i := 0; i < 12; i++ {
		selectedWords[i] = words[int(randBytes[i])%len(words)]
	}

	// Join the words with hyphens
	return strings.Join(selectedWords, "-")
}
