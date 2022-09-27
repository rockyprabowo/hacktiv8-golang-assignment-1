package helpers

import (
	"os"
)

// Returns true, if the length of the os.Args slice is greater than 1.
// Returns false otherwise.
func HasArguments() bool {
	return len(os.Args) > 1
}

// Returns true if the stdin got a stream of data waiting to be processed.
// Returns false if the stdin is from a terminal.
func HasDataInStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		PrintDebug("data is being piped to stdin", "HasDataInStdin")
		return true
	}
	PrintDebug("stdin is from a terminal", "HasDataInStdin")
	return false
}
