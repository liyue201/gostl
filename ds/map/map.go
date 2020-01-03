package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/comparator"
	. "github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
	gosync "sync"
)

var (
	defaultKeyComparator = BuiltinTypeComparator
	defaultLocker        sync.FakeLocker
)

type Option struct {
	keyCmp Comparator
	locker sync.Locker
}

type Options func(option *Option)

// Key comparator option
func WithKeyComparator(cmp Comparator) Options {
	return func(option *Option) {
		option.keyCmp = cmp
	}
}

// Note that iterators are not thread safe, and it is useless to turn on the setting option here.
// so don't use iterators in multi goroutines
func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

// Map uses RbTress for internal data structure, and every key can must bee unique.
type Map struct {
	tree   *rbtree.RbTree
	locker sync.Locker
}

// New new a map
func New(opts ...Options) *Map {
	option := Option{
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

//Insert inserts key-value to the map
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

//Get returns the value by key if found, or nil if not found
func (m *Map) Get(key interface{}) interface{} {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases node by key in the Map
func (m *Map) Erase(key interface{}) {
	m.locker.Lock()
	defer m.locker.Unlock()

	node := m.tree.FindNode(key)
	if node != nil {
		m.tree.Delete(node)
	}
}

//Erase erases node by iter in the Map
func (m *Map) EraseIter(iter ConstKvIterator) {
	m.locker.Lock()
	defer m.locker.Unlock()

	mpIter, ok := iter.(*MapIterator)
	if ok {
		m.tree.Delete(mpIter.node)
	}
}

//Begin returns the iterator related to value in the map, or an invalid iterator if not exist.
func (m *Map) Find(key interface{}) *MapIterator {
	m.locker.RUnlock()
	defer m.locker.RUnlock()

	node := m.tree.FindNode(key)
	return &MapIterator{node: node}
}

//LowerBound returns the first iterator that equal or greater than key in the Map
func (m *Map) LowerBound(key interface{}) *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	node := m.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the iterator with the minimum key in the Map, return nil if empty.
func (m *Map) Begin() *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return m.First()
}

//First returns the iterator with the minimum key in the Map, return nil if empty.
func (m *Map) First() *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator{node: m.tree.First()}
}

//Last returns the iterator with the maximum key in the Map, return nil if empty.
func (m *Map) Last() *MapIterator {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return &MapIterator{node: m.tree.Last()}
}

//Clear clears the Map
func (m *Map) Clear() {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.tree.Clear()
}

// Contains returns true if key in the Map. otherwise returns false.
func (m *Map) Contains(key interface{}) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()

	if m.tree.Find(key) != nil {
		return true
	}
	return false
}

// Contains returns the size of Map
func (m *Map) Size() int {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return m.tree.Size()
}

// Traversal traversals elements in map, it will not stop until to the end or visitor returns false
func (m *Map) Traversal(visitor visitor.KvVisitor) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	m.tree.Traversal(visitor)
}
