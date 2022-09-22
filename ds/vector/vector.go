package vector

import (
	"fmt"
	"github.com/liyue201/gostl/utils/iterator"
)

// Options holds the Vector's options
type Options struct {
	capacity int
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithCapacity is used to set the capacity of a Vector
func WithCapacity(capacity int) Option {
	return func(option *Options) {
		option.capacity = capacity
	}
}

// Vector is a linear data structure, the internal is a slice
type Vector[T any] struct {
	data []T
}

// New creates a new Vector
func New[T any](opts ...Option) *Vector[T] {
	option := Options{}
	for _, opt := range opts {
		opt(&option)
	}
	return &Vector[T]{
		data: make([]T, 0, option.capacity),
	}
}

// NewFromVector news a Vector from other Vector
func NewFromVector[T any](other *Vector[T]) *Vector[T] {
	v := &Vector[T]{data: make([]T, other.Size(), other.Capacity())}
	for i := range other.data {
		v.data[i] = other.data[i]
	}
	return v
}

// Size returns the size of the vector
func (v *Vector[T]) Size() int {
	return len(v.data)
}

// Capacity returns the capacity of the vector
func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

// Empty returns true if the vector is empty, otherwise returns false
func (v *Vector[T]) Empty() bool {
	if len(v.data) == 0 {
		return true
	}
	return false
}

// PushBack pushes val to the back of the vector
func (v *Vector[T]) PushBack(val T) {
	v.data = append(v.data, val)
}

// SetAt sets the value val to the vector at position pos
func (v *Vector[T]) SetAt(pos int, val T) {
	if pos < 0 || pos >= v.Size() {
		return
	}
	v.data[pos] = val
}

// InsertAt inserts the value val to the vector at position pos
func (v *Vector[T]) InsertAt(pos int, val T) {
	if pos < 0 || pos > v.Size() {
		return
	}
	v.data = append(v.data, val)
	for i := len(v.data) - 1; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[pos] = val
}

// EraseAt erases the value at position pos
func (v *Vector[T]) EraseAt(pos int) {
	v.EraseIndexRange(pos, pos+1)
}

// EraseIndexRange erases values at range[first, last)
func (v *Vector[T]) EraseIndexRange(first, last int) {
	if first > last {
		return
	}
	if first < 0 || last > v.Size() {
		return
	}

	left := v.data[:first]
	right := v.data[last:]
	v.data = append(left, right...)

}

// At returns the value at position pos, returns nil if pos is out off range .
func (v *Vector[T]) At(pos int) T {
	if pos < 0 || pos >= v.Size() {
		panic("ouf off range")
	}
	return v.data[pos]
}

//Front returns the first value in the vector, returns nil if the vector is empty.
func (v *Vector[T]) Front() T {
	return v.At(0)
}

//Back returns the last value in the vector, returns nil if the vector is empty.
func (v *Vector[T]) Back() T {
	return v.At(v.Size() - 1)
}

//PopBack returns the last value of the vector and erase it, returns nil if the vector is empty.
func (v *Vector[T]) PopBack() T {
	if v.Empty() {
		panic("out off range")
	}
	val := v.Back()
	v.data = v.data[:len(v.data)-1]
	return val
}

//Reserve makes a new space for the vector with passed capacity
func (v *Vector[T]) Reserve(capacity int) {
	if cap(v.data) >= capacity {
		return
	}
	data := make([]T, v.Size(), capacity)
	for i := 0; i < len(v.data); i++ {
		data[i] = v.data[i]
	}
	v.data = data
}

// ShrinkToFit shrinks the capacity of the vector to the fit size
func (v *Vector[T]) ShrinkToFit() {
	if len(v.data) == cap(v.data) {
		return
	}
	len := v.Size()
	data := make([]T, len, len)
	for i := 0; i < len; i++ {
		data[i] = v.data[i]
	}
	v.data = data
}

// Clear clears all data in the vector
func (v *Vector[T]) Clear() {
	v.data = v.data[:0]
}

// Data returns internal data of the vector
func (v *Vector[T]) Data() []T {
	return v.data
}

// Begin returns the first iterator of the vector
func (v *Vector[T]) Begin() *VectorIterator[T] {
	return v.First()
}

// End returns the end iterator of the vector
func (v *Vector[T]) End() *VectorIterator[T] {
	return v.IterAt(v.Size())
}

// First returns the first iterator of the vector
func (v *Vector[T]) First() *VectorIterator[T] {
	return v.IterAt(0)
}

// Last returns the last iterator of the vector
func (v *Vector[T]) Last() *VectorIterator[T] {
	return v.IterAt(v.Size() - 1)
}

// IterAt  returns the iterator at position of the vector
func (v *Vector[T]) IterAt(pos int) *VectorIterator[T] {
	return &VectorIterator[T]{vec: v, position: pos}
}

// Insert inserts a value val to the vector at the position of the iterator iter point to
func (v *Vector[T]) Insert(iter iterator.ConstIterator[T], val T) *VectorIterator[T] {
	index := iter.(*VectorIterator[T]).position
	v.InsertAt(index, val)
	return &VectorIterator[T]{vec: v, position: index}
}

// Erase erases the element of the iterator iter point to
func (v *Vector[T]) Erase(iter iterator.ConstIterator[T]) *VectorIterator[T] {
	index := iter.(*VectorIterator[T]).position
	v.EraseAt(index)
	return &VectorIterator[T]{vec: v, position: index}
}

// EraseRange erases all elements in the range[first, last)
func (v *Vector[T]) EraseRange(first, last iterator.ConstIterator[T]) *VectorIterator[T] {
	from := first.(*VectorIterator[T]).position
	to := last.(*VectorIterator[T]).position
	v.EraseIndexRange(from, to)
	return &VectorIterator[T]{vec: v, position: from}
}

// Resize resizes the size of the vector to the passed size
func (v *Vector[T]) Resize(size int) {
	if size >= v.Size() {
		return
	}
	v.data = v.data[:size]
}

// String returns a string representation of the vector
func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", v.data)
}
