package core

import (
	"GoAPIfy/tools/core/color"
	"fmt"
	"os"
)

func Help() {
	PrintLogo()
	fmt.Println(color.Colorize(color.Yellow, "Below is command of GoAPIfy command!\n"))
	fmt.Println(color.Colorize(color.Magenta, "   key"))
	fmt.Println(color.Colorize(color.Green, "     Command to generate your Application key with Base 64 string.\n"))
	fmt.Println(color.Colorize(color.Magenta, "   dev"))
	fmt.Println(color.Colorize(color.Green, "     Command to generate run a development server.\n"))
	fmt.Println(color.Colorize(color.Magenta, "   build"))
	fmt.Println(color.Colorize(color.Green, "     Command to build your production server.\n"))
	fmt.Println(color.Colorize(color.Magenta, "   clean"))
	fmt.Println(color.Colorize(color.Green, "     Command to build clean your builds.\n"))
	fmt.Println(color.Colorize(color.Magenta, "   run"))
	fmt.Println(color.Colorize(color.Green, "     Command to clean build, rebuild it, and run it.\n"))
	fmt.Println(color.Colorize(color.Magenta, "   docker [up/down]"))
	fmt.Println(color.Colorize(color.Green, "     Command to compose up and down the docker for your development.\n"))
	fmt.Println(color.Colorize(color.Magenta, "   rename"))
	fmt.Println(color.Colorize(color.Green, "     Command to rename your application based from .env\n     (Please make sure you have changed your APP_NAME in .env).\n"))
	fmt.Println(color.Colorize(color.Magenta, "   entity [name]"))
	fmt.Println(color.Colorize(color.Green, "     Create a new entity.\n     Entity is a package of controller and model.\n     It creates controllers file (input, handlers, and formatter) and model.\n     (Entity only contain alphanumeric no symbols and capital letters).\n"))
	os.Exit(0)
}
