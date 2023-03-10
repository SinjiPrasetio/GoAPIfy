package core

import (
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironment loads the .env file and prompts the user to restore from .env.example if it is corrupted.
func LoadEnvironment() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// Prompt user to restore .env file from example
		fmt.Println(color.Colorize(color.Red, "Failed to load .env file."))
		fmt.Println(color.Colorize(color.Yellow, "Would you like to restore from .env.example? (y/n)"))
		var input string
		fmt.Scanln(&input)
		if input == "y" || input == "Y" {
			// Copy contents of .env.example to .env file
			envExampleData, err := ioutil.ReadFile(".env.example")
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "Failed to read .env.example file. Please re-clone GoAPIfy"))
				os.Exit(0)
			}

			err = ioutil.WriteFile(".env", envExampleData, 0644)
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "Failed to write .env file."))
				os.Exit(0)
			}

			// Reload environment variables from .env file
			err = godotenv.Load()
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "Failed to load environment variables."))
				os.Exit(0)
			}
			fmt.Println(color.Colorize(color.Green, ".env file has been restored from .env.example."))
		} else {
			os.Exit(0)
		}
	}
}
