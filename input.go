package main

import (
	"bufio"
	"fmt"
	"os"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
	"strings"
)

// Reads from stdin, splits the input by spaces, and returns a slice of parsed uint64 numbers.
func StdinNumberInput() (numbers []uint64) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		h.PrintDebug(fmt.Sprintf("input is %v", text), "StdinNumberInput")

		if len(text) > 0 {
			input := strings.Split(text, " ")
			numbers = append(numbers, h.ProcessUint64Numbers(input)...)
		}
	}

	if err := scanner.Err(); err != nil {
		h.ExitWithError(err)
	}

	return
}

// Takes a list of numbers from the user seperated by spaces, and returns a slice of parsed uint64 numbers.
func InteractiveNumberInput() (numbers []uint64) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("[press " + h.GuessRuntimeSigtermKey() + " to exit]")
	fmt.Print("Attendance number(s): ")
	scanner.Scan()
	text := scanner.Text()
	h.PrintDebug(fmt.Sprintf("input is %v", text), "InteractiveNumberInput")
	if len(text) > 0 {
		input := strings.Split(text, " ")
		numbers = h.ProcessUint64Numbers(input)
	}

	if err := scanner.Err(); err != nil {
		h.ExitWithError(err)
	}

	return
}

// Get the numbers from the input source, whether it's from stdin, command line arguments, or
// interactive input.
func GetNumberFromInput() (numbers []uint64) {
	if !state.interactive {
		if state.fromStdin {
			numbers = append(numbers, StdinNumberInput()...)
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

	numbers = InteractiveNumberInput()
	h.PrintDebug(fmt.Sprintf("(interactive) %v", numbers), "GetNumberFromInput")
	return
}
