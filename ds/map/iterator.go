package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/iterator"
)

type MapIterator struct {
	node *rbtree.Node
}

func (this *MapIterator) IsValid() bool {
	if this.node != nil {
		return true
	}
	return false
}

func (this *MapIterator) Next() ConstIterator {
	if this.IsValid() {
		this.node = this.node.Next()
	}
	return this
}

func (this *MapIterator) Prev() ConstBidIterator {
	if this.IsValid() {
		this.node = this.node.Prev()
	}
	return this
}

func (this *MapIterator) Key() interface{} {
	return this.node.Key()
}

func (this *MapIterator) Value() interface{} {
	return this.node.Value()
}

func (this *MapIterator) SetValue(val interface{}) error {
	this.node.SetValue(val)
	return nil
}

func (this *MapIterator) Clone() ConstIterator {
	return &MapIterator{this.node}
}

func (this *MapIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*MapIterator)
	if !ok {
		return false
	}
	if otherIter.node == this.node {
		return true
	}
	return false
}
