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
	fmt.Println(color.Colorize(color.Magenta, "   rename"))
	fmt.Println(color.Colorize(color.Green, "     Command to rename your application based from .env\n     (Please make sure you have changed your APP_NAME in .env).\n"))
	fmt.Println(color.Colorize(color.Magenta, "   model [name]"))
	fmt.Println(color.Colorize(color.Green, "     Create a new model.\n     (Model only contain alphanumeric no symbols and capital letters).\n"))
	os.Exit(0)
}
