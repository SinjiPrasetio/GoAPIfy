package core

import (
	"fmt"
	"strings"
	"unicode"
)

const envFilePath = ".env"

// Update a variable in a .env file
func updateEnvVariable(envData string, variableName string, variableValue string) string {
	// Split file contents into lines
	lines := getLines(envData)

	// Loop through each line in the file
	for i, line := range lines {
		// Check if line is a variable assignment
		if line != "" && line[0] != '#' {
			// Check if line is the variable we want to update
			if key := getKey(line); key == variableName {
				// Split line into key-value pair
				parts := getParts(line)
				if len(parts) > 0 {
					// Update value of key
					lines[i] = fmt.Sprintf("%v=%v", variableName, variableValue)
				}
				break
			}
		}
	}

	// Join lines back into string and return
	return strings.Join(lines, "\n")
}

// Get the key part of a key-value pair
func getKey(line string) string {
	parts := getParts(line)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// Split a string into key-value pair parts
func getParts(line string) []string {
	return split(trim(line), "=")
}

// Split a string into lines
func getLines(data string) []string {
	return split(data, "\n")
}

// Split a string by a separator
func split(data string, sep string) []string {
	return strings.Split(data, sep)
}

// Trim whitespace from a string
func trim(s string) string {
	return strings.TrimSpace(s)
}

func containsWhitespace(s string) bool {
	return strings.ContainsAny(s, " \t\n\r")
}

func containsUppercase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	for _, r := range s {
		if unicode.IsSymbol(r) {
			return true
		}
	}
	return false
}

func containsOnlyLettersAndNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
