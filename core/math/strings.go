// Package math provides utility functions for mathematical operations, including generating
// random numbers and strings, and performing calculations on numbers.
package math

import (
	"encoding/hex"
	"math/rand"
)

// RandomString generates a random string of the specified length.
// It takes an integer length as input and returns a string containing a random
// assortment of characters from the character set [a-z, A-Z, 0-9].
func RandomString(length int) string {
	// Define the character set to use in generating the random string.
	var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")

	// Create a new slice of runes with the specified length.
	str := make([]rune, length)

	// Loop through each index in the slice and set its value to a random character from the character set.
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the slice of runes to a string and return it.
	return string(str)
}

// RandomNumber generates a random integer between min and max.
// It takes two integers, min and max, as input, and returns a random integer in the range [min, max].
func RandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomNumberString generates a random string of numbers with the specified length.
// It takes an integer length as input and returns a string containing a random
// assortment of digits from the set of digits [0-9].
func RandomNumberString(length int) (string, error) {
	// Create a buffer of bytes with the specified length.
	buffer := make([]byte, length)

	// Generate random bytes and store them in the buffer.
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Define the character set to use in generating the random number string.
	numSets := "1234567890"

	// Calculate the length of the numSets character set.
	numSetsLength := len(numSets)

	// Loop through each byte in the buffer and convert it to an index in the numSets character set.
	// Then set the corresponding character in the buffer to the selected numSets character.
	for i := 0; i < length; i++ {
		buffer[i] = numSets[int(buffer[i])%numSetsLength]
	}

	// Convert the buffer of bytes to a string and return it.
	return string(buffer), nil
}

// GenerateUUID generates a random UUID (v4) as a string.
func GenerateUUID() (string, error) {
	// Create a buffer to hold the random bytes.
	buffer := make([]byte, 16)

	// Read random bytes into the buffer.
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Set the version (bits 12-15) to 4 as per the UUID v4 specification.
	buffer[6] = (buffer[6] & 0x0f) | 0x40
	// Set the variant (bits 16-17) to the RFC 4122 variant as per the UUID v4 specification.
	buffer[8] = (buffer[8] & 0x3f) | 0x80

	// Encode the buffer as a hex string and return it.
	return hex.EncodeToString(buffer), nil
}
