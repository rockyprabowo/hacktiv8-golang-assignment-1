package helpers

// Set implementation using a map of struct{}s.
//
// The map is used to store the set's elements. The struct{} is an empty type, which means it takes up
// no space. It's used as a placeholder to indicate that the map's key is present.
//
// The Set type is generic, which means it can be used to create sets of any comparable type.
//
// @property sets - This is a map of type T to an empty struct.
type Set[T comparable] struct {
	sets map[T]struct{}
}

// Checks if the set has a value.
func (s *Set[T]) Has(v T) bool {
	_, ok := s.sets[v]
	return ok
}

// Adds a new element to the set.
func (s *Set[T]) Add(v T) {
	s.sets[v] = struct{}{}
}

// Removes a value from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.sets, v)
}

// Clears the set.
func (s *Set[T]) Clear() {
	s.sets = make(map[T]struct{})
}

// Returns the size of the set.
func (s *Set[T]) Size() int {
	return len(s.sets)
}

// Creates a new set of type T
func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.sets = make(map[T]struct{})
	return s
}
