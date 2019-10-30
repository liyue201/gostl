package priority_queue

import (
	"container/heap"
	. "github.com/liyue201/gostl/comparator"
)

var (
	defaultComparator = BuiltinTypeComparator
)

type ItemHolder struct {
	items  []interface{}
	cmpFun Comparator
}

func (this *ItemHolder) Push(item interface{}) {
	this.items = append(this.items, item)
}

func (this *ItemHolder) Pop() interface{} {
	if len(this.items) == 0 {
		return nil
	}
	item := this.items[this.Len()-1]
	this.items = this.items[:this.Len()-1]
	return item
}

func (this *ItemHolder) top() interface{} {
	if len(this.items) == 0 {
		return nil
	}
	return this.items[0]
}

func (this *ItemHolder) Size() int {
	return len(this.items)
}

func (this *ItemHolder) Len() int {
	return len(this.items)
}

func (this *ItemHolder) Less(i, j int) bool {
	if this.cmpFun(this.items[i], this.items[j]) < 0 {
		return true
	}
	return false
}

func (this *ItemHolder) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

type Option struct {
	cmp Comparator
}

type Options func(option *Option)

func WithComparator(cmp Comparator) Options {
	return func(option *Option) {
		option.cmp = cmp
	}
}

type PriorityQueue struct {
	holder *ItemHolder
}

func New(opts ...Options) *PriorityQueue {
	option := Option{
		cmp: defaultComparator,
	}
	for _, opt := range opts {
		opt(&option)
	}
	holder := &ItemHolder{
		items:  make([]interface{}, 0, 0),
		cmpFun: option.cmp,
	}
	return &PriorityQueue{
		holder: holder,
	}
}

func (this *PriorityQueue) Push(item interface{}) {
	heap.Push(this.holder, item)
}

func (this *PriorityQueue) Pop() interface{} {
	return heap.Pop(this.holder)
}

func (this *PriorityQueue) Top() interface{} {
	return this.holder.top()
}

func (this *PriorityQueue) Empty() bool {
	return this.holder.Size() == 0
}
