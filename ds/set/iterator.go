package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// SetIterator is an iterator implementation of set
type SetIterator struct {
	node *rbtree.Node
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *SetIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves the pointer of the iterator to the next node and returns itself
func (iter *SetIterator) Next() iterator.ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node and returns itself
func (iter *SetIterator) Prev() iterator.ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Value returns the element of the iterator point to
func (iter *SetIterator) Value() interface{} {
	return iter.node.Key()
}

// Clone clones the iterator into a new SetIterator
func (iter *SetIterator) Clone() iterator.ConstIterator {
	return &SetIterator{iter.node}
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SetIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*SetIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
