package array

import (
	. "github.com/liyue201/gostl/uitls/iterator"
)

//ArrayIterator is a SortableIterator
var _ SortableIterator = (*ArrayIterator)(nil)

type ArrayIterator struct {
	array    *Array
	position int
}

func (this *ArrayIterator) IsValid() bool {
	if this.position >= 0 && this.position < this.array.Size() {
		return true
	}
	return false
}

func (this *ArrayIterator) Value() interface{} {
	return this.array.At(this.position)
}

func (this *ArrayIterator) SetValue(val interface{}) error {
	return this.array.Set(this.position, val)
}

func (this *ArrayIterator) Next() ConstIterator {
	if this.position < this.array.Size() {
		this.position++
	}
	return this
}

func (this *ArrayIterator) Prev() ConstBidIterator {
	if this.position >= 0 {
		this.position--
	}
	return this
}

func (this *ArrayIterator) Clone() interface{} {
	return &ArrayIterator{array: this.array, position: this.position}
}

func (this *ArrayIterator) IteratorAt(position int) SortableIterator {
	return &ArrayIterator{array: this.array, position: position}
}

func (this *ArrayIterator) Position() int {
	return this.position
}
