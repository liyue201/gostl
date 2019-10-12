package gostl

import "errors"

var ErrArraySizeNotEqual = errors.New("array size are not equal")
var ErrOutOffRange = errors.New("out off range")

type Array struct {
	data []interface{}
}

func NewArray(size int) *Array {
	return &Array{data: make([]interface{}, size, size)}
}

func (this *Array) Fill(val interface{}) {
	for i := range this.data {
		this.data[i] = val
	}
}

func (this *Array) Assign(index int, val interface{}) error {
	if index < 0 || index > len(this.data) {
		return ErrOutOffRange
	}
	this.data[index] = val
	return nil
}

func (this *Array) At(index int) (interface{}, error) {
	if index < 0 || index > len(this.data) {
		return 0, ErrOutOffRange
	}
	return this.data[index], nil
}

func (this *Array) Front(index int) (interface{}, error) {
	return this.At(0)
}

func (this *Array) Back(index int) (interface{}, error) {
	return this.At(len(this.data) - 1)
}

func (this *Array) Size() int {
	return len(this.data)
}

func (this *Array) MaxSize() int {
	return len(this.data)
}

func (this *Array) Empty() bool {
	if len(this.data) > 0 {
		return false
	}
	return true
}

func (this *Array) Swap(other *Array) error {
	if this.Size() != other.Size() {
		return ErrArraySizeNotEqual
	}
	for i := range this.data {
		this.data[i], other.data[i] = other.data[i], this.data[i]
	}
	return nil
}

func (this *Array) Data() []interface{} {
	return this.data
}

func (this *Array) Begin() Iterator {
	return &ArrayIterator{array: this, curIndex: 0}
}

func (this *Array) End() Iterator {
	return &ArrayIterator{array: this, curIndex: len(this.data) - 1}
}

func (this *Array) RBegin() ReverseIterator {
	return &ArrayReverseIterator{array: this, curIndex: len(this.data) - 1}
}

func (this *Array) REnd() ReverseIterator {
	return &ArrayReverseIterator{array: this, curIndex: 0}
}

type ArrayIterator struct {
	array    *Array
	curIndex int
}

func (this *ArrayIterator) Next() Iterator {
	this.curIndex++
	return this
}

func (this *ArrayIterator) Data() interface{} {
	data, _ := this.array.At(this.curIndex)
	return data
}

func (this *ArrayIterator) Assign(val interface{}) error {
	return this.array.Assign(this.curIndex, val)
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

func (this *ArrayReverseIterator) Next() ReverseIterator {
	this.curIndex--
	return this
}

func (this *ArrayReverseIterator) Assign(val interface{}) error {
	return this.array.Assign(this.curIndex, val)
}

func (this *ArrayReverseIterator) Data() interface{} {
	data, _ := this.array.At(this.curIndex)
	return data
}

func (this *ArrayReverseIterator) Equal(other ReverseIterator) bool {
	otherItr, ok := other.(*ArrayReverseIterator)
	if !ok {
		return false
	}
	if this.array == otherItr.array && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}
