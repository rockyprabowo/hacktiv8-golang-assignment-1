package classmates

import (
	"fmt"
	h "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"
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
	if !collection.invalid.Has(attendanceNumber) {
		h.PrintDebug(
			fmt.Sprintf("%d is not marked as invalid (yet)", attendanceNumber),
			"GetByAttendanceNumber",
		)

		if classmate := collection.MemoGet(attendanceNumber); classmate != nil {
			h.PrintDebug(
				fmt.Sprintf("got classmate from memo %d to %p", attendanceNumber, classmate),
				"GetByAttendanceNumber",
			)
			return classmate, nil
		}

		h.PrintDebug(fmt.Sprintf("classmate %d is not memoised yet", attendanceNumber), "GetByAttendanceNumber")

		for _, classmate := range collection.classmates {
			if classmate.AttendanceNumber == attendanceNumber {
				collection.MemoSet(attendanceNumber, &classmate)
				h.PrintDebug(fmt.Sprintf("memoised %d to %p", attendanceNumber, &classmate), "GetByAttendanceNumber")
				return &classmate, nil
			}
		}

		h.PrintDebug(fmt.Sprintf("classmate %d is invalid", attendanceNumber), "GetByAttendanceNumber")
		collection.invalid.Add(attendanceNumber)
	}
	h.PrintDebug(fmt.Sprintf("classmate %d is marked as invalid", attendanceNumber), "GetByAttendanceNumber")
	return nil, fmt.Errorf("couldn't find a classmate with AttendanceNumber of %d", attendanceNumber)
}
