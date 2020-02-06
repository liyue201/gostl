package slice

import "github.com/liyue201/gostl/utils/iterator"

// SliceIterator is a implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator = (*SliceIterator)(nil)

// SliceIterator represents a slice iterator
type SliceIterator struct {
	s        ISlice
	position int
}

// IsValid returns trus if the iterator is valid, othterwise return false
func (iter *SliceIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.s.Len() {
		return true
	}
	return false
}

// Value returns the value of the iterator point to
func (iter *SliceIterator) Value() interface{} {
	return iter.s.At(iter.position)
}

// SetValue sets the value of the iterator point to
func (iter *SliceIterator) SetValue(val interface{}) {
	iter.s.Set(iter.position, val)
}

// Next moves the iterator's position to the next position, and returns itself
func (iter *SliceIterator) Next() iterator.ConstIterator {
	if iter.position < iter.s.Len() {
		iter.position++
	}
	return iter
}

// Prev move the iterator's position to the previous position, and return itself
func (iter *SliceIterator) Prev() iterator.ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones the iterator into a new one
func (iter *SliceIterator) Clone() iterator.ConstIterator {
	return &SliceIterator{s: iter.s, position: iter.position}
}

// IteratorAt creates an iterator with the passed position
func (iter *SliceIterator) IteratorAt(position int) iterator.RandomAccessIterator {
	return &SliceIterator{s: iter.s, position: position}
}

// Position returns the position of the iterator
func (iter *SliceIterator) Position() int {
	return iter.position
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SliceIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*SliceIterator)
	if !ok {
		return false
	}
	if otherIter.position == iter.position {
		return true
	}
	return false
}
