package array

import (
	"github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*ArrayIterator)(nil)

// ArrayIterator is an an implementation of Array iterator
type ArrayIterator struct {
	array    *Array
	position int
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *ArrayIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.array.Size() {
		return true
	}
	return false
}

// Value returns the value of array at the position of the iterator point to
func (iter *ArrayIterator) Value() interface{} {
	return iter.array.At(iter.position)
}

// SetValue sets the value of the array at the position of the iterator point to
func (iter *ArrayIterator) SetValue(val interface{}) {
	iter.array.Set(iter.position, val)
}

// Next moves the position of iterator to the next position and returns itself
func (iter *ArrayIterator) Next() iterator.ConstIterator {
	if iter.position < iter.array.Size() {
		iter.position++
	}
	return iter
}

// Prev moves the position of iterator to the previous position and returns itself
func (iter *ArrayIterator) Prev() iterator.ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator to a new iterator
func (iter *ArrayIterator) Clone() iterator.ConstIterator {
	return &ArrayIterator{array: iter.array, position: iter.position}
}

// IteratorAt creates a new iterator with position pos
func (iter *ArrayIterator) IteratorAt(pos int) iterator.RandomAccessIterator {
	return &ArrayIterator{array: iter.array, position: pos}
}

// Position returns the position of the iterator
func (iter *ArrayIterator) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ArrayIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*ArrayIterator)
	if !ok {
		return false
	}
	if otherIter.array == iter.array && otherIter.position == iter.position {
		return true
	}
	return false
}
