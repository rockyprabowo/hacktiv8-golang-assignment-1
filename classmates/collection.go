package classmates

import "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"

// A _ClassmateCollection is a collection of Classmates.
// @property {[]Classmate} data - This is the slice that will hold the data.
type _ClassmateCollection struct {
	classmates []Classmate
	memo       map[uint64]*Classmate
	invalid    *helpers.Set[uint64]
}

func (collection _ClassmateCollection) MemoGet(key uint64) *Classmate {
	return collection.memo[key]
}

func (collection _ClassmateCollection) MemoSet(key uint64, ptr *Classmate) {
	collection.memo[key] = ptr
}
