package classmates

import (
	"fmt"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
	"strings"
)

// Prints the classmate data with pretty decent formatting
func (classmate Classmate) Print() {
	textToPrint := []string{
		fmt.Sprintf("Attendance Number\t: %d\n", classmate.AttendanceNumber),
		fmt.Sprintf("Name\t\t\t: %s\n", classmate.Name),
		fmt.Sprintf("Address\t\t\t: %s\n", classmate.Address),
		fmt.Sprintf("Occupation\t\t: %s\n", classmate.Occupation),
		fmt.Sprintf("Reasons of Choosing Go\t: %s\n", classmate.ReasonsOfChoosingGo),
	}
	maxLen := h.MaxLenOnSlice(&textToPrint)

	fmt.Println(strings.Repeat("=", maxLen))
	for _, text := range textToPrint {
		fmt.Print(text)
	}
	fmt.Println(strings.Repeat("=", maxLen))
	fmt.Println()
}
