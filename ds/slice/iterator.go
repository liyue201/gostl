package slice

import . "github.com/liyue201/gostl/utils/iterator"

//SliceIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*SliceIterator)(nil)

type SliceIterator struct {
	s        ISlice
	position int
}

func (this *SliceIterator) IsValid() bool {
	if this.position >= 0 && this.position < this.s.Len() {
		return true
	}
	return false
}

func (this *SliceIterator) Value() interface{} {
	return this.s.At(this.position)
}

func (this *SliceIterator) SetValue(val interface{}) error {
	this.s.Set(this.position, val)
	return nil
}

func (this *SliceIterator) Next() ConstIterator {
	if this.position < this.s.Len() {
		this.position++
	}
	return this
}

func (this *SliceIterator) Prev() ConstBidIterator {
	if this.position >= 0 {
		this.position--
	}
	return this
}

func (this *SliceIterator) Clone() ConstIterator {
	return &SliceIterator{s: this.s, position: this.position}
}

func (this *SliceIterator) IteratorAt(position int) RandomAccessIterator {
	return &SliceIterator{s: this.s, position: position}
}

func (this *SliceIterator) Position() int {
	return this.position
}

func (this *SliceIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*SliceIterator)
	if !ok {
		return false
	}
	if otherIter.s == this.s && otherIter.position == this.position {
		return true
	}
	return false
}
