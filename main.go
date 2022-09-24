package main

import (
	"fmt"
	"os"
	"rockyprabowo/hacktiv8-assignments/assignment-1/classmates"
)

// It returns the first argument passed to the program, or an error if no argument was passed
func getIdFromArgument() (string, error) {
	if len(os.Args) == 1 {
		return "", fmt.Errorf("error: please provide an ID")
	}
	if len(os.Args) > 2 {
		fmt.Println("warning: only the first argument passed, the rest are ignored")
	}
	return os.Args[1], nil
}

// If there's an error, print it and exit the program
func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	// 1. Get the ID from command argument (os.Args)
	id, err := getIdFromArgument()
	if err != nil {
		exitWithError(err)
	}

	// 2. Get the classmate with the ID
	classmate, err := classmates.GetById(id)
	if err != nil {
		exitWithError(err)
	}
	fmt.Printf("Found the classmate with the ID = %s\n", id)

	// 3. Print the classmate data
	classmate.Print()
}
