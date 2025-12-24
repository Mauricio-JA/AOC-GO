package utils

import "runtime"

// GetEOL returns the end-of-line string for the current operating system.
func GetEOL() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
