package helpers

// Set represents the datastructure of a set.
type Set[T comparable] map[T]struct{}

// Clear clears the set.
func (s *Set[T]) Clear() {
	s = &Set[T]{}
}
