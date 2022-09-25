package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// SetIterator is an iterator implementation of set
type SetIterator[T any] struct {
	node *rbtree.Node[T, bool]
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *SetIterator[K]) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next moves the pointer of the iterator to the next node and returns itself
func (iter *SetIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node and returns itself
func (iter *SetIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Value returns the element of the iterator point to
func (iter *SetIterator[T]) Value() T {
	return iter.node.Key()
}

// Clone clones the iterator into a new SetIterator
func (iter *SetIterator[T]) Clone() iterator.ConstIterator[T] {
	return &SetIterator[T]{iter.node}
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SetIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	otherIter, ok := other.(*SetIterator[T])
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
