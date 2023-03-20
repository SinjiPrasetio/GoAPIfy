package core

import (
	"GoAPIfy/tools/core/color"
	"fmt"
)

func PrintLogo() {
	ClearConsole()
	fmt.Println(color.Colorize(color.Cyan, "   _____               _____ _____  __       "))
	fmt.Println(color.Colorize(color.Cyan, "  / ____|        /\\   |  __ \\_   _|/ _|      "))
	fmt.Println(color.Colorize(color.Cyan, " | |  __  ___   /  \\  | |__) || | | |_ _   _ "))
	fmt.Println(color.Colorize(color.Cyan, " | | |_ |/ _ \\ / /\\ \\ |  ___/ | | |  _| | | |"))
	fmt.Println(color.Colorize(color.Cyan, " | |__| | (_) / ____ \\| |    _| |_| | | |_| |"))
	fmt.Println(color.Colorize(color.Cyan, "  \\_____|\\___/_/    \\_\\_|   |_____|_|  \\__, |"))
	fmt.Println(color.Colorize(color.Cyan, "                                        __/ |"))
	fmt.Println(color.Colorize(color.Cyan, "                                       |___/ \n"))
}
