package simple_list

import . "github.com/liyue201/gostl/utils/iterator"

//ListIterator is an implementation Iterator
var _ Iterator = (*ListIterator)(nil)

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
func (iter *ListIterator) Next() ConstIterator {
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
func (iter *ListIterator) SetValue(value interface{}) error {
	if iter.node != nil {
		iter.node.Value = value
	}
	return nil
}

// Clone clones iter to a new ListIterator
func (iter *ListIterator) Clone() ConstIterator {
	return NewIterator(iter.node)
}

// Equal returns whether iter is equal to other
func (iter *ListIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*ListIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
