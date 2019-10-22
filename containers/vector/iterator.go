package vector

import (
	. "github.com/liyue201/gostl/iterator"
)
 
//ArrayIterator is a SortableIterator
var _ SortableIterator = (*VectorIterator)(nil)

type VectorIterator struct {
	vec      *Vector
	position int
}

func (this *VectorIterator) IsValid() bool {
	if this.position >= 0 && this.position < this.vec.Size() {
		return true
	}
	return false
}

func (this *VectorIterator) Value() interface{} {
	val := this.vec.At(this.position)
	return val
}

func (this *VectorIterator) SetValue(val interface{}) error {
	return this.vec.SetAt(this.position, val)
}

func (this *VectorIterator) Next() ConstIterator {
	if this.position < this.vec.Size() {
		this.position++
	}
	return this
}

func (this *VectorIterator) Prev() ConstBidIterator {
	if this.position >= 0 {
		this.position--
	}
	return this
}

func (this *VectorIterator) Clone() ConstIterator {
	return &VectorIterator{vec: this.vec, position: this.position}
}

func (this *VectorIterator) IteratorAt(position int) SortableIterator {
	return &VectorIterator{vec: this.vec, position: position}
}

func (this *VectorIterator) Position() int {
	return this.position
}
