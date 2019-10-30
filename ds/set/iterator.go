package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/iterator"
)

type SetIterator struct {
	node *rbtree.Node
}

func (this *SetIterator) IsValid() bool {
	if this.node != nil {
		return true
	}
	return false
}

func (this *SetIterator) Next() ConstIterator {
	if this.IsValid() {
		this.node = this.node.Next()
	}
	return this
}

func (this *SetIterator) Prev() ConstBidIterator {
	if this.IsValid() {
		this.node = this.node.Prev()
	}
	return this
}

func (this *SetIterator) Value() interface{} {
	return this.node.Key()
}

func (this *SetIterator) Clone() ConstIterator {
	return &SetIterator{this.node}
}

func (this *SetIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*SetIterator)
	if !ok {
		return false
	}
	if otherIter.node == this.node {
		return true
	}
	return false
}

