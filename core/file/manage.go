// Package file provides functions for working with files in a directory.
package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
