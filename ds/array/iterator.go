package array

import (
	"github.com/liyue201/gostl/utils/iterator"
)

type T any

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator[T] = (*ArrayIterator[T])(nil)

// ArrayIterator is an implementation of Array iterator
type ArrayIterator[T any] struct {
	array    *Array[T]
	position int
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *ArrayIterator[T]) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.array.Size() {
		return true
	}
	return false
}

// Value returns the value of array at the position of the iterator point to
func (iter *ArrayIterator[T]) Value() T {
	return iter.array.At(iter.position)
}

// SetValue sets the value of the array at the position of the iterator point to
func (iter *ArrayIterator[T]) SetValue(val T) {
	iter.array.Set(iter.position, val)
}

// Next moves the position of iterator to the next position and returns itself
func (iter *ArrayIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.position < iter.array.Size() {
		iter.position++
	}
	return iter
}

// Prev moves the position of iterator to the previous position and returns itself
func (iter *ArrayIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator to a new iterator
func (iter *ArrayIterator[T]) Clone() iterator.ConstIterator[T] {
	return &ArrayIterator[T]{array: iter.array, position: iter.position}
}

// IteratorAt creates a new iterator with position pos
func (iter *ArrayIterator[T]) IteratorAt(pos int) iterator.RandomAccessIterator[T] {
	return &ArrayIterator[T]{array: iter.array, position: pos}
}

// Position returns the position of the iterator
func (iter *ArrayIterator[T]) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ArrayIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	otherIter, ok := other.(*ArrayIterator[T])
	if !ok {
		return false
	}
	if otherIter.array == iter.array && otherIter.position == iter.position {
		return true
	}
	return false
}
