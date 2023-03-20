package core

import (
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Rename() {
	PrintLogo()
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()

	oldName := strings.TrimSpace(string(output))
	newName := os.Getenv("APP_NAME")
	if oldName == newName {
		fmt.Println(color.Colorize(color.Red, "The APP_NAME and the project name is the same, nothing would be done."))
		os.Exit(1)
	}
	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("Renaming project from %s to %s!", oldName, newName)))
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			content := string(data)
			if strings.Contains(content, oldName) && strings.Contains(content, `"`) {
				if strings.Contains(content, fmt.Sprintf(`"%s/`, oldName)) {
					newContent := strings.ReplaceAll(content, fmt.Sprintf(`"%s/`, oldName), fmt.Sprintf(`"%s/`, newName))
					err = ioutil.WriteFile(path, []byte(newContent), 0644)
					if err != nil {
						return err
					}
				} else if strings.Contains(content, fmt.Sprintf(` "%s"`, oldName)) {
					newContent := strings.ReplaceAll(content, fmt.Sprintf(` "%s"`, oldName), fmt.Sprintf(` "%s"`, newName))
					err = ioutil.WriteFile(path, []byte(newContent), 0644)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	cmd = exec.Command("go", "mod", "edit", "-module", os.Getenv("APP_NAME"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error running development server: %s", err)))
		os.Exit(0)
	}
	fmt.Println(color.Colorize(color.Magenta, fmt.Sprintf("Successfully renamed the project from %s to %s!", oldName, newName)))
}
