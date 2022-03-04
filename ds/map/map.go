package treemap

import (
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
type Map struct {
	tree   *rbtree.RbTree
	locker sync.Locker
}

// New creates a new map
func New(opts ...Option) *Map {
	option := Options{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &Map{tree: rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		locker: option.locker,
	}
}

//Insert inserts a key-value to the map
func (m *Map) Insert(key, value interface{}) {
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
func (m *Map) Get(key interface{}) interface{} {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases the node by the passed key from the map if the key in the Map
func (m *Map) Erase(key interface{}) {
	m.locker.Lock()
	defer m.locker.Unlock()

	node := m.tree.FindNode(key)
	if node != nil {
		m.tree.Delete(node)
	}
}

//EraseIter erases the node that iterator iter point to from the map
func (m *Map) EraseIter(iter iterator.ConstKvIterator) {
	m.locker.Lock()
	defer m.locker.Unlock()

	mpIter, ok := iter.(*MapIterator)
	if ok {
		m.tree.Delete(mpIter.node)
	}
}

//Find finds a node by the passed key and returns its iterator
func (m *Map) Find(key interface{}) *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindNode(key)
	return &MapIterator{node: node}
}

//LowerBound finds a node that its key is equal or greater than the passed key and returns its iterator
func (m *Map) LowerBound(key interface{}) *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//UpperBound finds a node that its key is greater than the passed key and returns its iterator
func (m *Map) UpperBound(key interface{}) *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindUpperBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the first node's iterator
func (m *Map) Begin() *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator{node: m.tree.First()}
}

//First returns the first node's iterator
func (m *Map) First() *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator{node: m.tree.First()}
}

//Last returns the last node's iterator
func (m *Map) Last() *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator{node: m.tree.Last()}
}

//Clear clears the map
func (m *Map) Clear() {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.tree.Clear()
}

// Contains returns true if the key is in the map. otherwise returns false.
func (m *Map) Contains(key interface{}) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()

	if m.tree.Find(key) != nil {
		return true
	}
	return false
}

// Size returns the amount of elements in the map
func (m *Map) Size() int {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return m.tree.Size()
}

// Traversal traversals elements in the map, it will not stop until to the end or the visitor returns false
func (m *Map) Traversal(visitor visitor.KvVisitor) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	m.tree.Traversal(visitor)
}
