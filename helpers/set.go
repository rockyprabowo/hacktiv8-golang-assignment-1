package helpers

type Set[T comparable] struct {
	sets map[T]struct{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.sets[v]
	return ok
}

func (s *Set[T]) Add(v T) {
	s.sets[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.sets, v)
}

func (s *Set[T]) Clear() {
	s.sets = make(map[T]struct{})
}

func (s *Set[T]) Size() int {
	return len(s.sets)
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.sets = make(map[T]struct{})
	return s
}
