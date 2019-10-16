package treemap

import (
	. "github.com/liyue201/gostl/container"
	"github.com/liyue201/gostl/container/rbtree"
)

type MapIterator struct {
	node *rbtree.Node
}

func (this *MapIterator) Next() ConstIterator {
	return &MapIterator{
		node: this.node.Next(),
	}
}

func (this *MapIterator) Value() interface{} {
	return this.node.Value
}

func (this *MapIterator) Equal(other ConstIterator) bool {
	otherItr, ok := other.(*MapIterator)
	if !ok {
		return false
	}
	if this.node == otherItr.node {
		return true
	}
	return false
}

type MapReverseIterator struct {
	node *rbtree.Node
}

func (this *MapReverseIterator) Next() ConstIterator {
	return &MapReverseIterator{
		node: this.node.Prev(),
	}
}

func (this *MapReverseIterator) Value() interface{} {
	return this.node.Value
}

func (this *MapReverseIterator) Equal(other ConstIterator) bool {
	otherItr, ok := other.(*MapReverseIterator)
	if !ok {
		return false
	}
	if this.node == otherItr.node {
		return true
	}
	return false
}
