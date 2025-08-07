//go:build go1.23

package deque

import "iter"

// Iter() provides a range-based immutable for loop. Example:
//
//	for v := range d.Iter() {
//		fmt.Println(v)
//	}
func (d *Deque[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range d.Size() {
			if !yield(d.At(i)) {
				break
			}
		}
	}
}

// IterMut() provides a range-based mutable for loop. Example:
//
//	for vPtr := range d.IterMut() {
//		*vPtr = newValue
//	}
func (d *Deque[T]) IterMut() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range d.Size() {
			seg, pos := d.pos(i)
			s := d.segmentAt(seg)
			ptr := &s.data[(pos+s.begin)%s.capacity()]
			if !yield(ptr) {
				break
			}
		}
	}
}

// Iter2() provides a range-based immutable for loop with index. Example:
//
//	for i, v := range d.Iter2() {
//		fmt.Println(i, v)
//	}
func (d *Deque[T]) Iter2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range d.Size() {
			seg, pos := d.pos(i)
			s := d.segmentAt(seg)
			if !yield(i, s.at(pos)) {
				break
			}
		}
	}
}

// IterMut2() provides a range-based mutable for loop with index. Example:
//
//	for i, vPtr := range d.IterMut2() {
//		*vPtr = newValue
//	}
func (d *Deque[T]) IterMut2() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range d.Size() {
			seg, pos := d.pos(i)
			s := d.segmentAt(seg)
			ptr := &s.data[(pos+s.begin)%s.capacity()]
			if !yield(i, ptr) {
				break
			}
		}
	}
}
