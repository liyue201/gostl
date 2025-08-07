//go:build go1.23

package array

import "iter"

// Iter() provides a range-based immutable for loop. Example:
//
//	for v := range a.Iter() {
//		fmt.Println(v)
//	}
func (a *Array[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range a.values {
			if !yield(v) {
				break
			}
		}
	}
}

// IterMut() provides a range-based mutable for loop. Example:
//
//	for vPtr := range a.IterMut() {
//		*vPtr = newValue
//	}
func (a *Array[T]) IterMut() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range a.values {
			if !yield(&a.values[i]) {
				break
			}
		}
	}
}

// Iter2() provides a range-based immutable for loop with index. Example:
//
//	for i, v := range a.Iter2() {
//		fmt.Println(i, v)
//	}
func (a *Array[T]) Iter2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range a.values {
			if !yield(i, v) {
				break
			}
		}
	}
}

// IterMut2() provides a range-based mutable for loop with index. Example:
//
//	for i, vPtr := range a.IterMut2() {
//		*vPtr = newValue
//	}
func (a *Array[T]) IterMut2() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range a.values {
			if !yield(i, &a.values[i]) {
				break
			}
		}
	}
}
