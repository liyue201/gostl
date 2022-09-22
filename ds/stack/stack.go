package stack

import (
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/ds/list/bidlist"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds the Stack's options
type Options[T any] struct {
	locker    sync.Locker
	container container.Container[T]
}

// Option is a function type used to set Options
type Option[T any] func(option *Options[T])

// WithGoroutineSafe is used to set a stack goroutine-safe
func WithGoroutineSafe[T any]() Option[T] {
	return func(option *Options[T]) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithContainer is used to set a stack's underlying container
func WithContainer[T any](c container.Container[T]) Option[T] {
	return func(option *Options[T]) {
		option.container = c
	}
}

// WithListContainer is used to set List for a stack's underlying container
func WithListContainer[T any]() Option[T] {
	return func(option *Options[T]) {
		option.container = bidlist.New[T]()
	}
}

//Stack is a last-in-first-out data structure
type Stack[T any] struct {
	container container.Container[T]
	locker    sync.Locker
}

// New creates a new stack
func New[T any](opts ...Option[T]) *Stack[T] {
	option := Options[T]{
		locker:    defaultLocker,
		container: deque.New[T](),
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Stack[T]{
		container: option.container,
		locker:    option.locker,
	}
}

// Size returns the amount of elements in the stack
func (s *Stack[T]) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Size()
}

// Empty returns true if the stack is empty, otherwise returns false
func (s *Stack[T]) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Empty()
}

// Push pushes a value to the stack
func (s *Stack[T]) Push(value T) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.PushBack(value)
}

// Top returns the top value in the stack
func (s *Stack[T]) Top() T {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Back()
}

// Pop removes the top value in the stack and returns it
func (s *Stack[T]) Pop() T {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.PopBack()
}

// Clear clears all elements in the stack
func (s *Stack[T]) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.Clear()
}

// String returns a string representation of the stack
func (s *Stack[T]) String() string {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.String()
}
