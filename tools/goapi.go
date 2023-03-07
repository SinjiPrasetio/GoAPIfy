package main

import (
	"GoAPIfy/tools/core"
	"fmt"
	"io/ioutil"
	"os"

	"GoAPIfy/tools/core/color"

	"github.com/joho/godotenv"
)

const envFilePath = ".env"

func main() {
	// Get argument from the command line.
	args1 := os.Args[1]

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
			fmt.Println(color.Colorize(color.Green, ".env file has been restocolor.red from .env.example."))
		} else {
			os.Exit(0)
		}
	}

	if args1 == "key" {
		core.KeyGenerate()
	}

	if args1 == "production" {
		core.ProductionCheck()
	}
}
