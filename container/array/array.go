package array

import (
	"errors"
	"fmt"
	. "github.com/liyue201/gostl/container"
)

var ErrArraySizeNotEqual = errors.New("array size are not equal")
var ErrOutOffRange = errors.New("out off range")

type Array struct {
	data []interface{}
}

func New(size int) *Array {
	return &Array{data: make([]interface{}, size, size)}
}

func NewFromArray(other *Array) *Array {
	this := &Array{data: make([]interface{}, other.Size(), other.Size())}
	for i := range other.data {
		this.data[i] = other.data[i]
	}
	return this
}

func (this *Array) Fill(val interface{}) {
	for i := range this.data {
		this.data[i] = val
	}
}

func (this *Array) Set(index int, val interface{}) error {
	if index < 0 || index > len(this.data) {
		return ErrOutOffRange
	}
	this.data[index] = val
	return nil
}

func (this *Array) At(index int) interface{} {
	if index < 0 || index > len(this.data) {
		return nil
	}
	return this.data[index]
}

func (this *Array) Front(index int) interface{} {
	return this.At(0)
}

func (this *Array) Back(index int) interface{} {
	return this.At(len(this.data) - 1)
}

func (this *Array) Size() int {
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
	this.data, other.data = other.data, this.data
	return nil
}

func (this *Array) Data() []interface{} {
	return this.data
}

func (this *Array) Begin() Iterator {
	return &ArrayIterator{array: this, curIndex: 0}
}

func (this *Array) End() Iterator {
	return &ArrayIterator{array: this, curIndex: len(this.data)}
}

func (this *Array) RBegin() ReverseIterator {
	return &ArrayReverseIterator{array: this, curIndex: len(this.data) - 1}
}

func (this *Array) REnd() ReverseIterator {
	return &ArrayReverseIterator{array: this, curIndex: -1}
}

func (this *Array) String() string {
	return fmt.Sprintf("%v", this.data)
}

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

func (this *ArrayReverseIterator) Next() ReverseIterator {
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
