//go:build go1.23

package slice

import "iter"

// Iter() provides a range-based immutable for loop. Example:
//
//	for v := range s.Iter() {
//		fmt.Println(v)
//	}
func (s *SliceWrapper[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range s.Len() {
			if !yield(s.slice[i]) {
				break
			}
		}
	}
}

// IterMut() provides a range-based mutable for loop. Example:
//
//	for vPtr := range s.IterMut() {
//		*vPtr = newValue
//	}
func (s *SliceWrapper[T]) IterMut() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range s.Len() {
			if !yield(&s.slice[i]) {
				break
			}
		}
	}
}

// Iter2() provides a range-based immutable for loop with index. Example:
//
//	for i, v := range s.Iter2() {
//		fmt.Println(i, v)
//	}
func (s *SliceWrapper[T]) Iter2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range s.Len() {
			if !yield(i, s.slice[i]) {
				break
			}
		}
	}
}

// IterMut2() provides a range-based mutable for loop with index. Example:
//
//	for i, vPtr := range s.IterMut2() {
//		*vPtr = newValue
//	}
func (s *SliceWrapper[T]) IterMut2() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range s.Len() {
			if !yield(i, &s.slice[i]) {
				break
			}
		}
	}
}
