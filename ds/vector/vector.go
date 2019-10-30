package vector

import (
	"errors"
	"fmt"
	. "github.com/liyue201/gostl/iterator"
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
	this := &Vector{data: make([]interface{}, other.Size(), other.Capacity())}
	for i := range other.data {
		this.data[i] = other.data[i]
	}
	return this
}

func (this *Vector) Size() int {
	return len(this.data)
}

func (this *Vector) Capacity() int {
	return cap(this.data)
}

func (this *Vector) Empty() bool {
	if len(this.data) == 0 {
		return true
	}
	return false
}

func (this *Vector) PushBack(val interface{}) {
	this.data = append(this.data, val)
}

func (this *Vector) SetAt(position int, val interface{}) error {
	if position < 0 || position >= this.Size() {
		return ErrOutOffRange
	}
	this.data[position] = val
	return nil
}

func (this *Vector) InsertAt(position int, val interface{}) error {
	if position < 0 || position > this.Size() {
		return ErrOutOffRange
	}
	this.data = append(this.data, val)
	for i := len(this.data) - 1; i > position; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[position] = val
	return nil
}

func (this *Vector) EraseAt(position int) error {
	return this.EraseIndexRange(position, position+1)
}

func (this *Vector) EraseIndexRange(first, last int) error {
	if first > last {
		return nil
	}
	if first < 0 || last > this.Size() {
		return ErrOutOffRange
	}

	left := this.data[:first]
	right := this.data[last:]
	this.data = append(left, right...)
	return nil
}

//At returns the value at position, returns nil if position out off range .
func (this *Vector) At(position int) interface{} {
	if position < 0 || position >= this.Size() {
		return nil
	}
	return this.data[position]
}

//At returns the first value of the vector, returns nil if the vector is empty.
func (this *Vector) Front() interface{} {
	return this.At(0)
}

//At returns the last value of the vector, returns nil if the vector is empty.
func (this *Vector) Back() interface{} {
	return this.At(this.Size() - 1)
}

//At returns the last value of the vector and erase it, returns nil if the vector is empty.
func (this *Vector) PopBack() interface{} {
	if this.Empty() {
		return nil
	}
	val := this.Back()
	this.data = this.data[:len(this.data)-1]
	return val
}

func (this *Vector) Reserve(capacity int) {
	if cap(this.data) >= capacity {
		return
	}
	data := make([]interface{}, this.Size(), capacity)
	for i := 0; i < len(this.data); i++ {
		data[i] = this.data[i]
	}
	this.data = data
}

func (this *Vector) ShrinkToFit() {
	if len(this.data) == cap(this.data) {
		return
	}
	len := this.Size()
	data := make([]interface{}, len, len)
	for i := 0; i < len; i++ {
		data[i] = this.data[i]
	}
	this.data = data
}

func (this *Vector) Clear() {
	this.data = this.data[:0]
}

func (this *Vector) Data() [] interface{} {
	return this.data
}

func (this *Vector) Begin() *VectorIterator {
	return this.First()
}

func (this *Vector) End() *VectorIterator {
	return this.IterAt(this.Size())
}

func (this *Vector) First() *VectorIterator {
	return this.IterAt(0)
}

func (this *Vector) Last() *VectorIterator {
	return this.IterAt(this.Size() - 1)
}

func (this *Vector) IterAt(position int) *VectorIterator {
	return &VectorIterator{vec: this, position: position}
}

func (this *Vector) Insert(iter ConstIterator, val interface{}) *VectorIterator {
	index := iter.(*VectorIterator).position
	this.InsertAt(index, val)
	return &VectorIterator{vec: this, position: index}
}

func (this *Vector) Erase(iter ConstIterator) *VectorIterator {
	index := iter.(*VectorIterator).position
	this.EraseAt(index)
	return &VectorIterator{vec: this, position: index}
}

func (this *Vector) EraseRange(first, last ConstIterator) *VectorIterator {
	from := first.(*VectorIterator).position
	to := last.(*VectorIterator).position
	this.EraseIndexRange(from, to)
	return &VectorIterator{vec: this, position: from}
}

func (this *Vector) Resize(size int) {
	if size >= this.Size() {
		return
	}
	this.data = this.data[:size]
}

func (this *Vector) String() string {
	return fmt.Sprintf("%v", this.data)
}
