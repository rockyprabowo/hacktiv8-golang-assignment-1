package main

import (
	"fmt"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
)

// A struct that holds the state of the application.
//
// The `interactive` field is a boolean that indicates whether the application is running in
// interactive mode.
//
// The `debugMode` field is a boolean that indicates whether the application is running in debug mode.
//
// The `fromStdin` field is a boolean that indicates whether the application is reading from stdin.
//
// The `fromArgs` field is a boolean that indicates whether the application is reading from command
// line arguments.
// @property {bool} interactive - If true, the app will run in interactive mode.
// @property {bool} debugMode - If true, the app will print debug information to the console.
// @property {bool} fromStdin - true if the input is from stdin
// @property {bool} fromArgs - true if the input is from command line arguments.
type _AppState struct {
	interactive bool
	debugMode   bool
	fromStdin   bool
	fromArgs    bool
}

// This variable holds the application states.
var state _AppState

// Sets up the application state and determines if the application should enter interactive
// mode
func Setup() {
	// Setup logging and application initial states.
	h.SetupLogging()

	state = _AppState{
		debugMode: h.IsDebug(),
		fromStdin: h.HasDataInStdin(),
		fromArgs:  h.HasArguments(),
	}
	// Application will enter interactive mode if there is no input from stdin or command arguments
	state.interactive = !(state.fromStdin || state.fromArgs)
}

// Prints the banner of the program
func ShowBanner() {
	fmt.Println("Assignment 1 - Classmates")
	fmt.Println("Made by Rocky Prabowo (149368582101-196)")
	fmt.Println()
}
