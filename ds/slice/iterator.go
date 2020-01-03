package slice

import . "github.com/liyue201/gostl/utils/iterator"

//SliceIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*SliceIterator)(nil)

type SliceIterator struct {
	s        ISlice
	position int
}

func (iter *SliceIterator) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.s.Len() {
		return true
	}
	return false
}

func (iter *SliceIterator) Value() interface{} {
	return iter.s.At(iter.position)
}

func (iter *SliceIterator) SetValue(val interface{}) error {
	iter.s.Set(iter.position, val)
	return nil
}

func (iter *SliceIterator) Next() ConstIterator {
	if iter.position < iter.s.Len() {
		iter.position++
	}
	return iter
}

func (iter *SliceIterator) Prev() ConstBidIterator {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

func (iter *SliceIterator) Clone() ConstIterator {
	return &SliceIterator{s: iter.s, position: iter.position}
}

func (iter *SliceIterator) IteratorAt(position int) RandomAccessIterator {
	return &SliceIterator{s: iter.s, position: position}
}

func (iter *SliceIterator) Position() int {
	return iter.position
}

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
