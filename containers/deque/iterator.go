package deque

import (
	. "github.com/liyue201/gostl/uitls/iterator"
)

type DequeIterator struct {
	dq       *Deque
	curIndex int
}
 
func (this *DequeIterator) IsValid() bool {
	if this.curIndex >= 0 && this.curIndex < this.dq.Size() {
		return true
	}
	return false
}

func (this *DequeIterator) Value() interface{} {
	return this.dq.At(this.curIndex)
}

func (this *DequeIterator) SetValue(val interface{}) error {
	return this.dq.Set(this.curIndex, val)
}

func (this *DequeIterator) Next() ConstIterator {
	if this.IsValid() {
		this.curIndex++
	}
	return this
}

func (this *DequeIterator) Prev() ConstBidIterator {
	if this.IsValid() {
		this.curIndex--
	}
	return this
}

func (this *DequeIterator) Clone() interface{} {
	return &DequeIterator{dq: this.dq, curIndex: this.curIndex}
}
