package classmates

import (
	"fmt"
	"strings"
)

// For each string in the slice, if the length of the current string is greater than the current maximum length,
// then set current maximum length to the length of the current string.
// Returns the maximum length of the strings from the slice
func longestStringLengthOnSlice(textSlice *[]string) int {
	maxLen := 0
	for _, text := range *textSlice {
		if currentLen := len(text); currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

// A method that is bound to the Classmate struct.
// Prints the Classmate data with pretty decent formatting
func (classmate Classmate) Print() {
	textToPrint := []string{
		fmt.Sprintf("ID\t\t\t: %s\n", classmate.ID),
		fmt.Sprintf("Name\t\t\t: %s\n", classmate.Name),
		fmt.Sprintf("Address\t\t\t: %s\n", classmate.Address),
		fmt.Sprintf("Occupation\t\t: %s\n", classmate.Occupation),
		fmt.Sprintf("Reasons of Choosing Go\t: %s\n", classmate.ReasonsOfChoosingGo),
	}
	maxLen := longestStringLengthOnSlice(&textToPrint)

	fmt.Println(strings.Repeat("=", maxLen))
	for _, text := range textToPrint {
		fmt.Print(text)
	}
	fmt.Println(strings.Repeat("=", maxLen))
}
