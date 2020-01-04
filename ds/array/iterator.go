package array

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*ArrayIterator)(nil)

// DequeIterator is an implementation of iterator for Array
type ArrayIterator struct {
	array    *Array
	position int
}

// IsValid returns whether iter is valid
func (iter *ArrayIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.array.Size() {
		return true
	}
	return false
}

// Value returns the internal value of iter
func (iter *ArrayIterator) Value() interface{} {
	return iter.array.At(iter.position)
}

// SetValue sets the value of iter
func (iter *ArrayIterator) SetValue(val interface{}) error {
	return iter.array.Set(iter.position, val)
}

// Next moves iter to next position and returns iter
func (iter *ArrayIterator) Next() ConstIterator {
	if iter.position < iter.array.Size() {
		iter.position++
	}
	return iter
}

// Prev moves iter to previous position and returns iter
func (iter *ArrayIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones iter to a new ArrayIterator
func (iter *ArrayIterator) Clone() ConstIterator {
	return &ArrayIterator{array: iter.array, position: iter.position}
}

// IteratorAt new and iterator with position at the passed position
func (iter *ArrayIterator) IteratorAt(position int) RandomAccessIterator {
	return &ArrayIterator{array: iter.array, position: position}
}


// Position return the position of iterator
func (iter *ArrayIterator) Position() int {
	return iter.position
}

// Equal returns whether iter is equal to other
func (iter *ArrayIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*ArrayIterator)
	if !ok {
		return false
	}
	if otherIter.array == iter.array && otherIter.position == iter.position {
		return true
	}
	return false
}
