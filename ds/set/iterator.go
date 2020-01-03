package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/iterator"
)

type SetIterator struct {
	node *rbtree.Node
}

func (iter *SetIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

func (iter *SetIterator) Next() ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

func (iter *SetIterator) Prev() ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

func (iter *SetIterator) Value() interface{} {
	return iter.node.Key()
}

func (iter *SetIterator) Clone() ConstIterator {
	return &SetIterator{iter.node}
}

func (iter *SetIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*SetIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
