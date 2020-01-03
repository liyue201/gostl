package queue

import (
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/ds/list/bid_list"
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

func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

func WithContainer(c container.Container) Options {
	return func(option *Option) {
		option.container = c
	}
}

func WithListContainer() Options {
	return func(option *Option) {
		option.container = bid_list.New()
	}
}

type Queue struct {
	container container.Container
	locker    sync.Locker
}

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

func (q *Queue) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Size()
}

func (q *Queue) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Empty()
}

func (q *Queue) Push(value interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.PushBack(value)
}

func (q *Queue) Front() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Front()
}

func (q *Queue) Back() interface{} {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Back()
}

func (q *Queue) Pop() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()

	return q.container.PopFront()
}

func (q *Queue) Clear() {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.Clear()
}

func (q *Queue) String() string {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.String()
}
