package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/iterator"
)

type MapIterator struct {
	node *rbtree.Node
}

func (iter *MapIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

func (iter *MapIterator) Next() ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

func (iter *MapIterator) Prev() ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

func (iter *MapIterator) Key() interface{} {
	return iter.node.Key()
}

func (iter *MapIterator) Value() interface{} {
	return iter.node.Value()
}

func (iter *MapIterator) SetValue(val interface{}) error {
	iter.node.SetValue(val)
	return nil
}

func (iter *MapIterator) Clone() ConstIterator {
	return &MapIterator{iter.node}
}

func (iter *MapIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*MapIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
