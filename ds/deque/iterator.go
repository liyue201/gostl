package deque

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*DequeIterator)(nil)

// DequeIterator is an implementation of iterator for Deque
type DequeIterator struct {
	dq       *Deque
	position int
}

// IsValid returns whether iter is valid
func (iter *DequeIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.dq.Size() {
		return true
	}
	return false
}

// Value returns the internal value of iter
func (iter *DequeIterator) Value() interface{} {
	return iter.dq.At(iter.position)
}

// SetValue sets the value of iter
func (iter *DequeIterator) SetValue(val interface{}) error {
	return iter.dq.Set(iter.position, val)
}

// Next moves iter to next position and returns iter
func (iter *DequeIterator) Next() ConstIterator {
	if iter.position < iter.dq.Size() {
		iter.position++
	}
	return iter
}

// Prev moves iter to previous position and returns iter
func (iter *DequeIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones iter to a new DequeIterator
func (iter *DequeIterator) Clone() ConstIterator {
	return &DequeIterator{dq: iter.dq, position: iter.position}
}

//  IteratorAt new and iterator with position at the passed position
func (iter *DequeIterator) IteratorAt(position int) RandomAccessIterator {
	return &DequeIterator{dq: iter.dq, position: position}
}

// Position return the position of iterator
func (iter *DequeIterator) Position() int {
	return iter.position
}

// Equal returns whether iter is equal to other
func (iter *DequeIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*DequeIterator)
	if !ok {
		return false
	}
	if otherIter.dq == iter.dq && otherIter.position == iter.position {
		return true
	}
	return false
}
