// Package math provides encryption and decryption functions using the AES block cipher in CFB mode.
package math

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Encrypt takes a plaintext string and returns the encrypted ciphertext as a base64-encoded string.
// It uses the AES block cipher in CFB mode with a random initialization vector and the APP_KEY environment variable as the encryption key.
func Encrypt(text string) (string, error) {
	plaintext := []byte(text)
	key := []byte(os.Getenv("APP_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	output := base64.StdEncoding.EncodeToString(ciphertext)
	return output, nil

}

// Decrypt takes a base64-encoded ciphertext string and returns the decrypted plaintext as a string.
// It uses the AES block cipher in CFB mode with the initialization vector embedded in the ciphertext and the APP_KEY environment variable as the decryption key.
func Decrypt(text string) (string, error) {
	ciphertext := []byte(text)
	key := []byte(os.Getenv("APP_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil

}

// Hash takes a plaintext password and returns the hashed password as a base64-encoded string.
// It uses the SHA-256 hash function to generate the hash.
func Hash(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	output := base64.StdEncoding.EncodeToString(hash[:])
	return output, nil
}

// HashChallenge takes a plaintext password and a base64-encoded hash challenge and returns true if the hash of the password matches the challenge.
// It uses the SHA-256 hash function to generate the hash.
func HashChallenge(password string, challenge string) bool {
	hash := sha256.Sum256([]byte(password))
	hashString := base64.StdEncoding.EncodeToString(hash[:])
	return hashString == challenge
}
