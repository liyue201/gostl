package simple_list

import . "github.com/liyue201/gostl/utils/iterator"

//ListIterator is an implementation Iterator
var _ Iterator = (*ListIterator)(nil)

type ListIterator struct {
	node *Node
}

func NewIterator(node *Node) *ListIterator {
	return &ListIterator{node: node}
}

func (iter *ListIterator) IsValid() bool {
	return iter.node != nil
}

func (iter *ListIterator) Next() ConstIterator {
	if iter.node != nil {
		iter.node = iter.node.Next()
	}
	return iter
}

func (iter *ListIterator) Value() interface{} {
	if iter.node == nil {
		return nil
	}
	return iter.node.Value
}

func (iter *ListIterator) SetValue(value interface{}) error {
	if iter.node != nil {
		iter.node.Value = value
	}
	return nil
}

func (iter *ListIterator) Clone() ConstIterator {
	return NewIterator(iter.node)
}

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
