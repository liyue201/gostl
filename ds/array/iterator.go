package array

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*ArrayIterator)(nil)

type ArrayIterator struct {
	array    *Array
	position int
}

func (iter *ArrayIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.array.Size() {
		return true
	}
	return false
}

func (iter *ArrayIterator) Value() interface{} {
	return iter.array.At(iter.position)
}

func (iter *ArrayIterator) SetValue(val interface{}) error {
	return iter.array.Set(iter.position, val)
}

func (iter *ArrayIterator) Next() ConstIterator {
	if iter.position < iter.array.Size() {
		iter.position++
	}
	return iter
}

func (iter *ArrayIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

func (iter *ArrayIterator) Clone() ConstIterator {
	return &ArrayIterator{array: iter.array, position: iter.position}
}

func (iter *ArrayIterator) IteratorAt(position int) RandomAccessIterator {
	return &ArrayIterator{array: iter.array, position: position}
}

func (iter *ArrayIterator) Position() int {
	return iter.position
}

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
