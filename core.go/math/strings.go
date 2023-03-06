package math

import "math/rand"

// The characters that can be used in generating a random number
const numSets = "1234567890"

// RandomString generates a random string of specified length
func RandomString(length int) string {
	// The character set that can be used in generating the random string
	var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")

	// Create a new slice of runes with the specified length
	str := make([]rune, length)

	// Loop through each index in the slice and set its value to a random character from the character set
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the slice of runes to a string and return it
	return string(str)
}

// RandomNumberString generates a random string of numbers with the specified length
func RandomNumberString(length int) (string, error) {
	// Create a buffer of bytes with the specified length
	buffer := make([]byte, length)

	// Generate random bytes and store them in the buffer
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Calculate the length of the number sets character set
	numSetsLength := len(numSets)

	// Loop through each byte in the buffer and convert it to an index in the number sets character set,
	// then set the corresponding character in the buffer to the selected number sets character
	for i := 0; i < length; i++ {
		buffer[i] = numSets[int(buffer[i])%numSetsLength]
	}

	// Convert the buffer of bytes to a string and return it
	return string(buffer), nil
}
