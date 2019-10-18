package vector

import (
	. "github.com/liyue201/gostl/container"
)

func (this *VectorIterator) IsValid() bool {
	if this.curIndex >= 0 && this.curIndex < this.vec.Size() {
		return true
	}
	return false
} 

func (this *VectorIterator) Value() interface{} {
	val := this.vec.At(this.curIndex)
	return val
}

func (this *VectorIterator) SetValue(val interface{}) error {
	return this.vec.SetAt(this.curIndex, val)
}

type VectorIterator struct {
	vec      *Vector
	curIndex int
}

func (this *VectorIterator) Next() ConstIterator {
	if this.IsValid() {
		this.curIndex++
	}
	return this
}

func (this *VectorIterator) Prev() ConstBidIterator {
	if this.IsValid() {
		this.curIndex--
	}
	return this
}

func (this *VectorIterator) Clone() interface{} {
	return &VectorIterator{vec: this.vec, curIndex: this.curIndex}
}
