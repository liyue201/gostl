package queue

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

type Queue struct {
	dq     *deque.Deque
	locker sync.Locker
}

func New(opts ...Options) *Queue {
	option := Option{
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Queue{
		dq:     deque.New(),
		locker: option.locker,
	}
}

func (this *Queue) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Size()
}

func (this *Queue) Empty() bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Empty()
}

func (this *Queue) Push(value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.dq.PushBack(value)
}

func (this *Queue) Front() interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Front()
}

func (this *Queue) Back() interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.Back()
}

func (this *Queue) Pop() interface{} {
	this.locker.Lock()
	defer this.locker.Unlock()

	return this.dq.PopFront()
}

func (this *Queue) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.dq.Clear()
}

func (this *Queue) String() string {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.dq.String()
}
