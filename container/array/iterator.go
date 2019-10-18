package array

import (
	. "github.com/liyue201/gostl/container"
)

type ArrayIterator struct {
	array    *Array
	curIndex int
}

func (this *ArrayIterator) IsValid() bool {
	if this.curIndex >= 0 && this.curIndex < this.array.Size() {
		return true
	}
	return false
}

func (this *ArrayIterator) Value() interface{} {
	return this.array.At(this.curIndex)
}

func (this *ArrayIterator) SetValue(val interface{}) error {
	return this.array.Set(this.curIndex, val)
}

func (this *ArrayIterator) Next() ConstIterator {
	if this.IsValid() {
		this.curIndex++
	}
	return this
}

func (this *ArrayIterator) Prev() ConstBidIterator {
	if this.IsValid() {
		this.curIndex--
	}
	return this
}

func (this *ArrayIterator) Clone() interface{} {
	return &ArrayIterator{array: this.array, curIndex: this.curIndex}
}

