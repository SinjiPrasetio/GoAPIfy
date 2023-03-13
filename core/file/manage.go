// Package file provides functions for working with files in a directory.
package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// Save : saves a file to the specified directory with a filename and contents.
// It creates the directory if it does not exist.
func Save(filename string, data []byte, directory string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(filepath.Join(directory, "public", "storage"), 0755)
	if err != nil {
		return err
	}

	// Open the file for writing
	file, err := os.Create(filepath.Join(directory, "public", "storage", filename))
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the file contents to disk
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Read : reads a file from the specified directory with a filename and returns its contents.
func Read(filename string, directory string) ([]byte, error) {
	// Open the file for reading
	file, err := os.Open(filepath.Join(directory, "public", "storage", filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents into memory
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete : deletes a file from the specified directory with a filename.
func Delete(filename string, directory string) error {
	// Remove the file from disk
	err := os.Remove(filepath.Join(directory, "public", "storage", filename))
	if err != nil {
		return err
	}

	return nil
}

const (
	TemporaryFileExpiration = 24 * time.Hour
)

func CreateTemporaryFile(data []byte, filename string) (string, error) {
	// Generate a UUID for the temporary file name
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	temporaryFilename := uuid.String() + "_" + filename

	// Construct the path for the temporary file
	temporaryDirectory := filepath.Join("public", "temporary")
	temporaryPath := filepath.Join(temporaryDirectory, temporaryFilename)

	// Create the temporary directory if it doesn't exist
	err = os.MkdirAll(temporaryDirectory, 0755)
	if err != nil {
		return "", err
	}

	// Write the file contents to disk
	err = ioutil.WriteFile(temporaryPath, data, 0644)
	if err != nil {
		return "", err
	}

	// Return the URL for the temporary file
	return GetTemporaryFileURL(temporaryFilename), nil
}

func GetTemporaryFileURL(filename string) string {
	// Get the protocol based on the production environment
	protocol := "http"
	if os.Getenv("APP_PRODUCTION") == "true" {
		protocol = "https"
	}

	// Get the domain name from the environment variable
	domain := os.Getenv("APP_DOMAIN")

	// Construct the URL for the temporary file
	return fmt.Sprintf("%s://%s/public/temporary/%s", protocol, domain, filename)
}
