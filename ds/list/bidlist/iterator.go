package bidlist

import (
	"github.com/liyue201/gostl/utils/iterator"
)

//ListIterator is an implementation of BidIterator
var _ iterator.BidIterator[T] = (*ListIterator[T])(nil)

// ListIterator is an implementation of list iterator
type ListIterator[T any] struct {
	node *Node[T]
}

// NewIterator creates a ListIterator
func NewIterator[T any](node *Node[T]) *ListIterator[T] {
	return &ListIterator[T]{node: node}
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *ListIterator[T]) IsValid() bool {
	return iter.node != nil
}

// Next moves the pointer of iterator to the next node and returns itself
func (iter *ListIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.node != nil {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of iterator to the previous node and returns itself
func (iter *ListIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.node != nil {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Value returns the node's value of the iterator point to
func (iter *ListIterator[T]) Value() T {
	if iter.node == nil {
		panic("invalid iterator")
	}
	return iter.node.Value
}

// SetValue sets the node's value of the iterator point to
func (iter *ListIterator[T]) SetValue(value T) {
	if iter.node != nil {
		iter.node.Value = value
	}
}

// Clone clones the iterator to a new iterator
func (iter *ListIterator[T]) Clone() iterator.ConstIterator[T] {
	return NewIterator[T](iter.node)
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ListIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	otherIter, ok := other.(*ListIterator[T])
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
