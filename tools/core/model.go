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

func Model(p string) {
	if containsWhitespace(p) || containsUppercase(p) || containsSymbol(p) || !containsOnlyLettersAndNumbers(p) {
		fmt.Println(color.Colorize(color.Red, "Model name cannot contain whitespace, uppercase, and symbol!"))
		os.Exit(0)
	}

	fileExist := storage.FileExists(fmt.Sprintf("./model/%s.go", p))
	if fileExist {
		fmt.Println(color.Colorize(color.Red, fmt.Sprintf("%s model is existed, cannot create new model!", p)))
		os.Exit(0)
	}

	modelName := stringable.Capitalize(p)
	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("Creating %s model!", p)))
	src, err := os.Open("./tools/templates/model/model.txt")
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	defer src.Close()
	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		panic(err)
	}

	// Replace the substring ${modelName} with "Example"
	modifiedBytes := []byte(strings.ReplaceAll(string(srcByte), "${modelName}", modelName))

	out, err := os.Create(fmt.Sprintf("./model/%s.go", p))
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

	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("%s model created!", p)))
}
