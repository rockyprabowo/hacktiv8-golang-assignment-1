package main

import (
	"fmt"
	"os"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
)

// Get the numbers from the input source, whether it's from stdin, command line arguments, or
// interactive input.
func GetNumberFromInput() (numbers []uint64) {
	if !state.interactive {
		if state.fromStdin {
			numbers = append(numbers, h.StdinNumberInput()...)
			h.PrintDebug(fmt.Sprintf("(stdin) %v", numbers), "GetNumberFromInput")
		}

		if state.fromArgs {
			args := os.Args[1:]
			h.PrintDebug(fmt.Sprintf("(args) %v", args), "GetNumberFromInput")
			numbers = append(numbers, h.ProcessUint64Numbers(args)...)
		}

		h.PrintDebug(fmt.Sprintf("(stdin + args) %v", numbers), "GetNumberFromInput")
		return
	}

	numbers = h.InteractiveNumberInput()
	h.PrintDebug(fmt.Sprintf("(interactive) %v", numbers), "GetNumberFromInput")
	return
}
