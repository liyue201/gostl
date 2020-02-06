package queue

import (
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/ds/list/bidlist"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultLocker    sync.FakeLocker
	defaultContainer = deque.New()
)

// Options holds Queue's options
type Options struct {
	locker    sync.Locker
	container container.Container
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set a Queue goroutine-safe
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithContainer is used to set a Queue's underlying container
func WithContainer(c container.Container) Option {
	return func(option *Options) {
		option.container = c
	}
}

// WithListContainer is used to set List as a Queue's underlying container
func WithListContainer() Option {
	return func(option *Options) {
		option.container = bidlist.New()
	}
}

// Queue is a first-in-first-out data structure
type Queue struct {
	container container.Container
	locker    sync.Locker
}

//New creates a new queue
func New(opts ...Option) *Queue {
	option := Options{
		locker:    defaultLocker,
		container: defaultContainer,
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Queue{
		container: option.container,
		locker:    option.locker,
	}
}

// Size returns the amount of elements in the queue
func (q *Queue) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Size()
}

// Empty returns true if the queue is empty, otherwise returns false
func (q *Queue) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Empty()
}

// Push pushes a value to the end of the queue
func (q *Queue) Push(value interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.PushBack(value)
}

// Front returns the front value in the queue
func (q *Queue) Front() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Front()
}

// Back returns the back value in the queue
func (q *Queue) Back() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Back()
}

// Pop removes the the front element in the queue, and returns its value
func (q *Queue) Pop() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()

	return q.container.PopFront()
}

// Clear clears all elements in the queue
func (q *Queue) Clear() {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.Clear()
}

// String returns a string representation of the queue
func (q *Queue) String() string {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.String()
}
