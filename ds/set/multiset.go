package set

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// MultiSet uses RbTress for internal data structure, and keys can bee repeated.
type MultiSet struct {
	tree   *rbtree.RbTree
	keyCmp comparator.Comparator
	locker sync.Locker
}

// NewMultiSet new a MultiSet
func NewMultiSet(opts ...Options) *MultiSet {
	option := Option{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiSet{
		tree:   rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

// Insert inserts element to the MultiSet
func (ms *MultiSet) Insert(element interface{}) {
	ms.locker.Lock()
	defer ms.locker.Unlock()

	ms.tree.Insert(element, Empty)
}

// Erase erases all node with element in ms MultiSet
func (ms *MultiSet) Erase(element interface{}) {
	ms.locker.Lock()
	defer ms.locker.Unlock()

	node := ms.tree.FindNode(element)
	for node != nil && ms.keyCmp(node.Key(), element) == 0 {
		nextNode := node.Next()
		ms.tree.Delete(node)
		node = nextNode
	}
}

// Find returns the iterator related to element in the MultiSet,or an invalid iterator if not exist.
func (ms *MultiSet) Find(element interface{}) *SetIterator {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	node := ms.tree.FindNode(element)
	return &SetIterator{node: node}
}

//LowerBound returns the first iterator that equal or greater than element in the MultiSet
func (ms *MultiSet) LowerBound(element interface{}) *SetIterator {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	node := ms.tree.FindLowerBoundNode(element)
	return &SetIterator{node: node}
}

// Begin returns the iterator with the minimum element in the Set, return nil if empty.
func (ms *MultiSet) Begin() *SetIterator {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return ms.First()
}

// First returns the iterator with the minimum element in the MultiSet, return nil if empty.
func (ms *MultiSet) First() *SetIterator {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return &SetIterator{node: ms.tree.First()}
}

//Last returns the iterator with the maximum element in the MultiSet, return nil if empty.
func (ms *MultiSet) Last() *SetIterator {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return &SetIterator{node: ms.tree.Last()}
}

// Clear clears the MultiSet
func (ms *MultiSet) Clear() {
	ms.locker.Lock()
	defer ms.locker.Unlock()

	ms.tree.Clear()
}

// Contains returns true if element in the MultiSet. otherwise returns false.
func (ms *MultiSet) Contains(element interface{}) bool {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	if ms.tree.Find(element) != nil {
		return true
	}
	return false
}

// Size returns the size of MultiSet
func (ms *MultiSet) Size() int {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return ms.tree.Size()
}

// Traversal traversals elements in MultiSet, it will not stop until to the end or visitor returns false
func (ms *MultiSet) Traversal(visitor visitor.Visitor) {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	for node := ms.tree.First(); node != nil; node = node.Next() {
		if !visitor(node.Key()) {
			break
		}
	}
}

// String returns the set's elements in string format
func (ms *MultiSet) String() string {
	str := "["
	ms.Traversal(func(value interface{}) bool {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", value)
		return true
	})
	str += "]"
	return str
}
