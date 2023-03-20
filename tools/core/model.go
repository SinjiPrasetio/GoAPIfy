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

	// update the migration
	migrationFile, err := os.OpenFile("./model/database.go", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}
	defer migrationFile.Close()

	migrationBytes, err := ioutil.ReadAll(migrationFile)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}
	// Check if the model already exists in the migration
	if strings.Contains(string(migrationBytes), fmt.Sprintf("&%s{}", modelName)) {
		fmt.Println(color.Colorize(color.Green, fmt.Sprintf("%s model already exists in migration", p)))
		return
	}
	// Add the new model to the migration
	migrationContent := strings.ReplaceAll(string(migrationBytes), "AutoMigrate(\n", fmt.Sprintf("AutoMigrate(\n\t\t&%s{},\n", modelName))
	_, err = migrationFile.Seek(0, 0)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}
	_, err = migrationFile.WriteString(migrationContent)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}
	fmt.Println(color.Colorize(color.Green, fmt.Sprintf("%s model added to migration", p)))
}
