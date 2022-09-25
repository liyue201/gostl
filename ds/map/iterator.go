package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// MapIterator is a map iterator
type MapIterator[K, V any] struct {
	node *rbtree.Node[K, V]
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *MapIterator[K, V]) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *MapIterator[K, V]) Next() iterator.ConstIterator[V] {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node, and returns itseft
func (iter *MapIterator[K, V]) Prev() iterator.ConstBidIterator[V] {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the node's key of the iterator point to
func (iter *MapIterator[K, V]) Key() K {
	return iter.node.Key()
}

// Value returns the node's value of the iterator point to
func (iter *MapIterator[K, V]) Value() V {
	return iter.node.Value()
}

// SetValue sets the node's value of the iterator point to
func (iter *MapIterator[K, V]) SetValue(val V) {
	iter.node.SetValue(val)
}

// Clone clones the iterator to a new MapIterator
func (iter *MapIterator[K, V]) Clone() iterator.ConstIterator[V] {
	return &MapIterator[K, V]{iter.node}
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *MapIterator[K, V]) Equal(other iterator.ConstIterator[V]) bool {
	otherIter, ok := other.(*MapIterator[K, V])
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
