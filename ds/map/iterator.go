package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// MapIterator is a map iterator
type MapIterator struct {
	node *rbtree.Node
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *MapIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *MapIterator) Next() iterator.ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node, and returns itseft
func (iter *MapIterator) Prev() iterator.ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the node's key of the iterator point to
func (iter *MapIterator) Key() interface{} {
	return iter.node.Key()
}

// Value returns the node's value of the iterator point to
func (iter *MapIterator) Value() interface{} {
	return iter.node.Value()
}

// SetValue sets the node's value of the iterator point to
func (iter *MapIterator) SetValue(val interface{}) {
	iter.node.SetValue(val)
}

// Clone clones the iterator to a new MapIterator
func (iter *MapIterator) Clone() iterator.ConstIterator {
	return &MapIterator{iter.node}
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *MapIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*MapIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
