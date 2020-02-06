package array

import (
	"fmt"
)

// Array is a fixed size slice
type Array struct {
	values []interface{}
}

// New creates a new array with passed size
func New(size int) *Array {
	return &Array{values: make([]interface{}, size, size)}
}

// NewFromArray creates a new array from another array, and copy its values
func NewFromArray(other *Array) *Array {
	a := &Array{values: make([]interface{}, other.Size(), other.Size())}
	for i := range other.values {
		a.values[i] = other.values[i]
	}
	return a
}

// Fill fills Array a with value val
func (a *Array) Fill(val interface{}) {
	for i := range a.values {
		a.values[i] = val
	}
}

// Set sets value val to the position pos of the array
func (a *Array) Set(pos int, val interface{}) {
	if pos < 0 || pos >= len(a.values) {
		return
	}
	a.values[pos] = val
}

// At returns the value at position pos in the array
func (a *Array) At(pos int) interface{} {
	if pos < 0 || pos >= len(a.values) {
		return nil
	}
	return a.values[pos]
}

// Front returns the first value in the array
func (a *Array) Front() interface{} {
	return a.At(0)
}

// Back returns the last value in the array
func (a *Array) Back() interface{} {
	return a.At(len(a.values) - 1)
}

// Size returns number of elements within the array
func (a *Array) Size() int {
	return len(a.values)
}

// Empty returns whether the array is empty or not
func (a *Array) Empty() bool {
	return len(a.values) == 0
}

// SwapArray swaps the values of two arrays
func (a *Array) SwapArray(other *Array) {
	if a.Size() != other.Size() {
		return
	}
	a.values, other.values = other.values, a.values
}

// Data returns the internal values of the array
func (a *Array) Data() []interface{} {
	return a.values
}

// Begin returns an iterator of the array with the first position
func (a *Array) Begin() *ArrayIterator {
	return a.First()
}

// End returns an iterator of the array with the position a.Size()
func (a *Array) End() *ArrayIterator {
	return a.IterAt(a.Size())
}

// First returns an iterator of the array with the first position
func (a *Array) First() *ArrayIterator {
	return a.IterAt(0)
}

// Last returns an iterator of the array with the last position
func (a *Array) Last() *ArrayIterator {
	return a.IterAt(a.Size() - 1)
}

// IterAt returns an iterator of the array with position pos
func (a *Array) IterAt(pos int) *ArrayIterator {
	return &ArrayIterator{array: a, position: pos}
}

// String returns a string representation of the array
func (a *Array) String() string {
	return fmt.Sprintf("%v", a.values)
}
