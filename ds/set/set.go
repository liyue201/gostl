package set

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
	gosync "sync"
)

// constants
const (
	Empty = 0
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultLocker        sync.FakeLocker
)

// Options holds Set's options
type Options struct {
	keyCmp comparator.Comparator
	locker sync.Locker
}

// Option is a function used to set Options
type Option func(option *Options)

// WithKeyComparator sets Key comparator option
func WithKeyComparator(cmp comparator.Comparator) Option {
	return func(option *Options) {
		option.keyCmp = cmp
	}
}

// WithThreadSafe set Map thread-safety,
// Note that iterators are not thread safe, and it is useless to turn on the setting option here.
// so don't use iterators in multi goroutines
func WithThreadSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// Set uses RbTress for internal data structure, and every key can must bee unique.
type Set struct {
	tree   *rbtree.RbTree
	keyCmp comparator.Comparator
	locker sync.Locker
}

// New news a set
func New(opts ...Option) *Set {
	option := Options{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &Set{
		tree:   rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

// Insert inserts element to the Set
func (s *Set) Insert(element interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	node := s.tree.FindNode(element)
	if node != nil {
		return
	}
	s.tree.Insert(element, Empty)
}

// Erase erases element in the Set
func (s *Set) Erase(element interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	node := s.tree.FindNode(element)
	if node != nil {
		s.tree.Delete(node)
	}
}

// Find returns the iterator related to element in the Set,or an invalid iterator if not exist.
func (s *Set) Find(element interface{}) *SetIterator {
	s.locker.RLock()
	defer s.locker.RUnlock()

	node := s.tree.FindNode(element)
	return &SetIterator{node: node}
}

// LowerBound returns the first iterator that equal or greater than element in the Set
func (s *Set) LowerBound(element interface{}) *SetIterator {
	s.locker.RLock()
	defer s.locker.RUnlock()

	node := s.tree.FindLowerBoundNode(element)
	return &SetIterator{node: node}
}

// Begin returns the iterator with the minimum element in the Set, return nil if empty.
func (s *Set) Begin() *SetIterator {
	return s.First()
}

// First returns the iterator with the minimum element in the Set, return nil if empty.
func (s *Set) First() *SetIterator {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return &SetIterator{node: s.tree.First()}
}

// Last returns the iterator with the maximum element in the Set, return nil if empty.
func (s *Set) Last() *SetIterator {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return &SetIterator{node: s.tree.Last()}
}

// Clear clears the Set
func (s *Set) Clear() {
	s.locker.Lock()
	defer s.locker.Lock()

	s.tree.Clear()
}

// Contains returns true if element in the Set. otherwise returns false.
func (s *Set) Contains(element interface{}) bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	if s.tree.Find(element) != nil {
		return true
	}
	return false
}

// Size returns the size of Set
func (s *Set) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.tree.Size()
}

// Traversal traversals elements in set, it will not stop until to the end or visitor returns false
func (s *Set) Traversal(visitor visitor.Visitor) {
	s.locker.RLock()
	defer s.locker.RUnlock()

	for node := s.tree.First(); node != nil; node = node.Next() {
		if !visitor(node.Key()) {
			break
		}
	}
}

// String returns the set's elements in string format
func (s *Set) String() string {
	str := "["
	s.Traversal(func(value interface{}) bool {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", value)
		return true
	})
	str += "]"
	return str
}

// Intersect returns a set with the common elements in s set and the other set
// Please ensure s set and other set uses the same keyCmp
func (s *Set) Intersect(other *Set) *Set {
	s.locker.RLock()
	defer s.locker.RUnlock()

	set := New(WithKeyComparator(s.keyCmp))
	sIter := s.tree.IterFirst()
	otherIter := other.tree.IterFirst()
	for sIter.IsValid() && otherIter.IsValid() {
		cmp := s.keyCmp(sIter.Key(), otherIter.Key())
		if cmp == 0 {
			set.tree.Insert(sIter.Key(), Empty)
			sIter.Next()
			otherIter.Next()
		} else if cmp < 0 {
			sIter.Next()
		} else {
			otherIter.Next()
		}
	}
	return set
}

// Union returns  a set with the all elements in s set and the other set
// Please ensure s set and other set uses the same keyCmp
func (s *Set) Union(other *Set) *Set {
	s.locker.RLock()
	defer s.locker.RUnlock()

	set := New(WithKeyComparator(s.keyCmp))
	sIter := s.tree.IterFirst()
	otherIter := other.tree.IterFirst()
	for sIter.IsValid() && otherIter.IsValid() {
		cmp := s.keyCmp(sIter.Key(), otherIter.Key())
		if cmp == 0 {
			set.tree.Insert(sIter.Key(), Empty)
			sIter.Next()
			otherIter.Next()
		} else if cmp < 0 {
			set.tree.Insert(sIter.Key(), Empty)
			sIter.Next()
		} else {
			set.tree.Insert(otherIter.Key(), Empty)
			otherIter.Next()
		}
	}
	for ; sIter.IsValid(); sIter.Next() {
		set.tree.Insert(sIter.Key(), Empty)
	}
	for ; otherIter.IsValid(); otherIter.Next() {
		set.tree.Insert(otherIter.Key(), Empty)
	}
	return set
}

// Diff returns a set with the elements in s set but not in the other set
// Please ensure s set and other set uses the same keyCmp
func (s *Set) Diff(other *Set) *Set {
	s.locker.RLock()
	defer s.locker.RUnlock()

	set := New(WithKeyComparator(s.keyCmp))
	sIter := s.tree.IterFirst()
	otherIter := other.tree.IterFirst()
	for sIter.IsValid() && otherIter.IsValid() {
		cmp := s.keyCmp(sIter.Key(), otherIter.Key())
		if cmp == 0 {
			sIter.Next()
			otherIter.Next()
		} else if cmp < 0 {
			set.tree.Insert(sIter.Key(), Empty)
			sIter.Next()
		} else {
			otherIter.Next()
		}
	}
	for ; sIter.IsValid(); sIter.Next() {
		set.tree.Insert(sIter.Key(), Empty)
	}
	return set
}
