package core

import (
	"GoAPIfy/core/math"
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"
)

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
