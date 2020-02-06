package simplelist

import "github.com/liyue201/gostl/utils/iterator"

//ListIterator is an implementation of Iterator
var _ iterator.Iterator = (*ListIterator)(nil)

// ListIterator is an iterator for list
type ListIterator struct {
	node *Node
}

// NewIterator news a ListIterator
func NewIterator(node *Node) *ListIterator {
	return &ListIterator{node: node}
}

// IsValid returns whether iter is valid
func (iter *ListIterator) IsValid() bool {
	return iter.node != nil
}

// Next returns the next iterator
func (iter *ListIterator) Next() iterator.ConstIterator {
	if iter.node != nil {
		iter.node = iter.node.Next()
	}
	return iter
}

// Value returns the internal value of iter
func (iter *ListIterator) Value() interface{} {
	if iter.node == nil {
		return nil
	}
	return iter.node.Value
}

// SetValue sets the value of iter
func (iter *ListIterator) SetValue(value interface{}) {
	if iter.node != nil {
		iter.node.Value = value
	}
}

// Clone clones iter to a new ListIterator
func (iter *ListIterator) Clone() iterator.ConstIterator {
	return NewIterator(iter.node)
}

// Equal returns whether iter is equal to other
func (iter *ListIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*ListIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
