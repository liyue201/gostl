package set

import (
	. "github.com/liyue201/gostl/container"
	"github.com/liyue201/gostl/container/rbtree"
)

type SetIterator struct {
	node *rbtree.Node
}

func (this *SetIterator) Next() ConstIterator {
	return &SetIterator{
		node: this.node.Next(),
	}
}

func (this *SetIterator) Value() interface{} {
	return this.node.Key
}

func (this *SetIterator) Equal(other ConstIterator) bool {
	otherItr, ok := other.(*SetIterator)
	if !ok {
		return false
	}
	if this.node == otherItr.node {
		return true
	}
	return false
}
