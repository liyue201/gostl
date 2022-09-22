package treemap

import (
	"errors"
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
	gosync "sync"
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultLocker        sync.FakeLocker
)

var ErrorNotFound = errors.New("not found")

// Options holds Map's options
type Options struct {
	keyCmp comparator.Comparator
	locker sync.Locker
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithKeyComparator is used to set the key comparator of map
func WithKeyComparator(cmp comparator.Comparator) Option {
	return func(option *Options) {
		option.keyCmp = cmp
	}
}

// WithGoroutineSafe is used to set a map goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterator in multi goroutines
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// Map uses RbTress for internal data structure, and every key can must bee unique.
type Map[K, V any] struct {
	tree   *rbtree.RbTree[K, V]
	locker sync.Locker
}

// New creates a new map
func New[K, V any](opts ...Option) *Map[K, V] {
	option := Options{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &Map[K, V]{tree: rbtree.New[K, V](rbtree.WithKeyComparator(option.keyCmp)),
		locker: option.locker,
	}
}

//Insert inserts a key-value to the map
func (m *Map[K, V]) Insert(key K, value V) {
	m.locker.Lock()
	defer m.locker.Unlock()

	node := m.tree.FindNode(key)
	if node != nil {
		node.SetValue(value)
		return
	}
	m.tree.Insert(key, value)
}

//Get returns the value of the passed key if the key is in the map, otherwise returns nil
func (m *Map[K, V]) Get(key K) (V, error) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindNode(key)
	if node != nil {
		return node.Value(), nil
	}
	return *new(V), ErrorNotFound
}

//Erase erases the node by the passed key from the map if the key in the Map
func (m *Map[K, V]) Erase(key K) {
	m.locker.Lock()
	defer m.locker.Unlock()

	node := m.tree.FindNode(key)
	if node != nil {
		m.tree.Delete(node)
	}
}

//EraseIter erases the node that iterator iter point to from the map
func (m *Map[K, V]) EraseIter(iter iterator.ConstKvIterator[K, V]) {
	m.locker.Lock()
	defer m.locker.Unlock()

	mpIter, ok := iter.(*MapIterator[K, V])
	if ok {
		m.tree.Delete(mpIter.node)
	}
}

//Find finds a node by the passed key and returns its iterator
func (m *Map[K, V]) Find(key K) *MapIterator[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindNode(key)
	return &MapIterator[K, V]{node: node}
}

//LowerBound finds a node that its key is equal or greater than the passed key and returns its iterator
func (m *Map[K, V]) LowerBound(key K) *MapIterator[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindLowerBoundNode(key)
	return &MapIterator[K, V]{node: node}
}

//UpperBound finds a node that its key is greater than the passed key and returns its iterator
func (m *Map[K, V]) UpperBound(key K) *MapIterator[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindUpperBoundNode(key)
	return &MapIterator[K, V]{node: node}
}

//Begin returns the first node's iterator
func (m *Map[K, V]) Begin() *MapIterator[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator[K, V]{node: m.tree.First()}
}

//First returns the first node's iterator
func (m *Map[K, V]) First() *MapIterator[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator[K, V]{node: m.tree.First()}
}

//Last returns the last node's iterator
func (m *Map[K, V]) Last() *MapIterator[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator[K, V]{node: m.tree.Last()}
}

//Clear clears the map
func (m *Map[K, V]) Clear() {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.tree.Clear()
}

// Contains returns true if the key is in the map. otherwise returns false.
func (m *Map[K, V]) Contains(key K) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()

	if _, err := m.tree.Find(key); err == nil {
		return true
	}
	return false
}

// Size returns the amount of elements in the map
func (m *Map[K, V]) Size() int {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return m.tree.Size()
}

// Traversal traversals elements in the map, it will not stop until to the end or the visitor returns false
func (m *Map[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	m.tree.Traversal(visitor)
}
