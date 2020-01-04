package priorityqueue

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

// ElementHolder is the holder of elements
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

// WithComparator sets the comparator option
func WithComparator(cmp Comparator) Options {
	return func(option *Option) {
		option.cmp = cmp
	}
}

// WithThreadSave sets the ThreadSave option
func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

// PriorityQueue is an implementation of priority queue
type PriorityQueue struct {
	holder *ElementHolder
	locker sync.Locker
}

// New news a PriorityQueue
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

// Push pushes an item to q
func (q *PriorityQueue) Push(item interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	heap.Push(q.holder, item)
}

// Pop pops an item from q
func (q *PriorityQueue) Pop() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()

	return heap.Pop(q.holder)
}

// Top returns the top item at q
func (q *PriorityQueue) Top() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.top()
}

// Top returns whether q is empty
func (q *PriorityQueue) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.Size() == 0
}
