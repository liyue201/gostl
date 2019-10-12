package vector

import (
	"errors"
	"fmt"
	. "github.com/liyue201/gostl/container"
)

var ErrOutOffRange = errors.New("out off range")
var ErrEmpty = errors.New("vector is empty")
var ErrInvalidIterator = errors.New("invalid iterator")

type Vector struct {
	data []interface{}
}

func New(capacity int) *Vector {
	return &Vector{data: make([]interface{}, 0, capacity)}
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

func (this *Vector) SetAt(index int, val interface{}) error {
	if index < 0 || index >= this.Size() {
		return ErrOutOffRange
	}
	this.data[index] = val
	return nil
}

func (this *Vector) InsertAt(index int, val interface{}) error {
	if index < 0 || index > this.Size() {
		return ErrOutOffRange
	}
	this.data = append(this.data, val)
	for i := len(this.data) - 1; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = val
	return nil
}

func (this *Vector) EraseAt(index int) error {
	return this.EraseIndexRange(index, index+1)
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

func (this *Vector) At(index int) (interface{}, error) {
	if index < 0 || index > this.Size() {
		return nil, ErrOutOffRange
	}
	return this.data[index], nil
}

func (this *Vector) Front() (interface{}, error) {
	return this.At(0)
}

func (this *Vector) Back() (interface{}, error) {
	return this.At(this.Size() - 1)
}

func (this *Vector) PopBack() (interface{}, error) {
	if this.Empty() {
		return nil, ErrEmpty
	}
	val, err := this.Back()
	this.data = this.data[:len(this.data)-1]
	return val, err
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

func (this *Vector) Begin() Iterator {
	return &VectorIterator{vec: this, curIndex: 0}
}

func (this *Vector) End() Iterator {
	return &VectorIterator{vec: this, curIndex: this.Size()}
}

func (this *Vector) RBegin() ReverseIterator {
	return &VectorReverseIterator{vec: this, curIndex: this.Size() - 1}
}

func (this *Vector) REnd() ReverseIterator {
	return &VectorReverseIterator{vec: this, curIndex: -1}
}

func (this *Vector) Insert(iter Iterator, val interface{}) Iterator {
	this.InsertAt(iter.(*VectorIterator).curIndex, val)
	return iter
}

func (this *Vector) Erase(iter Iterator) Iterator {
	this.EraseAt(iter.(*VectorIterator).curIndex)
	return iter
}

func (this *Vector) EraseRange(first, last Iterator) Iterator {
	from := first.(*VectorIterator).curIndex
	to := last.(*VectorIterator).curIndex
	this.EraseIndexRange(from, to)
	return &VectorIterator{vec: this, curIndex: from}
}

func (this *Vector) Resize(size int) {
	if size >= this.Size() {
		return
	}
	this.data = this.data[:size]
}

func (this *Vector) Swap(other *Vector) {
	this.data, other.data = other.data, this.data
}

func (this *Vector) String() string {
	return fmt.Sprintf("%v", this.data)
}

type VectorIterator struct {
	vec      *Vector
	curIndex int
}

func (this *VectorIterator) Next() Iterator {
	return &VectorIterator{vec: this.vec, curIndex: this.curIndex + 1}
}

func (this *VectorIterator) Value() interface{} {
	val, _ := this.vec.At(this.curIndex)
	return val
}

func (this *VectorIterator) Set(val interface{}) error {
	return this.vec.SetAt(this.curIndex, val)
}

func (this *VectorIterator) Equal(other Iterator) bool {
	otherItr, ok := other.(*VectorIterator)
	if !ok {
		return false
	}
	if this.vec == otherItr.vec && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}

type VectorReverseIterator struct {
	vec      *Vector
	curIndex int
}

func (this *VectorReverseIterator) Next() ReverseIterator {
	return &VectorReverseIterator{vec: this.vec, curIndex: this.curIndex - 1}
	return this
}

func (this *VectorReverseIterator) Set(val interface{}) error {
	return this.vec.SetAt(this.curIndex, val)
}

func (this *VectorReverseIterator) Value() interface{} {
	val, _ := this.vec.At(this.curIndex)
	return val
}

func (this *VectorReverseIterator) Equal(other ReverseIterator) bool {
	otherItr, ok := other.(*VectorReverseIterator)
	if !ok {
		return false
	}
	if this.vec == otherItr.vec && otherItr.curIndex == this.curIndex {
		return true
	}
	return false
}
