package main

import (
	"GoAPIfy/core/storage"
	"GoAPIfy/tools/core"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"GoAPIfy/tools/core/color"

	"github.com/joho/godotenv"
)

const envFilePath = ".env"

func main() {
	// Get argument from the command line.
	args := os.Args
	if len(args) == 1 {
		fmt.Println(color.Colorize(color.Red, "No command input, please use make help to view commands."))
		os.Exit(0)
	}

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

	if args[1] == "help" {
		core.Help()
	}

	if args[1] == "key" {
		core.PrintLogo()
		core.KeyGenerate()
	}

	if args[1] == "rename" {
		core.Rename()
	}

	if args[1] == "entity" {
		core.PrintLogo()

		if len(args) != 3 {
			fmt.Println(color.Colorize(color.Red, "Not enough arguments."))
			os.Exit(0)
		}
		p := args[2]
		folder := storage.FolderExists(fmt.Sprintf("./controller/%s", p))
		formatter := storage.FileExists(fmt.Sprintf("./controller/%s/formatter.go", p))
		handler := storage.FileExists(fmt.Sprintf("./controller/%s/handler.go", p))
		input := storage.FileExists(fmt.Sprintf("./controller/%s/input.go", p))
		model := storage.FileExists(fmt.Sprintf("./model/%s.go", p))
		if folder || formatter || handler || input || model {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("%s controller and model is existed, cannot create new controller!", p)))
			fmt.Println(color.Colorize(color.Yellow, "Would you like to delete and create new one?? (y/n)"))
			var input string
			fmt.Scanln(&input)
			if input == "y" || input == "Y" {
				err := os.RemoveAll(fmt.Sprintf("./controller/%s", p))
				if err != nil {
					fmt.Println(color.Colorize(color.Red, err.Error()))
				}
				err = os.Remove(fmt.Sprintf("./model/%s.go", p))
				if err != nil {
					fmt.Println(color.Colorize(color.Red, err.Error()))
				}
			}
			if input != "y" && input != "Y" {
				os.Exit(0)
			}
		}
		core.Model(args[2])
		core.Controller(args[2])
	}

	if args[1] == "docker" {
		if len(args) != 3 {
			fmt.Println(color.Colorize(color.Red, "Not enough arguments."))
			os.Exit(0)
		}
		core.Docker(args[2])
	}

	if args[1] == "dev" {
		core.PrintLogo()
		fmt.Println(color.Colorize(color.Green, "Starting Server..."))
		cmd := exec.Command("go", "run", "main.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error running development server: %s", err)))
			os.Exit(0)
		}
	}

	if args[1] == "build" {
		core.PrintLogo()
		fmt.Println(color.Colorize(color.Green, "Building Server..."))
		cmd := exec.Command("go", "build", "-o", "build/"+os.Getenv("APP_NAME"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error building server: %s", err)))
			os.Exit(0)
		}
		fmt.Println(color.Colorize(color.Green, "Have successfuly to build Server..."))
	}

	if args[1] == "clean" {
		core.PrintLogo()
		fmt.Println(color.Colorize(color.Green, "Cleaning build..."))

		if err := os.RemoveAll("build"); err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error removing build directory: %s", err)))
			os.Exit(0)
		}

		fmt.Println(color.Colorize(color.Green, "Build is cleaned..."))
	}

	if args[1] == "run" {
		core.PrintLogo()
		fmt.Println(color.Colorize(color.Green, "Clenaing up the build..."))

		if err := os.RemoveAll("build"); err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error removing build directory: %s", err)))
			os.Exit(0)
		}

		if err := os.MkdirAll("build", os.ModePerm); err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error creating build directory: %s", err)))
			os.Exit(0)
		}

		fmt.Println(color.Colorize(color.Green, "Building the server..."))

		cmd := exec.Command("go", "build", "-o", "build/"+os.Getenv("APP_NAME"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error building server: %s", err)))
			os.Exit(0)
		}

		fmt.Println(color.Colorize(color.Green, "Running the server..."))

		cmd = exec.Command("./build/" + os.Getenv("APP_NAME"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error running server: %s", err)))
			os.Exit(0)
		}
	}
}
