package rbtree

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator struct {
	node *Node
}

// NewIterator creates a RbTreeIterator from the passed node
func NewIterator(node *Node) *RbTreeIterator {
	return &RbTreeIterator{node: node}
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *RbTreeIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *RbTreeIterator) Next() iterator.ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node, and returns itself
func (iter *RbTreeIterator) Prev() iterator.ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the node's key of the iterator point to
func (iter *RbTreeIterator) Key() interface{} {
	return iter.node.Key()
}

// Value returns the node's value of the iterator point to
func (iter *RbTreeIterator) Value() interface{} {
	return iter.node.Value()
}

//SetValue sets the node's value of the iterator point to
func (iter *RbTreeIterator) SetValue(val interface{}) error {
	iter.node.SetValue(val)
	return nil
}

// Clone clones the iterator into a new RbTreeIterator
func (iter *RbTreeIterator) Clone() iterator.ConstIterator {
	return NewIterator(iter.node)
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *RbTreeIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*RbTreeIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
