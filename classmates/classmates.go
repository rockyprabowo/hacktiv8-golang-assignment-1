package classmates

import (
	"fmt"
)

// A Classmate is a struct with five fields.
// @property {uint64} AttendanceNumber - The number of the classmate.
// @property {string} Name - The name of the classmate.
// @property {string} Address - The address of the classmate.
// @property {string} Occupation - What do you do for a living?
// @property {string} ReasonsOfChoosingGo - Why did you choose Go?
type Classmate struct {
	AttendanceNumber    uint64
	Name                string
	Address             string
	Occupation          string
	ReasonsOfChoosingGo string
}

// GetByAttendanceNumber returns a pointer to a Classmate struct if it finds a classmate with the given Attendance Number.
// Otherwise, it returns an error.
func GetByAttendanceNumber(attendanceNumber uint64) (*Classmate, error) {
	for _, classmate := range classmates {
		if classmate.AttendanceNumber == attendanceNumber {
			return &classmate, nil
		}
	}
	return nil, fmt.Errorf("couldn't find a classmate with AttendanceNumber of %d", attendanceNumber)
}
