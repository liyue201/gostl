package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// MultiMap uses RbTress for internal data structure, and keys can bee repeated.
type MultiMap[K, V any] struct {
	tree   *rbtree.RbTree[K, V]
	locker sync.Locker
}

//NewMultiMap creates a new MultiMap
func NewMultiMap[K, V any](cmp comparator.Comparator[K], opts ...Option) *MultiMap[K, V] {
	option := Options{
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiMap[K, V]{tree: rbtree.New[K, V](cmp),
		locker: option.locker,
	}
}

//Insert inserts a key-value to the MultiMap
func (mm *MultiMap[K, V]) Insert(key K, value V) {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	mm.tree.Insert(key, value)
}

//Get returns the first node's value by the passed key if the key is in the MultiMap, otherwise returns nil
func (mm *MultiMap[K, V]) Get(key K) (V, error) {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindNode(key)
	if node != nil {
		return node.Value(), nil
	}
	return *new(V), ErrorNotFound
}

//Erase erases the key in the MultiMap
func (mm *MultiMap[K, V]) Erase(key K) {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	for {
		node := mm.tree.FindNode(key)
		if node == nil {
			break
		}
		mm.tree.Delete(node)
	}
}

//Find finds the node by the passed key in the MultiMap and returns its iterator
func (mm *MultiMap[K, V]) Find(key K) *MapIterator[K, V] {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindNode(key)
	return &MapIterator[K, V]{node: node}
}

//LowerBound find the first node that its key is equal or greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap[K, V]) LowerBound(key K) *MapIterator[K, V] {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindLowerBoundNode(key)
	return &MapIterator[K, V]{node: node}
}

//UpperBound find the first node that its key is greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap[K, V]) UpperBound(key K) *MapIterator[K, V] {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindUpperBoundNode(key)
	return &MapIterator[K, V]{node: node}
}

//Begin returns the first node's iterator
func (mm *MultiMap[K, V]) Begin() *MapIterator[K, V] {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator[K, V]{node: mm.tree.First()}
}

//First returns the first node's iterator
func (mm *MultiMap[K, V]) First() *MapIterator[K, V] {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator[K, V]{node: mm.tree.First()}
}

//Last returns the last node's iterator
func (mm *MultiMap[K, V]) Last() *MapIterator[K, V] {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator[K, V]{node: mm.tree.Last()}
}

//Clear clears the MultiMap
func (mm *MultiMap[K, V]) Clear() {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	mm.tree.Clear()
}

// Contains returns true if the passed value is in the MultiMap. otherwise returns false.
func (mm *MultiMap[K, V]) Contains(key K) bool {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	if _, err := mm.tree.Find(key); err == nil {
		return true
	}
	return false
}

// Size returns the amount of elements in the MultiMap
func (mm *MultiMap[K, V]) Size() int {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return mm.tree.Size()
}

// Traversal traversals elements in the MultiMap, it will not stop until to the end of the MultiMap or the visitor returns false
func (mm *MultiMap[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	mm.tree.Traversal(visitor)
}
