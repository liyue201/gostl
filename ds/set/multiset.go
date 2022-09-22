package set

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// MultiSet uses RbTress for internal data structure, and keys can bee repeated.
type MultiSet[T any] struct {
	tree   *rbtree.RbTree[T, bool]
	keyCmp comparator.Comparator
	locker sync.Locker
}

// NewMultiSet creates a new MultiSet
func NewMultiSet[T any](opts ...Option) *MultiSet[T] {
	option := Options{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiSet[T]{
		tree:   rbtree.New[T, bool](rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

// Insert inserts an element to the MultiSet
func (ms *MultiSet[T]) Insert(element T) {
	ms.locker.Lock()
	defer ms.locker.Unlock()

	ms.tree.Insert(element, Empty)
}

// Erase erases all node with passed element in the MultiSet
func (ms *MultiSet[T]) Erase(element T) {
	ms.locker.Lock()
	defer ms.locker.Unlock()

	for {
		node := ms.tree.FindNode(element)
		if node == nil {
			break
		}
		ms.tree.Delete(node)
	}
}

// Find finds the first element that is equal to the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) Find(element T) *SetIterator[T] {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	node := ms.tree.FindNode(element)
	return &SetIterator[T]{node: node}
}

//LowerBound finds the first element that is equal to or greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) LowerBound(element T) *SetIterator[T] {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	node := ms.tree.FindLowerBoundNode(element)
	return &SetIterator[T]{node: node}
}

//UpperBound finds the first element that is greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) UpperBound(element T) *SetIterator[T] {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	node := ms.tree.FindUpperBoundNode(element)
	return &SetIterator[T]{node: node}
}

// Begin returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet[T]) Begin() *SetIterator[T] {
	return ms.First()
}

// First returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet[T]) First() *SetIterator[T] {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return &SetIterator[T]{node: ms.tree.First()}
}

//Last returns the iterator with the maximum element in the MultiSet
func (ms *MultiSet[T]) Last() *SetIterator[T] {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return &SetIterator[T]{node: ms.tree.Last()}
}

// Clear clears all elements in the MultiSet
func (ms *MultiSet[T]) Clear() {
	ms.locker.Lock()
	defer ms.locker.Unlock()

	ms.tree.Clear()
}

// Contains returns true if the passed element is in the MultiSet. otherwise returns false.
func (ms *MultiSet[T]) Contains(element T) bool {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	if _, err := ms.tree.Find(element); err == nil {
		return true
	}
	return false
}

// Size returns the amount of elements in the MultiSet
func (ms *MultiSet[T]) Size() int {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	return ms.tree.Size()
}

// Traversal traversals elements in the MultiSet, it will not stop until to the end of the MultiSet or the visitor returns false
func (ms *MultiSet[T]) Traversal(visitor visitor.Visitor[T]) {
	ms.locker.RLock()
	defer ms.locker.RUnlock()

	for node := ms.tree.First(); node != nil; node = node.Next() {
		if !visitor(node.Key()) {
			break
		}
	}
}

// String returns s string representation of the MultiSet
func (ms *MultiSet[T]) String() string {
	str := "["
	ms.Traversal(func(value T) bool {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", value)
		return true
	})
	str += "]"
	return str
}
