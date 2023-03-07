package helper

import "fmt"

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

// Colorize formats a string with the specified color
func ColorizeCmd(color string, text string) string {
	return fmt.Sprintf(color, text, Black)
}
