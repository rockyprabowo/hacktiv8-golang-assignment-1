package classmates

import "rockyprabowo/hacktiv8-assignments/assignment-1/helpers"

// `_ClassmateCollection` is a struct that contains a slice of `Classmate` data, a map of memoised `Classmate`s, and
// a set of invalid `Classmate` AttendanceNumbers.
// @property {[]Classmate} classmates - This is the slice of Classmate objects that we'll be working
// with.
// @property memo - a map of classmate IDs to classmates. This is used to quickly look up classmates by
// ID.
// @property invalid - a set of invalid `Classmate`s AttendanceNumber
type _ClassmateCollection struct {
	classmates []Classmate
	memo       map[uint64]*Classmate
	invalid    *helpers.Set[uint64]
}

// Gets a classmate with its key in the collection memo.
func (collection _ClassmateCollection) MemoGet(key uint64) *Classmate {
	return collection.memo[key]
}

// Memoise a classmate in the collection with a key.
func (collection _ClassmateCollection) MemoSet(key uint64, ptr *Classmate) {
	collection.memo[key] = ptr
}
