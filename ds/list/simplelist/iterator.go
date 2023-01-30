package simplelist

import "github.com/liyue201/gostl/utils/iterator"

//ListIterator is an implementation of Iterator
var _ iterator.Iterator[any] = (*ListIterator[any])(nil)

// ListIterator is an iterator for list
type ListIterator[T any] struct {
	node *Node[T]
}

// NewIterator news a ListIterator
func NewIterator[T any](node *Node[T]) *ListIterator[T] {
	return &ListIterator[T]{node: node}
}

// IsValid returns whether iter is valid
func (iter *ListIterator[T]) IsValid() bool {
	return iter.node != nil
}

// Next returns the next iterator
func (iter *ListIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.node != nil {
		iter.node = iter.node.Next()
	}
	return iter
}

// Value returns the internal value of iter
func (iter *ListIterator[T]) Value() T {
	if iter.node == nil {
		panic("invalid iterator")
	}
	return iter.node.Value
}

// SetValue sets the value of iter
func (iter *ListIterator[T]) SetValue(value T) {
	if iter.node != nil {
		iter.node.Value = value
	}
}

// Clone clones iter to a new ListIterator
func (iter *ListIterator[T]) Clone() iterator.ConstIterator[T] {
	return NewIterator(iter.node)
}

// Equal returns whether iter is equal to other
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
