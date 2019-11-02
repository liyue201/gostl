package stack

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

type Stack struct {
	container container.Container
	locker    sync.Locker
}

func New(opts ...Options) *Stack {
	option := Option{
		locker:    defaultLocker,
		container: defaultContainer,
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Stack{
		container: option.container,
		locker:    option.locker,
	}
}

func (this *Stack) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Size()
}

func (this *Stack) Empty() bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Empty()
}

func (this *Stack) Push(value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.container.PushBack(value)
}

func (this *Stack) Top() interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.Back()
}

func (this *Stack) Pop() interface{} {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.container.PopBack()
}

func (this *Stack) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.container.Clear()
}

func (this *Stack) String() string {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.container.String()
}
