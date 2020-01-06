package vector

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*VectorIterator)(nil)

// VectorIterator is an implementation of iterator for Vector
type VectorIterator struct {
	vec      *Vector
	position int
}

// IsValid returns whether iter is valid or not
func (iter *VectorIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.vec.Size() {
		return true
	}
	return false
}

// Value returns the internal value of iter
func (iter *VectorIterator) Value() interface{} {
	val := iter.vec.At(iter.position)
	return val
}

// SetValue sets the value of iter
func (iter *VectorIterator) SetValue(val interface{}) error {
	return iter.vec.SetAt(iter.position, val)
}

// Next moves iter to next position and returns iter
func (iter *VectorIterator) Next() ConstIterator {
	if iter.position < iter.vec.Size() {
		iter.position++
	}
	return iter
}

// Prev moves iter to previous position and returns iter
func (iter *VectorIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones iter to a new DequeIterator
func (iter *VectorIterator) Clone() ConstIterator {
	return &VectorIterator{vec: iter.vec, position: iter.position}
}

// IteratorAt new and iterator with position at the passed position
func (iter *VectorIterator) IteratorAt(position int) RandomAccessIterator {
	return &VectorIterator{vec: iter.vec, position: position}
}

// Position return the position of iterator
func (iter *VectorIterator) Position() int {
	return iter.position
}

// Equal returns whether iter is equal to other
func (iter *VectorIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*VectorIterator)
	if !ok {
		return false
	}
	if otherIter.vec == iter.vec && otherIter.position == iter.position {
		return true
	}
	return false
}
