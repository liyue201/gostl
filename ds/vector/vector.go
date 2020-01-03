package vector

import (
	"errors"
	"fmt"
	. "github.com/liyue201/gostl/utils/iterator"
)

var ErrOutOffRange = errors.New("out off range")
var ErrEmpty = errors.New("vector is empty")
var ErrInvalidIterator = errors.New("invalid iterator")

type Option struct {
	capacity int
}

type Options func(option *Option)

func WithCapacity(capacity int) Options {
	return func(option *Option) {
		option.capacity = capacity
	}
}

type Vector struct {
	data []interface{}
}

func New(opts ...Options) *Vector {
	option := Option{}
	for _, opt := range opts {
		opt(&option)
	}
	return &Vector{
		data: make([]interface{}, 0, option.capacity),
	}
}

func NewFromVector(other *Vector) *Vector {
	v := &Vector{data: make([]interface{}, other.Size(), other.Capacity())}
	for i := range other.data {
		v.data[i] = other.data[i]
	}
	return v
}

func (v *Vector) Size() int {
	return len(v.data)
}

func (v *Vector) Capacity() int {
	return cap(v.data)
}

func (v *Vector) Empty() bool {
	if len(v.data) == 0 {
		return true
	}
	return false
}

func (v *Vector) PushBack(val interface{}) {
	v.data = append(v.data, val)
}

func (v *Vector) SetAt(position int, val interface{}) error {
	if position < 0 || position >= v.Size() {
		return ErrOutOffRange
	}
	v.data[position] = val
	return nil
}

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

func (v *Vector) EraseAt(position int) error {
	return v.EraseIndexRange(position, position+1)
}

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

func (v *Vector) Clear() {
	v.data = v.data[:0]
}

func (v *Vector) Data() []interface{} {
	return v.data
}

func (v *Vector) Begin() *VectorIterator {
	return v.First()
}

func (v *Vector) End() *VectorIterator {
	return v.IterAt(v.Size())
}

func (v *Vector) First() *VectorIterator {
	return v.IterAt(0)
}

func (v *Vector) Last() *VectorIterator {
	return v.IterAt(v.Size() - 1)
}

func (v *Vector) IterAt(position int) *VectorIterator {
	return &VectorIterator{vec: v, position: position}
}

func (v *Vector) Insert(iter ConstIterator, val interface{}) *VectorIterator {
	index := iter.(*VectorIterator).position
	v.InsertAt(index, val)
	return &VectorIterator{vec: v, position: index}
}

func (v *Vector) Erase(iter ConstIterator) *VectorIterator {
	index := iter.(*VectorIterator).position
	v.EraseAt(index)
	return &VectorIterator{vec: v, position: index}
}

func (v *Vector) EraseRange(first, last ConstIterator) *VectorIterator {
	from := first.(*VectorIterator).position
	to := last.(*VectorIterator).position
	v.EraseIndexRange(from, to)
	return &VectorIterator{vec: v, position: from}
}

func (v *Vector) Resize(size int) {
	if size >= v.Size() {
		return
	}
	v.data = v.data[:size]
}

func (v *Vector) String() string {
	return fmt.Sprintf("%v", v.data)
}
