package core

import (
	"GoAPI/tools/core/color"
	"GoAPIfy/core/math"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func KeyGenerate() {
	// Check if APP_KEY is already set
	if os.Getenv("APP_KEY") != "" {
		fmt.Println(color.Colorize(color.Green, "App Key has been specified."))
		os.Exit(0)
	}

	// Generate new APP_KEY
	fmt.Println(color.Colorize(color.Green, "Generating key..."))
	key := math.RandomString(50)

	// Read .env file
	envData, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Update APP_KEY value in .env file
	newEnvData := updateEnvVariable(string(envData), "APP_KEY", key)

	// Write new .env file
	if err := ioutil.WriteFile(envFilePath, []byte(newEnvData), 0644); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print new APP_KEY value
	fmt.Println(color.Colorize(color.Green, fmt.Sprintf("Key generated. Key: %s", key)))
	os.Exit(0)
}

func ProductionCheck() {

}
