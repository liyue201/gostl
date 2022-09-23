package priorityqueue

import (
	"github.com/liyue201/gostl/ds/heap"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultComparator = comparator.BuiltinTypeComparator
	defaultLocker     sync.FakeLocker
)

// ElementHolder holds elements of the PriorityQueue
type ElementHolder[T any] struct {
	elements []T
	cmpFun   comparator.Comparator
}

// Push pushes an element to the ElementHolder
func (h *ElementHolder[T]) Push(element T) {
	h.elements = append(h.elements, element)
}

// Pop pops an element from the ElementHolder
func (h *ElementHolder[T]) Pop() T {
	if len(h.elements) == 0 {
		panic("queue is empty")
	}
	item := h.elements[h.Len()-1]
	h.elements = h.elements[:h.Len()-1]
	return item
}

func (h *ElementHolder[T]) top() T {
	if len(h.elements) == 0 {
		panic("queue is empty")
	}
	return h.elements[0]
}

// Len returns the amount of elements in ElementHolder
func (h *ElementHolder[T]) Len() int {
	return len(h.elements)
}

// Len compare two elements at position i and j , and returns true if elements[i] < elements[j]
func (h *ElementHolder[T]) Less(i, j int) bool {
	if h.cmpFun(h.elements[i], h.elements[j]) < 0 {
		return true
	}
	return false
}

// Swap swaps two elements at position i and j
func (h *ElementHolder[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// Options holds PriorityQueue's options
type Options struct {
	cmp    comparator.Comparator
	locker sync.Locker
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithComparator is used to set the PriorityQueue's comparator
func WithComparator(cmp comparator.Comparator) Option {
	return func(option *Options) {
		option.cmp = cmp
	}
}

// WithGoroutineSafe is used to set the PriorityQueue goroutine-safe
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// PriorityQueue is an implementation of priority queue
type PriorityQueue[T any] struct {
	holder *ElementHolder[T]
	locker sync.Locker
}

// New creates a PriorityQueue
func New[T any](opts ...Option) *PriorityQueue[T] {
	option := Options{
		cmp:    defaultComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	holder := &ElementHolder[T]{
		elements: make([]T, 0, 0),
		cmpFun:   option.cmp,
	}
	return &PriorityQueue[T]{
		holder: holder,
		locker: option.locker,
	}
}

// Push pushes an element to the PriorityQueue
func (q *PriorityQueue[T]) Push(e T) {
	q.locker.Lock()
	defer q.locker.Unlock()

	heap.Push[T](q.holder, e)
}

// Pop pops an element from the PriorityQueue
func (q *PriorityQueue[T]) Pop() T {
	q.locker.Lock()
	defer q.locker.Unlock()

	return heap.Pop[T](q.holder)
}

// Top returns the top element in the PriorityQueue
func (q *PriorityQueue[T]) Top() T {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.top()
}

// Empty returns true if the PriorityQueue is empty, otherwise returns false
func (q *PriorityQueue[T]) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.Len() == 0
}

// Size returns the amount of elements in the queue
func (q *PriorityQueue[T]) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.holder.Len()
}
