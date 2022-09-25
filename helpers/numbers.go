package helpers

import (
	"fmt"
	"log"
	"strconv"
)

// It takes a string, and if it's a valid unsigned number, it returns the number as a uint64
func TryParseUint(text string) (uint64, error) {
	num, err := strconv.ParseUint(text, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("error: \"%s\" is not an unsigned number", text)
	}
	return num, nil
}

// It takes a slice of strings, tries to parse each string as a uint64, and returns a slice of uint64s
func ProcessUint64Numbers(input []string) (numbers []uint64) {
	for _, text := range input {
		num, err := TryParseUint(text)
		if err != nil {
			log.Println("debug: [ProcessUint64Numbers] " + err.Error())
			continue
		}
		numbers = append(numbers, num)
	}
	log.Printf("debug: [ProcessUint64Numbers] processed input as %v", numbers)

	return
}
