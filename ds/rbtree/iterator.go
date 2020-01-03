package rbtree

import . "github.com/liyue201/gostl/utils/iterator"

type RbTreeIterator struct {
	node *Node
}

func NewIterator(node *Node) *RbTreeIterator {
	return &RbTreeIterator{node: node}
}

func (iter *RbTreeIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

func (iter *RbTreeIterator) Next() ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

func (iter *RbTreeIterator) Prev() ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

func (iter *RbTreeIterator) Key() interface{} {
	return iter.node.Key()
}

func (iter *RbTreeIterator) Value() interface{} {
	return iter.node.Value()
}

func (iter *RbTreeIterator) SetValue(val interface{}) error {
	iter.node.SetValue(val)
	return nil
}

func (iter *RbTreeIterator) Clone() ConstIterator {
	return NewIterator(iter.node)
}

func (iter *RbTreeIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*RbTreeIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
