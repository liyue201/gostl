package array

import (
	"fmt"
)

// Array is a fixed size slice
type Array[T any] struct {
	values []T
}

// New creates a new array with passed size
func New[T any](size int) *Array[T] {
	return &Array[T]{values: make([]T, size, size)}
}

// NewFromArray creates a new array from another array, and copy its values
func NewFromArray[T any](other *Array[T]) *Array[T] {
	a := &Array[T]{values: make([]T, other.Size(), other.Size())}
	for i := range other.values {
		a.values[i] = other.values[i]
	}
	return a
}

// Fill fills Array a with value val
func (a *Array[T]) Fill(val T) {
	for i := range a.values {
		a.values[i] = val
	}
}

// Set sets value val to the position pos of the array
func (a *Array[T]) Set(pos int, val T) {
	if pos < 0 || pos >= len(a.values) {
		return
	}
	a.values[pos] = val
}

// At returns the value at position pos in the array
func (a *Array[T]) At(pos int) T {
	if pos < 0 || pos >= len(a.values) {
		panic("index out off range")
	}
	return a.values[pos]
}

// Front returns the first value in the array
func (a *Array[T]) Front() T {
	return a.At(0)
}

// Back returns the last value in the array
func (a *Array[T]) Back() T {
	return a.At(len(a.values) - 1)
}

// Size returns number of elements within the array
func (a *Array[T]) Size() int {
	return len(a.values)
}

// Empty returns whether the array is empty or not
func (a *Array[T]) Empty() bool {
	return len(a.values) == 0
}

// SwapArray swaps the values of two arrays
func (a *Array[T]) SwapArray(other *Array[T]) {
	if a.Size() != other.Size() {
		return
	}
	a.values, other.values = other.values, a.values
}

// Data returns the internal values of the array
func (a *Array[T]) Data() []T {
	return a.values
}

// Begin returns an iterator of the array with the first position
func (a *Array[T]) Begin() *ArrayIterator[T] {
	return a.First()
}

// End returns an iterator of the array with the position a.Size()
func (a *Array[T]) End() *ArrayIterator[T] {
	return a.IterAt(a.Size())
}

// First returns an iterator of the array with the first position
func (a *Array[T]) First() *ArrayIterator[T] {
	return a.IterAt(0)
}

// Last returns an iterator of the array with the last position
func (a *Array[T]) Last() *ArrayIterator[T] {
	return a.IterAt(a.Size() - 1)
}

// IterAt returns an iterator of the array with position pos
func (a *Array[T]) IterAt(pos int) *ArrayIterator[T] {
	return &ArrayIterator[T]{array: a, position: pos}
}

// String returns a string representation of the array
func (a *Array[T]) String() string {
	return fmt.Sprintf("%v", a.values)
}
