package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// MultiMap uses RbTress for internal data structure, and keys can bee repeated.
type MultiMap struct {
	tree   *rbtree.RbTree
	keyCmp comparator.Comparator
	locker sync.Locker
}

//NewMultiMap creates a new MultiMap
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

//Insert inserts a key-value to the MultiMap
func (mm *MultiMap) Insert(key, value any) {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	mm.tree.Insert(key, value)
}

//Get returns the first node's value by the passed key if the key is in the MultiMap, otherwise returns nil
func (mm *MultiMap) Get(key any) any {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases the key in the MultiMap
func (mm *MultiMap) Erase(key any) {
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
func (mm *MultiMap) Find(key any) *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindNode(key)
	return &MapIterator{node: node}
}

//LowerBound find the first node that its key is equal or greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap) LowerBound(key any) *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//UpperBound find the first node that its key is greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap) UpperBound(key any) *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	node := mm.tree.FindUpperBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the first node's iterator
func (mm *MultiMap) Begin() *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator{node: mm.tree.First()}
}

//First returns the first node's iterator
func (mm *MultiMap) First() *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator{node: mm.tree.First()}
}

//Last returns the last node's iterator
func (mm *MultiMap) Last() *MapIterator {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return &MapIterator{node: mm.tree.Last()}
}

//Clear clears the MultiMap
func (mm *MultiMap) Clear() {
	mm.locker.Lock()
	defer mm.locker.Unlock()

	mm.tree.Clear()
}

// Contains returns true if the passed value is in the MultiMap. otherwise returns false.
func (mm *MultiMap) Contains(value any) bool {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	if mm.tree.Find(value) != nil {
		return true
	}
	return false
}

// Size returns the amount of elements in the MultiMap
func (mm *MultiMap) Size() int {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	return mm.tree.Size()
}

// Traversal traversals elements in the MultiMap, it will not stop until to the end of the MultiMap or the visitor returns false
func (mm *MultiMap) Traversal(visitor visitor.KvVisitor) {
	mm.locker.RLock()
	defer mm.locker.RUnlock()

	mm.tree.Traversal(visitor)
}
