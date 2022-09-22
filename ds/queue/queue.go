package queue

import (
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/ds/list/bidlist"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds Queue's options
type Options[T any] struct {
	locker    sync.Locker
	container container.Container[T]
}

// Option is a function type used to set Options
type Option[T any] func(option *Options[T])

// WithGoroutineSafe is used to set a Queue goroutine-safe
func WithGoroutineSafe[T any]() Option[T] {
	return func(option *Options[T]) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithContainer is used to set a Queue's underlying container
func WithContainer[T any](c container.Container[T]) Option[T] {
	return func(option *Options[T]) {
		option.container = c
	}
}

// WithListContainer is used to set List as a Queue's underlying container
func WithListContainer[T any]() Option[T] {
	return func(option *Options[T]) {
		option.container = bidlist.New[T]()
	}
}

// Queue is a first-in-first-out data structure
type Queue[T any] struct {
	container container.Container[T]
	locker    sync.Locker
}

//New creates a new queue
func New[T any](opts ...Option[T]) *Queue[T] {
	option := Options[T]{
		locker:    defaultLocker,
		container: deque.New[T](),
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Queue[T]{
		container: option.container,
		locker:    option.locker,
	}
}

// Size returns the amount of elements in the queue
func (q *Queue[T]) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Size()
}

// Empty returns true if the queue is empty, otherwise returns false
func (q *Queue[T]) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Empty()
}

// Push pushes a value to the end of the queue
func (q *Queue[T]) Push(value T) {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.PushBack(value)
}

// Front returns the front value in the queue
func (q *Queue[T]) Front() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Front()
}

// Back returns the back value in the queue
func (q *Queue[T]) Back() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Back()
}

// Pop removes the the front element in the queue, and returns its value
func (q *Queue[T]) Pop() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()

	return q.container.PopFront()
}

// Clear clears all elements in the queue
func (q *Queue[T]) Clear() {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.Clear()
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.String()
}
