package helpers

import (
	"fmt"
	"log"
	"runtime"
)

// If the OS is darwin, return "cmd+c", otherwise return "ctrl+c".
func GuessRuntimeSigtermKey() string {
	if runtime.GOOS == "darwin" {
		return "cmd+c"
	}
	return "ctrl+c"
}

// If there's an error, print it and exit the program
func ExitWithError(err error) {
	fmt.Println(err)
	log.Fatalln(err)
}
