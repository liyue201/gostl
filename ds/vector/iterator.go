package vector

import (
	"github.com/liyue201/gostl/utils/iterator"
)

type T any

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator[T] = (*VectorIterator[T])(nil)

// VectorIterator represents a vector iterator
type VectorIterator[T any] struct {
	vec      *Vector[T]
	position int // the position of iterator point to
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *VectorIterator[T]) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.vec.Size() {
		return true
	}
	return false
}

// Value returns the value of the iterator point to
func (iter *VectorIterator[T]) Value() T {
	val := iter.vec.At(iter.position)
	return val
}

// SetValue sets the value of the iterator point to
func (iter *VectorIterator[T]) SetValue(val T) {
	iter.vec.SetAt(iter.position, val)
}

// Next moves the position of iterator to the next position and returns itself
func (iter *VectorIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.position < iter.vec.Size() {
		iter.position++
	}
	return iter
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *VectorIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator into a new iterator
func (iter *VectorIterator[T]) Clone() iterator.ConstIterator[T] {
	return &VectorIterator[T]{vec: iter.vec, position: iter.position}
}

// IteratorAt creates an iterator with the passed position
func (iter *VectorIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	return &VectorIterator[T]{vec: iter.vec, position: position}
}

// Position return the position of the iterator point to
func (iter *VectorIterator[T]) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *VectorIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	otherIter, ok := other.(*VectorIterator[T])
	if !ok {
		return false
	}
	if otherIter.vec == iter.vec && otherIter.position == iter.position {
		return true
	}
	return false
}
