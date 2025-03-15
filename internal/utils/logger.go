package utils

import "fmt"

// LogInfo logs general information messages.
func LogInfo(message string) {
	fmt.Println("[INFO] " + message)
}

// LogSuccess logs success messages.
func LogSuccess(message string) {
	fmt.Println("\033[32m[SUCCESS] " + message + "\033[0m") // Green color
}

// LogError logs error messages.
func LogError(message string) {
	fmt.Println("\033[31m[ERROR] " + message + "\033[0m") // Red color
}
