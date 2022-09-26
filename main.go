package main

import (
	"fmt"
	"os"
	"rockyprabowo/hacktiv8-assignments/assignment-1/classmates"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
)

// Get the attendance number from command argument (os.Args) or user input (fmt.Scanln), get the
// classmate with the Attendance Number, and print the classmate data.
func main() {
	ShowBanner()
	Setup()

	// do-while loop "hacks" for Go
	for interactive := true; interactive; interactive = state.interactive {
		counter, increment, incrementError := h.UseProcessCounter()

		// 1. Get the attendance number from command argument or user input (stdin or interactive).
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

		if len(numbers) == 0 {
			fmt.Println("error: invalid input. please try again.")
			fmt.Println()
			if !state.interactive {
				os.Exit(1)
			}
			continue
		}

		fmt.Printf("%d requested, %d processed, %d error(s)\n",
			counter.RequestedCount, counter.ProcessedCount, counter.ErrorCount)
	}
}
