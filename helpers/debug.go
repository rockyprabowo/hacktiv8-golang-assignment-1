package helpers

import (
	"io"
	"log"
	"os"
	"strings"
)

// PrintDebug prints a debug message to the console.
func PrintDebug(message string, source string) {
	log.Printf("debug: [%s] %s", source, message)
}

// If the environment variable `DEBUG` is set to `true` or `1`, then return `true`
func IsDebug() bool {
	debugEnv := strings.ToLower(os.Getenv("DEBUG"))
	return StringInSlice(debugEnv, []string{"true", "1"})
}

// If debugging is enabled, then log messages will be printed to the
// console. Otherwise, they will be discarded.
func SetupLogging() {
	if !IsDebug() {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
		return
	}
	PrintDebug("debug messages is active", "SetupLogging")
}
