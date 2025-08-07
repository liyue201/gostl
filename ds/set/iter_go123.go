//go:build go1.23

package set

import "iter"

// Iter() provides a range-based immutable for loop. Example:
//
//	for v := range s.Iter() {
//		fmt.Println(v)
//	}
func (s *Set[T]) Iter() iter.Seq[T] {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.tree.Iter()
}

func (s *MultiSet[T]) Iter() iter.Seq[T] {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.tree.Iter()
}
