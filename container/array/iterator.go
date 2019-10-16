package array

import (
	. "github.com/liyue201/gostl/container"
)

type ArrayIterator struct {
	array    *Array
	curIndex int
}

func (this *ArrayIterator) Next() Iterator {
	index := this.curIndex + 1
	if index > this.array.Size() {
		index = this.array.Size()
	}
	return &ArrayIterator{array: this.array, curIndex: index}
}

func (this *ArrayIterator) Value() interface{} {
	return this.array.At(this.curIndex)
}

func (this *ArrayIterator) Set(val interface{}) error {
	return this.array.Set(this.curIndex, val)
}

func (this *ArrayIterator) Equal(other Iterator) bool {
	otherItr, ok := other.(*ArrayIterator)
	if !ok {
		return false
	}
	if this.array == otherItr.array && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}

type ArrayReverseIterator struct {
	array    *Array
	curIndex int
}

func (this *ArrayReverseIterator) Next() Iterator {
	index := this.curIndex - 1
	if index < -1 {
		index = -1
	}
	return &ArrayReverseIterator{array: this.array, curIndex: index}
}

func (this *ArrayReverseIterator) Set(val interface{}) error {
	return this.array.Set(this.curIndex, val)
}

func (this *ArrayReverseIterator) Value() interface{} {
	return this.array.At(this.curIndex)
}

func (this *ArrayReverseIterator) Equal(other Iterator) bool {
	otherItr, ok := other.(*ArrayReverseIterator)
	if !ok {
		return false
	}
	if this.array == otherItr.array && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}
