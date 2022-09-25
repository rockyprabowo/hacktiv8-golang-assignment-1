# Hacktiv8 Scalable Web Service with Golang - Assignment 1

This is the first assignment for Scalable Web Service with Golang course by Hacktiv8.

This project is a CLI-based application where the user can see their classmates' data from this course.

This project applies the programming concepts on Go I learned from the course so far.

## Usage

* Clone this repository and change the current directory to the directory where you cloned repository.
* Run this command: `$ go run . [num]` where `[num]` is the attendance number you want to find.
* Or, you can build this application first with `$ go build` and then execute the application with `./assignment-1 [num]` on macOS or Linux, or `./assignment-1.exe [num]` on Windows.

## Changelog

### 2.0

* Now the application will find the item by the attendance number.
* Improved UX with interactive mode.
* Now can handle multiple space seperated input.
* Proper debugging when `DEBUG` environment value set to "true-ish" value like `1` or `true`.

### 1.0

* Initial project, with a single "classmate" data from myself. Adding data from the other classmates are WIP.
