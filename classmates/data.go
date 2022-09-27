package classmates

import "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"

// This variable holds the main `Classmate` collection.
// TODO: Use real data when it's possible.
var collection = _ClassmateCollection{
	memo:    map[uint64]*Classmate{},
	invalid: helpers.NewSet[uint64](),
	classmates: []Classmate{
		{
			AttendanceNumber:    1,
			Name:                "Rocky Prabowo",
			Address:             "Palembang",
			Occupation:          "Freelancer",
			ReasonsOfChoosingGo: "Learning a new language for future job(s)",
		},
		{
			AttendanceNumber:    2,
			Name:                "Prabowo",
			Address:             "Yogyakarta",
			Occupation:          "-",
			ReasonsOfChoosingGo: "I love programming",
		},
		{
			AttendanceNumber:    100,
			Name:                "Rocky P.",
			Address:             "Pekanbaru",
			Occupation:          "Engineer",
			ReasonsOfChoosingGo: "Finding a new job(s) opportunity",
		},
		{
			AttendanceNumber:    3,
			Name:                "R. Prabowo",
			Address:             "Jakarta",
			Occupation:          "Programmer",
			ReasonsOfChoosingGo: "Learning a new language",
		},
	},
}
