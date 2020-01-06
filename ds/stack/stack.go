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

// Options holds Stack's options
type Options struct {
	locker    sync.Locker
	container container.Container
}

// Option is a function used to set Options
type Option func(option *Options)

// WithThreadSafe uses ThreadSafe
func WithThreadSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithContainer uses c  for internal Container
func WithContainer(c container.Container) Option {
	return func(option *Options) {
		option.container = c
	}
}

// WithListContainer uses List for internal Container
func WithListContainer() Option {
	return func(option *Options) {
		option.container = bidlist.New()
	}
}

//Stack is a last-in-first-out data structure
type Stack struct {
	container container.Container
	locker    sync.Locker
}

// New news Stack
func New(opts ...Option) *Stack {
	option := Options{
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

// Size returns the size of s
func (s *Stack) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Size()
}

// Empty returns whether s is empty or not
func (s *Stack) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Empty()
}

// Push pushes value to s
func (s *Stack) Push(value interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.PushBack(value)
}

// Top returns the top value in s
func (s *Stack) Top() interface{} {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Back()
}

// Pop removes the the top item in s, and returns it's value
func (s *Stack) Pop() interface{} {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.PopBack()
}

// Clear clears all items in s
func (s *Stack) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.Clear()
}

// String returns s in string format
func (s *Stack) String() string {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.String()
}
