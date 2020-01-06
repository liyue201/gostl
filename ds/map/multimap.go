package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// Multimap uses RbTress for internal data structure, and keys can bee repeated.
type MultiMap struct {
	tree   *rbtree.RbTree
	keyCmp comparator.Comparator
	locker sync.Locker
}

func NewMultiMap(opts ...Option) *MultiMap {
	option := Options{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiMap{tree: rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

//Insert inserts key-value to the set
func (mm *MultiMap) Insert(key, value interface{}) {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	mm.tree.Insert(key, value)
}

//Get returns the first node's value by key if found, or nil if not found
func (mm *MultiMap) Get(key interface{}) interface{} {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases key in the Map
func (mm *MultiMap) Erase(key interface{}) {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	node := mm.tree.FindNode(key)
	for node != nil && mm.keyCmp(node.Key(), key) == 0 {
		nextNode := node.Next()
		mm.tree.Delete(node)
		node = nextNode
	}
}

//Begin returns the iterator related to key in the set, or an invalid iterator if not exist.
func (mm *MultiMap) Find(key interface{}) *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindNode(key)
	return &MapIterator{node: node}
}

//LowerBound returns the first iterator that equal or greater than key in the Map
func (mm *MultiMap) LowerBound(key interface{}) *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the iterator with the minimum key in the Map, return nil if empty.
func (mm *MultiMap) Begin() *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return mm.First()
}

//First returns the iterator with the minimum key in the Map, return nil if empty.
func (mm *MultiMap) First() *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator{node: mm.tree.First()}
}

//Last returns the iterator with the maximum key in the Map, return nil if empty.
func (mm *MultiMap) Last() *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator{node: mm.tree.Last()}
}

//Clear clears the Map
func (mm *MultiMap) Clear() {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	mm.tree.Clear()
}

// Contains returns true if value in the Map. otherwise returns false.
func (mm *MultiMap) Contains(value interface{}) bool {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	if mm.tree.Find(value) != nil {
		return true
	}
	return false
}

// Contains returns the size of Map
func (mm *MultiMap) Size() int {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return mm.tree.Size()
}

// Traversal traversals elements in the map, it will not stop until to the end or visitor returns false
func (mm *MultiMap) Traversal(visitor visitor.KvVisitor) {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	mm.tree.Traversal(visitor)
}
