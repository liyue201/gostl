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
	a := &Array{data: make([]interface{}, other.Size(), other.Size())}
	for i := range other.data {
		a.data[i] = other.data[i]
	}
	return a
}

func (a *Array) Fill(val interface{}) {
	for i := range a.data {
		a.data[i] = val
	}
}

func (a *Array) Set(position int, val interface{}) error {
	if position < 0 || position >= len(a.data) {
		return ErrOutOffRange
	}
	a.data[position] = val
	return nil
}

func (a *Array) At(position int) interface{} {
	if position < 0 || position >= len(a.data) {
		return nil
	}
	return a.data[position]
}

func (a *Array) Front() interface{} {
	return a.At(0)
}

func (a *Array) Back() interface{} {
	return a.At(len(a.data) - 1)
}

func (a *Array) Size() int {
	return len(a.data)
}

func (a *Array) Empty() bool {
	return len(a.data) == 0
}

func (a *Array) SwapArray(other *Array) error {
	if a.Size() != other.Size() {
		return ErrArraySizeNotEqual
	}
	a.data, other.data = other.data, a.data
	return nil
}

func (a *Array) Data() []interface{} {
	return a.data
}

func (a *Array) Begin() *ArrayIterator {
	return a.First()
}

func (a *Array) End() *ArrayIterator {
	return a.IterAt(a.Size())
}

func (a *Array) First() *ArrayIterator {
	return a.IterAt(0)
}

func (a *Array) Last() *ArrayIterator {
	return a.IterAt(a.Size() - 1)
}

func (a *Array) IterAt(position int) *ArrayIterator {
	return &ArrayIterator{array: a, position: position}
}

func (a *Array) String() string {
	return fmt.Sprintf("%v", a.data)
}
