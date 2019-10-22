package slice

import . "github.com/liyue201/gostl/iterator"


//SliceIterator is a SortableIterator
var _ SortableIterator = (*SliceIterator)(nil)

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

func (this *SliceIterator) IteratorAt(position int) SortableIterator {
	return &SliceIterator{s: this.s, position: position}
}

func (this *SliceIterator) Position() int {
	return this.position
}
