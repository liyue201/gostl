package priority_queue

import (
	"container/heap"
	. "github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultComparator = BuiltinTypeComparator
	defaultLocker     sync.FakeLocker
)

type ElementHolder struct {
	elements []interface{}
	cmpFun   Comparator
}

func (this *ElementHolder) Push(element interface{}) {
	this.elements = append(this.elements, element)
}

func (this *ElementHolder) Pop() interface{} {
	if len(this.elements) == 0 {
		return nil
	}
	item := this.elements[this.Len()-1]
	this.elements = this.elements[:this.Len()-1]
	return item
}

func (this *ElementHolder) top() interface{} {
	if len(this.elements) == 0 {
		return nil
	}
	return this.elements[0]
}

func (this *ElementHolder) Size() int {
	return len(this.elements)
}

func (this *ElementHolder) Len() int {
	return len(this.elements)
}

func (this *ElementHolder) Less(i, j int) bool {
	if this.cmpFun(this.elements[i], this.elements[j]) < 0 {
		return true
	}
	return false
}

func (this *ElementHolder) Swap(i, j int) {
	this.elements[i], this.elements[j] = this.elements[j], this.elements[i]
}

type Option struct {
	cmp    Comparator
	locker sync.Locker
}

type Options func(option *Option)

func WithComparator(cmp Comparator) Options {
	return func(option *Option) {
		option.cmp = cmp
	}
}

func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

type PriorityQueue struct {
	holder *ElementHolder
	locker sync.Locker
}

func New(opts ...Options) *PriorityQueue {
	option := Option{
		cmp:    defaultComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	holder := &ElementHolder{
		elements: make([]interface{}, 0, 0),
		cmpFun:   option.cmp,
	}
	return &PriorityQueue{
		holder: holder,
		locker: option.locker,
	}
}

func (this *PriorityQueue) Push(item interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	heap.Push(this.holder, item)
}

func (this *PriorityQueue) Pop() interface{} {
	this.locker.Lock()
	defer this.locker.Unlock()

	return heap.Pop(this.holder)
}

func (this *PriorityQueue) Top() interface{} {
	this.locker.RLock()
	defer this.locker.RLock()

	return this.holder.top()
}

func (this *PriorityQueue) Empty() bool {
	this.locker.RLock()
	defer this.locker.RLock()

	return this.holder.Size() == 0
}
