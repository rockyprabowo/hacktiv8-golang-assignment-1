package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// If the length of the os.Args slice is greater than 1, return true, otherwise return false.
func HasArguments() bool {
	return len(os.Args) > 1
}

// If the stdin is from a terminal, then it returns false, otherwise it returns true
func HasDataInStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		log.Println("debug: [HasDataInStdin] data is being piped to stdin")
		return true
	} else {
		log.Println("debug: [HasDataInStdin] stdin is from a terminal")
		return false
	}
}

// It reads from stdin, splits the input by spaces, and returns a slice of uint64 numbers
func StdinNumberInput() (numbers []uint64) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("debug: [StdinNumberInput] input is %v", text)

		if len(text) > 0 {
			input := strings.Split(text, " ")
			numbers = append(numbers, ProcessUint64Numbers(input)...)
		}
	}

	if err := scanner.Err(); err != nil {
		ExitWithError(err)
	}

	return
}

// It takes a string of numbers separated by spaces interactively, and returns a slice of uint64 numbers
func InteractiveNumberInput() (numbers []uint64) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("[press " + GuessRuntimeSigtermKey() + " to exit]")
	fmt.Print("Attendance number(s): ")
	scanner.Scan()
	text := scanner.Text()
	log.Printf("debug: [InteractiveNumberInput] input is %v", text)
	if len(text) > 0 {
		input := strings.Split(text, " ")
		numbers = ProcessUint64Numbers(input)
	}

	if err := scanner.Err(); err != nil {
		ExitWithError(err)
	}

	return
}
