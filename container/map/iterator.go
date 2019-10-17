package treemap

import (
	. "github.com/liyue201/gostl/container"
	"github.com/liyue201/gostl/container/rbtree"
)

type MapIterator struct {
	node *rbtree.Node
}

func (this *MapIterator) Next() ConstKvIterator {
	return &MapIterator{
		node: this.node.Next(),
	}
}

func (this *MapIterator) Key() interface{} {
	return this.node.Key()
}

func (this *MapIterator) Value() interface{} {
	return this.node.Value()
}

func (this *MapIterator) Equal(other ConstKvIterator) bool {
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

func (this *MapReverseIterator) Next() ConstKvIterator {
	return &MapReverseIterator{
		node: this.node.Prev(),
	}
}

func (this *MapReverseIterator) Key() interface{} {
	return this.node.Key()
}

func (this *MapReverseIterator) Value() interface{} {
	return this.node.Value()
}

func (this *MapReverseIterator) Equal(other ConstKvIterator) bool {
	otherItr, ok := other.(*MapReverseIterator)
	if !ok {
		return false
	}
	if this.node == otherItr.node {
		return true
	}
	return false
}
