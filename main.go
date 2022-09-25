package main

import (
	"fmt"
	"log"
	"rockyprabowo/hacktiv8-assignments/assignment-1/classmates"
)

// Get the attendance number from command argument (os.Args) or user input (fmt.Scanln), get the
// classmate with the Attendance Number, and print the classmate data.
func main() {
	var errorCount int
	ShowBanner()
	Setup()

	// do-while loop "hacks" for Go
	for interactive := true; interactive; interactive = state.interactive {
		// 1. Get the attendance number from command argument (os.Args) or user input (fmt.Scanln).
		numbers := GetNumberFromInput()

		for _, number := range numbers {
			// 2. Get the classmate with the Attendance Number.
			classmate, err := classmates.GetByAttendanceNumber(number)
			if err != nil {
				log.Println("debug: [main] " + err.Error())
				fmt.Println("error: " + err.Error())
				fmt.Println()
				errorCount++
				continue
			}
			log.Printf("debug: [main] found the Classmate with AttendanceNumber = %d\n", number)

			// 3. Print the classmate data
			classmate.Print()
		}

		if len(numbers) == 0 {
			fmt.Println("Invalid input. Please try again.")
			fmt.Println()
		}

		if !state.interactive {
			requestedCount := len(numbers)
			processedCount := len(numbers) - errorCount
			fmt.Printf("%d requested, %d processed, %d error(s)\n", requestedCount, processedCount, errorCount)
		}

	}
}
