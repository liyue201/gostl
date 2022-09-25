package deque

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// DequeIterator is an implementation of RandomAccessIterator
type T any

var _ iterator.RandomAccessIterator[T] = (*DequeIterator[T])(nil)

// DequeIterator is an implementation of Deque iterator
type DequeIterator[T any] struct {
	dq       *Deque[T]
	position int
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *DequeIterator[T]) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.dq.Size() {
		return true
	}
	return false
}

// Value returns the value of the deque at the position of the iterator point to
func (iter *DequeIterator[T]) Value() T {
	return iter.dq.At(iter.position)
}

// SetValue sets the value of the deque at the position of the iterator point to
func (iter *DequeIterator[T]) SetValue(val T) {
	iter.dq.Set(iter.position, val)
}

// Next moves the position of the iterator to the next position and returns itself
func (iter *DequeIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.position < iter.dq.Size() {
		iter.position++
	}
	return iter
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *DequeIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator to a new iterator
func (iter *DequeIterator[T]) Clone() iterator.ConstIterator[T] {
	return &DequeIterator[T]{dq: iter.dq, position: iter.position}
}

// IteratorAt creates a new iterator with the passed position
func (iter *DequeIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	return &DequeIterator[T]{dq: iter.dq, position: position}
}

// Position returns the position of iterator
func (iter *DequeIterator[T]) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *DequeIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	otherIter, ok := other.(*DequeIterator[T])
	if !ok {
		return false
	}
	if otherIter.dq == iter.dq && otherIter.position == iter.position {
		return true
	}
	return false
}
