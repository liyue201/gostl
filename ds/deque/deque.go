package deque

import (
	"errors"
	"fmt"
)

var ErrOutOffRange = errors.New("out off range")

type Deque struct {
	data  []interface{}
	begin int
	end   int
	size  int
}

type Option struct {
	capacity int
}

type Options func(option *Option)

func WithCapacity(capacity int) Options {
	return func(option *Option) {
		option.capacity = capacity
	}
}

func New(opts ...Options) *Deque {
	option := Option{}
	for _, opt := range opts {
		opt(&option)
	}
	return &Deque{
		data: make([]interface{}, option.capacity),
	}
}

//Size returns the size of deque
func (this *Deque) Size() int {
	return this.size
}

//Capacity returns the capacity of deque
func (this *Deque) Capacity() int {
	return len(this.data)
}

//Capacity returns true if the Deque is empty,otherwise returns false.
func (this *Deque) Empty() bool {
	if this.Size() == 0 {
		return true
	}
	return false
}

//expandIfNeeded expand the Deque if full.
func (this *Deque) expandIfNeeded() {
	if this.size == this.Capacity() {
		newCapacity := this.size * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		data := make([]interface{}, newCapacity, newCapacity)
		for i := 0; i < this.size; i++ {
			data[i] = this.data[(this.begin+i)%this.Capacity()]
		}
		this.data = data
		this.begin = 0
		this.end = this.size
	}
}

// shrinkIfNeeded shrink the Deque if is has too many unused space .
func (this *Deque) shrinkIfNeeded() {
	if int(float64(this.size*2)*1.2) < this.Capacity() {
		newCapacity := this.Capacity() / 2
		data := make([]interface{}, newCapacity, newCapacity)
		for i := 0; i < this.size; i++ {
			data[i] = this.data[(this.begin+i)%this.Capacity()]
		}
		this.data = data
		this.begin = 0
		this.end = this.size
	}
}

func (this *Deque) PushBack(value interface{}) {
	this.expandIfNeeded()
	this.data[this.end] = value
	this.end = this.nextIndex(this.end)
	this.size++
}

func (this *Deque) PushFront(value interface{}) {
	this.expandIfNeeded()
	this.begin = this.preIndex(this.begin)
	this.data[this.begin] = value
	this.size++
}

func (this *Deque) Insert(position int, value interface{}) error {
	if position < 0 || position > this.size {
		return ErrOutOffRange
	}
	if position == 0 {
		this.PushFront(value)
		return nil
	}
	if position == this.size {
		this.PushBack(value)
		return nil
	}

	this.expandIfNeeded()
	if position < this.size-position {
		//move the front pos items
		idx := this.preIndex(this.begin)
		for i := 0; i < position; i++ {
			this.data[idx] = this.data[this.nextIndex(idx)]
			idx = this.nextIndex(idx)
		}
		this.data[idx] = value
		this.begin = this.preIndex(this.begin)
	} else {
		//move the back pos items
		idx := this.end
		for i := 0; i < this.size-position; i++ {
			this.data[idx] = this.data[this.preIndex(idx)]
			idx = this.preIndex(idx)
		}
		this.data[idx] = value
		this.end = this.nextIndex(this.end)
	}
	this.size++
	return nil
}

func (this *Deque) PopBack() interface{} {
	if this.Empty() {
		return nil
	}
	index := this.preIndex(this.end)
	val := this.data[index]
	this.data[index] = nil
	this.end = index
	this.size--
	this.shrinkIfNeeded()
	return val
}

func (this *Deque) PopFront() interface{} {
	if this.Empty() {
		return nil
	}
	val := this.data[this.begin ]
	this.data[this.begin] = nil
	this.begin = this.nextIndex(this.begin)
	this.size--
	this.shrinkIfNeeded()
	return val
}

func (this *Deque) At(position int) interface{} {
	if position < 0 || position >= this.size {
		return nil
	}
	return this.data[(position+this.begin)%this.Capacity()]
}

func (this *Deque) Set(position int, val interface{}) error {
	if position < 0 || position >= len(this.data) {
		return ErrOutOffRange
	}
	this.data[(position+this.begin)%this.Capacity()] = val
	return nil
}

func (this *Deque) Back() interface{} {
	return this.At(this.size - 1)
}

func (this *Deque) Front() interface{} {
	return this.At(0)
}

func (this *Deque) nextIndex(index int) int {
	return (index + 1) % this.Capacity()
}

func (this *Deque) preIndex(index int) int {
	return (index - 1 + this.Capacity()) % this.Capacity()
}

func (this *Deque) Erase(pos int) error {
	return this.EraseRange(pos, pos+1)
}

func (this *Deque) Clear() {
	this.EraseRange(0, this.size)
}

//EraseRange erase the data in the range [firstPos, lastPos), not include lastPos.
func (this *Deque) EraseRange(firstPos, lastPos int) error {
	if firstPos < 0 || lastPos > this.size {
		return ErrOutOffRange
	}
	if firstPos >= lastPos {
		return nil
	}
	eraseNum := lastPos - firstPos
	leftNum := firstPos
	rightNum := this.size - lastPos

	if leftNum <= rightNum {
		//move left data
		idx := (this.begin + this.preIndex(lastPos)) % this.Capacity()
		for i := 0; i < leftNum; i++ {
			tempIndex := (idx - eraseNum + this.Capacity()) % this.Capacity()
			this.data[idx] = this.data[tempIndex]
			idx = this.preIndex(idx)
		}
		this.begin = this.nextIndex(idx)
		for i := 0; i < eraseNum; i++ {
			this.data[idx] = nil
			idx = this.preIndex(idx)
		}

	} else {
		idx := (this.begin + firstPos) % this.Capacity()
		for i := 0; i < rightNum; i++ {
			tempIndex := (idx + eraseNum + this.Capacity()) % this.Capacity()
			this.data[idx] = this.data[tempIndex]
			idx = this.nextIndex(idx)
		}
		this.end = idx
		for i := 0; i < eraseNum; i++ {
			this.data[idx] = nil
			idx = this.nextIndex(idx)
		}
	}
	this.size -= eraseNum
	this.shrinkIfNeeded()
	return nil
}

func (this *Deque) String() string {
	str := "["
	for i := 0; i < this.size; i++ {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprintf("%v", this.data[(this.begin+i)%this.Capacity()])
	}
	str += "]"
	return str
}

///////////////////////////////////////////////////
//iterator API
func (this *Deque) Begin() *DequeIterator {
	return this.First()
}

func (this *Deque) End() *DequeIterator {
	return this.IterAt(this.Size())
}

func (this *Deque) First() *DequeIterator {
	return this.IterAt(0)
}

func (this *Deque) Last() *DequeIterator {
	return this.IterAt(this.Size() - 1)
}

func (this *Deque) IterAt(position int) *DequeIterator {
	return &DequeIterator{dq: this, position: position}
}
