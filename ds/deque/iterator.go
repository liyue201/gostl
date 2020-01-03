package deque

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*DequeIterator)(nil)

type DequeIterator struct {
	dq       *Deque
	position int
}

func (iter *DequeIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.dq.Size() {
		return true
	}
	return false
}

func (iter *DequeIterator) Value() interface{} {
	return iter.dq.At(iter.position)
}

func (iter *DequeIterator) SetValue(val interface{}) error {
	return iter.dq.Set(iter.position, val)
}

func (iter *DequeIterator) Next() ConstIterator {
	if iter.position < iter.dq.Size() {
		iter.position++
	}
	return iter
}

func (iter *DequeIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

func (iter *DequeIterator) Clone() ConstIterator {
	return &DequeIterator{dq: iter.dq, position: iter.position}
}

func (iter *DequeIterator) IteratorAt(position int) RandomAccessIterator {
	return &DequeIterator{dq: iter.dq, position: position}
}

func (iter *DequeIterator) Position() int {
	return iter.position
}

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
