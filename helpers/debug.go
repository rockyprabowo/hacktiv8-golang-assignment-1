package helpers

import (
	"io"
	"log"
	"os"
	"strings"
)

// If the environment variable `DEBUG` is set to `true` or `1`, then return `true`
func IsDebug() bool {
	debugEnv := strings.ToLower(os.Getenv("DEBUG"))
	return StringInSlice(debugEnv, []string{"true", "1"})
}

// If debugging is enabled, then log messages will be printed to the
// console. Otherwise, they will be discarded.
func SetupLogging() {
	if IsDebug() {
		log.Println("debug: [SetupLogging] debug messages is active")
	} else {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
	}
}
