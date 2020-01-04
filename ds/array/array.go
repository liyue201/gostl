package array

import (
	"errors"
	"fmt"
)

//ErrArraySizeNotEqual defines "array size are not equal"
var ErrArraySizeNotEqual = errors.New("array size are not equal")

//ErrArraySizeNotEqual defines "out off range error"
var ErrOutOffRange = errors.New("out off range")

// Array internal is a slice
type Array struct {
	data []interface{}
}

// New news an array with size passed
func New(size int) *Array {
	return &Array{data: make([]interface{}, size, size)}
}

// NewFromArray news an array from other array
func NewFromArray(other *Array) *Array {
	a := &Array{data: make([]interface{}, other.Size(), other.Size())}
	for i := range other.data {
		a.data[i] = other.data[i]
	}
	return a
}

// Fill fills a with value val
func (a *Array) Fill(val interface{}) {
	for i := range a.data {
		a.data[i] = val
	}
}

// Set sets a with value val at position
func (a *Array) Set(position int, val interface{}) error {
	if position < 0 || position >= len(a.data) {
		return ErrOutOffRange
	}
	a.data[position] = val
	return nil
}

// At returns the value at position
func (a *Array) At(position int) interface{} {
	if position < 0 || position >= len(a.data) {
		return nil
	}
	return a.data[position]
}

// Front returns the first value of a
func (a *Array) Front() interface{} {
	return a.At(0)
}

// Back returns the last value of a
func (a *Array) Back() interface{} {
	return a.At(len(a.data) - 1)
}

// Size returns the size  of a
func (a *Array) Size() int {
	return len(a.data)
}

// Empty returns whether a is empty
func (a *Array) Empty() bool {
	return len(a.data) == 0
}

// SwapArray swaps the data of a with other
func (a *Array) SwapArray(other *Array) error {
	if a.Size() != other.Size() {
		return ErrArraySizeNotEqual
	}
	a.data, other.data = other.data, a.data
	return nil
}

// Data returns the internal data of a
func (a *Array) Data() []interface{} {
	return a.data
}

// Begin returns the first iterator of a
func (a *Array) Begin() *ArrayIterator {
	return a.First()
}

// End returns an end iterator of a
func (a *Array) End() *ArrayIterator {
	return a.IterAt(a.Size())
}

// First returns the first iterator of a
func (a *Array) First() *ArrayIterator {
	return a.IterAt(0)
}

// Last returns the last iterator of a
func (a *Array) Last() *ArrayIterator {
	return a.IterAt(a.Size() - 1)
}

// IterAt returns an iterator of a at position
func (a *Array) IterAt(position int) *ArrayIterator {
	return &ArrayIterator{array: a, position: position}
}

// String returns a with string format
func (a *Array) String() string {
	return fmt.Sprintf("%v", a.data)
}
