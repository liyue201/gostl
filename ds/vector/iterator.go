package vector

import (
	"github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*VectorIterator)(nil)

// VectorIterator represents a vector iterator
type VectorIterator struct {
	vec      *Vector
	position int // the position of iterator point to
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *VectorIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.vec.Size() {
		return true
	}
	return false
}

// Value returns the value of the iterator point to
func (iter *VectorIterator) Value() interface{} {
	val := iter.vec.At(iter.position)
	return val
}

// SetValue sets the value of the iterator point to
func (iter *VectorIterator) SetValue(val interface{}) {
	iter.vec.SetAt(iter.position, val)
}

// Next moves the position of iterator to the next position and returns itself
func (iter *VectorIterator) Next() iterator.ConstIterator {
	if iter.position < iter.vec.Size() {
		iter.position++
	}
	return iter
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *VectorIterator) Prev() iterator.ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator into a new iterator
func (iter *VectorIterator) Clone() iterator.ConstIterator {
	return &VectorIterator{vec: iter.vec, position: iter.position}
}

// IteratorAt creates an iterator with the passed position
func (iter *VectorIterator) IteratorAt(position int) iterator.RandomAccessIterator {
	return &VectorIterator{vec: iter.vec, position: position}
}

// Position return the position of the iterator point to
func (iter *VectorIterator) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *VectorIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*VectorIterator)
	if !ok {
		return false
	}
	if otherIter.vec == iter.vec && otherIter.position == iter.position {
		return true
	}
	return false
}
