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

func (h *ElementHolder) Push(element interface{}) {
	h.elements = append(h.elements, element)
}

func (h *ElementHolder) Pop() interface{} {
	if len(h.elements) == 0 {
		return nil
	}
	item := h.elements[h.Len()-1]
	h.elements = h.elements[:h.Len()-1]
	return item
}

func (h *ElementHolder) top() interface{} {
	if len(h.elements) == 0 {
		return nil
	}
	return h.elements[0]
}

func (h *ElementHolder) Size() int {
	return len(h.elements)
}

func (h *ElementHolder) Len() int {
	return len(h.elements)
}

func (h *ElementHolder) Less(i, j int) bool {
	if h.cmpFun(h.elements[i], h.elements[j]) < 0 {
		return true
	}
	return false
}

func (h *ElementHolder) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
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

func (q *PriorityQueue) Push(item interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	heap.Push(q.holder, item)
}

func (q *PriorityQueue) Pop() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()

	return heap.Pop(q.holder)
}

func (q *PriorityQueue) Top() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.top()
}

func (q *PriorityQueue) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.Size() == 0
}
