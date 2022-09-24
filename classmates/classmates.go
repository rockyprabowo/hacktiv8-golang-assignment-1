package classmates

import (
	"fmt"
)

type Classmate struct {
	ID                  string
	Name                string
	Address             string
	Occupation          string
	ReasonsOfChoosingGo string
}

// GetById returns a pointer to a Classmate struct if it finds a classmate with the given ID
// Otherwise, it returns an error.
func GetById(id string) (*Classmate, error) {
	for _, classmate := range classmates {
		if classmate.ID == id {
			return &classmate, nil
		}
	}
	return nil, fmt.Errorf("error: Couldn't find a classmate with ID of %s", id)
}
