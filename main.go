package main

import (
	"fmt"
	"os"
	"rockyprabowo/hacktiv8-assignments/assignment-1/classmates"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
)

// Main application loop.
// Firstly, get the attendance number from command arguments or user input (stdin or interactive), then get the
// classmate with its AttendanceNumber, and if the classmate exists, print the aforementioned classmate data.
func main() {
	ShowBanner()
	Setup()

	// do-while loop "hacks" for Go
	for interactive := true; interactive; interactive = state.interactive {
		counter, increment, incrementError := h.UseProcessCounter()

		// 1. Get the attendance number from command arguments or user input (stdin or interactive).
		numbers := GetNumberFromInput()

		for _, number := range numbers {
			// 2. Get the classmate with the Attendance Number.
			classmate, err := classmates.GetByAttendanceNumber(number)
			if err != nil {
				h.PrintDebug(err.Error(), "main")
				fmt.Println("error: " + err.Error())
				fmt.Println()
				incrementError()
				continue
			}
			h.PrintDebug(fmt.Sprintf("found the Classmate with AttendanceNumber = %d\n", number),
				"main")

			// 3. Print the classmate data
			classmate.Print()
			increment()
		}

		// Shows an error if input is not valid
		if len(numbers) == 0 {
			fmt.Println("error: invalid input. please try again.")
			fmt.Println()
			if !state.interactive {
				os.Exit(1)
			}
			continue
		}

		// Prints the process counter at the end of every input.
		fmt.Printf("%d requested, %d processed, %d error(s)\n",
			counter.RequestedCount, counter.ProcessedCount, counter.ErrorCount)
	}
}
