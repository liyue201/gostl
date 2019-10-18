package vector

import (
	"errors"
	"fmt"
	. "github.com/liyue201/gostl/container"
)

var ErrOutOffRange = errors.New("out off range")
var ErrEmpty = errors.New("vector is empty")
var ErrInvalidIterator = errors.New("invalid iterator")

type Less func(i, j int) bool

type Vector struct {
	data []interface{}
	less Less
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

//At returns the value at index, returns nil if index out off range .
func (this *Vector) At(index int) interface{} {
	if index < 0 || index >= this.Size() {
		return nil
	}
	return this.data[index]
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

func (this *Vector) Begin() BidIterator {
	return this.First()
}

func (this *Vector) First() BidIterator {
	return this.IterAt(0)
}

func (this *Vector) Last() BidIterator {
	return this.IterAt(this.Size() - 1)
}

func (this *Vector) IterAt(index int) BidIterator {
	return &VectorIterator{vec: this, curIndex: index}
}

func (this *Vector) Insert(iter ConstIterator, val interface{}) BidIterator {
	index := iter.(*VectorIterator).curIndex
	this.InsertAt(index, val)
	return &VectorIterator{vec: this, curIndex: index}
}

func (this *Vector) Erase(iter ConstIterator) BidIterator {
	index := iter.(*VectorIterator).curIndex
	this.EraseAt(index)
	return &VectorIterator{vec: this, curIndex: index}
}

func (this *Vector) EraseRange(first, last ConstIterator) BidIterator {
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

func (this *Vector) String() string {
	return fmt.Sprintf("%v", this.data)
}

/////////////////////////////////////////////////////////
//for sort.Sort API
func (this *Vector) SetLess(less Less) {
	this.less = less
}

//sort.Sort API
func (this *Vector) Len() int {
	return this.Size()
}

//sort.Sort API
func (this *Vector) Less(i, j int) bool {
	if this.less != nil {
		return this.less(i, j)
	}
	return false
}

//sort.Sort API
func (this *Vector) Swap(i, j int) {
	if i < 0 || j < 0 || i >= this.Size() || j >= this.Size() {
		return
	}
	this.data[i], this.data[j] = this.data[j], this.data[i]
}
