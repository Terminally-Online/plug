package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/scrypt"
)

var encryptionKey []byte

func init() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }
	
	if key := os.Getenv("ENCRYPTION_KEY"); key != "" {
		var err error
		encryptionKey, err = scrypt.Key([]byte(key), []byte("salt"), 32768, 8, 1, 32)
		if err != nil {
			panic(err)
		}
	}
}

func encrypt(text string) (string, error) {
	if encryptionKey == nil {
		return "", fmt.Errorf("encryption key not set")
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(text))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, []byte(text))

	return fmt.Sprintf("%x:%x", iv, ciphertext), nil
}

func decrypt(text string) (string, error) {
	if encryptionKey == nil {
		return "", fmt.Errorf("encryption key not set")
	}

	parts := strings.Split(text, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid encrypted text format")
	}

	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func encryptEnvFile() error {
	content, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	encrypted, err := encrypt(string(content))
	if err != nil {
		return err
	}

	if err := os.WriteFile(".env.encrypted", []byte(encrypted), 0644); err != nil {
		return err
	}

	fmt.Println("Encrypted .env file created as .env.encrypted")
	return nil
}

func decryptEnvFile() error {
	content, err := os.ReadFile(".env.encrypted")
	if err != nil {
		return err
	}

	decrypted, err := decrypt(string(content))
	if err != nil {
		return err
	}

	if err := os.WriteFile(".env", []byte(decrypted), 0644); err != nil {
		return err
	}

	fmt.Println(".env file created from decrypted content")
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [encrypt|decrypt]")
		return
	}

	key := os.Getenv("ENCRYPTION_KEY")
	if key == "" || key == "github-action" {
		fmt.Println("Skipping encryption and decryption.")
		return
	}

	command := os.Args[1]
	var err error

	switch command {
	case "encrypt":
		err = encryptEnvFile()
	case "decrypt":
		err = decryptEnvFile()
	default:
		fmt.Println("Usage: go run main.go [encrypt|decrypt]")
		return
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
