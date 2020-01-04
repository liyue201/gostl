package slice

import . "github.com/liyue201/gostl/utils/iterator"

// SliceIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*SliceIterator)(nil)

// SliceIterator is a iterator implementation of slice
type SliceIterator struct {
	s        ISlice
	position int
}

// IsValid returns whether iter is valid
func (iter *SliceIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.s.Len() {
		return true
	}
	return false
}

// Value returns the internal value of iter
func (iter *SliceIterator) Value() interface{} {
	return iter.s.At(iter.position)
}

// SetValue sets the value of iter
func (iter *SliceIterator) SetValue(val interface{}) error {
	iter.s.Set(iter.position, val)
	return nil
}

// Next returns the next iterator
func (iter *SliceIterator) Next() ConstIterator {
	if iter.position < iter.s.Len() {
		iter.position++
	}
	return iter
}

// Prev returns the previous iterator
func (iter *SliceIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone clones iter to a new SliceIterator
func (iter *SliceIterator) Clone() ConstIterator {
	return &SliceIterator{s: iter.s, position: iter.position}
}

// IteratorAt new and iterator with position at the position passed
func (iter *SliceIterator) IteratorAt(position int) RandomAccessIterator {
	return &SliceIterator{s: iter.s, position: position}
}

// Position return the position of iterator
func (iter *SliceIterator) Position() int {
	return iter.position
}

// Equal returns whether iter is equal to other
func (iter *SliceIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*SliceIterator)
	if !ok {
		return false
	} 
	if otherIter.position == iter.position {
		return true
	}
	return false
}
