package stack

import (
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultLocker sync.FakeLocker
)

type Option struct {
	locker sync.Locker
}

type Options func(option *Option)

func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

type Stack struct {
	dq     *deque.Deque
	locker sync.Locker
}

func New(opts ...Options) *Stack {
	option := Option{
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Stack{
		dq:     deque.New(),
		locker: option.locker,
	}
}

func (this *Stack) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Size()
}

func (this *Stack) Empty() bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Empty()
}

func (this *Stack) Push(value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.dq.PushBack(value)
}

func (this *Stack) Top() interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Back()
}

func (this *Stack) Pop() interface{} {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.dq.PopBack()
}

func (this *Stack) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.dq.Clear()
}

func (this *Stack) String() string {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.String()
}
