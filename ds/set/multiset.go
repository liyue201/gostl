package set

import (
	. "github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/iterator"
)

type MultiSet struct {
	tree   *rbtree.RbTree
	keyCmp Comparator
}

func NewMultiSet(opts ...Options) *MultiSet {
	option := Option{
		keyCmp: defaultKeyComparator,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiSet{
		tree:   rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
	}
}

// Insert inserts element to the MultiSet
func (this *MultiSet) Insert(element interface{}) {
	this.tree.Insert(element, Empty)
}

// Erase erases all node with element in this MultiSet
func (this *MultiSet) Erase(element interface{}) {
	node := this.tree.FindNode(element)
	for node != nil && this.keyCmp(node.Key(), element) == 0 {
		nextNode := node.Next()
		this.tree.Delete(node)
		node = nextNode
	}
}

// Begin returns the ConstIterator related to element in the MultiSet,or an invalid iterator if not exist.
func (this *MultiSet) Find(element interface{}) ConstIterator {
	node := this.tree.FindNode(element)
	return &SetIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than element in the MultiSet
func (this *MultiSet) LowerBound(element interface{}) ConstIterator {
	node := this.tree.FindLowerBoundNode(element)
	return &SetIterator{node: node}
}

// Begin returns the ConstIterator with the minimum element in the Set, return nil if empty.
func (this *MultiSet) Begin() ConstIterator {
	return this.First()
}

// First returns the ConstIterator with the minimum element in the MultiSet, return nil if empty.
func (this *MultiSet) First() ConstBidIterator {
	return &SetIterator{node: this.tree.First()}
}

//Last returns the ConstIterator with the maximum element in the MultiSet, return nil if empty.
func (this *MultiSet) Last() ConstBidIterator {
	return &SetIterator{node: this.tree.Last()}
}

// Clear clears the MultiSet
func (this *MultiSet) Clear() {
	this.tree.Clear()
}

// Contains returns true if element in the MultiSet. otherwise returns false.
func (this *MultiSet) Contains(element interface{}) bool {
	if this.tree.Find(element) != nil {
		return true
	}
	return false
}

// Contains returns the size of MultiSet
func (this *MultiSet) Size() int {
	return this.tree.Size()
}
