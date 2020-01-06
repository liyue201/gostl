package rbtree

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator struct {
	node *Node
}

// NewIterator news a RbTreeIterator from node
func NewIterator(node *Node) *RbTreeIterator {
	return &RbTreeIterator{node: node}
}

// IsValid returns whether iter is valid or not
func (iter *RbTreeIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves iter to next node and returns iter
func (iter *RbTreeIterator) Next() iterator.ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves iter to previous node and returns iter
func (iter *RbTreeIterator) Prev() iterator.ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the internal key of iter
func (iter *RbTreeIterator) Key() interface{} {
	return iter.node.Key()
}

// Value returns the internal value of iter
func (iter *RbTreeIterator) Value() interface{} {
	return iter.node.Value()
}

func (iter *RbTreeIterator) SetValue(val interface{}) error {
	iter.node.SetValue(val)
	return nil
}

// Clone clones iter to a new RbTreeIterator
func (iter *RbTreeIterator) Clone() iterator.ConstIterator {
	return NewIterator(iter.node)
}

// Equal returns whether iter is equal to other or not
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
