package cron

import (
	"GoAPIfy/service/appService"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/*
*
Defines a new interface for managing cron jobs.
@interface {Cron} - The interface for the cron job service
*/
type Cron interface {
	Start()
}

/*
*

	Defines a new type that implements the Cron interface.
	@typedef {cron} - The implementation of the Cron interface
	@property {appService} appService - The application service instance to be used in the cron job implementation
*/
type cron struct {
	appService appService.AppService
}

/*
*

	Creates a new instance of the cron job service.
	@param {appService} appService - The application service instance to be used in the cron job implementation
	@returns {cron} - A new instance of the cron job service
*/
func NewCron(appService appService.AppService) *cron {
	return &cron{appService}
}

const (
	TemporaryFileExpiration = 24 * time.Hour
)

func DeleteExpiredTemporaryFiles() {
	// Construct the path for the temporary directory
	temporaryDirectory := filepath.Join("public", "temporary")

	// Open the temporary directory
	dir, err := os.Open(temporaryDirectory)
	if err != nil {
		fmt.Printf("Error opening temporary directory: %s\n", err.Error())
		return
	}
	defer dir.Close()

	// Get the current time
	now := time.Now()

	// Iterate over the files in the temporary directory
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("Error reading temporary directory: %s\n", err.Error())
		return
	}
	for _, file := range files {
		if file.Mode().IsRegular() && now.After(file.ModTime().Add(TemporaryFileExpiration)) {
			// Delete the file if it's expired
			err := os.Remove(filepath.Join(temporaryDirectory, file.Name()))
			if err != nil {
				fmt.Printf("Error deleting temporary file %s: %s\n", file.Name(), err.Error())
			} else {
				fmt.Printf("Deleted temporary file %s\n", file.Name())
			}
		}
	}
}
