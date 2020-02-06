package deque

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// DequeIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*DequeIterator)(nil)

// DequeIterator is an implementation of Deque iterator
type DequeIterator struct {
	dq       *Deque
	position int
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *DequeIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.dq.Size() {
		return true
	}
	return false
}

// Value returns the value of the deque at the position of the iterator point to
func (iter *DequeIterator) Value() interface{} {
	return iter.dq.At(iter.position)
}

// SetValue sets the value of the deque at the position of the iterator point to
func (iter *DequeIterator) SetValue(val interface{}) {
	iter.dq.Set(iter.position, val)
}

// Next moves the position of the iterator to the next position and returns itself
func (iter *DequeIterator) Next() iterator.ConstIterator {
	if iter.position < iter.dq.Size() {
		iter.position++
	}
	return iter
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *DequeIterator) Prev() iterator.ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator to a new iterator
func (iter *DequeIterator) Clone() iterator.ConstIterator {
	return &DequeIterator{dq: iter.dq, position: iter.position}
}

// IteratorAt creates a new iterator with the passed position
func (iter *DequeIterator) IteratorAt(position int) iterator.RandomAccessIterator {
	return &DequeIterator{dq: iter.dq, position: position}
}

// Position returns the position of iterator
func (iter *DequeIterator) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *DequeIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*DequeIterator)
	if !ok {
		return false
	}
	if otherIter.dq == iter.dq && otherIter.position == iter.position {
		return true
	}
	return false
}
