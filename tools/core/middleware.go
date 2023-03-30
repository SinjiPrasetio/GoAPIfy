package core

import (
	"GoAPIfy/core/storage"
	"GoAPIfy/core/stringable"
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Middleware(p string) {
	if containsWhitespace(p) || containsUppercase(p) || containsSymbol(p) || !containsOnlyLettersAndNumbers(p) {
		fmt.Println(color.Colorize(color.Red, "Middleware name cannot contain whitespace, uppercase, and symbol!"))
		os.Exit(0)
	}

	fileExist := storage.FileExists(fmt.Sprintf("./middleware/%s.go", p))
	if fileExist {
		fmt.Println(color.Colorize(color.Red, fmt.Sprintf("%s middleware is existed, cannot create new middleware!", p)))
		os.Exit(0)
	}

	middlewareName := stringable.Capitalize(p)
	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("Creating %s middleware!", p)))
	src, err := os.Open("./tools/templates/middleware/middleware.txt")
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	defer src.Close()
	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		panic(err)
	}

	// Replace the substring ${middlewareName} with "Example"
	modifiedBytes := []byte(strings.ReplaceAll(string(srcByte), "${middlewareName}", middlewareName))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), "${middlewarePackage}", "middleware"))

	out, err := os.Create(fmt.Sprintf("./middleware/%s.go", p))
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}
	defer out.Close()

	_, err = out.Write(modifiedBytes)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}

	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("%s middleware created!", p)))
}
