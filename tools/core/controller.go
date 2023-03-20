package core

import (
	"GoAPIfy/core/stringable"
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"
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

	// Update the contents of the register file
	registerFile, err := os.OpenFile("./controller/register.go", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	defer registerFile.Close()
	registerBytes, err := ioutil.ReadAll(registerFile)
	if err != nil {
		panic(err)
	}
	registerContent := string(registerBytes)

	// Check if the handler is already registered
	if strings.Contains(registerContent, fmt.Sprintf("%sHandler *%s.%sHandler", controllerName, strings.ToLower(controllerName), controllerName)) {
		fmt.Println(color.Colorize(color.Green, "Handler already registered"))
		return
	}

	// Add the new handler to the Handlers struct
	handlerLine := fmt.Sprintf("\t%sHandler *%s.%sHandler\n", controllerName, strings.ToLower(controllerName), controllerName)
	newRegisterContent := strings.ReplaceAll(registerContent, "}\n", handlerLine+"}\n")

	// Update the RegisterHandler function to register the new handler
	registerHandlerLine := fmt.Sprintf("\t\t%sHandler: %s.New%sHandler(s, authService),\n", strings.ToLower(controllerName), strings.ToLower(controllerName), controllerName)
	newRegisterContent = strings.ReplaceAll(newRegisterContent, "// Initialize other handlers as needed\n", registerHandlerLine+"// Initialize other handlers as needed\n")

	// Write the modified contents of the register file
	registerFile.Seek(0, 0)
	if _, err := registerFile.WriteString(newRegisterContent); err != nil {
		panic(err)
	}

}

func CreateFormatter(p string, controllerName string) {
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
