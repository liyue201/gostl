package set

import (
	. "github.com/liyue201/gostl/container"
	"github.com/liyue201/gostl/container/rbtree"
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

//Insert inserts value to the Set
func (this *Set) Insert(value interface{}) {
	node := this.tree.FindNode(value)
	if node != nil {
		return
	}
	this.tree.Insert(value, Empty)
}

//Erase erases value in the Set
func (this *Set) Erase(value interface{}) {
	node := this.tree.FindNode(value)
	if node != nil {
		this.tree.Delete(node)
	}
}

//Begin returns the ConstIterator related to value in the Set, return nil if not exist.
func (this *Set) Find(value interface{}) ConstIterator {
	node := this.tree.FindNode(value)
	if node == nil {
		return nil
	}
	return &SetIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than value in the Set
func (this *Set) LowerBound(value interface{}) ConstIterator {
	node := this.tree.FindLowerBoundNode(value)
	return &SetIterator{node: node}
}

//Begin returns the ConstIterator with the min value in the Set, return nil if empty.
func (this *Set) Begin() ConstIterator {
	return &SetIterator{node: this.tree.Begin()}
}

//End returns ConstIterator with nil value in the Set
func (this *Set) End() ConstIterator {
	return &SetIterator{node: nil}
}

//Begin returns the ConstIterator with the max value in the Set, return nil if empty.
func (this *Set) RBegin() ConstIterator {
	return &SetReverseIterator{node: this.tree.RBegin()}
}

//REnd returns ConstIterator with nil value in the set
func (this *Set) REnd() ConstIterator {
	return &SetReverseIterator{node: nil}
}

//Clear clears the Set
func (this *Set) Clear() {
	this.tree.Clear()
}

// Contains returns true if value in the Set. otherwise returns false.
func (this *Set) Contains(value interface{}) bool {
	if this.tree.Find(value) != nil {
		return true
	}
	return false
}

// Contains returns the size of Set
func (this *Set) Size() int {
	return this.tree.Size()
}
