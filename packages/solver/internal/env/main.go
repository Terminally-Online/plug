package env

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

func Init() {
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

func Encrypt(text string) (string, error) {
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

func Decrypt(text string) (string, error) {
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

func EncryptEnvFile() error {
	content, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	encrypted, err := Encrypt(string(content))
	if err != nil {
		return err
	}

	if err := os.WriteFile(".env.encrypted", []byte(encrypted), 0644); err != nil {
		return err
	}

	fmt.Println("Encrypted .env file created as .env.encrypted")
	return nil
}

func DecryptEnvFile() error {
	content, err := os.ReadFile(".env.encrypted")
	if err != nil {
		return err
	}

	decrypted, err := Decrypt(string(content))
	if err != nil {
		return err
	}

	if err := os.WriteFile(".env", []byte(decrypted), 0644); err != nil {
		return err
	}

	fmt.Println(".env file created from decrypted content")
	return nil
}
