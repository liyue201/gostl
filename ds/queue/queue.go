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

func (this *Queue) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Size()
}

func (this *Queue) Empty() bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Empty()
}

func (this *Queue) Push(value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.container.PushBack(value)
}

func (this *Queue) Front() interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Front()
}

func (this *Queue) Back() interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Back()
}

func (this *Queue) Pop() interface{} {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.container.PopFront()
}

func (this *Queue) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.container.Clear()
}

func (this *Queue) String() string {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.String()
}
