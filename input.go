package main

import (
	"log"
	"os"
	"rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
)

// Get the numbers from the input source, whether it's from stdin, command line arguments, or
// interactive input.
func GetNumberFromInput() (numbers []uint64) {
	if !state.interactive {
		if state.fromStdin {
			numbers = helpers.StdinNumberInput()
			log.Printf("debug: [GetNumberFromInput] (stdin) %v", numbers)
			return
		}

		if state.fromArgs {
			args := os.Args[1:]
			log.Printf("debug: [GetNumberFromInput] (args) %v", args)
			numbers = helpers.ProcessUint64Numbers(args)
			return
		}
	}

	numbers = helpers.InteractiveNumberInput()
	log.Printf("debug: [GetNumberFromInput] (interactive) %v", numbers)
	return
}
