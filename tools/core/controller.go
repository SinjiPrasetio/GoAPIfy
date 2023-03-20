package core

import (
	"GoAPIfy/core/stringable"
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Controller(p string) {
	if containsWhitespace(p) || containsUppercase(p) || containsSymbol(p) || !containsOnlyLettersAndNumbers(p) {
		fmt.Println(color.Colorize(color.Red, "Model name cannot contain whitespace, uppercase, and symbol!"))
		os.Exit(0)
	}

	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("Creating %s controller!", p)))

	err := os.Mkdir(fmt.Sprintf("./controller/%s", p), 0755)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
	}

	controllerName := stringable.Capitalize(p)

	CreateHandler(p, controllerName)
	CreateFormatter(p, controllerName)
	CreateInput(p, controllerName)

	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("%s controller created!", p)))
}

func CreateHandler(p string, controllerName string) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()

	appName := strings.TrimSpace(string(output))
	src, err := os.Open("./tools/templates/controller/handler.txt")
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	defer src.Close()
	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		panic(err)
	}

	modifiedBytes := []byte(strings.ReplaceAll(string(srcByte), "${controllerName}", controllerName))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), "${controllerPackage}", p))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), "${AppName}", appName))

	out, err := os.Create(fmt.Sprintf("./controller/%s/handler.go", p))
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

}

func CreateFormatter(p string, controllerName string) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()

	appName := strings.TrimSpace(string(output))
	src, err := os.Open("./tools/templates/controller/formatter.txt")
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	defer src.Close()
	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		panic(err)
	}

	modifiedBytes := []byte(strings.ReplaceAll(string(srcByte), "${controllerName}", controllerName))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), "${controllerPackage}", p))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), "${AppName}", appName))

	out, err := os.Create(fmt.Sprintf("./controller/%s/formatter.go", p))
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
}

func CreateInput(p string, controllerName string) {
	src, err := os.Open("./tools/templates/controller/input.txt")
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	defer src.Close()
	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		panic(err)
	}

	modifiedBytes := []byte(strings.ReplaceAll(string(srcByte), "${controllerName}", controllerName))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), "${controllerPackage}", p))

	out, err := os.Create(fmt.Sprintf("./controller/%s/input.go", p))
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
}
