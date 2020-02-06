package bidlist

import "github.com/liyue201/gostl/utils/iterator"

//ListIterator is an implementation of BidIterator
var _ iterator.BidIterator = (*ListIterator)(nil)

// ListIterator is an implementation of list iterator
type ListIterator struct {
	node *Node
}

// NewIterator creates a ListIterator
func NewIterator(node *Node) *ListIterator {
	return &ListIterator{node: node}
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *ListIterator) IsValid() bool {
	return iter.node != nil
}

// Next moves the pointer of iterator to the next node and returns itself
func (iter *ListIterator) Next() iterator.ConstIterator {
	if iter.node != nil {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of iterator to the previous node and returns itself
func (iter *ListIterator) Prev() iterator.ConstBidIterator {
	if iter.node != nil {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Value returns the node's value of the iterator point to
func (iter *ListIterator) Value() interface{} {
	if iter.node == nil {
		return nil
	}
	return iter.node.Value
}

// SetValue sets the node's value of the iterator point to
func (iter *ListIterator) SetValue(value interface{}) {
	if iter.node != nil {
		iter.node.Value = value
	}
}

// Clone clones the iterator to a new iterator
func (iter *ListIterator) Clone() iterator.ConstIterator {
	return NewIterator(iter.node)
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
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
