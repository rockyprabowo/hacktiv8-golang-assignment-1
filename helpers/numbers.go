package helpers

import (
	"fmt"
	"strconv"
)

// Takes a string, and if it's a valid unsigned number, it returns the number as a uint64.
func TryParseUint(text string) (uint64, error) {
	num, err := strconv.ParseUint(text, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("error: \"%s\" is not an unsigned number", text)
	}
	return num, nil
}

// Takes a slice of strings, then try to parse each string as a uint64, and returns the slice of parsed uint64s.
func ProcessUint64Numbers(input []string) (numbers []uint64) {
	numbers = []uint64{}
	for _, text := range input {
		num, err := TryParseUint(text)
		if err != nil {
			PrintDebug(err.Error(), "ProcessUint64Numbers")
			continue
		}
		numbers = append(numbers, num)
	}
	PrintDebug(fmt.Sprintf("processed input as %v", numbers), "ProcessUint64Numbers")

	return
}
