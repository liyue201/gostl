package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// SetIterator is an iterator implementation of set
type SetIterator struct {
	node *rbtree.Node
}

// IsValid returns whether iter is valid or not
func (iter *SetIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves iter to next node and returns iter
func (iter *SetIterator) Next() iterator.ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves iter to previous node and returns iter
func (iter *SetIterator) Prev() iterator.ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Value returns the internal value of iter
func (iter *SetIterator) Value() interface{} {
	return iter.node.Key()
}

// Clone clones iter to a new SetIterator
func (iter *SetIterator) Clone() iterator.ConstIterator {
	return &SetIterator{iter.node}
}

// Equal returns whether iter is equal to other or not
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
