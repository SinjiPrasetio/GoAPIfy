// Package helper provides utility functions that can be used across the application.
package helper

import "fmt"

// Color constants represent ANSI escape codes for various colors.
const (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Blue    = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Cyan    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)

// ColorizeCmd returns a string with the specified color applied.
// It takes a color constant and a string as input, and returns a formatted string
// with the specified color applied.
func ColorizeCmd(color string, text string) string {
	return fmt.Sprintf(color, text)
}
