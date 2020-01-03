package vector

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*VectorIterator)(nil)

type VectorIterator struct {
	vec      *Vector
	position int
}

func (iter *VectorIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.vec.Size() {
		return true
	}
	return false
}

func (iter *VectorIterator) Value() interface{} {
	val := iter.vec.At(iter.position)
	return val
}

func (iter *VectorIterator) SetValue(val interface{}) error {
	return iter.vec.SetAt(iter.position, val)
}

func (iter *VectorIterator) Next() ConstIterator {
	if iter.position < iter.vec.Size() {
		iter.position++
	}
	return iter
}

func (iter *VectorIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

func (iter *VectorIterator) Clone() ConstIterator {
	return &VectorIterator{vec: iter.vec, position: iter.position}
}

func (iter *VectorIterator) IteratorAt(position int) RandomAccessIterator {
	return &VectorIterator{vec: iter.vec, position: position}
}

func (iter *VectorIterator) Position() int {
	return iter.position
}

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
