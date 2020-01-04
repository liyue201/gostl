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

type Option struct {
	locker    sync.Locker
	container container.Container
}

type Options func(option *Option)

// WithThreadSave uses  ThreadSave
func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithContainer uses c as for internal Container
func WithContainer(c container.Container) Options {
	return func(option *Option) {
		option.container = c
	}
}

// WithListContainer uses List for internal Container
func WithListContainer() Options {
	return func(option *Option) {
		option.container = bidlist.New()
	}
}

//Queue is a first in first out data structure
type Queue struct {
	container container.Container
	locker    sync.Locker
}

//New new a queue
func New(opts ...Options) *Queue {
	option := Option{
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

// Size returns the size of q
func (q *Queue) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Size()
}

// Size returns whether q is empty
func (q *Queue) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Empty()
}

// Push pushes value to q
func (q *Queue) Push(value interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.PushBack(value)
}

// Front returns the first value in q
func (q *Queue) Front() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Front()
}

// Front returns the last value in q
func (q *Queue) Back() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Back()
}

// Pop removes the the first item in q, and returns it's value
func (q *Queue) Pop() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()

	return q.container.PopFront()
}

// Clear clears all items in q
func (q *Queue) Clear() {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.Clear()
}

// String returns q in string format
func (q *Queue) String() string {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.String()
}
