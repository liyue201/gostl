package array

import (
	"errors"
	"fmt"
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

func (this *Array) Set(position int, val interface{}) error {
	if position < 0 || position >= len(this.data) {
		return ErrOutOffRange
	}
	this.data[position] = val
	return nil
}

func (this *Array) At(position int) interface{} {
	if position < 0 || position >= len(this.data) {
		return nil
	}
	return this.data[position]
}

func (this *Array) Front() interface{} {
	return this.At(0)
}

func (this *Array) Back() interface{} {
	return this.At(len(this.data) - 1)
}

func (this *Array) Size() int {
	return len(this.data)
}

func (this *Array) Empty() bool {
	return len(this.data) == 0
}

func (this *Array) SwapArray(other *Array) error {
	if this.Size() != other.Size() {
		return ErrArraySizeNotEqual
	}
	this.data, other.data = other.data, this.data
	return nil
}

func (this *Array) Data() []interface{} {
	return this.data
}

func (this *Array) Begin() *ArrayIterator {
	return this.First()
}

func (this *Array) End() *ArrayIterator {
	return this.IterAt(this.Size())
}

func (this *Array) First() *ArrayIterator {
	return this.IterAt(0)
}

func (this *Array) Last() *ArrayIterator {
	return this.IterAt(this.Size() - 1)
}

func (this *Array) IterAt(position int) *ArrayIterator {
	return &ArrayIterator{array: this, position: position}
}

func (this *Array) String() string {
	return fmt.Sprintf("%v", this.data)
}
