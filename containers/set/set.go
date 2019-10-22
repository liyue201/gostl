package set

import (
	"github.com/liyue201/gostl/containers/rbtree"
	. "github.com/liyue201/gostl/uitls/comparator"
	. "github.com/liyue201/gostl/iterator"
)
 
const (
	Empty = 0
)

type Set struct {
	tree *rbtree.RbTree
}

func New(cmp Comparator) *Set {
	return &Set{tree: rbtree.New(cmp)}
}

// Insert inserts element to the Set
func (this *Set) Insert(element interface{}) {
	node := this.tree.FindNode(element)
	if node != nil {
		return
	}
	this.tree.Insert(element, Empty)
}

// Erase erases element in the Set
func (this *Set) Erase(element interface{}) {
	node := this.tree.FindNode(element)
	if node != nil {
		this.tree.Delete(node)
	}
}

// Begin returns the ConstIterator related to element in the Set, return nil if not exist.
func (this *Set) Find(element interface{}) ConstIterator {
	node := this.tree.FindNode(element)
	if node == nil {
		return nil
	}
	return &SetIterator{node: node}
}

// LowerBound returns the first ConstIterator that equal or greater than element in the Set
func (this *Set) LowerBound(element interface{}) ConstIterator {
	node := this.tree.FindLowerBoundNode(element)
	return &SetIterator{node: node}
}

// Begin returns the ConstIterator with the minimum element in the Set, return nil if empty.
func (this *Set) Begin() ConstIterator {
	return this.First()
}

// First returns the ConstIterator with the minimum element in the Set, return nil if empty.
func (this *Set) First() ConstBidIterator {
	return &SetIterator{node: this.tree.First()}
}

// Last returns the ConstIterator with the maximum element in the Set, return nil if empty.
func (this *Set) Last() ConstBidIterator {
	return &SetIterator{node: this.tree.Last()}
}

// Clear clears the Set
func (this *Set) Clear() {
	this.tree.Clear()
}

// Contains returns true if element in the Set. otherwise returns false.
func (this *Set) Contains(element interface{}) bool {
	if this.tree.Find(element) != nil {
		return true
	}
	return false
}

// Contains returns the size of Set
func (this *Set) Size() int {
	return this.tree.Size()
}
