//go:build go1.23

package bidlist

import "iter"

// Iter() provides a range-based immutable for loop. Example:
//
//	for v := range l.Iter() {
//		fmt.Println(v)
//	}
func (l *List[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := l.FrontNode(); n != nil; n = n.Next() {
			if !yield(n.Value) {
				break
			}
		}
	}
}

// IterMut() provides a range-based mutable for loop. Example:
//
//	for vPtr := range l.IterMut() {
//		*vPtr = newValue
//	}
func (l *List[T]) IterMut() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for n := l.FrontNode(); n != nil; n = n.Next() {
			if !yield(&n.Value) {
				break
			}
		}
	}
}

// Iter2() provides a range-based immutable for loop with index. Example:
//
//	for i, v := range l.Iter2() {
//		fmt.Println(i, v)
//	}
func (l *List[T]) Iter2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		total := 0
		for n := l.FrontNode(); n != nil; n = n.Next() {
			if !yield(total, n.Value) {
				break
			}
			total++
		}
	}
}

// IterMut2() provides a range-based mutable for loop with index. Example:
//
//	for i, vPtr := range l.IterMut2() {
//		*vPtr = newValue
//	}
func (l *List[T]) IterMut2() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		total := 0
		for n := l.FrontNode(); n != nil; n = n.Next() {
			if !yield(total, &n.Value) {
				break
			}
			total++
		}
	}
}
