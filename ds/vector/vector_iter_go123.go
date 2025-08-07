//go:build go1.23

package vector

import "iter"

// Iter() provides a range-based immutable for loop. Example:
//
//	for v := range v.Iter() {
//		fmt.Println(v)
//	}
func (v *Vector[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range v.Size() {
			if !yield(v.data[i]) {
				break
			}
		}
	}
}

// IterMut() provides a range-based mutable for loop. Example:
//
//	for vPtr := range v.IterMut() {
//		*vPtr = newValue
//	}
func (v *Vector[T]) IterMut() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range v.Size() {
			if !yield(&v.data[i]) {
				break
			}
		}
	}
}

// Iter2() provides a range-based immutable for loop with index. Example:
//
//	for i, v := range v.Iter2() {
//		fmt.Println(i, v)
//	}
func (v *Vector[T]) Iter2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range v.Size() {
			if !yield(i, v.data[i]) {
				break
			}
		}
	}
}

// IterMut2() provides a range-based mutable for loop with index. Example:
//
//	for i, vPtr := range v.IterMut2() {
//		*vPtr = newValue
//	}
func (v *Vector[T]) IterMut2() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range v.Size() {
			if !yield(i, &v.data[i]) {
				break
			}
		}
	}
}
