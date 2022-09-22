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

// Options holds the Stack's options
type Options struct {
	locker    sync.Locker
	container container.Container
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set a stack goroutine-safe
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithContainer is used to set a stack's underlying container
func WithContainer(c container.Container) Option {
	return func(option *Options) {
		option.container = c
	}
}

// WithListContainer is used to set List for a stack's underlying container
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

// New creates a new stack
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

// Size returns the amount of elements in the stack
func (s *Stack) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Size()
}

// Empty returns true if the stack is empty, otherwise returns false
func (s *Stack) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Empty()
}

// Push pushes a value to the stack
func (s *Stack) Push(value any) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.PushBack(value)
}

// Top returns the top value in the stack
func (s *Stack) Top() any {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Back()
}

// Pop removes the the top value in the stack and returns it
func (s *Stack) Pop() any {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.PopBack()
}

// Clear clears all elements in the stack
func (s *Stack) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.Clear()
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.String()
}
