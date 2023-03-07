package storage

import "os"

// FileExists checks if the specified file exists
func FileExists(filename string) bool {
	// Attempt to retrieve information about the file
	info, err := os.Stat(filename)

	// If the file doesn't exist, return false
	if os.IsNotExist(err) {
		return false
	}

	// If the file exists and is not a directory, return true
	return !info.IsDir()
}
