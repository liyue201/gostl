package slice

import "github.com/liyue201/gostl/utils/iterator"

type T any

// SliceIterator is an implementation of RandomAccessIterator

var _ iterator.RandomAccessIterator[T] = (*SliceIterator[T])(nil)

// SliceIterator represents a slice iterator
type SliceIterator[T any] struct {
	s        ISlice[T]
	position int
}

// IsValid returns trus if the iterator is valid, othterwise return false
func (iter *SliceIterator[T]) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.s.Len() {
		return true
	}
	return false
}

// Value returns the value of the iterator point to
func (iter *SliceIterator[T]) Value() T {
	return iter.s.At(iter.position)
}

// SetValue sets the value of the iterator point to
func (iter *SliceIterator[T]) SetValue(val T) {
	iter.s.Set(iter.position, val)
}

// Next moves the iterator's position to the next position, and returns itself
func (iter *SliceIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.position < iter.s.Len() {
		iter.position++
	}
	return iter
}

// Prev move the iterator's position to the previous position, and return itself
func (iter *SliceIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator into a new one
func (iter *SliceIterator[T]) Clone() iterator.ConstIterator[T] {
	return &SliceIterator[T]{s: iter.s, position: iter.position}
}

// IteratorAt creates an iterator with the passed position
func (iter *SliceIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	return &SliceIterator[T]{s: iter.s, position: position}
}

// Position returns the position of the iterator
func (iter *SliceIterator[T]) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SliceIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	otherIter, ok := other.(*SliceIterator[T])
	if !ok {
		return false
	}
	if otherIter.position == iter.position {
		return true
	}
	return false
}
