package set

import (
	"github.com/liyue201/gostl/containers/rbtree"
	. "github.com/liyue201/gostl/uitls/iterator"
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

func (this *SetIterator) Clone() interface{} {
	return &SetIterator{this.node}
}
