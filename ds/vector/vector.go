package vector

import (
	"errors"
	"fmt"
	. "github.com/liyue201/gostl/utils/iterator"
)

// Define some errors
var (
	ErrOutOffRange = errors.New("out off range")
	ErrEmpty = errors.New("vector is empty")
 	ErrInvalidIterator = errors.New("invalid iterator")
)

// Options holds Vector's options
type Options struct {
	capacity int
}

// Option is a function used to set Options
type Option func(option *Options)

// WithCapacity sets the capacity of Vector
func WithCapacity(capacity int) Option {
	return func(option *Options) {
		option.capacity = capacity
	}
}

// Vector is a linear data structure, internal is a slice
type Vector struct {
	data []interface{}
}

// New news a Vector
func New(opts ...Option) *Vector {
	option := Options{}
	for _, opt := range opts {
		opt(&option)
	}
	return &Vector{
		data: make([]interface{}, 0, option.capacity),
	}
}

// NewFromVector news a Vector from other Vector
func NewFromVector(other *Vector) *Vector {
	v := &Vector{data: make([]interface{}, other.Size(), other.Capacity())}
	for i := range other.data {
		v.data[i] = other.data[i]
	}
	return v
}

// Size returns the size of v
func (v *Vector) Size() int {
	return len(v.data)
}

// Capacity returns the capacity of v
func (v *Vector) Capacity() int {
	return cap(v.data)
}

// Empty returns whether v is empty or not
func (v *Vector) Empty() bool {
	if len(v.data) == 0 {
		return true
	}
	return false
}

// PushBack pushes val to the back of v
func (v *Vector) PushBack(val interface{}) {
	v.data = append(v.data, val)
}

// SetAt sets val to v at position
func (v *Vector) SetAt(position int, val interface{}) error {
	if position < 0 || position >= v.Size() {
		return ErrOutOffRange
	}
	v.data[position] = val
	return nil
}

// InsertAt inserts val to v at position
func (v *Vector) InsertAt(position int, val interface{}) error {
	if position < 0 || position > v.Size() {
		return ErrOutOffRange
	}
	v.data = append(v.data, val)
	for i := len(v.data) - 1; i > position; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[position] = val
	return nil
}

// EraseAt erases the value at position
func (v *Vector) EraseAt(position int) error {
	return v.EraseIndexRange(position, position+1)
}

// EraseAt erases values at range[first, last)
func (v *Vector) EraseIndexRange(first, last int) error {
	if first > last {
		return nil
	}
	if first < 0 || last > v.Size() {
		return ErrOutOffRange
	}

	left := v.data[:first]
	right := v.data[last:]
	v.data = append(left, right...)
	return nil
}

//At returns the value at position, returns nil if position out off range .
func (v *Vector) At(position int) interface{} {
	if position < 0 || position >= v.Size() {
		return nil
	}
	return v.data[position]
}

//Front returns the first value of the vector, returns nil if the vector is empty.
func (v *Vector) Front() interface{} {
	return v.At(0)
}

//Back returns the last value of the vector, returns nil if the vector is empty.
func (v *Vector) Back() interface{} {
	return v.At(v.Size() - 1)
}

//PopBack returns the last value of the vector and erase it, returns nil if the vector is empty.
func (v *Vector) PopBack() interface{} {
	if v.Empty() {
		return nil
	}
	val := v.Back()
	v.data = v.data[:len(v.data)-1]
	return val
}

//Reserve make a new space with total capacity.
func (v *Vector) Reserve(capacity int) {
	if cap(v.data) >= capacity {
		return
	}
	data := make([]interface{}, v.Size(), capacity)
	for i := 0; i < len(v.data); i++ {
		data[i] = v.data[i]
	}
	v.data = data
}

//ShrinkToFit shrinks the capacity of v to fit size
func (v *Vector) ShrinkToFit() {
	if len(v.data) == cap(v.data) {
		return
	}
	len := v.Size()
	data := make([]interface{}, len, len)
	for i := 0; i < len; i++ {
		data[i] = v.data[i]
	}
	v.data = data
}

// Clear clear all data of v
func (v *Vector) Clear() {
	v.data = v.data[:0]
}

// Data returns internal data of v
func (v *Vector) Data() []interface{} {
	return v.data
}

// Begin returns the first iterator of v
func (v *Vector) Begin() *VectorIterator {
	return v.First()
}

// End returns the end iterator of v
func (v *Vector) End() *VectorIterator {
	return v.IterAt(v.Size())
}

// First returns the first iterator of v
func (v *Vector) First() *VectorIterator {
	return v.IterAt(0)
}

// Last returns the last iterator of v
func (v *Vector) Last() *VectorIterator {
	return v.IterAt(v.Size() - 1)
}

// Last returns the iterator at position of v
func (v *Vector) IterAt(position int) *VectorIterator {
	return &VectorIterator{vec: v, position: position}
}

// Insert inserts val at the position of iter
func (v *Vector) Insert(iter ConstIterator, val interface{}) *VectorIterator {
	index := iter.(*VectorIterator).position
	v.InsertAt(index, val)
	return &VectorIterator{vec: v, position: index}
}

// Erase erases val at the position of iter
func (v *Vector) Erase(iter ConstIterator) *VectorIterator {
	index := iter.(*VectorIterator).position
	v.EraseAt(index)
	return &VectorIterator{vec: v, position: index}
}

// EraseRange erases all val in the range[first, last)
func (v *Vector) EraseRange(first, last ConstIterator) *VectorIterator {
	from := first.(*VectorIterator).position
	to := last.(*VectorIterator).position
	v.EraseIndexRange(from, to)
	return &VectorIterator{vec: v, position: from}
}

// Resize resize the size of v to size  passed
func (v *Vector) Resize(size int) {
	if size >= v.Size() {
		return
	}
	v.data = v.data[:size]
}

// String returns v in string format
func (v *Vector) String() string {
	return fmt.Sprintf("%v", v.data)
}
