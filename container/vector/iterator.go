package vector

import (
	. "github.com/liyue201/gostl/container"
)

type VectorIterator struct {
	vec      *Vector
	curIndex int
}

func (this *VectorIterator) Next() Iterator {
	index := this.curIndex + 1
	if index > this.vec.Size() {
		index = this.vec.Size()
	}
	return &VectorIterator{vec: this.vec, curIndex: index}
}

func (this *VectorIterator) Value() interface{} {
	val := this.vec.At(this.curIndex)
	return val
}

func (this *VectorIterator) Set(val interface{}) error {
	return this.vec.SetAt(this.curIndex, val)
}

func (this *VectorIterator) Equal(other Iterator) bool {
	otherItr, ok := other.(*VectorIterator)
	if !ok {
		return false
	}
	if this.vec == otherItr.vec && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}

type VectorReverseIterator struct {
	vec      *Vector
	curIndex int
}

func (this *VectorReverseIterator) Next() Iterator {
	index := this.curIndex - 1
	if index < -1 {
		index = -1
	}
	return &VectorReverseIterator{vec: this.vec, curIndex: index}
}

func (this *VectorReverseIterator) Set(val interface{}) error {
	return this.vec.SetAt(this.curIndex, val)
}

func (this *VectorReverseIterator) Value() interface{} {
	return this.vec.At(this.curIndex)
}

func (this *VectorReverseIterator) Equal(other Iterator) bool {
	otherItr, ok := other.(*VectorReverseIterator)
	if !ok {
		return false
	}
	if this.vec == otherItr.vec && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}
