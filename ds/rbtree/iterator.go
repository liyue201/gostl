package rbtree

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator[K, V any] struct {
	node *Node[K, V]
}

// NewIterator creates a RbTreeIterator from the passed node
func NewIterator[K, V any](node *Node[K, V]) *RbTreeIterator[K, V] {
	return &RbTreeIterator[K, V]{node: node}
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *RbTreeIterator[K, V]) IsValid() bool {
	return iter.node != nil
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *RbTreeIterator[K, V]) Next() iterator.ConstIterator[V] {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node, and returns itself
func (iter *RbTreeIterator[K, V]) Prev() iterator.ConstBidIterator[V] {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the node's key of the iterator point to
func (iter *RbTreeIterator[K, V]) Key() K {
	return iter.node.Key()
}

// Value returns the node's value of the iterator point to
func (iter *RbTreeIterator[K, V]) Value() V {
	return iter.node.Value()
}

// SetValue sets the node's value of the iterator point to
func (iter *RbTreeIterator[K, V]) SetValue(val V) error {
	iter.node.SetValue(val)
	return nil
}

// Clone clones the iterator into a new RbTreeIterator
func (iter *RbTreeIterator[K, V]) Clone() iterator.ConstIterator[V] {
	return NewIterator(iter.node)
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *RbTreeIterator[K, V]) Equal(other iterator.ConstIterator[V]) bool {
	otherIter, ok := other.(*RbTreeIterator[K, V])
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
