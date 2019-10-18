package array

import (
	"errors"
	"fmt"
	"github.com/liyue201/gostl/uitls/comparator"
	. "github.com/liyue201/gostl/uitls/iterator"
)

var ErrArraySizeNotEqual = errors.New("array size are not equal")
var ErrOutOffRange = errors.New("out off range")

type Array struct {
	data    []interface{}
	cmpFunc comparator.Comparator
}

func New(size int) *Array {
	return &Array{data: make([]interface{}, size, size)}
}

func NewFromArray(other *Array) *Array {
	this := &Array{data: make([]interface{}, other.Size(), other.Size())}
	for i := range other.data {
		this.data[i] = other.data[i]
	}
	return this
}

func (this *Array) Fill(val interface{}) {
	for i := range this.data {
		this.data[i] = val
	}
}

func (this *Array) Set(index int, val interface{}) error {
	if index < 0 || index >= len(this.data) {
		return ErrOutOffRange
	}
	this.data[index] = val
	return nil
}

func (this *Array) At(index int) interface{} {
	if index < 0 || index >= len(this.data) {
		return nil
	}
	return this.data[index]
}

func (this *Array) Front() interface{} {
	return this.At(0)
}

func (this *Array) Back() interface{} {
	return this.At(len(this.data) - 1)
}

func (this *Array) Size() int {
	return len(this.data)
}

func (this *Array) Empty() bool {
	if len(this.data) > 0 {
		return false
	}
	return true
}

func (this *Array) SwapArray(other *Array) error {
	if this.Size() != other.Size() {
		return ErrArraySizeNotEqual
	}
	this.data, other.data = other.data, this.data
	return nil
}

func (this *Array) Data() []interface{} {
	return this.data
}

func (this *Array) Begin() BidIterator {
	return this.First()
}

func (this *Array) First() BidIterator {
	return this.IterAt(0)
}

func (this *Array) Last() BidIterator {
	return this.IterAt(this.Size() - 1)
}

func (this *Array) IterAt(index int) BidIterator {
	return &ArrayIterator{array: this, curIndex: index}
}

func (this *Array) String() string {
	return fmt.Sprintf("%v", this.data)
}

/////////////////////////////////////////////////////////
//for sort.Sort API
func (this *Array) SetComparator(cmp comparator.Comparator) {
	this.cmpFunc = cmp
}

//sort.Sort API
func (this *Array) Len() int {
	return this.Size()
}

//sort.Sort API
func (this *Array) Less(i, j int) bool {
	if this.cmpFunc != nil {
		if this.cmpFunc(this.At(i), this.At(j)) < 0 {
			return true
		}
	}
	return false
}

//sort.Sort API
func (this *Array) Swap(i, j int) {
	if i < 0 || j < 0 || i >= this.Size() || j >= this.Size() {
		return
	}
	this.data[i], this.data[j] = this.data[j], this.data[i]
}
