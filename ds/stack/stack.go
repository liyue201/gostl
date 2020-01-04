package stack

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
		option.container = bidlist.New()
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

func (s *Stack) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Size()
}

func (s *Stack) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Empty()
}

func (s *Stack) Push(value interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.PushBack(value)
}

func (s *Stack) Top() interface{} {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Back()
}

func (s *Stack) Pop() interface{} {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.PopBack()
}

func (s *Stack) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.Clear()
}

func (s *Stack) String() string {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.String()
}
