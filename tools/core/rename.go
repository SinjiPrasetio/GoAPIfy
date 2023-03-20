package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Rename(oldName string, newName string) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
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
					fmt.Printf("%s has been replaced with %s in file %s\n", oldName, newName, path)
				} else if strings.Contains(content, fmt.Sprintf(` "%s"`, oldName)) {
					newContent := strings.ReplaceAll(content, fmt.Sprintf(` "%s"`, oldName), fmt.Sprintf(` "%s"`, newName))
					err = ioutil.WriteFile(path, []byte(newContent), 0644)
					if err != nil {
						return err
					}
					fmt.Printf("%s has been replaced with %s in file %s\n", oldName, newName, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
